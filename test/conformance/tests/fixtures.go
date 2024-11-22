package tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/k8snetworkplumbingwg/sriov-network-operator/test/util/clean"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/test/util/cluster"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/test/util/discovery"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/test/util/namespaces"
)

var sriovInfos *cluster.EnabledNodes

var _ = BeforeSuite(func() {
	err := clean.All()
	Expect(err).NotTo(HaveOccurred())

	isSingleNode, err := cluster.IsSingleNode(clients)
	Expect(err).ToNot(HaveOccurred())
	if isSingleNode {
		disableDrainState, err := cluster.GetNodeDrainState(clients, operatorNamespace)
		Expect(err).ToNot(HaveOccurred())
		if discovery.Enabled() {
			if !disableDrainState {
				Fail("SriovOperatorConfig DisableDrain property must be enabled in a single node environment")
			}
		}
		snoTimeoutMultiplier = 1
		if !disableDrainState {
			err = cluster.SetDisableNodeDrainState(clients, operatorNamespace, true)
			Expect(err).ToNot(HaveOccurred())
			clean.RestoreNodeDrainState = true
		}
	}

	err = namespaces.Create(namespaces.Test, clients)
	Expect(err).ToNot(HaveOccurred())
	err = namespaces.Clean(operatorNamespace, namespaces.Test, clients, discovery.Enabled())
	Expect(err).ToNot(HaveOccurred())
	WaitForSRIOVStable()
	Eventually(func(g Gomega) {
		sriovInfos, err = cluster.DiscoverSriov(clients, operatorNamespace)
		g.Expect(err).ToNot(HaveOccurred())
	}).WithTimeout(5*time.Minute).WithPolling(10*time.Second).Should(Succeed())
})

var _ = AfterSuite(func() {
	err := clean.All()
	Expect(err).NotTo(HaveOccurred())
})
