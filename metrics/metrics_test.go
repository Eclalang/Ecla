package metrics

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime/metrics"
	"testing"
	"time"
	"unsafe"
)

// TestMetricsCreation tests the creation of a new Metrics.
func TestMetricsCreation(t *testing.T) {
	m := NewMetrics()
	if m == nil {
		t.Errorf("NewMetrics() returned nil")
	}
}

// TestTimerFunctions tests all timer-related functions.
func TestTimerFunctions(t *testing.T) {
	m := NewMetrics()

	// Start and stop total execution timer
	m.StartTimers()
	time.Sleep(10 * time.Millisecond)
	m.StopTotalTimer()
	if m.TotalExecTime <= 0 {
		t.Errorf("TotalExecTime was not set correctly")
	}

	// Start and stop lexer timer
	m.StartLexerTimer()
	time.Sleep(10 * time.Millisecond)
	m.StopLexerTimer()
	if m.LexerExecTime <= 0 {
		t.Errorf("LexerExecTime was not set correctly")
	}

	// Start and stop parser timer
	m.StartParserTimer()
	time.Sleep(10 * time.Millisecond)
	m.StopParserTimer()
	if m.ParserExecTime <= 0 {
		t.Errorf("ParserExecTime was not set correctly")
	}

	// Start and stop interpreter timer
	m.StartInterpreterTimer()
	time.Sleep(10 * time.Millisecond)
	m.StopInterpreterTimer()
	if m.InterpreterExecTime <= 0 {
		t.Errorf("InterpreterExecTime was not set correctly")
	}
}

// TestMeasure tests the Measure function.
func TestMeasure(t *testing.T) {
	m := NewMetrics()
	// redirecting the standard output to test the printing of metrics

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	m.Measure() // This should cover the printing of different metric types
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	if len(out) == 0 {
		t.Errorf("No metrics were printed")
	}
}

// TestMedianBucket tests the medianBucket function.
func TestMedianBucket(t *testing.T) {
	// Test with a histogram that has an even distribution
	hEven := &metrics.Float64Histogram{
		Counts:  []uint64{2, 2, 2, 2, 2},
		Buckets: []float64{0.1, 0.2, 0.3, 0.4, 0.5},
	}
	medianEven := medianBucket(hEven)
	if medianEven != 0.3 {
		t.Errorf("Expected median of 0.3 for even distribution, got %f", medianEven)
	}

	// Test with a histogram that has an odd distribution
	hOdd := &metrics.Float64Histogram{
		Counts:  []uint64{1, 2, 3, 2, 1},
		Buckets: []float64{0.1, 0.2, 0.3, 0.4, 0.5},
	}
	medianOdd := medianBucket(hOdd)
	if medianOdd != 0.3 {
		t.Errorf("Expected median of 0.3 for odd distribution, got %f", medianOdd)
	}
}

// TestMedianBucketPanic tests the medianBucket function to trigger the panic.
func TestMedianBucketPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// Creating a Float64Histogram with empty or zero counts
	h := &metrics.Float64Histogram{
		Counts:  []uint64{}, // or use []uint64{0, 0, 0} to simulate zero counts
		Buckets: []float64{0.1, 0.2, 0.3},
	}

	// This should trigger the panic
	medianBucket(h)
}

func TestMeasureWithKindBad(t *testing.T) {
	fmt.Println(metrics.Value{})

	m := &Metrics{
		reader: &MockMetricsReader{
			SamplesToReturn: []metrics.Sample{
				{Name: "test.bad", Value: metrics.Value{}},
			},
		},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for metrics.KindBad, but it did not occur")
		}
	}()
	m.Measure()
}

func TestMeasureWithUnknownKind(t *testing.T) {
	mValue := metrics.Value{}

	// Using unsafe to modify the unexported 'kind' field in metrics.Value
	valField := reflect.ValueOf(&mValue).Elem()
	kindField := valField.FieldByName("kind")
	kindField = reflect.NewAt(kindField.Type(), unsafe.Pointer(kindField.UnsafeAddr())).Elem()
	kindField.SetInt(int64(9999))

	m := &Metrics{
		reader: &MockMetricsReader{
			SamplesToReturn: []metrics.Sample{
				{Name: "test.unknown", Value: mValue},
			},
		},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for unknown kind, but it did not occur")
		}
	}()
	m.Measure()
}
