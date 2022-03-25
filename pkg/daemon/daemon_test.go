package daemon

import (
	"context"
	"testing"
	"time"

	//"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1/fake"

	v1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	snclientset "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/fake"
	mcclientset "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/fake"

	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/utils"
	k8sfake "k8s.io/client-go/kubernetes/fake"
)

// k8sfake "k8s.io/client-go/fake"

func TestDaemon(t *testing.T) {

	t.Setenv("NAMESPACE", "test-namespace")

	stopCh := make(chan struct{})
	defer close(stopCh)

	exitCh := make(chan error)
	defer close(exitCh)

	syncCh := make(chan struct{})
	defer close(syncCh)

	refreshCh := make(chan Message)
	defer close(refreshCh)

	kubeCs := k8sfake.NewSimpleClientset()
	sriovCs := snclientset.NewSimpleClientset()
	machineConfigCs := mcclientset.NewSimpleClientset()

	//versionedFakeClientset := fake.NewSimpleClientset()
	sut := New("test-node",
		sriovCs,
		kubeCs,
		machineConfigCs,
		exitCh,
		stopCh,
		syncCh,
		refreshCh,
		utils.Baremetal,
	)

	sut.Run(stopCh, exitCh)

	sriovCs.SriovnetworkV1().SriovNetworkNodeStates("test-namespace").
		Create(context.Background(),
			&v1.SriovNetworkNodeState{
				Spec: v1.SriovNetworkNodeStateSpec{},
				Status: v1.SriovNetworkNodeStateStatus{
					Interfaces: []v1.InterfaceExt{
						{
							VFs: []v1.VirtualFunction{
								{},
							},
							DeviceID:   "158b",
							Driver:     "i40e",
							Mtu:        1500,
							Name:       "ens803f0",
							PciAddress: "0000:86:00.0",
							Vendor:     "8086",
							NumVfs:     4,
							TotalVfs:   64,
						},
						{
							VFs: []v1.VirtualFunction{
								{},
							},
							DeviceID:   "158b",
							Driver:     "i40e",
							Mtu:        1500,
							Name:       "ens803f1",
							PciAddress: "0000:86:00.1",
							Vendor:     "8086",
							NumVfs:     4,
							TotalVfs:   64,
						},
						{
							VFs: []v1.VirtualFunction{
								{},
							},
							DeviceID:   "1015",
							Driver:     "i40e",
							Mtu:        1500,
							Name:       "ens803f2",
							PciAddress: "0000:86:00.2",
							Vendor:     "8086",
							NumVfs:     4,
							TotalVfs:   64,
						},
					},
				},
			}, metav1.CreateOptions{})

	time.Sleep(30 * time.Second)
}
