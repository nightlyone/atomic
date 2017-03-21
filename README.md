# atomic
Atomic operations as methods on types

[![GoDoc](https://godoc.org/github.com/nightlyone/atomic?status.svg)](https://godoc.org/github.com/nightlyone/atomic)
[![Build Status](https://secure.travis-ci.org/nightlyone/atomic.png)](https://travis-ci.org/nightlyone/atomic)

# example usage
```go
package main

import (
	"fmt"

	"github.com/nightlyone/atomic"
)

type Service struct {
	Health atomic.Bool
}

func main() {
	service := new(Service)
	isHealthy := service.Health.Value()
	fmt.Printf("service is healthy? %t\n", isHealthy)
	// Output: service is healthy? false
}
```
# LICENSE
BSD-3-Clause

# Contributing
Pull requests and github issues welcome.
