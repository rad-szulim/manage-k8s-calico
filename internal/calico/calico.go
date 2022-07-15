package calico

import (
	"context"
	"log"

	calicoVersion "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"github.com/projectcalico/api/pkg/client/clientset_generated/clientset"

	// calicoVersion "github.com/projectcalico/libcalico-go/lib/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	rest "k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// type Clientset struct {
// 	Calico clientset.New
// }

func GetIppool(c *rest.Config) error {
	cs, err := clientset.NewForConfig(c)
	if err != nil {
		return err
	}
	// var result runtime.Object
	// if err := cs.ProjectcalicoV3().RESTClient().Get().Resource("ippools").Do(context.TODO()).Into(result); err != nil {
	// 	return err
	// }
	ipp, err := cs.ProjectcalicoV3().IPPools().List(context.TODO(), v1.ListOptions{})
	//Get(context.TODO(), "default-ipv4-ippool", v1.GetOptions{})
	if err != nil {
		log.Println("Error in cs ProjectcalicoV3 .IPPools().List call")
		return err
	}
	// log.Printf("Calico IPPools : %+v", result)
	log.Printf("Calico IPPools Items: %+v", ipp.Items)
	log.Printf("Calico IPPools APIVersion: %+v", ipp.APIVersion)
	log.Printf("Calico IPPools Kind: %+v", ipp.Kind)
	return nil
}

func GetClient() (client.Client, error) {
	scheme := runtime.NewScheme()
	calicoVersion.AddToScheme(scheme)
	// kubeconfig := ctrl.GetConfigOrDie()
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

// TODO: change client to be a method !!!
func GetIppool2(c client.Client) error {
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

// func ListBGP(c *rest.Config) error {
// 	cs, err := clientset.NewForConfig(c)
// 	if err != nil {
// 		return err
// 	}
// 	bgp, err := cs.ProjectcalicoV3().BGPConfigurations().List(context.TODO(), v1.ListOptions{})
// 	//Get(context.TODO(), "default-ipv4-ippool", v1.GetOptions{})
// 	if err != nil {
// 		log.Println("Error in cs ProjectcalicoV3().BGPConfigurations().List call")
// 		return err
// 	}
// 	log.Printf("Calico IPPools ListMeata: %+v", bgp.ListMeta)
// 	log.Printf("Calico IPPools Items: %+v", bgp.Items)
// 	return nil
// }
