package examples

import (
	"fmt"
	"github.com/jeanmolossi/go-http"
	"sync"
)

func ConcurrentRequests() {
	defer intervaler("concurrent requests")()

	// The requests

	client := gohttp.New(gohttp.WithBaseURL(BaseURL))

	workers := 100
	wg := &sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()

			path := status(500)
			if i > 2 {
				path = GetPath
			}

			response, err := client.Get(path)
			if err != nil {
				fmt.Println("[ERROR]", err)
				return
			}

			if response.Success() {
				fmt.Println("[INFO] response is ok", response.Status())
			}

			var target Tracer
			err = response.JSON(&target)
			if err != nil {
				fmt.Println("[ERROR]", err)
				return
			}

			fmt.Printf("[INFO] %v\n", target.Headers.AmznTraceID)
		}(i)
	}

	wg.Wait()
}
