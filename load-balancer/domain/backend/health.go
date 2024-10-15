package backend

import (
	"net/http"
	"time"
)

func NewHealthCheck(bm *Manager, interval time.Duration) *HealthCheck {
	return &HealthCheck{BackendManager: bm, Interval: interval}
}

func (hc *HealthCheck) StartHealthCheck() {
	ticker := time.NewTicker(hc.Interval)
	go func() {
		for range ticker.C {
			for _, backend := range hc.BackendManager.backends {
				resp, err := http.Get(backend.URL + "/health")
				if err != nil || resp.StatusCode != http.StatusOK {
					hc.BackendManager.MarkBackendUnhealthy(backend.URL)
				}
			}
		}
	}()
}
