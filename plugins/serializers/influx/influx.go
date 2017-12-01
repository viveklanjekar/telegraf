package influx

import (
	"bytes"
	"errors"
	"io"
	"strconv"

	"github.com/influxdata/telegraf"
)

const MaxInt = int(^uint(0) >> 1)

var (
	ErrNeedMoreSpace = errors.New("need more space")
	ErrInvalidName   = errors.New("invalid name")
	ErrInvalidField  = errors.New("invalid field")
)

// Serializer is a serializer for line protocol.
type Serializer struct {
	maxLineBytes int
	bytesWritten int

	buf    bytes.Buffer
	header []byte
	footer []byte
	pair   []byte
}

func NewSerializer(maxLineBytes int) *Serializer {
	serializer := &Serializer{
		maxLineBytes: maxLineBytes,
		header:       make([]byte, 0, 50),
		footer:       make([]byte, 0, 21),
		pair:         make([]byte, 0, 50),
	}
	return serializer
}

// Serialize writes the telegraf.Metric to a byte slice.  May produce multiple
// lines of output if longer than maximum line length.  Lines are terminated
// with a newline (LF) char.
//
// Always returns the same buffer!?  Do I need to fix this?
func (s *Serializer) Serialize(m telegraf.Metric) ([]byte, error) {
	s.buf.Reset()
	err := s.writeMetric(&s.buf, m)
	if err != nil {
		return nil, err
	}
	return s.buf.Bytes(), nil
}

func (s *Serializer) Write(w io.Writer, m telegraf.Metric) (int, error) {
	return s.bytesWritten, s.writeMetric(w, m)
}

func (s *Serializer) writeString(w io.Writer, str string) error {
	n, err := io.WriteString(w, str)
	s.bytesWritten += n
	return err
}

func (s *Serializer) write(w io.Writer, b []byte) error {
	n, err := w.Write(b)
	s.bytesWritten += n
	return err
}

func (s *Serializer) buildHeader(m telegraf.Metric) error {
	s.header = s.header[:0]

	name := nameEscape(m.Name())
	if name == "" {
		return ErrInvalidName
	}

	s.header = append(s.header, name...)

	for _, tag := range m.TagList() {
		key := escape(tag.Key)
		value := escape(tag.Value)

		// Some keys and values are not encodeable as line protocol, such as
		// those with a trailing '\' or empty strings.
		if key == "" || value == "" {
			continue
		}

		s.header = append(s.header, ',')
		s.header = append(s.header, key...)
		s.header = append(s.header, '=')
		s.header = append(s.header, value...)
	}

	s.header = append(s.header, ' ')
	return nil
}

func (s *Serializer) buildFooter(m telegraf.Metric) {
	s.footer = s.footer[:0]
	s.footer = append(s.footer, ' ')
	s.footer = strconv.AppendInt(s.footer, m.Time().UnixNano(), 10)
	s.footer = append(s.footer, '\n')
}

func (s *Serializer) buildFieldPair(key string, value interface{}) error {
	s.pair = s.pair[:0]
	key = escape(key)

	// Some keys are not encodeable as line protocol, such as those with a
	// trailing '\' or empty strings.
	if key == "" {
		return ErrInvalidField
	}

	s.pair = append(s.pair, key...)
	s.pair = append(s.pair, '=')
	s.pair = appendFieldValue(s.pair, value)
	return nil
}

func (s *Serializer) writeMetric(w io.Writer, m telegraf.Metric) error {
	var err error

	err = s.buildHeader(m)
	if err != nil {
		return err
	}

	err = s.write(w, s.header)
	if err != nil {
		return err
	}

	s.buildFooter(m)

	first := true
	for _, field := range m.FieldList() {
		err = s.buildFieldPair(field.Key, field.Value)
		if err != nil {
			continue
		}

		bytesNeeded := len(s.pair) + len(s.footer) + len(s.header)
		if !first {
			bytesNeeded += 1
		}

		if s.maxLineBytes > 0 && bytesNeeded > s.maxLineBytes {
			// Need at least one field per line
			if first {
				return ErrNeedMoreSpace
			}

			err = s.write(w, s.footer)
			if err != nil {
				return err
			}

			err = s.write(w, s.header)
			if err != nil {
				return err
			}

			first = true
		}

		if !first {
			err = s.writeString(w, ",")
			if err != nil {
				return err
			}
		}
		first = false
		s.write(w, s.pair)
	}

	return s.write(w, s.footer)
}

func appendFieldValue(buf []byte, value interface{}) []byte {
	switch v := value.(type) {
	case int64:
		return appendIntField(buf, v)
	case float64:
		return appendFloatField(buf, v)
	case uint64:
		return appendUintField(buf, v)
	case string:
		return appendStringField(buf, v)
	case bool:
		return appendBoolField(buf, v)
	}
	return buf
}

func appendIntField(buf []byte, value int64) []byte {
	return append(strconv.AppendInt(buf, value, 10), 'i')
}

func appendUintField(buf []byte, value uint64) []byte {
	if value > uint64(MaxInt) {
		value = uint64(MaxInt)
	}
	// No support for unsigned int yet
	return append(strconv.AppendUint(buf, value, 10), 'i')
}

func appendFloatField(buf []byte, value float64) []byte {
	return strconv.AppendFloat(buf, value, 'g', -1, 64)
}

func appendBoolField(buf []byte, value bool) []byte {
	return strconv.AppendBool(buf, value)
}

func appendStringField(buf []byte, value string) []byte {
	buf = append(buf, '"')
	buf = append(buf, stringFieldEscape(value)...)
	buf = append(buf, '"')
	return buf
}
