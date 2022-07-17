package main

import (
	"context"
	"intel/rad-szulim/manage-k8s-calico/internal/calico"
	"log"
	"time"
)

func main() {
	// creates runtime client for k8s
	c, err := calico.GetClient()
	if err != nil {
		panic(err.Error())
	}
	cli := calico.ClientManager{Client: c}
	for {
		pools, err := cli.ListIppool(context.TODO())
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Calico IPpools %v \n", pools)
		bgp, err := cli.ListBGP(context.TODO())
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Calico BGP Config %v \n", bgp)
		time.Sleep(30 * time.Second)
	}
}
