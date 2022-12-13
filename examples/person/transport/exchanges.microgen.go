// Code generated by microgen 0.10.0. DO NOT EDIT.

package transport

import api "github.com/valerylobachev/microgen/examples/person/api"

type (
	CreatePersonRequest struct {
		Payload api.CreatePersonPayload `json:"payload"`
	}
	// Formal exchange type, please do not delete.
	CreatePersonResponse struct{}

	UpdatePersonRequest struct {
		Payload api.UpdatePersonPayload `json:"payload"`
	}
	// Formal exchange type, please do not delete.
	UpdatePersonResponse struct{}

	DeletePersonRequest struct {
		Payload api.DeletePersonPayload `json:"payload"`
	}
	// Formal exchange type, please do not delete.
	DeletePersonResponse struct{}

	GetPersonByIdRequest struct {
		Id     string `json:"id"`
		Source string `json:"source"`
	}
	GetPersonByIdResponse struct {
		Res *api.Person `json:"res"`
	}

	GetPersonsByIdRequest struct {
		Ids    []string `json:"ids"`
		Source string   `json:"source"`
	}
	GetPersonsByIdResponse struct {
		Res []api.Person `json:"res"`
	}

	FindPersonsRequest struct {
		Query api.FindPersonQuery `json:"query"`
	}
	FindPersonsResponse struct {
		Res []api.Person `json:"res"`
	}
)