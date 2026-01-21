# Go REST Client

This client should simplify rest client creation by encapsulating common functionality required for all REST requests.

## Usage

First define objects for the api requests and responses:

    type DogList struct {
        Data []Dog `json:"data,omitempty"`
    }

    type Dog struct {
        Id    string `json:"id,omitempty"`
        Breed string `json:"breed,omitempty"`
    }

The declare a client and 

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