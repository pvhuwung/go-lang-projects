package traffic

import "load-balancer/domain/backend"

type Handler struct {
	BackendManager *backend.Manager
}

func NewTrafficHandler(backendManager *backend.Manager) *Handler {
	return &Handler{BackendManager: backendManager}
}
