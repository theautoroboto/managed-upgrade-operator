package availabilitychecks

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type HTTPAvailabilityChecker struct {
	Targets []string
	Timeout time.Duration
}

type HTTPConfig struct {
	Targets []string
	Timeout time.Duration
}

// GetHTTPAvailabilityChecker returns a HTTP implementation of the
// AvailabilityChecker interface.
func GetHTTPAvailabilityChecker(c Config) (AvailabilityChecker, error) {
	if data, ok := c.(HTTPConfig); ok {
		return &HTTPAvailabilityChecker{
			Targets: data.Targets,
			Timeout: data.Timeout,
		}, nil
	}
	return &HTTPAvailabilityChecker{}, fmt.Errorf("Attempt to get HTTP implementation failed ascertation as HTTPConfig")
}

// HTTPAvailabilityCheck accepts a slice of HTTP targets accompanied by a timeout
// and asynchronously checks the targets to deem them available. If any of the targets
// are deemed unhealthy, all running routines are cancelled and an error is returned.
func (h HTTPAvailabilityChecker) AvailabilityCheck() error {
	client := http.Client{
		Timeout: (h.Timeout * time.Second),
	}

	var HTTPAvailabilityError error

	// WaitGroup to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// Set the WaitGroup counter to the length of Targets. The WaitGroup
	// can then block completion of the call until all goroutines have completed
	// by calling wg.Done
	wg.Add(len(h.Targets))

	for _, url := range h.Targets {

		go func(url string) {
			err := retry(3, time.Second, func() error {
				resp, err := client.Get(url)

				if err != nil {
					return err
				}

				if resp != nil && resp.Body != nil {
					defer resp.Body.Close()
				}

				s := resp.StatusCode

				switch {
				case s >= 500:
					return fmt.Errorf("Server error: %v", s)
				case s >= 400:
					return stop{fmt.Errorf("Client error: %v", s)}
				default:
					return nil
				}
			})

			if err != nil {
				HTTPAvailabilityError = err
			}

			// Call this to drop the wg Coutner by 1
			wg.Done()

		}(url)
	}
	// Waits for WaitGroup counter to be 0.
	wg.Wait()
	return HTTPAvailabilityError
}
