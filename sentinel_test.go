package sentinel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	v := Valid[any]()
	valid, err := v.Validate(nil)
	require.True(t, valid)
	require.Nil(t, err)
}

func TestInvalid(t *testing.T) {
	v := Invalid[any]()
	valid, err := v.Validate(nil)
	require.False(t, valid)
	require.NotNil(t, err)
}

func TestNot(t *testing.T) {
	t.Run("not valid is always false", func(t *testing.T) {
		v := Not(Valid[any]())
		valid, _ := v.Validate(nil)
		require.False(t, valid)
	})

	t.Run("not invalid is always true", func(t *testing.T) {
		v := Not(Invalid[any]())
		valid, _ := v.Validate(nil)
		require.True(t, valid)
	})
}

func TestAnd(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		v Validator[any]
	}{
		{
			name: "valid valid",
			want: true,

			v: And(
				Valid[any](),
				Valid[any](),
			),
		},
		{
			name: "valid invalid",
			want: false,

			v: And(
				Valid[any](),
				Invalid[any](),
			),
		},
		{
			name: "invalid valid",
			want: false,

			v: And(
				Invalid[any](),
				Valid[any](),
			),
		},
		{
			name: "invalid invalid",
			want: false,

			v: And(
				Invalid[any](),
				Invalid[any](),
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

		v Validator[any]
	}{
		{
			name: "valid valid",
			want: true,

			v: Or(
				Valid[any](),
				Valid[any](),
			),
		},
		{
			name: "valid invalid",
			want: true,

			v: Or(
				Valid[any](),
				Invalid[any](),
			),
		},
		{
			name: "invalid valid",
			want: true,

			v: Or(
				Invalid[any](),
				Valid[any](),
			),
		},
		{
			name: "invalid invalid",
			want: false,

			v: Or(
				Invalid[any](),
				Invalid[any](),
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
