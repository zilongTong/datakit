// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD:vendor/golang.org/x/sys/plan9/asm_plan9_386.s
#include "textflag.h"

//
// System call support for 386, Plan 9
//
=======
//go:build (darwin || freebsd || netbsd || openbsd) && gc
// +build darwin freebsd netbsd openbsd
// +build gc

#include "textflag.h"

// System call support for 386 BSD
>>>>>>> dev:vendor/golang.org/x/sys/unix/asm_bsd_386.s

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

TEXT	·Syscall(SB),NOSPLIT,$0-32
	JMP	syscall·Syscall(SB)

TEXT	·Syscall6(SB),NOSPLIT,$0-44
	JMP	syscall·Syscall6(SB)

<<<<<<< HEAD:vendor/golang.org/x/sys/plan9/asm_plan9_386.s
TEXT ·RawSyscall(SB),NOSPLIT,$0-28
=======
TEXT	·Syscall9(SB),NOSPLIT,$0-52
	JMP	syscall·Syscall9(SB)

TEXT	·RawSyscall(SB),NOSPLIT,$0-28
>>>>>>> dev:vendor/golang.org/x/sys/unix/asm_bsd_386.s
	JMP	syscall·RawSyscall(SB)

TEXT ·RawSyscall6(SB),NOSPLIT,$0-40
	JMP	syscall·RawSyscall6(SB)

TEXT ·seek(SB),NOSPLIT,$0-36
	JMP	syscall·seek(SB)

TEXT ·exit(SB),NOSPLIT,$4-4
	JMP	syscall·exit(SB)
