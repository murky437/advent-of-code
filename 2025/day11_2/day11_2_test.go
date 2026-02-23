package day11_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(2), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(332052564714990), Run("input.txt"))
}
