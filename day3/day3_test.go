package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(357), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(17332), Run("input.txt"))
}
