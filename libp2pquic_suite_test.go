package libp2pquic

import (
	mrand "math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLibp2pQuicTransport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "libp2p QUIC Transport Suite")
}

var _ = BeforeSuite(func() {
	mrand.Seed(GinkgoRandomSeed())
})
