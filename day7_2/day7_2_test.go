package day7_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunWithSmallInput(t *testing.T) {
	require.Equal(t, int64(40), Run("small_input.txt"))
}

func TestRun(t *testing.T) {
	require.Equal(t, int64(23607984027985), Run("input.txt"))
}
