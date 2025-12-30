package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(3), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(640), Run("input.txt"))
}
