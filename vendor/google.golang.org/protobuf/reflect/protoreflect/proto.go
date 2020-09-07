// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package protoreflect provides interfaces to dynamically manipulate messages.
//
// This package includes type descriptors which describe the structure of types
// defined in proto source files and value interfaces which provide the
// ability to examine and manipulate the contents of messages.
//
//
// Protocol Buffer Descriptors
//
// Protobuf descriptors (e.g., EnumDescriptor or MessageDescriptor)
// are immutable objects that represent protobuf type information.
// They are wrappers around the messages declared in descriptor.proto.
// Protobuf descriptors alone lack any information regarding Go types.
//
// Enums and messages generated by this module implement Enum and ProtoMessage,
// where the Descriptor and ProtoReflect.Descriptor accessors respectively
// return the protobuf descriptor for the values.
//
// The protobuf descriptor interfaces are not meant to be implemented by
// user code since they might need to be extended in the future to support
// additions to the protobuf language.
// The "google.golang.org/protobuf/reflect/protodesc" package converts between
// google.protobuf.DescriptorProto messages and protobuf descriptors.
//
//
// Go Type Descriptors
//
// A type descriptor (e.g., EnumType or MessageType) is a constructor for
// a concrete Go type that represents the associated protobuf descriptor.
// There is commonly a one-to-one relationship between protobuf descriptors and
// Go type descriptors, but it can potentially be a one-to-many relationship.
//
// Enums and messages generated by this module implement Enum and ProtoMessage,
// where the Type and ProtoReflect.Type accessors respectively
// return the protobuf descriptor for the values.
//
// The "google.golang.org/protobuf/types/dynamicpb" package can be used to
// create Go type descriptors from protobuf descriptors.
//
//
// Value Interfaces
//
// The Enum and Message interfaces provide a reflective view over an
// enum or message instance. For enums, it provides the ability to retrieve
// the enum value number for any concrete enum type. For messages, it provides
// the ability to access or manipulate fields of the message.
//
// To convert a proto.Message to a protoreflect.Message, use the
// former's ProtoReflect method. Since the ProtoReflect method is new to the
// v2 message interface, it may not be present on older message implementations.
// The "github.com/golang/protobuf/proto".MessageReflect function can be used
// to obtain a reflective view on older messages.
//
//
// Relationships
//
// The following diagrams demonstrate the relationships between
// various types declared in this package.
//
//
//	                       ┌───────────────────────────────────┐
//	                       V                                   │
//	   ┌────────────── New(n) ─────────────┐                   │
//	   │                                   │                   │
//	   │      ┌──── Descriptor() ──┐       │  ┌── Number() ──┐ │
//	   │      │                    V       V  │              V │
//	╔════════════╗  ╔════════════════╗  ╔════════╗  ╔════════════╗
//	║  EnumType  ║  ║ EnumDescriptor ║  ║  Enum  ║  ║ EnumNumber ║
//	╚════════════╝  ╚════════════════╝  ╚════════╝  ╚════════════╝
//	      Λ           Λ                   │ │
//	      │           └─── Descriptor() ──┘ │
//	      │                                 │
//	      └────────────────── Type() ───────┘
//
// • An EnumType describes a concrete Go enum type.
// It has an EnumDescriptor and can construct an Enum instance.
//
// • An EnumDescriptor describes an abstract protobuf enum type.
//
// • An Enum is a concrete enum instance. Generated enums implement Enum.
//
//
//	  ┌──────────────── New() ─────────────────┐
//	  │                                        │
//	  │         ┌─── Descriptor() ─────┐       │   ┌── Interface() ───┐
//	  │         │                      V       V   │                  V
//	╔═════════════╗  ╔═══════════════════╗  ╔═════════╗  ╔══════════════╗
//	║ MessageType ║  ║ MessageDescriptor ║  ║ Message ║  ║ ProtoMessage ║
//	╚═════════════╝  ╚═══════════════════╝  ╚═════════╝  ╚══════════════╝
//	       Λ           Λ                      │ │  Λ                  │
//	       │           └──── Descriptor() ────┘ │  └─ ProtoReflect() ─┘
//	       │                                    │
//	       └─────────────────── Type() ─────────┘
//
// • A MessageType describes a concrete Go message type.
// It has a MessageDescriptor and can construct a Message instance.
//
// • A MessageDescriptor describes an abstract protobuf message type.
//
// • A Message is a concrete message instance. Generated messages implement
// ProtoMessage, which can convert to/from a Message.
//
//
//	      ┌── TypeDescriptor() ──┐    ┌───── Descriptor() ─────┐
//	      │                      V    │                        V
//	╔═══════════════╗  ╔═════════════════════════╗  ╔═════════════════════╗
//	║ ExtensionType ║  ║ ExtensionTypeDescriptor ║  ║ ExtensionDescriptor ║
//	╚═══════════════╝  ╚═════════════════════════╝  ╚═════════════════════╝
//	      Λ                      │   │ Λ                      │ Λ
//	      └─────── Type() ───────┘   │ └─── may implement ────┘ │
//	                                 │                          │
//	                                 └────── implements ────────┘
//
// • An ExtensionType describes a concrete Go implementation of an extension.
// It has an ExtensionTypeDescriptor and can convert to/from
// abstract Values and Go values.
//
// • An ExtensionTypeDescriptor is an ExtensionDescriptor
// which also has an ExtensionType.
//
// • An ExtensionDescriptor describes an abstract protobuf extension field and
// may not always be an ExtensionTypeDescriptor.
package protoreflect

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/pragma"
)

