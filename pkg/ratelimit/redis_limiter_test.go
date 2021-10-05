// Copyright 2021 Comcast Cable Communications Management, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build integration
// +build integration

package ratelimit_test

import (
	"github.com/xmidt-org/ears/pkg/ratelimit/redis"
	"github.com/xmidt-org/ears/pkg/tenant"
	"testing"
)

func TestRedisBackendLimiter(t *testing.T) {
	limiter := redis.NewRedisRateLimiter(
		tenant.Id{"myOrg", "myUnitTestApp"},
		"localhost:6379",
		0,
	)

	testBackendLimiter(limiter, t)
}
