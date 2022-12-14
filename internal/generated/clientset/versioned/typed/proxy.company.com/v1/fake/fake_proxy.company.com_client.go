// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/antgubarev/xds-controller/internal/generated/clientset/versioned/typed/proxy.company.com/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeProxyV1 struct {
	*testing.Fake
}

func (c *FakeProxyV1) Sidecars(namespace string) v1.SidecarInterface {
	return &FakeSidecars{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProxyV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
