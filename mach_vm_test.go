// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mach

import (
	"os"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestMachVMMirroredBuffer(t *testing.T) {
	bufferSize := os.Getpagesize()
	address, err := Mach_vm_allocate(Mach_task_self(), 0, uint(bufferSize*2), VM_FLAGS_ANYWHERE)
	require.NoError(t, err)

	_, err = Mach_vm_allocate(Mach_task_self(), address, uint(bufferSize), VM_FLAGS_FIXED|VM_FLAGS_OVERWRITE)
	require.NoError(t, err)

	remapped_address, _, _, err := Mach_vm_remap(Mach_task_self(), address+uintptr(bufferSize), uint(bufferSize), 0, VM_FLAGS_FIXED|VM_FLAGS_OVERWRITE, Mach_task_self(), address, false, VM_INHERIT_NONE)
	require.NoError(t, err)
	require.Equal(t, address+uintptr(bufferSize), remapped_address)

	buffer := unsafe.Slice((*byte)(unsafe.Pointer(address)), bufferSize)
	mirror := unsafe.Slice((*byte)(unsafe.Pointer(address+uintptr(bufferSize))), bufferSize)
	expected := []byte("foobarbaz")
	copy(buffer, expected)
	require.Equal(t, expected, mirror[:len(expected)])

	err = Mach_vm_deallocate(Mach_task_self(), address, uint(bufferSize*2))
	require.NoError(t, err)
}
