package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Maxim-hash/go-musthave-metrics-tpl/internal/agent"
)

func main() {
	ParseFlags()

	var (
		pollInterval   = time.Duration(flagPollInterval) * time.Second
		reportInterval = time.Duration(flagReportInterval) * time.Second
		serverAddr     = flagServerAddr
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
			log.Println("Sending metrics to server...")
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
