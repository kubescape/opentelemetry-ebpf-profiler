//go:build amd64 && !dummy && external_trigger
// +build amd64,!dummy,external_trigger

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package support // import "go.opentelemetry.io/ebpf-profiler/support"

import (
	_ "embed"
)

//go:embed ebpf/tracer.ebpf.release.external.amd64
var tracerData []byte