type doNotImplement pragma.DoNotImplement

// ProtoMessage is the top-level interface that all proto messages implement.
// This is declared in the protoreflect package to avoid a cyclic dependency;
// use the proto.Message type instead, which aliases this type.
type ProtoMessage interface{ ProtoReflect() Message }

// Syntax is the language version of the proto file.
type Syntax syntax

type syntax int8 // keep exact type opaque as the int type may change

const (
	Proto2 Syntax = 2
	Proto3 Syntax = 3
)

// IsValid reports whether the syntax is valid.
func (s Syntax) IsValid() bool {
	switch s {
	case Proto2, Proto3:
		return true
	default:
		return false
	}
}

// String returns s as a proto source identifier (e.g., "proto2").
func (s Syntax) String() string {
	switch s {
	case Proto2:
		return "proto2"
	case Proto3:
		return "proto3"
	default:
		return fmt.Sprintf("<unknown:%d>", s)
	}
}

// GoString returns s as a Go source identifier (e.g., "Proto2").
func (s Syntax) GoString() string {
	switch s {
	case Proto2:
		return "Proto2"
	case Proto3:
		return "Proto3"
	default:
		return fmt.Sprintf("Syntax(%d)", s)
	}
}

// Cardinality determines whether a field is optional, required, or repeated.
type Cardinality cardinality

type cardinality int8 // keep exact type opaque as the int type may change

// Constants as defined by the google.protobuf.Cardinality enumeration.
const (
	Optional Cardinality = 1 // appears zero or one times
	Required Cardinality = 2 // appears exactly one time; invalid with Proto3
	Repeated Cardinality = 3 // appears zero or more times
)

// IsValid reports whether the cardinality is valid.
func (c Cardinality) IsValid() bool {
	switch c {
	case Optional, Required, Repeated:
		return true
	default:
		return false
	}
}

// String returns c as a proto source identifier (e.g., "optional").
func (c Cardinality) String() string {
	switch c {
	case Optional:
		return "optional"
	case Required:
		return "required"
	case Repeated:
		return "repeated"
	default:
		return fmt.Sprintf("<unknown:%d>", c)
	}
}

// GoString returns c as a Go source identifier (e.g., "Optional").
func (c Cardinality) GoString() string {
	switch c {
	case Optional:
		return "Optional"
	case Required:
		return "Required"
	case Repeated:
		return "Repeated"
	default:
		return fmt.Sprintf("Cardinality(%d)", c)
	}
}

// Kind indicates the basic proto kind of a field.
type Kind kind

type kind int8 // keep exact type opaque as the int type may change

// Constants as defined by the google.protobuf.Field.Kind enumeration.
const (
	BoolKind     Kind = 8
	EnumKind     Kind = 14
	Int32Kind    Kind = 5
	Sint32Kind   Kind = 17
	Uint32Kind   Kind = 13
	Int64Kind    Kind = 3
	Sint64Kind   Kind = 18
	Uint64Kind   Kind = 4
	Sfixed32Kind Kind = 15
	Fixed32Kind  Kind = 7
	FloatKind    Kind = 2
	Sfixed64Kind Kind = 16
	Fixed64Kind  Kind = 6
	DoubleKind   Kind = 1
	StringKind   Kind = 9
	BytesKind    Kind = 12
	MessageKind  Kind = 11
	GroupKind    Kind = 10
)

// IsValid reports whether the kind is valid.
func (k Kind) IsValid() bool {
	switch k {
	case BoolKind, EnumKind,
		Int32Kind, Sint32Kind, Uint32Kind,
		Int64Kind, Sint64Kind, Uint64Kind,
		Sfixed32Kind, Fixed32Kind, FloatKind,
		Sfixed64Kind, Fixed64Kind, DoubleKind,
		StringKind, BytesKind, MessageKind, GroupKind:
		return true
	default:
		return false
	}
}

// String returns k as a proto source identifier (e.g., "bool").
func (k Kind) String() string {
	switch k {
	case BoolKind:
		return "bool"
	case EnumKind:
		return "enum"
	case Int32Kind:
		return "int32"
	case Sint32Kind:
		return "sint32"
	case Uint32Kind:
		return "uint32"
	case Int64Kind:
		return "int64"
	case Sint64Kind:
		return "sint64"
	case Uint64Kind:
		return "uint64"
	case Sfixed32Kind:
		return "sfixed32"
	case Fixed32Kind:
		return "fixed32"
	case FloatKind:
		return "float"
	case Sfixed64Kind:
		return "sfixed64"
	case Fixed64Kind:
		return "fixed64"
	case DoubleKind:
		return "double"
	case StringKind:
		return "string"
	case BytesKind:
		return "bytes"
	case MessageKind:
		return "message"
	case GroupKind:
		return "group"
	default:
		return fmt.Sprintf("<unknown:%d>", k)
	}
}

