package day2_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(4174379265), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(69564213293), Run("input.txt"))
}
