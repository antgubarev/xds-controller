package crd

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	restclient "k8s.io/client-go/rest"

	v1 "github.com/antgubarev/xds-controller/internal/apis/proxy.company.com/v1"
	versionedclientset "github.com/antgubarev/xds-controller/internal/generated/clientset/versioned"
	proxyinformers "github.com/antgubarev/xds-controller/internal/generated/informers/externalversions"
	"github.com/antgubarev/xds-controller/internal/xds"
	"k8s.io/client-go/tools/cache"
)

type CrdServer struct {
	logger     *logrus.Logger
	envoyCache *xds.XdsCache
	client     *versionedclientset.Clientset
}

func NewCrdServer(logger *logrus.Logger, envoyCache *xds.XdsCache, cfg *restclient.Config) *CrdServer {
	proxyVersionedClient, err := versionedclientset.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf("Error building example clientset: %s", err.Error())
	}

	return &CrdServer{
		logger:     logger,
		envoyCache: envoyCache,
		client:     proxyVersionedClient,
	}
}

func (crdServer *CrdServer) Run(ctx context.Context) {
	proxyInformerFactory := proxyinformers.NewSharedInformerFactory(crdServer.client, time.Second*30)
	sidecarsInformer := proxyInformerFactory.Proxy().V1().Sidecars().Informer()
	sidecarsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			sidecar, ok := obj.(*v1.Sidecar)
			if !ok {
				crdServer.logger.Errorf("Can't parse added object: %v", obj)
			}
			if err := crdServer.envoyCache.Update(ctx, &xds.CbConfig{
				MaxRetries:     uint32(sidecar.Spec.Cb.Tries),
				ConnectTimeout: time.Duration(sidecar.Spec.Cb.Timeout * int(time.Millisecond)),
			}); err != nil {
				crdServer.logger.Errorf("Sidecar added, updating envoy cache: %v", err)
				return
			}
			crdServer.logger.Infof("Sidecar %s is added, envoy cache updated", sidecar.Name)
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			switch newObj.(type) {
			case *v1.Sidecar:
				newSidecar, ok := newObj.(*v1.Sidecar)
				if !ok {
					crdServer.logger.Errorf("Can't parse updated object: %v", newSidecar)
				}
				if oldObj != newObj {
					if err := crdServer.envoyCache.Update(ctx, &xds.CbConfig{
						MaxRetries:     uint32(newSidecar.Spec.Cb.Tries),
						ConnectTimeout: time.Duration(newSidecar.Spec.Cb.Timeout * int(time.Millisecond)),
					}); err != nil {
						crdServer.logger.Errorf("Sidecar %s updated, updating envoy cache: %v", newSidecar.Name, err)
						return
					}
					crdServer.logger.Infof("Sidecar %s updated, envoy cache updated \n", newSidecar.Name)
				}
			default:
				crdServer.logger.Errorf("Update event: unknown object %v", newObj)
			}
		},
		DeleteFunc: func(obj interface{}) {
			sidecar, ok := obj.(*v1.Sidecar)
			if !ok {
				crdServer.logger.Errorf("Can't parse deleted object: %v", obj)
			}
			if err := crdServer.envoyCache.Update(ctx, GetDefaultSidecarConfig()); err != nil {
				crdServer.logger.Errorf("sidecar %s deleted, updating envoy cache by default values: %v", sidecar.Name, err)
				return
			}
		},
	})

	stop := make(chan struct{})
	defer close(stop)

	proxyInformerFactory.Start(stop)
	if !cache.WaitForCacheSync(stop, sidecarsInformer.HasSynced) {
		crdServer.logger.Info("Failed to sync cache")
	}

	<-stop
}

func GetDefaultSidecarConfig() *xds.CbConfig {
	return &xds.CbConfig{
		MaxRetries:     3,
		ConnectTimeout: 3000,
	}
}
