package calico

import (
	"context"
	"log"

	calicoVersion "github.com/projectcalico/api/pkg/apis/projectcalico/v3"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetClient() (client.Client, error) {
	scheme := runtime.NewScheme()
	calicoVersion.AddToScheme(scheme)
	kubeconfig, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}
	controllerClient, err := client.New(kubeconfig, client.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}
	return controllerClient, nil
}

// TODO: change client to be a method for testing purposes!!!
func ListIppool(c client.Client) error {
	pools := &calicoVersion.IPPoolList{}
	log.Printf("Calico IPPools APIVersion: %+v", pools.APIVersion)
	log.Printf("Calico IPPools Group: %+v", pools.TypeMeta.GroupVersionKind().Group)
	log.Printf("Calico IPPools TypeMeta: %+v", pools.TypeMeta)
	if err := c.List(context.TODO(), pools,
		&client.ListOptions{Raw: &v1.ListOptions{
			ResourceVersion: "0", // 0 for get means any version
		}}); err != nil {
		return err
	}
	log.Printf("Calico IPPools: %+v", pools)
	log.Printf("Calico Name: %+v", pools.Items[0].Name)
	log.Printf("Calico CIDR: %+v", pools.Items[0].Spec.CIDR)
	return nil
}

func ListBGP(c client.Client) error {
	bgp := &calicoVersion.BGPConfigurationList{}
	if err := c.List(context.TODO(), bgp,
		&client.ListOptions{Raw: &v1.ListOptions{
			ResourceVersion: "0", // 0 for get means any version
		}}); err != nil {
		return err
	}
	log.Printf("Calico BGPConfigurationList: %+v", bgp)
	return nil
}
