// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is generated by mkmach.go

#include "textflag.h"

TEXT	libc_mach_vm_allocate_trampoline<>(SB),NOSPLIT,$0
	MOVD	8(R0), R1	// address
	MOVD	16(R0), R2	// size
	MOVW	24(R0), R3	// flags
	MOVW	0(R0), R0	// target
	BL	libc_mach_vm_allocate(SB)
	RET

GLOBL	·libc_mach_vm_allocate_trampoline_addr(SB), RODATA, $8
DATA	·libc_mach_vm_allocate_trampoline_addr(SB)/8, $libc_mach_vm_allocate_trampoline<>(SB)

TEXT	libc_mach_vm_deallocate_trampoline<>(SB),NOSPLIT,$0
	MOVD	8(R0), R1	// address
	MOVD	16(R0), R2	// size
	MOVW	0(R0), R0	// target
	BL	libc_mach_vm_deallocate(SB)
	RET

GLOBL	·libc_mach_vm_deallocate_trampoline_addr(SB), RODATA, $8
DATA	·libc_mach_vm_deallocate_trampoline_addr(SB)/8, $libc_mach_vm_deallocate_trampoline<>(SB)

TEXT	libc_mach_vm_remap_trampoline<>(SB),NOSPLIT,$0
	SUB	$32, RSP
	MOVD	64(R0), R1
	MOVD	R1, 0(RSP)	// cur_protection
	MOVD	72(R0), R1
	MOVD	R1, 8(RSP)	// max_protection
	MOVD	80(R0), R1
	MOVD	R1, 16(RSP)	// inheritance
	MOVD	8(R0), R1	// target_address
	MOVD	16(R0), R2	// size
	MOVD	24(R0), R3	// mask
	MOVW	32(R0), R4	// flags
	MOVW	40(R0), R5	// src_task
	MOVD	48(R0), R6	// src_address
	MOVD	56(R0), R7	// copy
	MOVW	0(R0), R0	// target_task
	BL	libc_mach_vm_remap(SB)
	ADD	$32, RSP
	RET

GLOBL	·libc_mach_vm_remap_trampoline_addr(SB), RODATA, $8
DATA	·libc_mach_vm_remap_trampoline_addr(SB)/8, $libc_mach_vm_remap_trampoline<>(SB)