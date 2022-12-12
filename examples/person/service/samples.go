package service

import (
	"github.com/valerylobachev/microgen/examples/person/api"
	"time"
)

var data = []api.Person{
	api.Person{
		Id:         "P0001",
		Lastname:   "Fisher",
		Firstname:  "Kristina",
		Middlename: "",
		CategoryId: "PERSON",
		Email:      "kristina.fisher@example.com",
		Phone:      "(754)-566-6120",
		Attributes: map[string]string{
			"gender": "female",
		},
		UpdatedAt: time.Now(),
		UpdatedBy: "person~P0001",
	},
	api.Person{
		Id:         "P0002",
		Lastname:   "Martin",
		Firstname:  "Leah",
		Middlename: "",
		CategoryId: "PERSON",
		Email:      "leah.martin@example.com",
		Phone:      "(257)-953-6866",
		Attributes: map[string]string{
			"gender": "female",
		},
		UpdatedAt: time.Now(),
		UpdatedBy: "person~P0001",
	},
	api.Person{
		Id:         "P0003",
		Lastname:   "Holland",
		Firstname:  "Patrick",
		Middlename: "",
		CategoryId: "PERSON",
		Email:      "patrick.holland@example.com",
		Phone:      "(369)-562-7084",
		Attributes: map[string]string{
			"gender": "male",
		},
		UpdatedAt: time.Now(),
		UpdatedBy: "person~P0001",
	},
}

func Samples() map[string]api.Person {
	res := make(map[string]api.Person)
	for _, d := range data {
		res[d.Id] = d
	}
	return res
}
