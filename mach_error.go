// Copyright 2024 Vladimir Kochnev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mach

type kern_return_err kern_return_t

func (err kern_return_err) Error() string {
	return Mach_error_string(mach_error_t(err))
}

//machnb	Mach_error_string(error_value mach_error_t) string
//machnb	Mach_error(str string, error_value mach_error_t)
