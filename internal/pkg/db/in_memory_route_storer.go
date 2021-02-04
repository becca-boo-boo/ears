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

package db

import (
	"context"
	"github.com/xmidt-org/ears/pkg/route"
	"sync"
	"time"
)

type InMemoryRouteStorer struct {
	routes map[string]*route.Config
	lock   *sync.RWMutex
}

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

func NewInMemoryRouteStorer(config Config) *InMemoryRouteStorer {
	return &InMemoryRouteStorer{
		routes: make(map[string]*route.Config),
		lock:   &sync.RWMutex{},
	}
}

func (s *InMemoryRouteStorer) GetRoute(ctx context.Context, id string) (route.Config, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	empty := route.Config{}

	r, ok := s.routes[id]
	if !ok {
		return empty, &route.RouteNotFoundError{RouteId: id}
	}

	newCopy := *r
	return newCopy, nil
}

func (s *InMemoryRouteStorer) GetAllRoutes(ctx context.Context) ([]route.Config, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	routes := make([]route.Config, 0)
	for _, r := range s.routes {
		routes = append(routes, *r)
	}
	return routes, nil
}

func (s *InMemoryRouteStorer) setRoute(r route.Config) {
	r.Modified = time.Now().Unix()
	if existing, ok := s.routes[r.Id]; !ok {
		r.Created = r.Modified
	} else {
		r.Created = existing.Created
	}
	s.routes[r.Id] = &r
}

func (s *InMemoryRouteStorer) SetRoute(ctx context.Context, r route.Config) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.setRoute(r)
	return nil
}

func (s *InMemoryRouteStorer) SetRoutes(ctx context.Context, routes []route.Config) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, r := range routes {
		s.setRoute(r)
	}
	return nil
}

func (s *InMemoryRouteStorer) DeleteRoute(ctx context.Context, id string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.routes, id)
	return nil
}

func (s *InMemoryRouteStorer) DeleteRoutes(ctx context.Context, ids []string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, id := range ids {
		delete(s.routes, id)
	}
	return nil
}
