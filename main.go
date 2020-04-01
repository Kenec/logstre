package main

import (
	"fmt"
	"github.com/Kenec/logstre/metrics"
	"time"
)

func main() {
	fmt.Println("Starting.....")

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Metrics at", t)
				metrics.CpuUsage()
			}
		}
	}()

	select{}
}
