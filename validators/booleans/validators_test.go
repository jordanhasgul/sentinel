package booleans_test

import (
	"testing"

	"github.com/jordanhasgul/sentinel/validators/booleans"
	"github.com/stretchr/testify/require"
)

func TestIsTrue(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t bool
	}{
		{
			name: "true is true",
			want: true,

			t: true,
		},
		{
			name: "false is true",
			want: false,

			t: false,
		},
	}
	for _, testCase := range testCases {
		v := booleans.IsTrue[bool]()

		got, _ := v.Validate(testCase.t)
		require.Equal(t, testCase.want, got)
	}
}

func TestIsFalse(t *testing.T) {
	testCases := []struct {
		name string
		want bool

		t bool
	}{
		{
			name: "true is false",
			want: false,

			t: true,
		},
		{
			name: "false is false",
			want: true,

			t: false,
		},
	}
	for _, testCase := range testCases {
		v := booleans.IsFalse[bool]()

		got, _ := v.Validate(testCase.t)
		require.Equal(t, testCase.want, got)
	}
}
