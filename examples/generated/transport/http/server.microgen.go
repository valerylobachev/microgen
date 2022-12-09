// Code generated by microgen 0.9.0. DO NOT EDIT.

package transporthttp

import (
	transport "github.com/valerylobachev/microgen/examples/generated/transport"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/uppercase").Handler(
		http.NewServer(
			endpoints.UppercaseEndpoint,
			_Decode_Uppercase_Request,
			_Encode_Uppercase_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "Uppercase", logger)))...))
	mux.Methods("GET").Path("/count/{text}/{symbol}").Handler(
		http.NewServer(
			endpoints.CountEndpoint,
			_Decode_Count_Request,
			_Encode_Count_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "Count", logger)))...))
	mux.Methods("POST").Path("/test-case").Handler(
		http.NewServer(
			endpoints.TestCaseEndpoint,
			_Decode_TestCase_Request,
			_Encode_TestCase_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "TestCase", logger)))...))
	mux.Methods("POST").Path("/dummy-method").Handler(
		http.NewServer(
			endpoints.DummyMethodEndpoint,
			_Decode_DummyMethod_Request,
			_Encode_DummyMethod_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "DummyMethod", logger)))...))
	return mux
}
