// This file was automatically generated by "microgen 0.6.0" utility.
// Please, do not edit.
package transportgrpc

import (
	generated "github.com/devimteam/microgen/example/generated"
	protobuf "github.com/devimteam/microgen/example/generated/transport/converter/protobuf"
	stringsvc "github.com/devimteam/protobuf/stringsvc"
	grpc "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"
)

type stringServiceServer struct {
	uppercase grpc.Handler
	count     grpc.Handler
	testCase  grpc.Handler
}

func NewGRPCServer(endpoints *generated.Endpoints, opts ...grpc.ServerOption) stringsvc.StringServiceServer {
	return &stringServiceServer{
		count: grpc.NewServer(
			endpoints.CountEndpoint,
			protobuf.DecodeCountRequest,
			protobuf.EncodeCountResponse,
			opts...,
		),
		testCase: grpc.NewServer(
			endpoints.TestCaseEndpoint,
			protobuf.DecodeTestCaseRequest,
			protobuf.EncodeTestCaseResponse,
			opts...,
		),
		uppercase: grpc.NewServer(
			endpoints.UppercaseEndpoint,
			protobuf.DecodeUppercaseRequest,
			protobuf.EncodeUppercaseResponse,
			opts...,
		),
	}
}

func (S *stringServiceServer) Uppercase(ctx context.Context, req *stringsvc.UppercaseRequest) (*stringsvc.UppercaseResponse, error) {
	_, resp, err := S.uppercase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*stringsvc.UppercaseResponse), nil
}

func (S *stringServiceServer) Count(ctx context.Context, req *stringsvc.CountRequest) (*stringsvc.CountResponse, error) {
	_, resp, err := S.count.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*stringsvc.CountResponse), nil
}

func (S *stringServiceServer) TestCase(ctx context.Context, req *stringsvc.TestCaseRequest) (*stringsvc.TestCaseResponse, error) {
	_, resp, err := S.testCase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*stringsvc.TestCaseResponse), nil
}