package sentinel

import (
	"reflect"

	"github.com/jordanhasgul/multierr"
	"github.com/jordanhasgul/sentinel/constraints"
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
	// todo: what if len(vs) < 2

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
	// todo: what if len(vs) < 2

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

// Equal returns a Validator that returns true if t1 == t2, where t1 is an
// instance of type T.
func Equal[T constraints.Equated](t2 T) Validator[T] {
	return ValidateFunc[T](func(t1 T) (bool, error) {
		return t1 == t2, nil
	})
}

// NotEqual returns a Validator that returns true if t1 != t2, where t1 is
// an instance of type T.
func NotEqual[T constraints.Equated](t2 T) Validator[T] {
	return Not(Equal[T](t2))
}

// EqualFunc returns a Validator that returns true if f(t1, t2) == true,
// where t1 is an instance of type T.
func EqualFunc[T any](f func(T, T) bool, t2 T) Validator[T] {
	return ValidateFunc[T](func(t1 T) (bool, error) {
		return f(t1, t2), nil
	})
}

// NotEqualFunc returns a Validator that returns true if f(t1, t2) != true,
// where t1 is an instance of type T.
func NotEqualFunc[T any](f func(T, T) bool, t2 T) Validator[T] {
	return Not(EqualFunc[T](f, t2))
}

// Less returns a Validator that returns true if t1 < t2, where t1 is an
// instance of type T.
func Less[T constraints.Ordered](t2 T) Validator[T] {
	return ValidateFunc[T](func(t1 T) (bool, error) {
		return t1 < t2, nil
	})
}

// LessFunc returns a Validator that returns true if f(t1, t2) < 0, where
// t1 is an instance of type T.
func LessFunc[T any](f func(T, T) int, t2 T) Validator[T] {
	return ValidateFunc[T](func(t1 T) (bool, error) {
		return f(t1, t2) < 0, nil
	})
}

// LessOrEqual returns a Validator that returns true if t1 <= t2, where t1
// is an instance of type T.
func LessOrEqual[T constraints.Ordered](t2 T) Validator[T] {
	return Or(Less[T](t2), Equal[T](t2))
}

// LessOrEqualFunc returns a Validator that returns true if f(t1, t2) <= 0,
// where t1 is an instance of type T.
func LessOrEqualFunc[T any](f func(T, T) int, t2 T) Validator[T] {
	eq := func(a, b T) bool { return f(a, b) == 0 }
	return Or(LessFunc[T](f, t2), EqualFunc[T](eq, t2))
}

// Greater returns a Validator that returns true if t1 > t2, where t1 is an
// instance of type T.
func Greater[T constraints.Ordered](t2 T) Validator[T] {
	return ValidateFunc[T](func(t1 T) (bool, error) {
		return t1 > t2, nil
	})
}

// GreaterFunc returns a Validator that returns true if f(t1, t2) < 0,
// where t1 is an instance of type T.
func GreaterFunc[T any](f func(T, T) int, t2 T) Validator[T] {
	return ValidateFunc[T](func(t1 T) (bool, error) {
		return f(t1, t2) > 0, nil
	})
}

// GreaterOrEqual returns a Validator that returns true if t1 >= t2, where
// t1 is an instance of type T.
func GreaterOrEqual[T constraints.Ordered](t2 T) Validator[T] {
	return Or(Greater[T](t2), Equal[T](t2))
}

// GreaterOrEqualFunc returns a Validator that returns true if f(t1, t2) <= 0,
//
//	where t1 is an instance of type T.
func GreaterOrEqualFunc[T any](f func(T, T) int, t2 T) Validator[T] {
	eq := func(a, b T) bool { return f(a, b) == 0 }
	return Or(GreaterFunc[T](f, t2), EqualFunc[T](eq, t2))
}

// Nil returns a Validator that returns true if T is nillable and t == nil,
// where t is an instance of type T.
func Nil[T any]() Validator[T] {
	return ValidateFunc[T](func(t T) (bool, error) {
		var (
			value = reflect.ValueOf(t)
			kind  = value.Kind()

			nillable = kind == reflect.Ptr || kind == reflect.UnsafePointer ||
				kind == reflect.Func || kind == reflect.Map || kind == reflect.Slice ||
				kind == reflect.Chan || kind == reflect.Interface
		)
		return nillable && value.IsNil(), nil
	})
}

// NotNil returns a Validator that returns true if T is nillable and
// t != nil, where t is an instance of type T.
func NotNil[T any]() Validator[T] {
	return Not(Nil[T]())
}
