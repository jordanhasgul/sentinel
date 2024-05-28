package sentinel

import (
	"fmt"

	"github.com/jordanhasgul/multierr"
)

// Validator validates instances of type T.
type Validator[T any] interface {
	Validate(t T) (bool, error)
}

// ValidateFunc is an adapter that allows the use of ordinary Go functions
// as validators. If f is a function with the appropriate signature, then
// ValidateFunc(f) is a Validator that calls f.
type ValidateFunc[T any] func(t T) (bool, error)

// Validate calls f(t).
func (f ValidateFunc[T]) Validate(t T) (bool, error) {
	return f(t)
}

// WithValue returns a Validator that returns true if v returns true when
// validating instances of T under the application of f.
func WithValue[T, U any](f func(T) U, v Validator[U]) ValidateFunc[T] {
	return func(t T) (bool, error) {
		return v.Validate(f(t))
	}
}

func WithValues[T, U any](f, g func(T) U, h func(U) Validator[U]) ValidateFunc[T] {
	return func(t T) (bool, error) {
		return h(g(t)).Validate(f(t))
	}
}

// Not returns a Validator that negates the result of v when validating
// instances of T.
func Not[T any](v Validator[T]) ValidateFunc[T] {
	return func(t T) (bool, error) {
		_, err := v.Validate(t)
		if err == nil {
			errorStr := "not-ing: %s"
			return false, fmt.Errorf(errorStr, "")
		}

		return true, nil
	}
}

// And returns a Validator that returns true if all the validators in vs
// return true when validating instances of T.
func And[T any](vs ...Validator[T]) ValidateFunc[T] {
	if len(vs) < 2 {
		panicStr := "sentinel: 'And' must be called with at least 2 validators but was given %d."
		panic(fmt.Sprintf(panicStr, len(vs)))
	}

	return func(t T) (bool, error) {
		var e *multierr.Error
		for _, v := range vs {
			_, err := v.Validate(t)
			if err != nil {
				e = multierr.Append(e, err)
			}
		}
		if e.Len() != 0 {
			errorStr := "and-ing: %w"
			return false, fmt.Errorf(errorStr, e)
		}

		return true, nil
	}
}

// Or returns a Validator that returns true if any of the validators in vs
// return true when validating instances of T.
func Or[T any](vs ...Validator[T]) ValidateFunc[T] {
	if len(vs) < 2 {
		panicStr := "sentinel: 'Or' must be called with at least 2 validators but was given %d."
		panic(fmt.Sprintf(panicStr, len(vs)))
	}

	return func(t T) (bool, error) {
		var e *multierr.Error
		for _, v := range vs {
			_, err := v.Validate(t)
			if err != nil {
				e = multierr.Append(e, err)
			}
		}
		if e.Len() == len(vs) {
			errorStr := "or-ing: %w"
			return false, fmt.Errorf(errorStr, e)
		}

		return true, nil
	}
}
