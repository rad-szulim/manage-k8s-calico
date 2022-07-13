package main

import (
	"intel/rad-szulim/manage-k8s-calico/internal/calico"
	"time"

	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	for {
		if err := calico.GetIppool(config); err != nil {
			panic(err.Error())
		}
		time.Sleep(30 * time.Second)
	}
}
