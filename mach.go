// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mach

import "unsafe"

//go:linkname libcCall runtime.libcCall

func libcCall(fn, arg unsafe.Pointer) int32

//go:cgo_import_dynamic libc_mach_task_self_ mach_task_self_ "/usr/lib/libSystem.B.dylib"

func Mach_task_self() mach_port_t

func makeCString(str string) *byte {
	bytes := make([]byte, len(str)+1)
	for i, c := range []byte(str) {
		bytes[i] = c
	}
	return unsafe.SliceData(bytes)
}

func makeStringFromCString(from *byte) string {
	var n int
	ptr := unsafe.Pointer(from)
	for n = 0; *(*byte)(ptr) != 0; n++ {
		ptr = unsafe.Add(ptr, 1)
	}
	bytes := make([]byte, n)
	copy(bytes, unsafe.Slice(from, n))
	return string(bytes)
}

func boolUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
