package api

import "context"

// @microgen middleware, logging,  http, recovering, error-logging, tracing, caching, metrics
// @http-path api/persons/v1
type PersonService interface {
	// @http-method DELETE
	// @http-path CreatePerson/{p1}/{p2}
	// @http-query-vars q1 q2
	// @http-body payload
	CreatePerson(
		ctx context.Context,
		payload CreatePersonPayload, // body
		p1, p2, q1, q2 string,
	) (err error)
	UpdatePerson(ctx context.Context, payload UpdatePersonPayload) (err error)
	DeletePerson(ctx context.Context, payload DeletePersonPayload) (err error)
	// @http-method GET
	// @http-path getPersonById/{id}
	// @http-query-vars source
	GetPersonById(ctx context.Context, id string, source string) (res *Person, err error)
	// @http-query-vars source
	// @http-body ids
	GetPersonsById(ctx context.Context, ids []string, source string) (res []Person, err error)
	// @http-body query
	FindPersons(ctx context.Context, query FindPersonQuery) (res []Person, err error)
}
