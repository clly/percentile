package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var nilFloat []float64

func TestParse(t *testing.T) {
	testcases := []struct {
		Test   []string
		Result []float64
		Err    bool
	}{
		{
			Test:   []string{},
			Result: []float64{.50, .95},
			Err:    false,
		},
		{
			Test:   []string{"10", "20"},
			Result: []float64{.10, .20},
			Err:    false,
		},
		{
			Test:   []string{"101"},
			Result: nilFloat,
			Err:    true,
		},
		{
			Test:   []string{"-1"},
			Result: nilFloat,
			Err:    true,
		},
		{
			Test:   []string{"foo"},
			Result: nilFloat,
			Err:    true,
		},
	}

	for _, test := range testcases {
		percentiles, err := parse(test.Test)
		require.Equal(t, test.Result, percentiles, "Failed on test case %+v", test)
		if test.Err {
			require.Error(t, err, "Failed on test case %+v", test)
		} else {
			require.NoError(t, err, "Failed on test case %+v", test)
		}
	}
}

func TestReadToFloat(t *testing.T) {
	testcases := []struct {
		Test   *strings.Reader
		Result []float64
		Err    bool
	}{
		{
			Test:   strings.NewReader("1\n2\n3\n4\n5\n10\n9\n8\n7\n6\n"),
			Result: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Err:    false,
		},
		{
			Test:   strings.NewReader("foo\nbar"),
			Result: nilFloat,
			Err:    true,
		},
	}

	for _, test := range testcases {
		percentiles, err := readToFloat(test.Test)
		require.Equal(t, test.Result, percentiles, "Failed on test case %+v", test)
		if test.Err {
			require.Error(t, err, "Failed on test case %+v", test)
		} else {
			require.NoError(t, err, "Failed on test case %+v", test)
		}
	}
}
