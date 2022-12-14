// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/antgubarev/xds-controller/internal/apis/proxy.company.com/v1"
	scheme "github.com/antgubarev/xds-controller/internal/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SidecarsGetter has a method to return a SidecarInterface.
// A group's client should implement this interface.
type SidecarsGetter interface {
	Sidecars(namespace string) SidecarInterface
}

// SidecarInterface has methods to work with Sidecar resources.
type SidecarInterface interface {
	Create(ctx context.Context, sidecar *v1.Sidecar, opts metav1.CreateOptions) (*v1.Sidecar, error)
	Update(ctx context.Context, sidecar *v1.Sidecar, opts metav1.UpdateOptions) (*v1.Sidecar, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Sidecar, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SidecarList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sidecar, err error)
	SidecarExpansion
}

// sidecars implements SidecarInterface
type sidecars struct {
	client rest.Interface
	ns     string
}

// newSidecars returns a Sidecars
func newSidecars(c *ProxyV1Client, namespace string) *sidecars {
	return &sidecars{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sidecar, and returns the corresponding sidecar object, and an error if there is any.
func (c *sidecars) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Sidecar, err error) {
	result = &v1.Sidecar{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sidecars").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sidecars that match those selectors.
func (c *sidecars) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SidecarList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SidecarList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sidecars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sidecars.
func (c *sidecars) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sidecars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sidecar and creates it.  Returns the server's representation of the sidecar, and an error, if there is any.
func (c *sidecars) Create(ctx context.Context, sidecar *v1.Sidecar, opts metav1.CreateOptions) (result *v1.Sidecar, err error) {
	result = &v1.Sidecar{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sidecars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sidecar).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sidecar and updates it. Returns the server's representation of the sidecar, and an error, if there is any.
func (c *sidecars) Update(ctx context.Context, sidecar *v1.Sidecar, opts metav1.UpdateOptions) (result *v1.Sidecar, err error) {
	result = &v1.Sidecar{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sidecars").
		Name(sidecar.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sidecar).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sidecar and deletes it. Returns an error if one occurs.
func (c *sidecars) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sidecars").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sidecars) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sidecars").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sidecar.
func (c *sidecars) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sidecar, err error) {
	result = &v1.Sidecar{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sidecars").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
