package runner

import (
	"context"
	"fmt"
	"net"
	"testing"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	cache "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"google.golang.org/grpc"

	"github.com/envoyproxy/xds-conformance/pkg/conformance"
)

type hasher struct{}

func (hasher) ID(*core.Node) string {
	return ""
}

// TestGoRunner uses the go control plane to validate we can properly use go tests to run the conformance tests
// This is not meant to test go control plane for conformance.
func TestGoRunner(t *testing.T) {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 0))
	if err != nil {
		t.Fatal(err)
	}

	snapshots := cache.NewSnapshotCache(false, hasher{}, nil)

	server := serverv3.NewServer(context.Background(), snapshots, nil)
	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	snapshot := cache.Snapshot{}
	snapshot.Resources[types.Cluster] = cache.Resources{Version: "test"}
	if err := snapshots.SetSnapshot("", snapshot); err != nil {
		t.Fatal(err)
	}

	go func() {
		_ = grpcServer.Serve(lis)
	}()
	defer grpcServer.GracefulStop()

	// TODO select only the minimal subset
	RunGoTest(t, nil, conformance.TestInput{Address: lis.Addr().String()})
}
