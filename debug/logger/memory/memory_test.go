package memory

import (
	"letui-micro/debug/logger"
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {
	// set size to some value
	size := 100
	// override the global logger
	lg := NewLog(logger.Size(size))
	// make sure we have the right size of the logger ring buffer
	if lg.(*memoryLog).Size() != size {
		t.Errorf("expected buffer size: %d, got: %d", size, lg.(*memoryLog).Size())
	}

	// Log some cruft
	lg.Write(logger.Record{Message: "foobar"})
	lg.Write(logger.Record{Message: "foo bar"})

	// Check if the logs are stored in the logger ring buffer
	expected := []string{"foobar", "foo bar"}
	entries, _ := lg.Read(logger.Count(len(expected)))
	for i, entry := range entries {
		if !reflect.DeepEqual(entry.Message, expected[i]) {
			t.Errorf("expected %s, got %s", expected[i], entry.Message)
		}
	}
}
