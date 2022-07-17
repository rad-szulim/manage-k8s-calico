package calico

import (
	"context"

	calicoVersion "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"github.com/projectcalico/api/pkg/lib/numorstring"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ClientManager struct {
	Client client.Client
}

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

func (c ClientManager) ListIppool(ctx context.Context) ([]calicoVersion.IPPool, error) {
	pools := &calicoVersion.IPPoolList{}
	if err := c.Client.List(ctx, pools,
		&client.ListOptions{Raw: &v1.ListOptions{
			ResourceVersion: "0", // 0 for get means any version
		}}); err != nil {
		return nil, err
	}
	return pools.Items, nil
}

func (c ClientManager) ListBGP(ctx context.Context) ([]calicoVersion.BGPConfiguration, error) {
	bgp := &calicoVersion.BGPConfigurationList{}
	if err := c.Client.List(ctx, bgp,
		&client.ListOptions{Raw: &v1.ListOptions{
			ResourceVersion: "0", // 0 for get means any version
		}}); err != nil {
		return nil, err
	}
	return bgp.Items, nil
}

func (c ClientManager) CreateBGP(ctx context.Context,
	name, asnumber string, subnets []string) error {
	cfg := calicoVersion.NewBGPConfiguration()

	asnum, err := numorstring.ASNumberFromString(asnumber)
	if err != nil {
		return err
	}
	cfg.Name = name
	cfg.Spec.ASNumber = &asnum
	cidrs := make([]calicoVersion.ServiceClusterIPBlock, len(subnets))
	for i, oneCidr := range subnets {
		cidrs[i].CIDR = oneCidr
	}
	cfg.Spec.ServiceClusterIPs = cidrs

	if err := c.Client.Create(ctx, cfg, &client.CreateOptions{}); err != nil {
		return err
	}
	return nil
}

func (c ClientManager) DeleteBGP(ctx context.Context, name string) error {
	cfg := calicoVersion.NewBGPConfiguration()
	cfg.Name = name
	if err := c.Client.Delete(ctx, cfg,
		&client.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}
