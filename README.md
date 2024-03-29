# Requester

[![Coverage Status](https://coveralls.io/repos/github/dogukanayd/requester/badge.svg?branch=main)](https://coveralls.io/github/dogukanayd/requester?branch=main)

Request is a HTTP request library for Go with interfaces and mocks for unit tests.


The reason for having this package is using the interfaces to mock the HTTP requests easily on the codebase. Before
creating this repository I was copying and pasting this package to my projects. So I decided to move this package to the
separated package from my projects and use it with all of my projects.

## How to use?

In your codebase, you should create a package called "requester" or with a name anything that you want. After that, you 
have to create exactly the same interfaces that the package provides. Right after that you can mock the methods and 
use them inside your unit tests.

Example usage of the package:
```go
type ExampleRequestBody struct {
	Name    string
	Surname string
}

func example() {
	erb := ExampleRequestBody{
		Name:    "Dogukan",
		Surname: "Aydogdu",
	}

	erbj, _ := json.Marshal(erb)

	response, err := (&Request{
		Timeout: 60,
		Headers: []map[string]interface{}{
			{
				"Content-Type": "application/json",
			},
		},
		Endpoint: "https://www.example.com",
		Body:     erbj,
	}).Post()
}
```

Example for mock the interface your own codebase:

```go
package your_request_package

import (
	"github.com/dogukanayd/requester"
	"net/http"
)

type Requester interface {
	Get() (*http.Response, error)
	Post() (*http.Response, error)
	Put() (*http.Response, error)
	Delete() (*http.Response, error)
}

```
