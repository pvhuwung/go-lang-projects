package traffic

import (
	"net/http"
)

func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	backend, err := h.BackendManager.GetAvailableBackend()
	if err != nil {
		http.Error(w, "No available backend", http.StatusServiceUnavailable)
		return
	}

	proxy := &http.Transport{}
	req, _ := http.NewRequest(r.Method, backend.URL+r.RequestURI, r.Body)
	resp, err := proxy.RoundTrip(req)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, value := range resp.Header {
		w.Header().Set(key, value[0])
	}
	w.WriteHeader(resp.StatusCode)
	resp.Write(w)
}
