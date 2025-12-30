package day5_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(14), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(365804144481581), Run("input.txt"))
}
