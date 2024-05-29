package main

import (
	"context"
	"log"

	"github.com/nireitdev/go-ntpd-server-statistics/config"
	"github.com/nireitdev/go-ntpd-server-statistics/traffic"
)

// Global vars
var (
	cfg *config.Config
	ctx context.Context
)

func main() {
	ctx = context.Background()

	cfg = config.ReadConfig()

	//Captura de trafico:
	log.Println("Starting...")
	captures := traffic.NewCapture(traffic.Traffic{
		Ctx:       ctx,
		Device:    cfg.Config.Device,
		Ip:        cfg.Config.Ip,
		Timerange: cfg.Config.Timerange,
	})

	for cap := range captures {
		log.Printf("Count: %d \n", cap.Count)

	}
}
