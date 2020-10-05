package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/xmidt-org/ears/internal"

	"github.com/rs/zerolog/log"
)

var (
	ROUTE = `
	{
		"org_id" : "comcast",
		"app_id" : "xfi",
		"user_id" : "boris",
		"src_type" : "debug",
		"src_params" : {},
		"dst_type" : "debug",
		"dst_params" : {},
		"filter_chain" : [
			{
				"type" : "match",
				"params" : {}
			},
			{
				"type" : "filter",
				"params" : {}
			},
			{
				"type" : "transform",
				"params" : {}
			}
		],
		"delivery_mode" : "at_least_once"
	}
	`
)

func main() {
	ctx := context.Background()
	var rtmgr internal.RoutingTableManager
	rtmgr = internal.NewInMemoryRoutingTableManager()
	log.Debug().Msg(fmt.Sprintf("ears has %d routes", rtmgr.GetRouteCount(ctx)))
	var rte internal.RoutingTableEntry
	err := json.Unmarshal([]byte(ROUTE), &rte)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	err = rtmgr.AddRoute(ctx, &rte)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("ears has %d routes", rtmgr.GetRouteCount(ctx)))
	allRoutes, err := rtmgr.GetAllRoutes(ctx)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	if len(allRoutes) > 0 {
		log.Debug().Msg(fmt.Sprintf("first route has hash %s", allRoutes[0].Hash(ctx)))
	}
}
