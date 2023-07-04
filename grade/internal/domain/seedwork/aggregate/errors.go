package aggregate

import "errors"

var ErrConcurrency = errors.New(
	"aggregate modified concurrently",
)
