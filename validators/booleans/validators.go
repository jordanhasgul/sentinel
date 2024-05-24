package booleans

import (
	"github.com/jordanhasgul/sentinel"
	"github.com/jordanhasgul/sentinel/constraints"
	"github.com/jordanhasgul/sentinel/validators"
)

// IsTrue returns a Validator that returns true if b == true, where b is an
// instance of type B.
func IsTrue[B constraints.Boolean]() sentinel.ValidateFunc[B] {
	return validators.Equal[B](true)
}

// IsFalse returns a Validator that returns true if b == false, where b is an
// instance of type B.
func IsFalse[B constraints.Boolean]() sentinel.ValidateFunc[B] {
	return validators.Equal[B](false)
}
