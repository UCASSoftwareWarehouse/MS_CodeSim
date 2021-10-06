package server

import (
	"code_sim/internal"
	"code_sim/pb_gen"
	"context"
)

type codeSimServer struct {
	*pb_gen.UnimplementedCodeSimServer
}

func newCodeSimServer() *codeSimServer {
	return &codeSimServer{}
}

func (c *codeSimServer) HelloWorld(ctx context.Context, request *pb_gen.CodeSimHelloWorldRequest) (*pb_gen.CodeSimHelloWorldResponse, error) {
	return &pb_gen.CodeSimHelloWorldResponse{ThanksText: request.GetHelloText() + ", thx."}, nil
}

func (c *codeSimServer) Search(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	return internal.Search(ctx, req)
}

func (c *codeSimServer) Upload(stream pb_gen.CodeSim_UploadServer) error {
	return internal.Upload(stream)
}
