package day8_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(25272), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(3200955921), Run("input.txt"))
}
