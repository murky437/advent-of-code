package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(7), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(438), Run("input.txt"))
}
