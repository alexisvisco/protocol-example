package main

import (
	"context"
	"fmt"
	"net/http"
	commonpb "protocol-example/gen/go/protos/common/v1"
	dummypb "protocol-example/gen/go/protos/dummy/v1"
)

type DummyServer struct{}

func (s *DummyServer) Health(ctx context.Context, request *commonpb.HealthRequest) (*commonpb.HealthResponse, error) {
	return &commonpb.HealthResponse{Status: commonpb.HealthStatus_HEALTH_STATUS_SERVING}, nil
}

func (s *DummyServer) Log(ctx context.Context, request *dummypb.LogRequest) (*dummypb.LogResponse, error) {
	return &dummypb.LogResponse{
		Response: request.Name,
	}, nil
}

// Run the implementation in a local server
func main() {
	handler := dummypb.NewDummyServiceServer(&DummyServer{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(handler.PathPrefix(), handler)
	fmt.Println("listen :8080")
	http.ListenAndServe(":8080", mux)
}
