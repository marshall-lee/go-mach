# Copyright 2024 Vladimir Kochnev. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: all
all:
	go run mkmach.go
	gofmt -w zmach_*.go
	go test
