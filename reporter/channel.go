package reporter

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/ebpf-profiler/libpf"
)

type Symbol struct {
	lineNumber     libpf.SourceLineno
	functionOffset uint32
	functionName   string
	filePath       string
}

type SymbolizedTrace struct {
	libpf.Trace
	symbolMap map[libpf.AddressOrLineno]*Symbol
}

type CompleteTrace struct {
	Trace *SymbolizedTrace
	Meta  *TraceEventMeta
}

type ChannelReporter struct {
	name      string
	traces    chan *CompleteTrace
	symbolMap map[libpf.AddressOrLineno]*Symbol
}

func NewChannelReporter(name string, traces chan *CompleteTrace) *ChannelReporter {
	return &ChannelReporter{
		name:      name,
		traces:    traces,
		symbolMap: make(map[libpf.AddressOrLineno]*Symbol),
	}
}

var _ Reporter = (*ChannelReporter)(nil)

func (r *ChannelReporter) ReportTraceEvent(trace *libpf.Trace, meta *TraceEventMeta) {
	symbolTrace := &SymbolizedTrace{}
	symbolTrace.Trace = *trace
	symbolTrace.symbolMap = nil
	for i := range trace.Linenos {
		if symbol, exists := r.symbolMap[trace.Linenos[i]]; exists {
			symbolTrace.symbolMap = make(map[libpf.AddressOrLineno]*Symbol)
			symbolTrace.symbolMap[trace.Linenos[i]] = symbol
			delete(r.symbolMap, trace.Linenos[i])
		}
	}

	new_trace := &CompleteTrace{
		Trace: symbolTrace,
		Meta:  meta,
	}

	r.traces <- new_trace
}

func (r *ChannelReporter) FrameMetadata(fileID libpf.FileID, addressOrLine libpf.AddressOrLineno, lineNumber libpf.SourceLineno, functionOffset uint32, functionName, filePath string) {
	r.symbolMap[addressOrLine] = &Symbol{
		lineNumber:     lineNumber,
		functionOffset: functionOffset,
		functionName:   functionName,
		filePath:       filePath,
	}
}

func (r *ChannelReporter) SpendChannel() {
	for {
		new_trace := <-r.traces
		if new_trace.Trace.symbolMap != nil {
			fmt.Println("Comm: ", new_trace.Meta.Comm, "Pid:", new_trace.Meta.PID, "Tid:", new_trace.Meta.TID, "Timestamp:", new_trace.Meta.Timestamp)
			for addressOrLine, symbol := range new_trace.Trace.symbolMap {
				fmt.Println("AddressOrLine: ", addressOrLine, "Symbol: ", symbol.functionName, ":", symbol.lineNumber, ":", symbol.filePath)
			}
			fmt.Println('-' * 50)
		}
	}
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

func (r *ChannelReporter) ExecutableMetadata(args *ExecutableMetadataArgs) {}

func (r *ChannelReporter) ReportHostMetadata(metadataMap map[string]string) {}
func (r *ChannelReporter) ReportHostMetadataBlocking(ctx context.Context, metadataMap map[string]string, maxRetries int, waitRetry time.Duration) error {
	return nil
}

func (r *ChannelReporter) ReportMetrics(timestamp uint32, ids []uint32, values []int64) {}
