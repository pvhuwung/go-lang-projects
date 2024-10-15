package infra

import (
	"fmt"
	"load-balancer/domain/traffic"
	"net/http"
)

func StartHTTPServer(handler *traffic.Handler, port int) {
	http.HandleFunc("/", handler.HandleRequest)
	address := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
