package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	require.Equal(t, 524, Run("input.txt"))
}
