package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Maxim-hash/go-musthave-metrics-tpl/internal/agent"
)

func main() {
	const (
		pollInterval   = 2 * time.Second
		reportInterval = 10 * time.Second
		serverAddr     = "http://localhost:8080"
	)

	collector := agent.NewCollector()
	sender := agent.NewSender(serverAddr)

	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	var last agent.Record

	for {
		select {
		case <-pollTicker.C:
			last = collector.Collect()
		case <-reportTicker.C:
			for name, value := range last.Gauges {
				if err := sender.SendMetric("gauge", name, fmt.Sprintf("%f", value)); err != nil {
					log.Printf("send gauge %s failed: %v", name, err)
				}
			}
			for name, value := range last.Counters {
				if err := sender.SendMetric("counter", name, fmt.Sprintf("%d", value)); err != nil {
					log.Printf("send counter %s failed: %v", name, err)
				}
			}

		}
	}
}
