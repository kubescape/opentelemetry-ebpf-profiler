package reporter

import "go.opentelemetry.io/ebpf-profiler/libpf"

type ChannelReporter struct {
	name   string
	traces chan *libpf.Trace
}

func NewChannelReporter(name string, traces chan *libpf.Trace) *ChannelReporter {
	return &ChannelReporter{
		name:   name,
		traces: traces,
	}
}

func (r *ChannelReporter) ReportFramesForTrace(_ *libpf.Trace) {}

func (r *ChannelReporter) ReportCountForTrace(_ *libpf.Trace) {}

func (r *ChannelReporter) ReportTraceEvent(trace *libpf.Trace, meta *TraceEventMeta) {}

func (r *ChannelReporter) SupportsReportTraceEvent() bool {
	return true
}
