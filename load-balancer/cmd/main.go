package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"load-balancer/domain/backend"
	"load-balancer/domain/traffic"
	"load-balancer/infra"
	"os"
	"time"
)

type Config struct {
	LoadBalancer struct {
		Port        int      `json:"port"`
		Backends    []string `json:"backends"`
		HealthCheck struct {
			IntervalSeconds int `json:"intervalSeconds"`
		} `json:"healthCheck"`
		Algorithm string `json:"algorithm"`
	} `json:"loadBalancer"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	backendManager := backend.NewBackendManager(config.LoadBalancer.Backends)

	healthCheckInterval := time.Duration(config.LoadBalancer.HealthCheck.IntervalSeconds) * time.Second
	healthCheck := backend.NewHealthCheck(backendManager, healthCheckInterval)
	healthCheck.StartHealthCheck()

	trafficHandler := traffic.NewTrafficHandler(backendManager)

	port := config.LoadBalancer.Port
	if port == 0 {
		port = 8080
	}
	fmt.Printf("Starting load balancer on port %d...\n", port)
	infra.StartHTTPServer(trafficHandler, port)
}
