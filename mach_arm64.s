// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

TEXT ·Mach_task_self(SB),0,$0-4
	MOVD	$libc_mach_task_self_(SB), R0
	MOVD	(R0), R0
	MOVW	R0, ret+0(FP)
	RET
