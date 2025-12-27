package day1_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, 6, Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, 6558, Run("input.txt"))
}
