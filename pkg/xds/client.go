package xds

import (
	"context"
	"fmt"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"google.golang.org/grpc"
)

type AdsClient struct {
	discovery.AggregatedDiscoveryService_StreamAggregatedResourcesClient
	cancel context.CancelFunc
}

func (c *AdsClient) Cleanup() {
	if c.cancel != nil {
		c.cancel()
	}
}

func ConnectAds(ctx context.Context, url string) (*AdsClient, error) {
	ctx, cancel := context.WithCancel(ctx)
	conn, err := grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		cancel()
		return nil, fmt.Errorf("gRPC dial failed: %s", err)
	}

	xds := discovery.NewAggregatedDiscoveryServiceClient(conn)
	client, err := xds.StreamAggregatedResources(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("stream resources failed: %s", err)
	}

	return &AdsClient{client, cancel}, nil
}
