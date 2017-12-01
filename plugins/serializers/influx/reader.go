package influx

import (
	"errors"
	"io"

	"github.com/influxdata/telegraf"
)

type buffer struct {
	buf  []byte
	rpos int
	wpos int
}

func (b *buffer) Write(p []byte) (int, error) {
	n := copy(b.buf[b.wpos:], p)
	b.wpos += n
	if n != len(p) {
		return n, errors.New("buffer full")
	}
	return n, nil
}

func (b *buffer) WriteString(s string) (int, error) {
	n := copy(b.buf[b.wpos:], s)
	b.wpos += n
	if n != len(s) {
		return n, errors.New("buffer full")
	}
	return n, nil
}

func (b *buffer) Reset() {
	b.rpos = 0
	b.wpos = 0
}

// reader is an io.Reader for line protocol.
type reader struct {
	serializer *Serializer
	metrics    []telegraf.Metric
	offset     int
	buffer     *buffer
}

func NewReader(metrics []telegraf.Metric, maxLine int) io.Reader {
	return &reader{
		serializer: &Serializer{},
		metrics:    metrics,
		buffer:     &buffer{buf: make([]byte, 0, maxLine)},
	}
}

// Read up to len(p) bytes into p.
//
// May read partial metrics.
func (r *reader) Read(p []byte) (int, error) {
	unread := r.metrics[r.offset:]

	var total int
	for _, metric := range unread {
		n, err := r.serializer.Write(r.buffer, metric)
		total += n
		if err != nil {
			return total, err
		}

		read := copy(p, r.buffer.buf[:r.buffer.wpos])
		if read != r.buffer.wpos {
			r.buffer.rpos = read
		}
	}
	return total, nil
}
