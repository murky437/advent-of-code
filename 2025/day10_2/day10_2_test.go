package day10_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, 33, Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, 16463, Run("input.txt"))
}
