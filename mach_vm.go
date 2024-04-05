// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mach

type vm_map_t = uint32
type vm_prot_t = int32
type vm_inherit_t = uint32

const (
	VM_FLAGS_FIXED              = 0x00000000
	VM_FLAGS_ANYWHERE           = 0x00000001
	VM_FLAGS_PURGABLE           = 0x00000002
	VM_FLAGS_4GB_CHUNK          = 0x00000004
	VM_FLAGS_RANDOM_ADDR        = 0x00000008
	VM_FLAGS_NO_CACHE           = 0x00000010
	VM_FLAGS_RESILIENT_CODESIGN = 0x00000020
	VM_FLAGS_RESILIENT_MEDIA    = 0x00000040
	VM_FLAGS_PERMANENT          = 0x00000080
	VM_FLAGS_TPRO               = 0x00001000
	VM_FLAGS_OVERWRITE          = 0x00004000

	VM_INHERIT_SHARE       = 0
	VM_INHERIT_COPY        = 1
	VM_INHERIT_NONE        = 2
	VM_INHERIT_DONATE_COPY = 3
	VM_INHERIT_DEFAULT     = VM_INHERIT_COPY
	VM_INHERIT_LAST_VALID  = VM_INHERIT_NONE
)

//mach	mach_vm_allocate(target vm_map_t, address *uintptr, size uint, flags int32) error

func Mach_vm_allocate(target vm_map_t, address uintptr, size uint, flags int32) (out_address uintptr, err error) {
	err = mach_vm_allocate(target, &address, size, flags)
	return address, err
}

//mach	Mach_vm_deallocate(target vm_map_t, address uintptr, size uint) error

func Mach_vm_remap(target_task vm_map_t, target_address uintptr, size uint, mask uint, flags int32, src_task vm_map_t, src_address uintptr, copy bool, inheritance vm_inherit_t) (out_target_address uintptr, cur_protection vm_prot_t, max_protection vm_prot_t, err error) {
	err = mach_vm_remap(target_task, &target_address, size, mask, flags, src_task, src_address, copy, &cur_protection, &max_protection, inheritance)
	return target_address, cur_protection, max_protection, err
}

//mach	mach_vm_remap(target_task vm_map_t, target_address *uintptr, size uint, mask uint, flags int32, src_task vm_map_t, src_address uintptr, copy bool, cur_protection *vm_prot_t, max_protection *vm_prot_t, inheritance vm_inherit_t) error
