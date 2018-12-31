package main

import (
	"fmt"

	//"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

func main() {
	spec, err := loads.Spec("swagger.yaml")
	if err != nil {
		fmt.Println("Failed to load spec")
	}

	if err := validate.Spec(spec, strfmt.Default); err != nil {
		fmt.Println("Spec is invalid")
	} else {
		fmt.Println("Everything is good")
	}
}