// GoString returns k as a Go source identifier (e.g., "BoolKind").
func (k Kind) GoString() string {
	switch k {
	case BoolKind:
		return "BoolKind"
	case EnumKind:
		return "EnumKind"
	case Int32Kind:
		return "Int32Kind"
	case Sint32Kind:
		return "Sint32Kind"
	case Uint32Kind:
		return "Uint32Kind"
	case Int64Kind:
		return "Int64Kind"
	case Sint64Kind:
		return "Sint64Kind"
	case Uint64Kind:
		return "Uint64Kind"
	case Sfixed32Kind:
		return "Sfixed32Kind"
	case Fixed32Kind:
		return "Fixed32Kind"
	case FloatKind:
		return "FloatKind"
	case Sfixed64Kind:
		return "Sfixed64Kind"
	case Fixed64Kind:
		return "Fixed64Kind"
	case DoubleKind:
		return "DoubleKind"
	case StringKind:
		return "StringKind"
	case BytesKind:
		return "BytesKind"
	case MessageKind:
		return "MessageKind"
	case GroupKind:
		return "GroupKind"
	default:
		return fmt.Sprintf("Kind(%d)", k)
	}
}

// FieldNumber is the field number in a message.
type FieldNumber = protowire.Number

// FieldNumbers represent a list of field numbers.
type FieldNumbers interface {
	// Len reports the number of fields in the list.
	Len() int
	// Get returns the ith field number. It panics if out of bounds.
	Get(i int) FieldNumber
	// Has reports whether n is within the list of fields.
	Has(n FieldNumber) bool

	doNotImplement
}

// FieldRanges represent a list of field number ranges.
type FieldRanges interface {
	// Len reports the number of ranges in the list.
	Len() int
	// Get returns the ith range. It panics if out of bounds.
	Get(i int) [2]FieldNumber // start inclusive; end exclusive
	// Has reports whether n is within any of the ranges.
	Has(n FieldNumber) bool

	doNotImplement
}

// EnumNumber is the numeric value for an enum.
type EnumNumber int32

// EnumRanges represent a list of enum number ranges.
type EnumRanges interface {
	// Len reports the number of ranges in the list.
	Len() int
	// Get returns the ith range. It panics if out of bounds.
	Get(i int) [2]EnumNumber // start inclusive; end inclusive
	// Has reports whether n is within any of the ranges.
	Has(n EnumNumber) bool

	doNotImplement
}

// Name is the short name for a proto declaration. This is not the name
// as used in Go source code, which might not be identical to the proto name.
type Name string // e.g., "Kind"

// IsValid reports whether s is a syntactically valid name.
// An empty name is invalid.
func (s Name) IsValid() bool {
	return consumeIdent(string(s)) == len(s)
}

// Names represent a list of names.
type Names interface {
	// Len reports the number of names in the list.
	Len() int
	// Get returns the ith name. It panics if out of bounds.
	Get(i int) Name
	// Has reports whether s matches any names in the list.
	Has(s Name) bool

	doNotImplement
}

// FullName is a qualified name that uniquely identifies a proto declaration.
// A qualified name is the concatenation of the proto package along with the
// fully-declared name (i.e., name of parent preceding the name of the child),
// with a '.' delimiter placed between each Name.
//
// This should not have any leading or trailing dots.
type FullName string // e.g., "google.protobuf.Field.Kind"

// IsValid reports whether s is a syntactically valid full name.
// An empty full name is invalid.
func (s FullName) IsValid() bool {
	i := consumeIdent(string(s))
	if i < 0 {
		return false
	}
	for len(s) > i {
		if s[i] != '.' {
			return false
		}
		i++
		n := consumeIdent(string(s[i:]))
		if n < 0 {
			return false
		}
		i += n
	}
	return true
}

func consumeIdent(s string) (i int) {
	if len(s) == 0 || !isLetter(s[i]) {
		return -1
	}
	i++
	for len(s) > i && isLetterDigit(s[i]) {
		i++
	}
	return i
}
func isLetter(c byte) bool {
	return c == '_' || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}
func isLetterDigit(c byte) bool {
	return isLetter(c) || ('0' <= c && c <= '9')
}

// Name returns the short name, which is the last identifier segment.
// A single segment FullName is the Name itself.
func (n FullName) Name() Name {
	if i := strings.LastIndexByte(string(n), '.'); i >= 0 {
		return Name(n[i+1:])
	}
	return Name(n)
}

// Parent returns the full name with the trailing identifier removed.
// A single segment FullName has no parent.
func (n FullName) Parent() FullName {
	if i := strings.LastIndexByte(string(n), '.'); i >= 0 {
		return n[:i]
	}
	return ""
}

// Append returns the qualified name appended with the provided short name.
//
// Invariant: n == n.Parent().Append(n.Name()) // assuming n is valid
func (n FullName) Append(s Name) FullName {
	if n == "" {
		return FullName(s)
	}
	return n + "." + FullName(s)
}
