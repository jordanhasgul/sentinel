package sentinel

import (
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

// Valid returns a Validator that always returns true when validating
// instances of type T.
func Valid[T any]() Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		return true, nil
	})
}

// Invalid returns a Validator that always returns false when validating
// instances of T.
func Invalid[T any]() Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		return false, nil
	})
}

// Not returns a Validator that negates the result of v when validating
// instances of T.
func Not[T any](v Validator[T]) Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		ok, err := v.Validate(t)
		if ok || err != nil {
			return false, err
		}

		return true, nil
	})
}

// And returns a Validator that returns true if all the validators in vs
// return true when validating instances of T.
func And[T any](vs ...Validator[T]) Validator[T] {
	if len(vs) == 0 {
		return Valid[T]()
	}

	return ValidateFunc[T](func(t T) (bool, error) {
		var e *multierr.Error
		for _, v := range vs {
			ok, err := v.Validate(t)
			if !ok || err != nil {
				e = multierr.Append(e, err)
			}
		}
		if e.Len() != 0 {
			return false, e
		}

		return true, nil
	})
}

// Or returns a Validator that returns true if any of the validators in vs
// return true when validating instances of T.
func Or[T any](vs ...Validator[T]) Validator[T] {
	if len(vs) == 0 {
		return Valid[T]()
	}

	return ValidateFunc[T](func(t T) (bool, error) {
		var e *multierr.Error
		for _, v := range vs {
			ok, err := v.Validate(t)
			if !ok || err != nil {
				e = multierr.Append(e, err)
			}
		}
		if e.Len() == len(vs) {
			return false, e
		}

		return true, nil
	})
}

// WithValue returns a Validator that returns true if v returns true when
// validating instances of T under the application of f.
func WithValue[T, U any](f func(T) U, v Validator[U]) Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		return v.Validate(f(t))
	})
}
