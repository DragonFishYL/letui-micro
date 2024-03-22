// Package memory provides an in memory log buffer
package memory

import (
	"fmt"
	"letui-micro/debug/logger"

	"letui-micro/util/ring"
)

var (
	// DefaultSize of the logger buffer
	DefaultSize = 1024
)

// memoryLog is default micro log
type memoryLog struct {
	*ring.Buffer
}

// NewLog returns default Logger with
func NewLog(opts ...logger.Option) logger.Log {
	// get default options
	options := logger.DefaultOptions()

	// apply requested options
	for _, o := range opts {
		o(&options)
	}

	return &memoryLog{
		Buffer: ring.New(options.Size),
	}
}

// Write writes logs into logger
func (l *memoryLog) Write(r logger.Record) error {
	l.Buffer.Put(fmt.Sprint(r.Message))
	return nil
}

// Read reads logs and returns them
func (l *memoryLog) Read(opts ...logger.ReadOption) ([]logger.Record, error) {
	options := logger.ReadOptions{}
	// initialize the read options
	for _, o := range opts {
		o(&options)
	}

	var entries []*ring.Entry
	// if Since options ha sbeen specified we honor it
	if !options.Since.IsZero() {
		entries = l.Buffer.Since(options.Since)
	}

	// only if we specified valid count constraint
	// do we end up doing some serious if-else kung-fu
	// if since constraint has been provided
	// we return *count* number of logs since the given timestamp;
	// otherwise we return last count number of logs
	if options.Count > 0 {
		switch len(entries) > 0 {
		case true:
			// if we request fewer logs than what since constraint gives us
			if options.Count < len(entries) {
				entries = entries[0:options.Count]
			}
		default:
			entries = l.Buffer.Get(options.Count)
		}
	}

	records := make([]logger.Record, 0, len(entries))
	for _, entry := range entries {
		record := logger.Record{
			Timestamp: entry.Timestamp,
			Message:   entry.Value,
		}
		records = append(records, record)
	}

	return records, nil
}

// Stream returns channel for reading log records
// along with a stop channel, close it when done
func (l *memoryLog) Stream() (logger.Stream, error) {
	// get stream channel from ring buffer
	stream, stop := l.Buffer.Stream()
	// make a buffered channel
	records := make(chan logger.Record, 128)
	// get last 10 records
	last10 := l.Buffer.Get(10)

	// stream the log records
	go func() {
		// first send last 10 records
		for _, entry := range last10 {
			records <- logger.Record{
				Timestamp: entry.Timestamp,
				Message:   entry.Value,
				Metadata:  make(map[string]string),
			}
		}
		// now stream continuously
		for entry := range stream {
			records <- logger.Record{
				Timestamp: entry.Timestamp,
				Message:   entry.Value,
				Metadata:  make(map[string]string),
			}
		}
	}()

	return &logStream{
		stream: records,
		stop:   stop,
	}, nil
}
