package main

import (
	"intel/rad-szulim/manage-k8s-calico/internal/calico"
	"time"
)

func main() {
	// creates runtime client for k8s
	c, err := calico.GetClient()
	if err != nil {
		panic(err.Error())
	}
	for {
		if err := calico.GetIppool2(c); err != nil {
			panic(err.Error())
		}
		time.Sleep(30 * time.Second)
	}
}
