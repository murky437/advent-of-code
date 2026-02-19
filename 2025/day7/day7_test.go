package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(21), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(1619), Run("input.txt"))
}
