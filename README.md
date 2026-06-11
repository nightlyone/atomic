# atomic
Atomic operations as methods on types

**Note**: As of Go 1.19 the standard library package [`sync/atomic`](https://pkg.go.dev/sync/atomic)
implements similar data types, so this package here will not see further
development. It will stay archived.

[![Go Reference](https://pkg.go.dev/badge/github.com/nightlyone/atomic.svg)](https://pkg.go.dev/github.com/nightlyone/atomic)

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
