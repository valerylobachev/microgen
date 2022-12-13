package api

import "context"

// @microgen middleware, logging,  http, recovering, error-logging, tracing, caching, metrics
// @http-path api/persons/v1
type PersonService interface {
	// @http-path сreatePerson
	// @http-body payload
	CreatePerson(ctx context.Context, payload CreatePersonPayload) (err error)
	// @http-path updatePerson
	// @http-body payload
	UpdatePerson(ctx context.Context, payload UpdatePersonPayload) (err error)
	// @http-path deletePerson
	// @http-body payload
	DeletePerson(ctx context.Context, payload DeletePersonPayload) (err error)
	// @http-method GET
	// @http-path getPersonById/{id}
	// @http-query-vars source
	GetPersonById(ctx context.Context, id string, source string) (res *Person, err error)
	// @http-path getPersonsById
	// @http-query-vars source
	// @http-body ids
	GetPersonsById(ctx context.Context, ids []string, source string) (res []Person, err error)
	// @http-path findPersons
	// @http-body query
	FindPersons(ctx context.Context, query FindPersonQuery) (res []Person, err error)
}
