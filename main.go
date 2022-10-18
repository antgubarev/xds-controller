package main

import (
	"context"
	"path/filepath"

	"github.com/antgubarev/xds-controller/internal/crd"
	"github.com/antgubarev/xds-controller/internal/xds"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	envoyCache "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	envoyServer "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	envoyTest "github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	ctx := context.Background()

	snapCache := envoyCache.NewSnapshotCache(false, envoyCache.IDHash{}, logger)
	xdsCache := xds.NewXdsCache(logger, snapCache)
	if err := xdsCache.Update(ctx, crd.GetDefaultSidecarConfig()); err != nil {
		logger.Fatal(err)
	}

	logger.Infof("Starting envoy grpc server on port 18000 \n")
	srv := envoyServer.NewServer(ctx, snapCache, &envoyTest.Callbacks{})
	go func() {
		if err := xds.RunServer(ctx, srv, 18000); err != nil {
			logger.Errorf("xds server couldn't start: %v", err)
		}
	}()

	logger.Infof("Starting crd server \n")
	cfg, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		logger.Fatalf("BuildConfigFromFlags: %v", err)
	}

	crdServer := crd.NewCrdServer(logger, xdsCache, cfg)
	crdServer.Run(ctx)
}
