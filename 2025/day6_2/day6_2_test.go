package day6_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(3263827), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(11299263623062), Run("input.txt"))
}
