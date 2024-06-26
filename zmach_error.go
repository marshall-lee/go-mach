// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is generated by mkmach.go

package mach

import "unsafe"

var libc_mach_error_string_trampoline_addr uintptr

//go:cgo_import_dynamic libc_mach_error_string mach_error_string "/usr/lib/libSystem.B.dylib"

func Mach_error_string(error_value mach_error_t) string {
	args := struct {
		error_value uintptr
		ret         uintptr
	}{
		error_value: uintptr(error_value),
	}
	libcCall(unsafe.Pointer(libc_mach_error_string_trampoline_addr), unsafe.Pointer(&args))
	return makeStringFromCString((*byte)(unsafe.Pointer(args.ret)))
}

var libc_mach_error_trampoline_addr uintptr

//go:cgo_import_dynamic libc_mach_error mach_error "/usr/lib/libSystem.B.dylib"

func Mach_error(str string, error_value mach_error_t) {
	args := struct {
		str         uintptr
		error_value uintptr
	}{
		str:         uintptr(unsafe.Pointer(makeCString(str))),
		error_value: uintptr(error_value),
	}
	libcCall(unsafe.Pointer(libc_mach_error_trampoline_addr), unsafe.Pointer(&args))
}
