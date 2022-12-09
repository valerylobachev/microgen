package api

import (
	"time"
)

type PersonId string

func (c PersonId) String() string {
	return string(c)
}

type Person struct {
	Id         string            `json:"id"`
	Lastname   string            `json:"lastname"`
	Firstname  string            `json:"firstname"`
	Middlename string            `json:"middlename,omitempty"`
	CategoryId string            `json:"categoryId"`
	Phone      string            `json:"phone,omitempty"`
	Email      string            `json:"email,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	UpdatedBy  string            `json:"updatedBy"`
}

type CreatePersonPayload struct {
	Id         string            `json:"id"`
	Lastname   string            `json:"lastname"`
	Firstname  string            `json:"firstname"`
	Middlename string            `json:"middlename,omitempty"`
	CategoryId string            `json:"categoryId"`
	Phone      string            `json:"phone,omitempty"`
	Email      string            `json:"email,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	UpdatedBy  string            `json:"updatedBy"`
}

type UpdatePersonPayload struct {
	Id         string            `json:"id"`
	Lastname   string            `json:"lastname"`
	Firstname  string            `json:"firstname"`
	Middlename string            `json:"middlename,omitempty"`
	CategoryId string            `json:"categoryId"`
	Phone      string            `json:"phone,omitempty"`
	Email      string            `json:"email,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	UpdatedBy  string            `json:"updatedBy"`
}

type DeletePersonPayload struct {
	Id        string `json:"id"`
	UpdatedBy string `json:"updatedBy" validate:"required"`
}

type FindPersonQuery struct {
	Offset int    `json:"offset"`
	Size   int    `json:"size" validate:"required"`
	Filter string `json:"filter,omitempty"`
}
