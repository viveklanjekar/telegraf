package influx

import (
	"fmt"

	"github.com/influxdata/telegraf"
)

type Handler interface {
	SetMeasurement(name []byte)
	AddTag(key []byte, value []byte)
	AddInt(key []byte, value []byte)
	AddFloat(key []byte, value []byte)
	AddString(key []byte, value []byte)
	AddBool(key []byte, value []byte)
	SetTimestamp(tm []byte)
	Reset()
}

type ParseError struct {
	Offset int
	msg    string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s at position %d", e.msg, e.Offset)
}

type Parser struct {
	*machine
	handler *MetricHandler
}

func NewParser(handler *MetricHandler) *Parser {
	return &Parser{
		machine: NewMachine(handler),
		handler: handler,
	}
}

func (p *Parser) Parse(input []byte) ([]telegraf.Metric, error) {
	metrics := make([]telegraf.Metric, 0)
	p.machine.SetData(input)

	for p.machine.ParseLine() {
		err := p.machine.Err()
		if err != nil {
			return nil, &ParseError{
				Offset: p.machine.Position(),
				msg:    err.Error(),
			}
		}

		metric, err := p.handler.Metric()
		if err != nil {
			return nil, err
		}
		p.handler.Reset()
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

func (p *Parser) ParseLine(line string) (telegraf.Metric, error) {
	panic("not implemented")
}

func (p *Parser) SetDefaultTags(tags map[string]string) {
	panic("not implemented")
}
