package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(50), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(4750176210), Run("input.txt"))
}
