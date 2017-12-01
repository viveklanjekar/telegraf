package telegraf

import (
	"time"
)

// ValueType is an enumeration of metric types that represent a simple value.
type ValueType int

// Possible values for the ValueType enum.
const (
	_ ValueType = iota
	Counter
	Gauge
	Untyped
	Summary
	Histogram
)

// todo: aggregation hints
type AggregationType int

const (
	_ AggregationType = iota
	Min
	Max
	Sum
	Mean
)

type Tag struct {
	Key   string
	Value string
}

func (t *Tag) String() string {
	return t.Key + "=" + t.Value
}

type Field struct {
	Key   string
	Value interface{}
}

type Hint struct {
	Key       string
	ValueType ValueType
}

type Metric interface {
	// Getting data structure functions
	Name() string
	Tags() map[string]string
	TagList() []*Tag
	Fields() map[string]interface{}
	FieldList() []*Field
	Time() time.Time

	// Name functions
	SetName(name string)
	AddPrefix(prefix string)
	AddSuffix(suffix string)

	// Tag functions
	HasTag(key string) bool
	AddTag(key, value string)
	RemoveTag(key string)

	// Field functions
	HasField(key string) bool
	AddField(key string, value interface{})
	RemoveField(key string)

	// HashID returns an unique identifier for the series.
	HashID() uint64

	// Copy returns a deep copy of the Metric.
	Copy() Metric
}
