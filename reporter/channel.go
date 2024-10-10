package reporter

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/ebpf-profiler/libpf"
)

type CompleteTrace struct {
	Trace *libpf.Trace
	Meta  *TraceEventMeta
}

type ChannelReporter struct {
	name   string
	traces chan *CompleteTrace
}

func NewChannelReporter(name string, traces chan *CompleteTrace) *ChannelReporter {
	return &ChannelReporter{
		name:   name,
		traces: traces,
	}
}

var _ Reporter = (*ChannelReporter)(nil)

func (r *ChannelReporter) ReportTraceEvent(trace *libpf.Trace, meta *TraceEventMeta) {
	fmt.Println("ChannelReporter: ReportTraceEvent")
	new_trace := &CompleteTrace{
		Trace: trace,
		Meta:  meta,
	}

	r.traces <- new_trace
}

func (r *ChannelReporter) SupportsReportTraceEvent() bool {
	return true
}

func (r *ChannelReporter) ReportFramesForTrace(trace *libpf.Trace) {}

func (r *ChannelReporter) ReportCountForTrace(traceHash libpf.TraceHash, count uint16, meta *TraceEventMeta) {
}

func (r *ChannelReporter) Stop() {}

func (r *ChannelReporter) GetMetrics() Metrics {
	return Metrics{}
}

func (r *ChannelReporter) ReportFallbackSymbol(frameID libpf.FrameID, symbol string) {}
func (r *ChannelReporter) ExecutableMetadata(args *ExecutableMetadataArgs)           {}
func (r *ChannelReporter) FrameMetadata(fileID libpf.FileID, addressOrLine libpf.AddressOrLineno, lineNumber libpf.SourceLineno, functionOffset uint32, functionName, filePath string) {
}
func (r *ChannelReporter) ReportHostMetadata(metadataMap map[string]string) {}
func (r *ChannelReporter) ReportHostMetadataBlocking(ctx context.Context, metadataMap map[string]string, maxRetries int, waitRetry time.Duration) error {
}
func (r *ChannelReporter) ReportMetrics(timestamp uint32, ids []uint32, values []int64) {}
