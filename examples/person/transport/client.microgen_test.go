package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/valerylobachev/microgen/examples/person/api"
	"testing"
)

func TestEndpointsSet_CreatePerson(t *testing.T) {
	type fields struct {
		CreatePersonEndpoint   endpoint.Endpoint
		UpdatePersonEndpoint   endpoint.Endpoint
		DeletePersonEndpoint   endpoint.Endpoint
		GetPersonByIdEndpoint  endpoint.Endpoint
		GetPersonsByIdEndpoint endpoint.Endpoint
		FindPersonsEndpoint    endpoint.Endpoint
	}
	type args struct {
		arg0 context.Context
		arg1 api.CreatePersonPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := EndpointsSet{
				CreatePersonEndpoint:   tt.fields.CreatePersonEndpoint,
				UpdatePersonEndpoint:   tt.fields.UpdatePersonEndpoint,
				DeletePersonEndpoint:   tt.fields.DeletePersonEndpoint,
				GetPersonByIdEndpoint:  tt.fields.GetPersonByIdEndpoint,
				GetPersonsByIdEndpoint: tt.fields.GetPersonsByIdEndpoint,
				FindPersonsEndpoint:    tt.fields.FindPersonsEndpoint,
			}
			if err := set.CreatePerson(tt.args.arg0, tt.args.arg1); (err != nil) != tt.wantErr {
				t.Errorf("CreatePerson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
