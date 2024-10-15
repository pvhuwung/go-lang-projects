package backend

import (
	"errors"
)

func NewBackendManager(backends []string) *Manager {
	b := make([]*Backend, len(backends))
	for i, url := range backends {
		b[i] = &Backend{URL: url, Healthy: true}
	}
	return &Manager{backends: b}
}

func (bm *Manager) GetAvailableBackend() (*Backend, error) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	for i := 0; i < len(bm.backends); i++ {
		bm.current = (bm.current + 1) % len(bm.backends)
		if bm.backends[bm.current].Healthy {
			return bm.backends[bm.current], nil
		}
	}
	return nil, errors.New("no healthy backends available")
}

func (bm *Manager) MarkBackendUnhealthy(url string) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	for _, b := range bm.backends {
		if b.URL == url {
			b.Healthy = false
		}
	}
}
