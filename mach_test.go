package mach

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMachTaskSelf(t *testing.T) {
	require.NotZero(t, Mach_task_self())
}
