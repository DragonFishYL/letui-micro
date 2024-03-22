package noop

import "letui-micro/debug/logger"

type noop struct{}

func (n *noop) Read(...logger.ReadOption) ([]logger.Record, error) {
	return nil, nil
}

func (n *noop) Write(logger.Record) error {
	return nil
}

func (n *noop) Stream() (logger.Stream, error) {
	return nil, nil
}

func NewLog(opts ...logger.Option) logger.Log {
	return new(noop)
}
