package examples

import (
	"fmt"
	"github.com/jeanmolossi/go-http"
)

func SimpleRequestExample() {
	defer intervaler("simple request example")()

	// The request

	client := gohttp.New(gohttp.WithBaseURL(BaseURL))
	response, err := client.Get(GetPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !response.Success() {
		fmt.Println(fmt.Errorf("response has ended without success"))
		return
	}

	var result *Tracer
	err = response.JSON(&result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", result)
}
