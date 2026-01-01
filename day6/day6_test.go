package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(4277556), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(5316572080628), Run("input.txt"))
}
