package sentinel

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	v := Valid[any]()
	valid, _ := v.Validate(nil)
	require.True(t, valid)
}

func TestInvalid(t *testing.T) {
	v := Invalid[any]()
	valid, _ := v.Validate(nil)
	require.False(t, valid)
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

func TestEqual(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 int
		t2 int
	}{
		{
			name: "0 equals 0",
			want: true,

			t1: 0,
			t2: 0,
		},
		{
			name: "0 equals 1",
			want: false,

			t1: 0,
			t2: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := Equal(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestNotEqual(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 int
		t2 int
	}{
		{
			name: "0 not equals 0",
			want: false,

			t1: 0,
			t2: 0,
		},
		{
			name: "0 not equals 1",
			want: true,

			t1: 0,
			t2: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := NotEqual(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestEqualFunc(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 []byte
		t2 []byte
	}{
		{
			name: "'test' equals 'test'",
			want: true,

			t1: []byte("test"),
			t2: []byte("test"),
		},
		{
			name: "'test' equals 'tset'",
			want: false,

			t1: []byte("test"),
			t2: []byte("tset"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := EqualFunc(bytes.Equal, testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestNotEqualFunc(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 []byte
		t2 []byte
	}{
		{
			name: "'test' not equals 'test'",
			want: false,

			t1: []byte("test"),
			t2: []byte("test"),
		},
		{
			name: "'test' not equals 'tset'",
			want: true,

			t1: []byte("test"),
			t2: []byte("tset"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := NotEqualFunc(bytes.Equal, testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}
