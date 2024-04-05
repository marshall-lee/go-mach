// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mach

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestMachVMMirroredBuffer(t *testing.T) {
	address, err := Mach_vm_allocate(Mach_task_self(), 0, 8192*2, VM_FLAGS_ANYWHERE)
	require.NoError(t, err)

	_, err = Mach_vm_allocate(Mach_task_self(), address, 8192, VM_FLAGS_FIXED|VM_FLAGS_OVERWRITE)
	require.NoError(t, err)

	remapped_address, _, _, err := Mach_vm_remap(Mach_task_self(), address+8192, 8192, 0, VM_FLAGS_FIXED|VM_FLAGS_OVERWRITE, Mach_task_self(), address, false, VM_INHERIT_NONE)
	require.NoError(t, err)
	require.Equal(t, address+8192, remapped_address)

	buffer := unsafe.Slice((*byte)(unsafe.Pointer(address)), 8192)
	mirror := unsafe.Slice((*byte)(unsafe.Pointer(address+8192)), 8192)
	expected := []byte("foobarbaz")
	copy(buffer, expected)
	require.Equal(t, expected, mirror[:len(expected)])

	err = Mach_vm_deallocate(Mach_task_self(), address, 8192*2)
	require.NoError(t, err)
}
