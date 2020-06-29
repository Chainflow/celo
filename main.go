package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"

	"github.com/chainflow/celo/config"
	"github.com/chainflow/celo/server"
)

func main() {
	cfg, err := config.ReadFromFile()
	if err != nil {
		log.Fatal(err)
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://localhost:%s", cfg.InfluxDB.Port),
		Username: cfg.InfluxDB.Username,
		Password: cfg.InfluxDB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	// Calling go routine to send alerts of new proposals
	go func() {
		for {
			if err := server.NewProposalAlert(cfg,c); err != nil {
				fmt.Println("Error while sending new proposal alert", err)
			}
			time.Sleep(3 * time.Second)
		}
	}()

	wg.Wait()
}
