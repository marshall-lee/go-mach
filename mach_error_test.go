// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mach

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMachErrorString(t *testing.T) {
	require.Equal(t, "(ipc/send) invalid destination port", Mach_error_string(0x10000003))
	require.Equal(t, "(ipc/send) invalid destination port", kern_return_err(0x10000003).Error())
}
