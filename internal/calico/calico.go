package calico

import (
	"context"
	"fmt"

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

// Get Client returns new k8s cluster client connection ready to interact with
// Calico API.
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

// ListIPPool lists all Calico IP Pools.
func (c ClientManager) ListIPPool(ctx context.Context) ([]calicoVersion.IPPool, error) {
	pools := &calicoVersion.IPPoolList{}
	if err := c.Client.List(ctx, pools,
		&client.ListOptions{}); err != nil {
		return nil, err
	}
	return pools.Items, nil
}

// CreateIPPool adds a Calico IP Pool with the specified name.
func (c ClientManager) CreateIPPool(ctx context.Context, name string) error {
	// This method is only used in unit tests.  Default IP Pool is created
	// when adding Calico as the default CNI.
	pool := &calicoVersion.IPPool{}
	pool.ObjectMeta.Name = name
	return c.Client.Create(ctx, pool,
		&client.CreateOptions{})
}

// UpdateIPPool updates specified IP Pool.
func (c ClientManager) UpdateIPPool(ctx context.Context, ippool *calicoVersion.IPPool) error {
	return c.Client.Update(ctx, ippool,
		&client.UpdateOptions{})
}

// DisableBGPExportForIPPool disables BGP Export for the IP Pool identified by name.
func (c ClientManager) DisableBGPExportForIPPool(ctx context.Context, name string) error {
	pools, err := c.ListIPPool(ctx)
	if err != nil {
		return err
	}
	for _, pool := range pools {
		if pool.ObjectMeta.Name == name {
			pool.Spec.DisableBGPExport = true
			if err := c.UpdateIPPool(ctx, &pool); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("could not find requested IP Pool %s when attemping to Disable BGP export", name)
}

// ListBGP lists Calico BGP Configurations.
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

// CreateBGP creates a Calico BGP Configuration.
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

// DeleteBGP deletes a Calico BGP Configuration.
func (c ClientManager) DeleteBGP(ctx context.Context, name string) error {
	cfg := calicoVersion.NewBGPConfiguration()
	cfg.Name = name
	if err := c.Client.Delete(ctx, cfg,
		&client.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

// ListPeer lists all Calico BGP Peers.
func (c ClientManager) ListPeer(ctx context.Context) ([]calicoVersion.BGPPeer, error) {
	peer := &calicoVersion.BGPPeerList{}
	if err := c.Client.List(ctx, peer,
		&client.ListOptions{Raw: &v1.ListOptions{
			ResourceVersion: "0", // 0 for get means any version
		}}); err != nil {
		return nil, err
	}
	return peer.Items, nil
}

// CreatePeer creates one Calico BGP Peer.
func (c ClientManager) CreatePeer(ctx context.Context,
	name, asnumber, ip string) error {
	p := calicoVersion.NewBGPPeer()
	asnum, err := numorstring.ASNumberFromString(asnumber)
	if err != nil {
		return err
	}
	p.Name = name
	p.Spec.ASNumber = asnum
	p.Spec.PeerIP = ip

	if err := c.Client.Create(ctx, p, &client.CreateOptions{}); err != nil {
		return err
	}
	return nil
}

// DeletePeer deletes one Calico BGP Peer.
func (c ClientManager) DeletePeer(ctx context.Context, name string) error {
	p := calicoVersion.NewBGPPeer()
	p.Name = name
	if err := c.Client.Delete(ctx, p,
		&client.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}
