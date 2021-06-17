/**
 *  Copyright (c) 2020  Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package app

import (
	"context"
	"fmt"
	"github.com/lightstep/otel-launcher-go/launcher"
	"github.com/rs/zerolog"
	"github.com/xmidt-org/ears/internal/pkg/config"
	"github.com/xmidt-org/ears/internal/pkg/rtsemconv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"go.uber.org/fx"
	"net/http"
)

func NewMux(a *APIManager, middleware []func(next http.Handler) http.Handler) (http.Handler, error) {
	for _, m := range middleware {
		a.muxRouter.Use(m)
	}
	return a.muxRouter, nil
}

func SetupAPIServer(lifecycle fx.Lifecycle, config config.Config, logger *zerolog.Logger, mux http.Handler) error {
	port := config.GetInt("ears.api.port")
	if port < 1 {
		err := &InvalidOptionError{fmt.Sprintf("invalid port value %d", port)}
		logger.Error().Msg(err.Error())
		return err
	}

	server := &http.Server{
		Addr:    ":" + config.GetString("ears.api.port"),
		Handler: mux,
	}

	var ls launcher.Launcher
	var traceProvider *sdktrace.TracerProvider
	ctx := context.Background() // long lived context

	if config.GetBool("ears.opentelemetry.lightstep.active") {
		ls = launcher.ConfigureOpentelemetry(
			launcher.WithServiceName(rtsemconv.EARSServiceName),
			launcher.WithAccessToken(config.GetString("ears.opentelemetry.lightstep.accessToken")),
			launcher.WithServiceVersion("1.0"),
		)
		logger.Info().Str("telemetryexporter", "lightstep").Msg("started")
	}
	if config.GetBool("ears.opentelemetry.datadog.active") {
		exporter, err := otlp.NewExporter(ctx,
			otlpgrpc.NewDriver(
				otlpgrpc.WithEndpoint("localhost:55680"),
				otlpgrpc.WithInsecure(),
			),
		)
		if err != nil {
			return err
		}
		traceProvider = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(
				resource.NewWithAttributes(
					semconv.ServiceNameKey.String(rtsemconv.EARSServiceName),
				),
			),
		)
		otel.SetTracerProvider(traceProvider)
		//global.SetMeterProvider(pusher.MeterProvider())
		propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
		otel.SetTextMapPropagator(propagator)
		logger.Info().Str("telemetryexporter", "otel").Msg("started")
	}
	if config.GetBool("ears.opentelemetry.stdout.active") {
		exporter, err := stdout.NewExporter(
			stdout.WithPrettyPrint(),
		)
		if err != nil {
			return err
		}
		bsp := sdktrace.NewBatchSpanProcessor(exporter)
		traceProvider = sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))
		otel.SetTracerProvider(traceProvider)
		//global.SetMeterProvider(pusher.MeterProvider())
		propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
		otel.SetTextMapPropagator(propagator)
		logger.Info().Str("telemetryexporter", "stdout").Msg("started")
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go server.ListenAndServe()
				logger.Info().Str("port", fmt.Sprintf("%d", port)).Msg("API Server Started")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				err := server.Shutdown(ctx)
				if err != nil {
					logger.Error().Str("op", "SetupAPIServer.OnStop").Msg(err.Error())
				} else {
					logger.Info().Msg("API Server Stopped")
				}
				if config.GetBool("ears.opentelemetry.lightstep.active") {
					ls.Shutdown()
					logger.Info().Msg("lightstep exporter stopped")
				}
				if config.GetBool("ears.opentelemetry.datadog.active") {
					traceProvider.Shutdown(ctx)
					logger.Info().Msg("otel exporter stopped")
				}
				return nil
			},
		},
	)
	return nil
}
