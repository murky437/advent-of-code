package day4_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(43), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(8538), Run("input.txt"))
}
