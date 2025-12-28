package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(1227775554), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(52316131093), Run("input.txt"))
}
