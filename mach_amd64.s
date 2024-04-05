// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

TEXT ·Mach_task_self(SB),0,$0-4
	MOVQ	$libc_mach_task_self_(SB), AX
	MOVQ	(AX), AX
	MOVL	AX, ret+0(FP)
	RET
