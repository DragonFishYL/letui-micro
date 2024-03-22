package memory

import "letui-micro/debug/logger"

type logStream struct {
	stream <-chan logger.Record
	stop   chan bool
}

func (l *logStream) Chan() <-chan logger.Record {
	return l.stream
}

func (l *logStream) Stop() error {
	select {
	case <-l.stop:
		return nil
	default:
		close(l.stop)
	}
	return nil
}
