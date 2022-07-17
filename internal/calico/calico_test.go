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
	It("Add, Delete, Get Calico k8s BGP Configuration", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Create the fake client.
		By("Setting up fake Calico client for k8s")
		cli := calico.ClientManager{Client: getFakeClient()}

		By("Creating BGP Configuration")
		err := cli.CreateBGP(ctx, "my-name", "65000", []string{
			"192.168.100.0/24", "192.168.200.0/24"})
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
			},
		))

		By("Deleting BGP Configuration")
		err = cli.DeleteBGP(ctx, "my-name")
		Expect(err).ToNot(HaveOccurred())

		By("Listing BGP Configuration for the second time")
		b2, err := cli.ListBGP(ctx)
		Expect(err).ToNot(HaveOccurred())
		Expect(b2).To(HaveLen(0))
	})
})

func getFakeClient() client.WithWatch {
	scheme := runtime.NewScheme()
	calicoVersion.AddToScheme(scheme)
	return fake.NewFakeClientWithScheme(scheme)
}
