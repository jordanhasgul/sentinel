package validators

import (
	"cmp"
	"fmt"
	"reflect"

	"github.com/jordanhasgul/sentinel"
	"github.com/jordanhasgul/sentinel/constraints"
)

// Valid returns a sentinel.Validator that always returns true when validating
// instances of type T.
func Valid[T any]() sentinel.ValidateFunc[T] {
	return func(t T) (bool, error) {
		return true, nil
	}
}

// Invalid returns a sentinel.Validator that always returns false when validating
// instances of T.
func Invalid[T any]() sentinel.ValidateFunc[T] {
	return func(t T) (bool, error) {
		errString := "'%#v' of type '%s' is always invalid"
		return false, fmt.Errorf(errString, t, reflect.TypeOf(t))
	}
}

// Equal returns a sentinel.Validator that returns true if t1 == t2, where t1 is an
// instance of type T.
func Equal[T constraints.Equated](t2 T) sentinel.ValidateFunc[T] {
	eq := func(a, b T) bool {
		return a == b
	}
	return EqualFunc(eq)(t2)
}

func EqualFunc[T any](eq func(T, T) bool) func(T) sentinel.ValidateFunc[T] {
	return func(t2 T) sentinel.ValidateFunc[T] {
		return func(t1 T) (bool, error) {
			if !eq(t1, t2) {
				errString := "'%#v' is not equal to '%#v' (both of type '%s')"
				return false, fmt.Errorf(errString, t1, t2, reflect.TypeOf(t1))
			}

			return true, nil
		}
	}
}

// NotEqual returns a sentinel.Validator that returns true if t1 != t2, where t1 is
// an instance of type T.
func NotEqual[T constraints.Equated](t2 T) sentinel.ValidateFunc[T] {
	return sentinel.Not(Equal[T](t2))
}

func NotEqualFunc[T any](eq func(T, T) bool) func(T) sentinel.ValidateFunc[T] {
	return func(t2 T) sentinel.ValidateFunc[T] {
		return sentinel.Not(EqualFunc[T](eq)(t2))
	}
}

// Less returns a sentinel.Validator that returns true if t1 < t2, where t1 is an
// instance of type T.
func Less[T constraints.Ordered](t2 T) sentinel.ValidateFunc[T] {
	return LessFunc(cmp.Compare[T])(t2)
}

func LessFunc[T any](cmp func(T, T) int) func(T) sentinel.ValidateFunc[T] {
	return func(t2 T) sentinel.ValidateFunc[T] {
		return func(t1 T) (bool, error) {
			if !(cmp(t1, t2) < 0) {
				errString := "'%#v' is not less than '%#v' (both of type '%s')"
				return false, fmt.Errorf(errString, t1, t2, reflect.TypeOf(t1))
			}

			return true, nil
		}
	}
}

// LessOrEqual returns a sentinel.Validator that returns true if t1 <= t2, where t1
// is an instance of type T.
func LessOrEqual[T constraints.Ordered](t2 T) sentinel.ValidateFunc[T] {
	return sentinel.Or(
		Less[T](t2),
		Equal[T](t2),
	)
}

func LessOrEqualFunc[T any](cmp func(T, T) int) func(T) sentinel.ValidateFunc[T] {
	return func(t2 T) sentinel.ValidateFunc[T] {
		eq := func(a, b T) bool {
			return cmp(a, b) == 0
		}
		return sentinel.Or(
			LessFunc[T](cmp)(t2),
			EqualFunc[T](eq)(t2),
		)
	}
}

// Greater returns a sentinel.Validator that returns true if t1 > t2, where t1 is an
// instance of type T.
func Greater[T constraints.Ordered](t2 T) sentinel.ValidateFunc[T] {
	return GreaterFunc(cmp.Compare[T])(t2)
}

func GreaterFunc[T any](cmp func(T, T) int) func(T) sentinel.ValidateFunc[T] {
	return func(t2 T) sentinel.ValidateFunc[T] {
		return func(t1 T) (bool, error) {
			if !(cmp(t1, t2) > 0) {
				errString := "'%#v' is not greater than '%#v' (both of type '%s')"
				return false, fmt.Errorf(errString, t1, t2, reflect.TypeOf(t1))
			}

			return true, nil
		}
	}
}

// GreaterOrEqual returns a sentinel.Validator that returns true if t1 >= t2, where
// t1 is an instance of type T.
func GreaterOrEqual[T constraints.Ordered](t2 T) sentinel.ValidateFunc[T] {
	return sentinel.Or(
		Greater[T](t2),
		Equal[T](t2),
	)
}

func GreaterOrEqualFunc[T any](cmp func(T, T) int) func(T) sentinel.ValidateFunc[T] {
	return func(t2 T) sentinel.ValidateFunc[T] {
		eq := func(a, b T) bool {
			return cmp(a, b) == 0
		}
		return sentinel.Or(
			GreaterFunc[T](cmp)(t2),
			EqualFunc[T](eq)(t2),
		)
	}
}

// Nil returns a sentinel.Validator that returns true if T is nillable and t == nil,
// where t is an instance of type T.
func Nil[T any]() sentinel.ValidateFunc[T] {
	return func(t T) (bool, error) {
		var (
			value = reflect.ValueOf(t)
			kind  = value.Kind()

			nillable = kind == reflect.Ptr || kind == reflect.UnsafePointer ||
				kind == reflect.Func || kind == reflect.Map || kind == reflect.Slice ||
				kind == reflect.Chan || kind == reflect.Interface
		)
		if !nillable || !value.IsNil() {
			errString := "'%#v' of type '%s' is not nil"
			return false, fmt.Errorf(errString, t, reflect.TypeOf(value))
		}

		return true, nil
	}
}

// NotNil returns a sentinel.Validator that returns true if T is nillable and
// t != nil, where t is an instance of type T.
func NotNil[T any]() sentinel.ValidateFunc[T] {
	return sentinel.Not(Nil[T]())
}

func Positive[R constraints.Real]() sentinel.ValidateFunc[R] {
	return Greater[R](0)
}

func Negative[R constraints.Real]() sentinel.ValidateFunc[R] {
	return Less[R](0)
}
