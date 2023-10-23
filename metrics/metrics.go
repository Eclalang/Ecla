package metrics

import (
	"fmt"
	"runtime/metrics"
	"time"
)

type Metrics struct {
	// metrics variables
	StartExecTime, StartLexerTime, StartParserTime, StartInterpreterTime time.Time
	TotalExecTime, LexerExecTime, ParserExecTime, InterpreterExecTime    time.Duration
	FCalls, StackSize, HeapSize, MinMem, MaxMem, MinCpu, MaxCpu          int
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) StartTimers() {
	m.StartExecTime = time.Now()
}

func (m *Metrics) StopTotalTimer() {
	m.TotalExecTime = time.Since(m.StartExecTime)
	m.Measure()
}

func (m *Metrics) StartLexerTimer() {
	m.StartLexerTime = time.Now()
}

func (m *Metrics) StopLexerTimer() {
	m.LexerExecTime = time.Since(m.StartLexerTime)
}

func (m *Metrics) StartParserTimer() {
	m.StartParserTime = time.Now()
}

func (m *Metrics) StopParserTimer() {
	m.ParserExecTime = time.Since(m.StartParserTime)
}

func (m *Metrics) StartInterpreterTimer() {
	m.StartInterpreterTime = time.Now()
}

func (m *Metrics) StopInterpreterTimer() {
	m.InterpreterExecTime = time.Since(m.StartInterpreterTime)
}

func (m *Metrics) Measure() {
	// Get descriptions for all supported metrics.
	descs := metrics.All()

	// Create a sample for each metric.
	samples := make([]metrics.Sample, len(descs))
	for i := range samples {
		samples[i].Name = descs[i].Name
	}

	// Sample the metrics. Re-use the samples slice if you can!
	metrics.Read(samples)

	// Iterate over all results.
	for _, sample := range samples {
		// Pull out the name and value.
		name, value := sample.Name, sample.Value

		// Handle each sample.
		switch value.Kind() {
		case metrics.KindUint64:
			fmt.Printf("%s: %d\n", name, value.Uint64())
		case metrics.KindFloat64:
			fmt.Printf("%s: %f\n", name, value.Float64())
		case metrics.KindFloat64Histogram:
			// The histogram may be quite large, so let's just pull out
			// a crude estimate for the median for the sake of this example.
			fmt.Printf("%s: %f\n", name, medianBucket(value.Float64Histogram()))
		case metrics.KindBad:
			// This should never happen because all metrics are supported
			// by construction.
			panic("bug in runtime/metrics package!")
		default:
			// This may happen as new metrics get added.
			//
			// The safest thing to do here is to simply log it somewhere
			// as something to look into, but ignore it for now.
			// In the worst case, you might temporarily miss out on a new metric.
			fmt.Printf("%s: unexpected metric Kind: %v\n", name, value.Kind())
		}
	}
}

func medianBucket(h *metrics.Float64Histogram) float64 {
	total := uint64(0)
	for _, count := range h.Counts {
		total += count
	}
	thresh := total / 2
	total = 0
	for i, count := range h.Counts {
		total += count
		if total >= thresh {
			return h.Buckets[i]
		}
	}
	panic("should not happen")
}
