package day9_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(24), Run("small_input.txt"))
}

func TestRunWithSmallInputModified(t *testing.T) {
	require.Equal(t, int64(36), Run("small_input_modified.txt"))
}

func TestRunWithSmallInputModified2(t *testing.T) {
	require.Equal(t, int64(12), Run("small_input_modified_2.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(1574684850), Run("input.txt"))
}
