package backend

import (
	"sync"
	"time"
)

type Backend struct {
	URL     string
	Healthy bool
}

type Manager struct {
	backends []*Backend
	mu       sync.Mutex
	current  int
}
type HealthCheck struct {
	BackendManager *Manager
	Interval       time.Duration
}
