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

// Valid returns a Validator that always returns true, without any error.
func Valid[T any]() Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		return true, nil
	})
}

// Invalid returns a Validator that always returns false, without any error.
func Invalid[T any]() Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		return false, nil
	})
}
