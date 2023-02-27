package examples

import (
	"fmt"
	"github.com/jeanmolossi/go-http"
	"time"
)

func DelayedRequest() {
	defer intervaler("delayed request example")()

	// The request

	client := gohttp.New(
		gohttp.WithBaseURL(BaseURL),
		gohttp.WithTimeout(time.Second*2),
	)

	response, err := client.Get(delay(10))
	if err != nil {
		fmt.Println(err)
		return
	}

	if !response.Success() {
		fmt.Println(fmt.Errorf("response is not ok"))
		return
	}

	var result Tracer
	err = response.JSON(&result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[INFO] %+v\n", result.Headers.AmznTraceID)
}
