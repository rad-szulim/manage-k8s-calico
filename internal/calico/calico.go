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
	ipp, err := cs.ProjectcalicoV3().IPPools().Get(context.TODO(), "default-ipv4-ippool", v1.GetOptions{})
	if err != nil {
		return err
	}
	log.Printf("Calico IPPools are: %+v", ipp.Spec.CIDR)
	return nil
}
