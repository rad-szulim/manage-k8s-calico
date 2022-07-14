package calico

import (
	"context"
	"log"

	"github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// type Clientset struct {
// 	Calico clientset.New
// }

func GetIppool(c *rest.Config) error {
	cs, err := clientset.NewForConfig(c)
	if err != nil {
		return err
	}
	ipp, err := cs.ProjectcalicoV3().IPPools().List(context.TODO(), v1.ListOptions{})
	//Get(context.TODO(), "default-ipv4-ippool", v1.GetOptions{})
	if err != nil {
		log.Println("Error in cs ProjectcalicoV3 .IPPools().List call")
		return err
	}
	log.Printf("Calico IPPools ListMeata: %+v", ipp.ListMeta)
	log.Printf("Calico IPPools Items: %+v", ipp.Items)
	return nil
}

func ListBGP(c *rest.Config) error {
	cs, err := clientset.NewForConfig(c)
	if err != nil {
		return err
	}
	bgp, err := cs.ProjectcalicoV3().BGPConfigurations().List(context.TODO(), v1.ListOptions{})
	//Get(context.TODO(), "default-ipv4-ippool", v1.GetOptions{})
	if err != nil {
		log.Println("Error in cs ProjectcalicoV3().BGPConfigurations().List call")
		return err
	}
	log.Printf("Calico IPPools ListMeata: %+v", bgp.ListMeta)
	log.Printf("Calico IPPools Items: %+v", bgp.Items)
	return nil
}
