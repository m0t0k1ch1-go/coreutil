package coreutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/coreutil"
)

func TestPtr(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
		}{
			{
				name: "int",
				in:   42,
			},
			{
				name: "string",
				in:   "foo",
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				require.Equal(t, &tc.in, coreutil.Ptr(tc.in))
			})
		}
	})
}
