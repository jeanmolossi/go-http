# go-http

package to pollyfill some thing in golang requests

# Installing

```sh
go get github.com/jeanmolossi/go-http
```

# Usage

Basic usage

```go
package main

import (
    "fmt"
    "github.com/jeanmolossi/go-http"
)

type (
    TraceHeaders struct {
        AmznTraceID string `json:"X-Amzn-Trace-Id"`
    }

    Tracer struct {
        Headers TraceHeaders `json:"headers"`
    }
)

func main() {
    // instantiate a new client
    client := gohttp.New(
        // defines the base url to requests
        gohttp.WithBaseURL("https://httpbin.org"),
    )

    // perform a get request
    response, err := client.Get("/get")
    if err != nil {
        fmt.Println(err)
        return
    }

    // if response has not success status code
    // statusCode < 300
    if !response.Success() {
        fmt.Println(fmt.Errorf("response has ended without success"))
        return
    }

    var result *Tracer
    // unmarshal as json the result into *Tracer struct
    err = response.JSON(&result)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("%+v\n", result) // &{Headers:{AmznTraceID:Root=1-00aa0aa0-00aa00aa0a0a00aa0a00000a}}
}
```

You can look more in [examples](./examples) folder

- [Simple request](./examples/simple_request.go)
- [Concurrent requests](./examples/concurrent_requests.go)
- [Request with timeout](./examples/delayed_request.go)