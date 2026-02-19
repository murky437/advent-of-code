package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(13), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(1508), Run("input.txt"))
}
