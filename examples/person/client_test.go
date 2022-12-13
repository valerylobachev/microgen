package person

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"github.com/valerylobachev/microgen/examples/person/api"
	http "github.com/valerylobachev/microgen/examples/person/transport/http"
	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func NewClient() api.PersonService {
	rand.Seed(time.Now().UnixNano())
	u, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err.Error())
	}
	set := http.NewHTTPClient(u)

	return set
}

func Test_CreatePerson(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	id := strconv.Itoa(rand.Intn(10000000))
	payload := api.CreatePersonPayload{
		Id:         id,
		Lastname:   "Doe",
		Firstname:  "John",
		Middlename: "Max",
		CategoryId: "USER",
		Phone:      "+1-234-567-8901",
		Email:      "jd@example.com",
		Attributes: map[string]string{
			"gender": "male",
		},
		UpdatedBy: "person~P0001",
	}

	err := client.CreatePerson(ctx, payload)
	if err != nil {
		t.Errorf("failed to create person: %s", err.Error())
	}
	person, err := client.GetPersonById(ctx, id, "write")
	if err != nil {
		t.Errorf("failed to get person: %s", err.Error())
	}
	if (person.Lastname != payload.Lastname) ||
		(person.Firstname != payload.Firstname) ||
		(person.Middlename != payload.Middlename) {
		fmt.Println(payload)
		fmt.Println(*person)
		t.Error("persons not equal")
	}

}

func Test_GetPersonById(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	person, err := client.GetPersonById(ctx, "P00011", "write")
	if err != nil {
		t.Errorf("failed to get persons: %s", err.Error())
	}
	if person == nil {
		t.Error("result is empty")
	}
}

func Test_GetPersonsById(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	persons, err := client.GetPersonsById(ctx, []string{"P0001", "P0002", "P0003"}, "write")
	if err != nil {
		t.Errorf("failed to get persons: %s", err.Error())
	}
	lo.ForEach(persons, func(item api.Person, _ int) {
		fmt.Println(item)
	})
	if len(persons) == 0 {
		t.Error("result is empty")
	}
}

func Test_FindPersons(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	persons, err := client.FindPersons(ctx, api.FindPersonQuery{
		Offset: 0,
		Size:   100,
		Filter: "kri",
	})
	if err != nil {
		t.Errorf("failed to find person: %s", err.Error())
	}
	lo.ForEach(persons, func(item api.Person, _ int) {
		fmt.Println(item)
	})
	if len(persons) == 0 {
		t.Error("result is empty")
	}
}
