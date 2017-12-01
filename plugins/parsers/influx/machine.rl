package influx

import (
	"fmt"
)

var (
	ErrFieldParse = fmt.Errorf("expected field")
	ErrTagParse = fmt.Errorf("expected tag")
	ErrTimestampParse = fmt.Errorf("expected timestamp")
	ErrParse = fmt.Errorf("parse error")
)

%%{
machine LineProtocol;

action begin {
	m.pb = m.p
}

action yield {
	yield = true
	fbreak;
}

action field_error {
	m.err = ErrFieldParse
	fhold;
	fnext nextline;
	fbreak;
}

action tagset_error {
	m.err = ErrTagParse
	fhold;
	fnext nextline;
	fbreak;
}

action timestamp_error {
	m.err = ErrTimestampParse
	fhold;
	fnext nextline;
	fbreak;
}

action parse_error {
	m.err = ErrParse
	fhold;
	fnext nextline;
	fbreak;
}

action recover {
	fgoto main;
}

action name {
	m.handler.SetMeasurement(m.text())
}

action tagkey {
	key = m.text()
}

action tagvalue {
	m.handler.AddTag(key, m.text())
}

action fieldkey {
	key = m.text()
}

action integer {
	m.handler.AddInt(key, m.text())
}

action float {
	m.handler.AddFloat(key, m.text())
}

action bool {
	m.handler.AddBool(key, m.text())
}

action string {
	m.handler.AddString(key, m.text())
}

action timestamp {
	m.handler.SetTimestamp(m.text())
}

ws =
	[\t\v\f ];

non_zero_digit =
	[1-9];

integer =
	'-'? ( digit | ( non_zero_digit digit* ) );

number =
	( integer ( '.' digit* )? ) | ( '.' digit* );

timestamp =
	('-'? digit{1,19}) >begin %timestamp $err(timestamp_error);

fieldkeychar =
	[^\t\n\f\r ,=\\] | ( '\\' [^\t\n\f\r] );

fieldkey =
	fieldkeychar+ >begin %fieldkey;

# TODO: scientific notation
fieldfloat =
	number >begin %float;

fieldinteger =
	(integer 'i') >begin %integer;

false =
	"false" | "FALSE" | "False" | "F" | "f";

true =
	"true" | "TRUE" | "True" | "T" | "t";

fieldbool =
	(true | false) >begin %bool;

fieldstringchar =
	[^\\"] | '\\' [\\"];

fieldstring =
	fieldstringchar* >begin %string;

fieldstringquoted =
	'"' fieldstring '"';

fieldvalue = fieldinteger | fieldfloat | fieldstringquoted | fieldbool;

field =
	fieldkey '=' fieldvalue $err(field_error);

fieldset =
	field ( ',' field )*;

tagchar =
	[^\t\n\f\r ,=\\] | ( '\\' [^\t\n\f\r] );

tagkey =
	tagchar+ >begin %tagkey;

tagvalue =
	tagchar+ >begin %tagvalue;

tagset =
	(',' (tagkey '=' tagvalue) $err(tagset_error))*;

measurement_chars =
	[^\t\n\f\r ,\\] | ( '\\' [^\t\n\f\r\\] );

measurement =
	measurement_chars+ >begin %name %eof(name);

newline =
	[\r\n]+;

line =
	ws* measurement tagset (ws+ fieldset) $err(field_error) (ws+ timestamp)? ws* %yield;

nextline :=
	[^\r\n]* [\r\n]+ @recover;

main :=
	(line ( newline line )* newline?)?;

write data;
}%%

type machine struct {
	data       []byte
	cs         int
	p, pe, eof int
	pb         int
	handler    Handler
	err        error

	count int
}

func NewMachine(handler Handler) *machine {
	m := &machine{
		handler: handler,
	}

	%% access m.;
	%% variable p m.p;
	%% variable pe m.pe;
	%% variable eof m.eof;
	%% variable data m.data;
	%% write init;

	return m
}

func (m *machine) SetData(data []byte) {
	m.data = data
	m.p = 0
	m.pb = 0
	m.pe = len(data)
	m.eof = len(data)
	m.cs = 0
	m.err = nil

	%% write init;
}

// ParseLine parses a line of input and returns true if more data can be
// parsed.
func (m *machine) ParseLine() bool {
	if m.data == nil || m.p >= m.pe {
		return false
	}

	m.err = nil
	var key []byte
	var yield bool

	%% write exec;

	// Even though there was an error, return true. On the next call to this
	// function we will attempt to scan to the next line of input and recover.
	if m.err != nil {
		return true
	}

	// Don't check the error state in the case that we just yielded, because
	// the yield indicates we just completed parsing a line.
	if !yield && m.cs == LineProtocol_error {
		m.err = ErrParse
		return true
	}

	return true
}

// Err returns the error that occurred on the last call to ParseLine.  If the
// result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
	return m.err
}

// Position returns the current position into the input.
func (m *machine) Position() int {
	return m.p
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}
