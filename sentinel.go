package sentinel

import "github.com/jordanhasgul/multierr"

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

// Valid returns a Validator that always returns true.
func Valid[T any]() Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		return true, nil
	})
}

// Invalid returns a [Validator] that always returns false.
func Invalid[T any]() Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		return false, nil
	})
}

// Not takes a single [Validator] and returns a [Validator] that negates
// the result of the original validator.
func Not[T any](v Validator[T]) Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		ok, err := v.Validate(t)
		if ok || err != nil {
			return false, err
		}

		return true, nil
	})
}

// And takes multiple [Validator] and returns a [Validator] that returns true
// if all the provided validators succeed.
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

// Or takes multiple [Validator] and returns a [Validator] that returns true
// if any of the provided validators succeed.
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
