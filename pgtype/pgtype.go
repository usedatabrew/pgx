package pgtype

import (
	"errors"
	"io"
)

// PostgreSQL oids for common types
const (
	BoolOID             = 16
	ByteaOID            = 17
	CharOID             = 18
	NameOID             = 19
	Int8OID             = 20
	Int2OID             = 21
	Int4OID             = 23
	TextOID             = 25
	OIDOID              = 26
	TidOID              = 27
	XIDOID              = 28
	CIDOID              = 29
	JSONOID             = 114
	CidrOID             = 650
	CidrArrayOID        = 651
	Float4OID           = 700
	Float8OID           = 701
	UnknownOID          = 705
	InetOID             = 869
	BoolArrayOID        = 1000
	Int2ArrayOID        = 1005
	Int4ArrayOID        = 1007
	TextArrayOID        = 1009
	ByteaArrayOID       = 1001
	VarcharArrayOID     = 1015
	Int8ArrayOID        = 1016
	Float4ArrayOID      = 1021
	Float8ArrayOID      = 1022
	AclItemOID          = 1033
	AclItemArrayOID     = 1034
	InetArrayOID        = 1041
	VarcharOID          = 1043
	DateOID             = 1082
	TimestampOID        = 1114
	TimestampArrayOID   = 1115
	DateArrayOID        = 1182
	TimestamptzOID      = 1184
	TimestamptzArrayOID = 1185
	RecordOID           = 2249
	UUIDOID             = 2950
	JSONBOID            = 3802
)

type Status byte

const (
	Undefined Status = iota
	Null
	Present
)

type InfinityModifier int8

const (
	Infinity         InfinityModifier = 1
	None             InfinityModifier = 0
	NegativeInfinity InfinityModifier = -Infinity
)

type Value interface {
	ConvertFrom(src interface{}) error
	AssignTo(dst interface{}) error
}

type BinaryDecoder interface {
	DecodeBinary(src []byte) error
}

type TextDecoder interface {
	DecodeText(src []byte) error
}

// BinaryEncoder is implemented by types that can encode themselves into the
// PostgreSQL binary wire format.
type BinaryEncoder interface {
	// EncodeBinary should encode the binary format of self to w. If self is the
	// SQL value NULL then write nothing and return (true, nil). The caller of
	// EncodeBinary is responsible for writing the correct NULL value or the
	// length of the data written.
	EncodeBinary(w io.Writer) (null bool, err error)
}

// TextEncoder is implemented by types that can encode themselves into the
// PostgreSQL text wire format.
type TextEncoder interface {
	// EncodeText should encode the text format of self to w. If self is the SQL
	// value NULL then write nothing and return (true, nil). The caller of
	// EncodeText is responsible for writing the correct NULL value or the length
	// of the data written.
	EncodeText(w io.Writer) (null bool, err error)
}

var errUndefined = errors.New("cannot encode status undefined")