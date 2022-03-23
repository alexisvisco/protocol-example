package main

import (
	"context"
	"fmt"
	"github.com/twitchtv/twirp"
	"net/http"
	commonpb "protocol-example/gen/go/protos/common/v1"
	dummypb "protocol-example/gen/go/protos/dummy/v1"
	leadspb "protocol-example/gen/go/protos/leads/v1"
)

type Server struct{}

func (s *Server) Create(ctx context.Context, request *leadspb.CreateRequest) (*leadspb.CreateResponse, error) {
	return &leadspb.CreateResponse{}, nil
}

func (s *Server) Update(ctx context.Context, request *leadspb.UpdateRequest) (*leadspb.UpdateResponse, error) {
	return &leadspb.UpdateResponse{}, nil

}

func (s *Server) Fetch(ctx context.Context, request *leadspb.FetchRequest) (*leadspb.FetchResponse, error) {
	return &leadspb.FetchResponse{}, nil
}

func (s *Server) Previous(ctx context.Context, request *leadspb.PreviousRequest) (*leadspb.PreviousResponse, error) {
	return &leadspb.PreviousResponse{}, nil
}

func (s *Server) Health(ctx context.Context, request *commonpb.HealthRequest) (*commonpb.HealthResponse, error) {
	return &commonpb.HealthResponse{Status: commonpb.HealthStatus_HEALTH_STATUS_SERVING}, nil
}

func (s *Server) Log(ctx context.Context, request *dummypb.LogRequest) (*dummypb.LogResponse, error) {
	return &dummypb.LogResponse{
		Response: request.Name,
	}, nil
}

// Run the implementation in a local server
func main() {
	//
	client := dummypb.NewDummyServiceProtobufClient("http://localhost:8080", http.DefaultClient)
	log, err := client.Log(context.Background(), &dummypb.LogRequest{Name: "hello"})
	fmt.Println(log, err.(twirp.Error).MetaMap())
}
