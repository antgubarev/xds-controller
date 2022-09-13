package main

import (
	"path/filepath"
	"time"

	v1 "github.com/antgubarev/xds-controller/internal/apis/proxy.company.com/v1"
	versionedclientset "github.com/antgubarev/xds-controller/internal/generated/clientset/versioned"
	proxyinformers "github.com/antgubarev/xds-controller/internal/generated/informers/externalversions"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()

	cfg, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		logger.Fatalf("BuildConfigFromFlags: %v", err)
	}

	_, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	proxyVersionedClient, err := versionedclientset.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf("Error building example clientset: %s", err.Error())
	}

	proxyInformerFactory := proxyinformers.NewSharedInformerFactory(proxyVersionedClient, time.Second*60)

	sidecarsInformer := proxyInformerFactory.Proxy().V1().Sidecars().Informer()
	sidecarsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			sidecar, ok := obj.(*v1.Sidecar)
			if !ok {
				logger.Errorf("Can't parse added object: %v", obj)
			}
			logger.Infof("Sidecar %s is added", sidecar.Name)
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			switch newObj.(type) {
			case *v1.Sidecar:
				newSidecar, ok := newObj.(*v1.Sidecar)
				if !ok {
					logger.Errorf("Can't parse updated object: %v", newSidecar)
				}
				if oldObj != newObj {
					logger.Infof("Sidecar %s is updated", newSidecar.Name)
				}
			default:
				logger.Errorf("update event: unknown object %v", newObj)
			}
		},
		DeleteFunc: func(obj interface{}) {
			sidecar, ok := obj.(*v1.Sidecar)
			if !ok {
				logger.Errorf("Can't parse deleted object: %v", obj)
			}
			logger.Infof("Sidecar %s is deleted", sidecar.Name)
		},
	})

	stop := make(chan struct{})
	defer close(stop)

	proxyInformerFactory.Start(stop)
	if !cache.WaitForCacheSync(stop, sidecarsInformer.HasSynced) {
		logger.Info("Failed to sync cache")
	}

	<-stop
}
