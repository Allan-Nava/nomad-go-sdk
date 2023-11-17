package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/nomad/api"
)

func main() {
	// ...
	// Create a new Nomad client
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// Query information about the Nomad cluster
	clusterInfo, err := client.Status().Leader()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nomad Cluster Leader: %s\n", clusterInfo)
}
