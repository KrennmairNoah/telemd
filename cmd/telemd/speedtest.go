package main

import (
	"fmt"

	"github.com/kylegrantlucas/speedtest"
)

func main() {
	client, err := speedtest.NewDefaultClient()
	if err != nil {
		fmt.Printf("error creating client: %v", err)
	}

	// Pass an empty string to select the fastest server
	server, err := client.GetServer("")
	if err != nil {
		fmt.Printf("error getting server: %v", err)
	}

	dmbps, err := client.Download(server)
	if err != nil {
		fmt.Printf("error getting download: %v", err)
	}

	umbps, err := client.Upload(server)
	if err != nil {
		fmt.Printf("error getting upload: %v", err)
	}

	fmt.Printf("Ping: %3.2f ms | Download: %3.2f Mbps | Upload: %3.2f Mbps\n", server.Latency, dmbps, umbps)
}
