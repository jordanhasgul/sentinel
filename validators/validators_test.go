package validators_test

import (
	"bytes"
	"testing"

	"github.com/jordanhasgul/sentinel/validators"
	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	v := validators.Valid[any]()
	valid, _ := v.Validate(nil)
	require.True(t, valid)
}

func TestInvalid(t *testing.T) {
	v := validators.Invalid[any]()
	valid, _ := v.Validate(nil)
	require.False(t, valid)
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
			v := validators.Equal(testCase.t2)

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
			v := validators.EqualFunc(bytes.Equal)(testCase.t2)

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
			v := validators.NotEqual(testCase.t2)

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
			v := validators.NotEqualFunc(bytes.Equal)(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestLess(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 int
		t2 int
	}{
		{
			name: "0 < 0",
			want: false,

			t1: 0,
			t2: 0,
		},
		{
			name: "1 < 0",
			want: false,

			t1: 1,
			t2: 0,
		},
		{
			name: "0 < 1",
			want: true,

			t1: 0,
			t2: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.Less(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestLessFunc(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 []byte
		t2 []byte
	}{
		{
			name: "'test' < 'test'",
			want: false,

			t1: []byte("test"),
			t2: []byte("test"),
		},
		{
			name: "'tset' < 'test'",
			want: false,

			t1: []byte("tset"),
			t2: []byte("test"),
		},
		{
			name: "'test' < 'tset'",
			want: true,

			t1: []byte("test"),
			t2: []byte("tset"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.LessFunc(bytes.Compare)(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestLessOrEqual(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 int
		t2 int
	}{
		{
			name: "0 <= 0",
			want: true,

			t1: 0,
			t2: 0,
		},
		{
			name: "1 <= 0",
			want: false,

			t1: 1,
			t2: 0,
		},
		{
			name: "0 <= 1",
			want: true,

			t1: 0,
			t2: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.LessOrEqual(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestLessOrEqualFunc(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 []byte
		t2 []byte
	}{
		{
			name: "'test' <= 'test'",
			want: true,

			t1: []byte("test"),
			t2: []byte("test"),
		},
		{
			name: "'tset' <= 'test'",
			want: false,

			t1: []byte("tset"),
			t2: []byte("test"),
		},
		{
			name: "'test' <= 'tset'",
			want: true,

			t1: []byte("test"),
			t2: []byte("tset"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.LessOrEqualFunc(bytes.Compare)(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestGreater(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 int
		t2 int
	}{
		{
			name: "0 > 0",
			want: false,

			t1: 0,
			t2: 0,
		},
		{
			name: "1 > 0",
			want: true,

			t1: 1,
			t2: 0,
		},
		{
			name: "0 > 1",
			want: false,

			t1: 0,
			t2: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.Greater(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestGreaterFunc(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 []byte
		t2 []byte
	}{
		{
			name: "'test' > 'test'",
			want: false,

			t1: []byte("test"),
			t2: []byte("test"),
		},
		{
			name: "'tset' > 'test'",
			want: true,

			t1: []byte("tset"),
			t2: []byte("test"),
		},
		{
			name: "'test' > 'tset'",
			want: false,

			t1: []byte("test"),
			t2: []byte("tset"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.GreaterFunc(bytes.Compare)(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestGreaterOrEqual(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 int
		t2 int
	}{
		{
			name: "0 >= 0",
			want: true,

			t1: 0,
			t2: 0,
		},
		{
			name: "1 >= 0",
			want: true,

			t1: 1,
			t2: 0,
		},
		{
			name: "0 >= 1",
			want: false,

			t1: 0,
			t2: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.GreaterOrEqual(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestGreaterOrEqualFunc(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t1 []byte
		t2 []byte
	}{
		{
			name: "'test' >= 'test'",
			want: true,

			t1: []byte("test"),
			t2: []byte("test"),
		},
		{
			name: "'tset' >= 'test'",
			want: true,

			t1: []byte("tset"),
			t2: []byte("test"),
		},
		{
			name: "'test' >= 'tset'",
			want: false,

			t1: []byte("test"),
			t2: []byte("tset"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.GreaterOrEqualFunc(bytes.Compare)(testCase.t2)

			got, _ := v.Validate(testCase.t1)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestNil(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t any
	}{
		{
			name: "non-nillable type",
			want: false,

			t: 1,
		},
		{
			name: "nillable type non-nil value",
			want: false,

			t: []int{1},
		},
		{
			name: "nillable type nil value",
			want: true,

			t: []int(nil),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.Nil[any]()

			got, _ := v.Validate(testCase.t)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestNotNil(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t any
	}{
		{
			name: "non-nillable type",
			want: true,

			t: 1,
		},
		{
			name: "nillable type with non-nil value",
			want: true,

			t: []int{1},
		},
		{
			name: "nillable type with nil value",
			want: false,

			t: []int(nil),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			v := validators.NotNil[any]()

			got, _ := v.Validate(testCase.t)
			require.Equal(t, testCase.want, got)
		})
	}
}
