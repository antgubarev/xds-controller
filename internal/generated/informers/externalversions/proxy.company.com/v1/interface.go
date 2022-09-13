// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/antgubarev/xds-controller/internal/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Sidecars returns a SidecarInformer.
	Sidecars() SidecarInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Sidecars returns a SidecarInformer.
func (v *version) Sidecars() SidecarInformer {
	return &sidecarInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}