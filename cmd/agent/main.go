package main

import (
	"time"

	"github.com/Maxim-hash/go-musthave-metrics-tpl/internal/agent"
)

func main() {
	flags := agent.ParseFlags()

	var (
		pollInterval   = time.Duration(flags.FlagPollInterval) * time.Second
		reportInterval = time.Duration(flags.FlagReportInterval) * time.Second
		serverAddr     = flags.FlagServerAddr
	)

	collector := agent.NewCollector()
	sender := agent.NewSender(serverAddr)

	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	agent := agent.NewAgent(sender, collector)

	agent.Run(pollTicker, reportTicker)

}
