package agent

import (
	"fmt"
	"log"
	"time"
)

type Agent struct {
	Sender    *Sender
	Collector *Collector
}

func NewAgent(sender *Sender, collector *Collector) *Agent {
	return &Agent{
		Sender:    sender,
		Collector: collector,
	}
}

func (a *Agent) Run(pollTicker, reportTicker *time.Ticker) {
	log.Println("Agent started")
	var last Record

	for {
		select {
		case <-pollTicker.C:
			last = a.Collector.Collect()
		case <-reportTicker.C:
			log.Println("Sending metrics to server...")
			for name, value := range last.Gauges {
				if err := a.Sender.SendMetric("gauge", name, fmt.Sprintf("%f", value)); err != nil {
					log.Printf("send gauge %s failed: %v", name, err)
				}
			}
			for name, value := range last.Counters {
				if err := a.Sender.SendMetric("counter", name, fmt.Sprintf("%d", value)); err != nil {
					log.Printf("send counter %s failed: %v", name, err)
				}
			}

		}
	}
}
