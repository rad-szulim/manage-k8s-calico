package calico_test

import (
	"context"
	"intel/rad-szulim/manage-k8s-calico/internal/calico"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"k8s.io/apimachinery/pkg/runtime"

	calicoVersion "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("Calico k8s test", func() {
	It("Add, Get, Delete Calico k8s BGP Configuration", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		By("Setting up fake Calico client for k8s")
		cli := calico.ClientManager{Client: getFakeClient()}

		By("Creating BGP Configurations")
		err := cli.CreateBGP(ctx, "my-name", "65000", []string{
			"192.168.100.0/24", "192.168.200.0/24"})
		Expect(err).ToNot(HaveOccurred())

		err = cli.CreateBGP(ctx, "my-name-2", "65001", []string{
			"192.168.101.0/24", "192.168.201.0/24"})
		Expect(err).ToNot(HaveOccurred())

		By("Listing BGP Configuration")
		b, err := cli.ListBGP(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(b).ToNot(BeNil())
		cidrFn := func(element interface{}) string {
			return element.(calicoVersion.ServiceClusterIPBlock).CIDR
		}
		Expect(b).To(MatchAllElements(
			func(element interface{}) string {
				return element.(calicoVersion.BGPConfiguration).Name
			},
			Elements{
				"my-name": MatchFields(IgnoreExtras, Fields{
					"ObjectMeta": MatchFields(IgnoreExtras, Fields{
						"Name": Equal("my-name"),
					}),
					"Spec": MatchFields(IgnoreExtras, Fields{
						"ASNumber": PointTo(BeEquivalentTo(65000)),
						"ServiceClusterIPs": MatchAllElements(
							cidrFn, Elements{
								"192.168.100.0/24": MatchFields(IgnoreExtras, Fields{
									"CIDR": Equal("192.168.100.0/24"),
								}),
								"192.168.200.0/24": MatchFields(IgnoreExtras, Fields{
									"CIDR": Equal("192.168.200.0/24"),
								}),
							},
						),
					}),
				}),
				"my-name-2": MatchFields(IgnoreExtras, Fields{
					"ObjectMeta": MatchFields(IgnoreExtras, Fields{
						"Name": Equal("my-name-2"),
					}),
					"Spec": MatchFields(IgnoreExtras, Fields{
						"ASNumber": PointTo(BeEquivalentTo(65001)),
						"ServiceClusterIPs": MatchAllElements(
							cidrFn, Elements{
								"192.168.101.0/24": MatchFields(IgnoreExtras, Fields{
									"CIDR": Equal("192.168.101.0/24"),
								}),
								"192.168.201.0/24": MatchFields(IgnoreExtras, Fields{
									"CIDR": Equal("192.168.201.0/24"),
								}),
							},
						),
					}),
				}),
			},
		))

		By("Deleting 1st BGP Configuration")
		err = cli.DeleteBGP(ctx, "my-name")
		Expect(err).ToNot(HaveOccurred())

		By("Deleting 2nd BGP Configuration")
		err = cli.DeleteBGP(ctx, "my-name-2")
		Expect(err).ToNot(HaveOccurred())

		By("Listing BGP Configuration for the second time")
		b2, err := cli.ListBGP(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(b2).To(HaveLen(0))
	})

	It("Add, Get, Delete Calico k8s BGP Peers", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		By("Setting up fake Calico client for k8s")
		cli := calico.ClientManager{Client: getFakeClient()}

		By("Creating BGP Peers")
		err := cli.CreatePeer(ctx, "my-peer", "65000", "10.10.10.1")
		Expect(err).ToNot(HaveOccurred())

		err = cli.CreatePeer(ctx, "my-peer-2", "65001", "10.10.10.2")
		Expect(err).ToNot(HaveOccurred())

		p, err := cli.ListPeer(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(p).To(HaveLen(2))

		Expect(p).To(MatchAllElements(
			func(element interface{}) string {
				return element.(calicoVersion.BGPPeer).Name
			},
			Elements{
				"my-peer": MatchFields(IgnoreExtras, Fields{
					"ObjectMeta": MatchFields(IgnoreExtras, Fields{
						"Name": Equal("my-peer"),
					}),
					"Spec": MatchFields(IgnoreExtras, Fields{
						"ASNumber": BeEquivalentTo(65000),
						"PeerIP":   Equal("10.10.10.1"),
					}),
				}),
				"my-peer-2": MatchFields(IgnoreExtras, Fields{
					"ObjectMeta": MatchFields(IgnoreExtras, Fields{
						"Name": Equal("my-peer-2"),
					}),
					"Spec": MatchFields(IgnoreExtras, Fields{
						"ASNumber": BeEquivalentTo(65001),
						"PeerIP":   Equal("10.10.10.2"),
					}),
				}),
			},
		))

		By("Deleting 1st BGP Peer")
		err = cli.DeletePeer(ctx, "my-peer")
		Expect(err).ToNot(HaveOccurred())

		By("Deleting 2nd BGP Peer")
		err = cli.DeletePeer(ctx, "my-peer-2")
		Expect(err).ToNot(HaveOccurred())

		By("Listing BGP Peers for the second time")
		p2, err := cli.ListPeer(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(p2).To(HaveLen(0))
	})
})

func getFakeClient() client.WithWatch {
	scheme := runtime.NewScheme()
	calicoVersion.AddToScheme(scheme)
	return fake.NewFakeClientWithScheme(scheme)
}
