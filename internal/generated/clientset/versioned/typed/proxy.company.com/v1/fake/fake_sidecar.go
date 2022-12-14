// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	proxycompanycomv1 "github.com/antgubarev/xds-controller/internal/apis/proxy.company.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSidecars implements SidecarInterface
type FakeSidecars struct {
	Fake *FakeProxyV1
	ns   string
}

var sidecarsResource = schema.GroupVersionResource{Group: "proxy.company.com", Version: "v1", Resource: "sidecars"}

var sidecarsKind = schema.GroupVersionKind{Group: "proxy.company.com", Version: "v1", Kind: "Sidecar"}

// Get takes name of the sidecar, and returns the corresponding sidecar object, and an error if there is any.
func (c *FakeSidecars) Get(ctx context.Context, name string, options v1.GetOptions) (result *proxycompanycomv1.Sidecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sidecarsResource, c.ns, name), &proxycompanycomv1.Sidecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*proxycompanycomv1.Sidecar), err
}

// List takes label and field selectors, and returns the list of Sidecars that match those selectors.
func (c *FakeSidecars) List(ctx context.Context, opts v1.ListOptions) (result *proxycompanycomv1.SidecarList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sidecarsResource, sidecarsKind, c.ns, opts), &proxycompanycomv1.SidecarList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &proxycompanycomv1.SidecarList{ListMeta: obj.(*proxycompanycomv1.SidecarList).ListMeta}
	for _, item := range obj.(*proxycompanycomv1.SidecarList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sidecars.
func (c *FakeSidecars) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sidecarsResource, c.ns, opts))

}

// Create takes the representation of a sidecar and creates it.  Returns the server's representation of the sidecar, and an error, if there is any.
func (c *FakeSidecars) Create(ctx context.Context, sidecar *proxycompanycomv1.Sidecar, opts v1.CreateOptions) (result *proxycompanycomv1.Sidecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sidecarsResource, c.ns, sidecar), &proxycompanycomv1.Sidecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*proxycompanycomv1.Sidecar), err
}

// Update takes the representation of a sidecar and updates it. Returns the server's representation of the sidecar, and an error, if there is any.
func (c *FakeSidecars) Update(ctx context.Context, sidecar *proxycompanycomv1.Sidecar, opts v1.UpdateOptions) (result *proxycompanycomv1.Sidecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sidecarsResource, c.ns, sidecar), &proxycompanycomv1.Sidecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*proxycompanycomv1.Sidecar), err
}

// Delete takes name of the sidecar and deletes it. Returns an error if one occurs.
func (c *FakeSidecars) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sidecarsResource, c.ns, name, opts), &proxycompanycomv1.Sidecar{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSidecars) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sidecarsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &proxycompanycomv1.SidecarList{})
	return err
}

// Patch applies the patch and returns the patched sidecar.
func (c *FakeSidecars) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *proxycompanycomv1.Sidecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sidecarsResource, c.ns, name, pt, data, subresources...), &proxycompanycomv1.Sidecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*proxycompanycomv1.Sidecar), err
}
