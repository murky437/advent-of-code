package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(5), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(670), Run("input.txt"))
}
