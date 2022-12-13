package service

import (
	"context"
	"errors"
	"github.com/valerylobachev/microgen/examples/person/api"
	"strings"
	"time"
)

type inMemoryPersonService struct {
	Persons map[string]api.Person
}

func NewPersonService() api.PersonService {
	//persons := make(map[api.PersonId]api.Person)
	persons := Samples()
	return &inMemoryPersonService{
		Persons: persons,
	}
}

func (p *inMemoryPersonService) CreatePerson(ctx context.Context, payload api.CreatePersonPayload) (err error) {
	_, ok := p.Persons[payload.Id]
	if ok {
		return errors.New("person exists")
	}
	person := api.Person{
		Id:         payload.Id,
		Lastname:   payload.Lastname,
		Firstname:  payload.Firstname,
		Middlename: payload.Middlename,
		CategoryId: payload.CategoryId,
		Phone:      payload.Phone,
		Email:      payload.Email,
		Attributes: payload.Attributes,
		UpdatedAt:  time.Now(),
		UpdatedBy:  payload.UpdatedBy,
	}
	p.Persons[payload.Id] = person
	return nil
}

func (p *inMemoryPersonService) UpdatePerson(ctx context.Context, payload api.UpdatePersonPayload) (err error) {
	_, ok := p.Persons[payload.Id]
	if !ok {
		return errors.New("person not found")
	}
	person := api.Person{
		Id:         payload.Id,
		Lastname:   payload.Lastname,
		Firstname:  payload.Firstname,
		Middlename: payload.Middlename,
		CategoryId: payload.CategoryId,
		Phone:      payload.Phone,
		Email:      payload.Email,
		Attributes: payload.Attributes,
		UpdatedAt:  time.Now(),
		UpdatedBy:  payload.UpdatedBy,
	}
	p.Persons[payload.Id] = person
	return nil
}

func (p *inMemoryPersonService) DeletePerson(ctx context.Context, payload api.DeletePersonPayload) (err error) {
	_, ok := p.Persons[payload.Id]
	if !ok {
		return errors.New("person not found")
	}
	delete(p.Persons, payload.Id)
	return nil
}

func (p *inMemoryPersonService) GetPersonById(ctx context.Context, id string, source string) (person *api.Person, err error) {
	res, ok := p.Persons[id]
	if ok {
		return &res, nil
	}
	return nil, errors.New("person not found")
}

func (p *inMemoryPersonService) GetPersonsById(ctx context.Context, ids []string, source string) (persons []api.Person, err error) {
	result := make([]api.Person, 0, len(ids))
	for _, id := range ids {
		res, ok := p.Persons[id]
		if ok {
			result = append(result, res)
		}
	}
	return result, nil
}

func (p *inMemoryPersonService) FindPersons(ctx context.Context, query api.FindPersonQuery) (result []api.Person, err error) {
	hits := make([]api.Person, 0)
	filter := strings.ToUpper(query.Filter)
	for _, t := range p.Persons {
		if strings.Contains(strings.ToUpper(t.Lastname), filter) ||
			strings.Contains(strings.ToUpper(t.Firstname), filter) ||
			strings.Contains(strings.ToUpper(t.Middlename), filter) {
			hits = append(hits, t)
		}
	}
	return hits, nil
}
