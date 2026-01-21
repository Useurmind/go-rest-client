package main

import (
	"fmt"

	"github.com/Useurmind/go-rest-client/pkg/client"
	"github.com/go-logr/stdr"
)

type DogList struct {
	Data []Dog `json:"data,omitempty"`
}

type Dog struct {
	Id    string `json:"id,omitempty"`
	Breed string `json:"breed,omitempty"`
}

func main() {
	logger := stdr.New(nil)
	// stdr.SetVerbosity(4)
	cli := client.NewRestClient("https://dogapi.dog/api/v2")
	cli.SetLogger(&logger)

	ctx := client.NewRequestContextJson[interface{}, DogList](cli)

	dogList := DogList{}
	status, statusCode, err := ctx.Get("breeds", nil, &dogList)
	fmt.Println(status, statusCode)
	if err != nil {
		panic(err)
	}

	status, statusCode, err = ctx.Get("400", nil, nil)
	fmt.Println(status, statusCode)
	if err != nil {
		panic(err)
	}
}
