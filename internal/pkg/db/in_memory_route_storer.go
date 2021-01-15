package db

import (
	"context"
	"github.com/xmidt-org/ears/internal/pkg/app"
	"github.com/xmidt-org/ears/pkg/route"
	"sync"
	"time"
)

type InMemoryRouteStorer struct {
	routes map[string]*route.Config
	lock   *sync.RWMutex
}

func NewInMemoryRouteStorer(config app.Config) *InMemoryRouteStorer {
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
		return empty, &route.RouteNotFoundError{id}
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
