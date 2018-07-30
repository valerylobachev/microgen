// Code generated by microgen 1.0.0-beta. DO NOT EDIT.

package transport

import (
	"context"
	"errors"
	generated "github.com/devimteam/microgen/examples/generated"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// TraceClientEndpoints is used for tracing endpoints on client side.
func TraceClientEndpoints(endpoints EndpointsSet, tracer opentracinggo.Tracer) EndpointsSet {
	return EndpointsSet{
		CountEndpoint:     opentracing.TraceClient(tracer, "Count")(endpoints.CountEndpoint),
		TestCaseEndpoint:  opentracing.TraceClient(tracer, "TestCase")(endpoints.TestCaseEndpoint),
		UppercaseEndpoint: opentracing.TraceClient(tracer, "Uppercase")(endpoints.UppercaseEndpoint),
	}
}

func (set EndpointsSet) Uppercase(arg0 context.Context, arg1 map[string]string) (res0 string, res1 error) {
	request := UppercaseRequest{StringsMap: arg1}
	response, res1 := set.UppercaseEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*UppercaseResponse).Ans, res1
}

func (set EndpointsSet) Count(arg0 context.Context, arg1 string, arg2 string) (res0 int, res1 []int, res2 error) {
	request := CountRequest{
		Symbol: arg2,
		Text:   arg1,
	}
	response, res2 := set.CountEndpoint(arg0, &request)
	if res2 != nil {
		if e, ok := status.FromError(res2); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res2 = errors.New(e.Message())
		}
		return
	}
	return response.(*CountResponse).Count, response.(*CountResponse).Positions, res2
}

func (set EndpointsSet) TestCase(arg0 context.Context, arg1 []*generated.Comment) (res0 map[string]int, res1 error) {
	request := TestCaseRequest{Comments: arg1}
	response, res1 := set.TestCaseEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*TestCaseResponse).Tree, res1
}

func (set EndpointsSet) IgnoredMethod() {
	return
}

func (set EndpointsSet) IgnoredErrorMethod() (res0 error) {
	return
}
