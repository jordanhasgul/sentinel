package sentinel_test

import (
	"testing"

	"github.com/jordanhasgul/sentinel"
	"github.com/jordanhasgul/sentinel/validators"
	"github.com/stretchr/testify/require"
)

func TestNot(t *testing.T) {
	t.Run("not valid is always false", func(t *testing.T) {
		v := sentinel.Not(validators.Valid[any]())
		valid, _ := v.Validate(nil)
		require.False(t, valid)
	})

	t.Run("not invalid is always true", func(t *testing.T) {
		v := sentinel.Not(validators.Invalid[any]())
		valid, _ := v.Validate(nil)
		require.True(t, valid)
	})
}

func TestAnd(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		v sentinel.Validator[any]
	}{
		{
			name: "valid valid",
			want: true,

			v: sentinel.And(
				validators.Valid[any](),
				validators.Valid[any](),
			),
		},
		{
			name: "valid invalid",
			want: false,

			v: sentinel.And(
				validators.Valid[any](),
				validators.Invalid[any](),
			),
		},
		{
			name: "invalid valid",
			want: false,

			v: sentinel.And(
				validators.Invalid[any](),
				validators.Valid[any](),
			),
		},
		{
			name: "invalid invalid",
			want: false,

			v: sentinel.And(
				validators.Invalid[any](),
				validators.Invalid[any](),
			),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, _ := testCase.v.Validate(nil)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestOr(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		v sentinel.Validator[any]
	}{
		{
			name: "valid valid",
			want: true,

			v: sentinel.Or(
				validators.Valid[any](),
				validators.Valid[any](),
			),
		},
		{
			name: "valid invalid",
			want: true,

			v: sentinel.Or(
				validators.Valid[any](),
				validators.Invalid[any](),
			),
		},
		{
			name: "invalid valid",
			want: true,

			v: sentinel.Or(
				validators.Invalid[any](),
				validators.Valid[any](),
			),
		},
		{
			name: "invalid invalid",
			want: false,

			v: sentinel.Or(
				validators.Invalid[any](),
				validators.Invalid[any](),
			),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, _ := testCase.v.Validate(nil)
			require.Equal(t, testCase.want, got)
		})
	}
}
