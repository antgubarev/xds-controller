// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	proxycompanycomv1 "github.com/antgubarev/xds-controller/internal/apis/proxy.company.com/v1"
	versioned "github.com/antgubarev/xds-controller/internal/generated/clientset/versioned"
	internalinterfaces "github.com/antgubarev/xds-controller/internal/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/antgubarev/xds-controller/internal/generated/listers/proxy.company.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SidecarInformer provides access to a shared informer and lister for
// Sidecars.
type SidecarInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SidecarLister
}

type sidecarInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSidecarInformer constructs a new informer for Sidecar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSidecarInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSidecarInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSidecarInformer constructs a new informer for Sidecar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSidecarInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProxyV1().Sidecars(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProxyV1().Sidecars(namespace).Watch(context.TODO(), options)
			},
		},
		&proxycompanycomv1.Sidecar{},
		resyncPeriod,
		indexers,
	)
}

func (f *sidecarInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSidecarInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sidecarInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&proxycompanycomv1.Sidecar{}, f.defaultInformer)
}

func (f *sidecarInformer) Lister() v1.SidecarLister {
	return v1.NewSidecarLister(f.Informer().GetIndexer())
}
