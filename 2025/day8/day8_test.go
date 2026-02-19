package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(40), Run("small_input.txt", 10))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(175440), Run("input.txt", 1000))
}
