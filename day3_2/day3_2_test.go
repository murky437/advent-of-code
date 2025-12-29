package day3_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(3121910778619), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(172516781546707), Run("input.txt"))
}
