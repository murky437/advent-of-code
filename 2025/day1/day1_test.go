package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, 3, Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, 1135, Run("input.txt"))
}
