package sentinel

// Validator validates instances of a type, T.
type Validator[T any] interface {
	Validate(t T) (bool, error)
}

// ValidateFunc is an adapter that allows the use of ordinary
// Go functions as validators. If f is a function with the appropriate
// signature, then ValidateFunc(f) is a [Validator] that calls f.
type ValidateFunc[T any] func(t T) (bool, error)

func (f ValidateFunc[T]) Validate(t T) (bool, error) {
	return f(t)
}
