// Code generated by microgen 1.0.0-beta. DO NOT EDIT.

package transportgrpc

import (
	transport "github.com/devimteam/microgen/examples/addsvc/transport"
	pb "github.com/devimteam/microgen/examples/protobuf"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpckit "github.com/go-kit/kit/transport/grpc"
	opentracinggo "github.com/opentracing/opentracing-go"
	grpc "google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn, addr string, opts ...grpckit.ClientOption) transport.EndpointsSet {
	return transport.EndpointsSet{
		ConcatEndpoint: grpckit.NewClient(
			conn, addr, "Concat",
			_Encode_Concat_Request,
			_Decode_Concat_Response,
			pb.ConcatResponse{},
			opts...,
		).Endpoint(),
		SumEndpoint: grpckit.NewClient(
			conn, addr, "Sum",
			_Encode_Sum_Request,
			_Decode_Sum_Response,
			pb.SumResponse{},
			opts...,
		).Endpoint(),
	}
}

func TracingGRPCClientOptions(tracer opentracinggo.Tracer, logger log.Logger) func([]grpckit.ClientOption) []grpckit.ClientOption {
	return func(opts []grpckit.ClientOption) []grpckit.ClientOption {
		return append(opts, grpckit.ClientBefore(
			opentracing.ContextToGRPC(tracer, logger),
		))
	}
}
