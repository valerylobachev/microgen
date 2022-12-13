// Code generated by microgen 0.10.0. DO NOT EDIT.

package transporthttp

import (
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	httpkit "github.com/go-kit/kit/transport/http"
	opentracinggo "github.com/opentracing/opentracing-go"
	transport "github.com/valerylobachev/microgen/examples/person/transport"
	"net/url"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {
	return transport.EndpointsSet{
		CreatePersonEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreatePerson_Request,
			_Decode_CreatePerson_Response,
			opts...,
		).Endpoint(),
		DeletePersonEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_DeletePerson_Request,
			_Decode_DeletePerson_Response,
			opts...,
		).Endpoint(),
		FindPersonsEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_FindPersons_Request,
			_Decode_FindPersons_Response,
			opts...,
		).Endpoint(),
		GetPersonByIdEndpoint: httpkit.NewClient(
			"GET", u,
			_Encode_GetPersonById_Request,
			_Decode_GetPersonById_Response,
			opts...,
		).Endpoint(),
		GetPersonsByIdEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetPersonsById_Request,
			_Decode_GetPersonsById_Response,
			opts...,
		).Endpoint(),
		UpdatePersonEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_UpdatePerson_Request,
			_Decode_UpdatePerson_Response,
			opts...,
		).Endpoint(),
	}
}

func TracingHTTPClientOptions(tracer opentracinggo.Tracer, logger log.Logger) func([]httpkit.ClientOption) []httpkit.ClientOption {
	return func(opts []httpkit.ClientOption) []httpkit.ClientOption {
		return append(opts, httpkit.ClientBefore(
			opentracing.ContextToHTTP(tracer, logger),
		))
	}
}
