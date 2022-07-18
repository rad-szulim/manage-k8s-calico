package main

import (
	"context"
	"intel/rad-szulim/manage-k8s-calico/internal/calico"
	"log"
	"time"
)

// this name must be used when setting up cluster-wide bgp configuration
// as opposed to node-specific one.
const defaultBGPConfig = "default"

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
		if len(bgp) != 0 {
			log.Printf("Calico default BGP Config already exists\n")
		} else {
			if err := cli.CreateBGP(context.TODO(), defaultBGPConfig, "65000", []string{"10.240.0.0/16"}); err != nil {
				panic(err.Error())
			}
		}
		bgp2, err := cli.ListBGP(context.TODO())
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Calico BGP Config %v \n", bgp2)

		peer, err := cli.ListPeer(context.TODO())
		if err != nil {
			panic(err.Error())
		}
		if len(peer) != 0 {
			log.Printf("Calico peer already exists\n")
		} else {
			if err := cli.CreatePeer(context.TODO(), "my-peer-1", "65000", "10.10.10.1"); err != nil {
				panic(err.Error())
			}
		}
		peer2, err := cli.ListPeer(context.TODO())
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Calico BGP peers %v \n", peer2)

		time.Sleep(30 * time.Second)
	}
}
