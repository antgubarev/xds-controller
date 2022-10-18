package xds

import (
	"context"
	"fmt"
	"time"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/sirupsen/logrus"

	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	router "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	envoyCache "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"

	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type CbConfig struct {
	MaxRetries     uint32
	ConnectTimeout time.Duration
}

type XdsCache struct {
	snapCache envoyCache.SnapshotCache
}

func NewXdsCache(logger *logrus.Logger, snapCache envoyCache.SnapshotCache) *XdsCache {
	return &XdsCache{
		snapCache: snapCache,
	}
}

func (ec *XdsCache) Update(ctx context.Context, cbConfig *CbConfig) error {
	snap := generateSnapshot(NodeName, cbConfig)
	if err := snap.Consistent(); err != nil {
		return fmt.Errorf("inconsistent snapshot: %v", err)
	}
	if err := ec.snapCache.SetSnapshot(ctx, NodeName, snap); err != nil {
		return fmt.Errorf("snapshot error: %v", err)
	}

	return nil
}

func generateSnapshot(nodeName string, cbConfig *CbConfig) *cache.Snapshot {
	snap, _ := cache.NewSnapshot(NodeName,
		map[resource.Type][]types.Resource{
			resource.ClusterType:  {makeCluster(cbConfig)},
			resource.RouteType:    {makeRoute()},
			resource.ListenerType: {makeListener()},
		},
	)

	return snap
}

func makeCluster(cbConfig *CbConfig) *cluster.Cluster {
	return &cluster.Cluster{
		Name:                 ClusterName,
		ConnectTimeout:       durationpb.New(cbConfig.ConnectTimeout),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_STATIC},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		CircuitBreakers: &cluster.CircuitBreakers{
			Thresholds: []*cluster.CircuitBreakers_Thresholds{{
				MaxRetries: wrapperspb.UInt32(cbConfig.MaxRetries),
			}},
		},
		LoadAssignment: &endpoint.ClusterLoadAssignment{
			ClusterName: ClusterName,
			Endpoints: []*endpoint.LocalityLbEndpoints{{
				LbEndpoints: []*endpoint.LbEndpoint{{
					HostIdentifier: &endpoint.LbEndpoint_Endpoint{
						Endpoint: &endpoint.Endpoint{
							Address: &corev3.Address{
								Address: &corev3.Address_SocketAddress{
									SocketAddress: &corev3.SocketAddress{
										Protocol: corev3.SocketAddress_TCP,
										Address:  "127.0.0.1",
										PortSpecifier: &corev3.SocketAddress_PortValue{
											PortValue: 8080,
										},
									},
								},
							},
						},
					},
				}},
			}},
		},
	}
}

func makeRoute() *route.RouteConfiguration {
	return &route.RouteConfiguration{
		Name: "local_route",
		VirtualHosts: []*route.VirtualHost{{
			Name:    "local_service",
			Domains: []string{"*"},
			Routes: []*route.Route{{
				Match: &route.RouteMatch{
					PathSpecifier: &route.RouteMatch_Prefix{
						Prefix: "/",
					},
				},
				Action: &route.Route_Route{
					Route: &route.RouteAction{
						ClusterSpecifier: &route.RouteAction_Cluster{
							Cluster: "local_service",
						},
					},
				},
			}},
		}},
	}
}

func makeListener() *listener.Listener {
	routerConfig, _ := anypb.New(&router.Router{})

	source := &corev3.ConfigSource{}
	source.ResourceApiVersion = resource.DefaultAPIVersion
	source.ConfigSourceSpecifier = &corev3.ConfigSource_ApiConfigSource{
		ApiConfigSource: &corev3.ApiConfigSource{
			TransportApiVersion:       resource.DefaultAPIVersion,
			ApiType:                   corev3.ApiConfigSource_GRPC,
			SetNodeOnFirstMessageOnly: true,
			GrpcServices: []*corev3.GrpcService{{
				TargetSpecifier: &corev3.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &corev3.GrpcService_EnvoyGrpc{ClusterName: XdsClusterName},
				},
			}},
		},
	}

	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "ingress_http",
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				ConfigSource:    source,
				RouteConfigName: "local_route",
			},
		},
		HttpFilters: []*hcm.HttpFilter{{
			Name:       wellknown.Router,
			ConfigType: &hcm.HttpFilter_TypedConfig{TypedConfig: routerConfig},
		}},
	}
	pbst, err := anypb.New(manager)
	if err != nil {
		panic(err)
	}

	return &listener.Listener{
		Address: &corev3.Address{
			Address: &corev3.Address_SocketAddress{
				SocketAddress: &corev3.SocketAddress{
					Protocol: corev3.SocketAddress_TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &corev3.SocketAddress_PortValue{
						PortValue: 80,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{{
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	}
}
