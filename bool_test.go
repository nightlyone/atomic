package atomic_test

import (
	"fmt"

	"github.com/nightlyone/atomic"
)

// Service is a flaky service
type Service struct {
	health atomic.Bool
}

func ExampleBool() {
	service := new(Service)
	isHealthy := service.health.Value()
	fmt.Printf("service is healthy? %t\n", isHealthy)
	// Output: service is healthy? false
}
