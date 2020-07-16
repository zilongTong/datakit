// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-protos. DO NOT EDIT.

package fieldnum

// Field numbers for google.protobuf.FileDescriptorSet.
const (
	FileDescriptorSet_File = 1 // repeated google.protobuf.FileDescriptorProto
)

// Field numbers for google.protobuf.FileDescriptorProto.
const (
	FileDescriptorProto_Name             = 1  // optional string
	FileDescriptorProto_Package          = 2  // optional string
	FileDescriptorProto_Dependency       = 3  // repeated string
	FileDescriptorProto_PublicDependency = 10 // repeated int32
	FileDescriptorProto_WeakDependency   = 11 // repeated int32
	FileDescriptorProto_MessageType      = 4  // repeated google.protobuf.DescriptorProto
	FileDescriptorProto_EnumType         = 5  // repeated google.protobuf.EnumDescriptorProto
	FileDescriptorProto_Service          = 6  // repeated google.protobuf.ServiceDescriptorProto
	FileDescriptorProto_Extension        = 7  // repeated google.protobuf.FieldDescriptorProto
	FileDescriptorProto_Options          = 8  // optional google.protobuf.FileOptions
	FileDescriptorProto_SourceCodeInfo   = 9  // optional google.protobuf.SourceCodeInfo
	FileDescriptorProto_Syntax           = 12 // optional string
)

// Field numbers for google.protobuf.DescriptorProto.
const (
	DescriptorProto_Name           = 1  // optional string
	DescriptorProto_Field          = 2  // repeated google.protobuf.FieldDescriptorProto
	DescriptorProto_Extension      = 6  // repeated google.protobuf.FieldDescriptorProto
	DescriptorProto_NestedType     = 3  // repeated google.protobuf.DescriptorProto
	DescriptorProto_EnumType       = 4  // repeated google.protobuf.EnumDescriptorProto
	DescriptorProto_ExtensionRange = 5  // repeated google.protobuf.DescriptorProto.ExtensionRange
	DescriptorProto_OneofDecl      = 8  // repeated google.protobuf.OneofDescriptorProto
	DescriptorProto_Options        = 7  // optional google.protobuf.MessageOptions
	DescriptorProto_ReservedRange  = 9  // repeated google.protobuf.DescriptorProto.ReservedRange
	DescriptorProto_ReservedName   = 10 // repeated string
)

// Field numbers for google.protobuf.DescriptorProto.ExtensionRange.
const (
	DescriptorProto_ExtensionRange_Start   = 1 // optional int32
	DescriptorProto_ExtensionRange_End     = 2 // optional int32
	DescriptorProto_ExtensionRange_Options = 3 // optional google.protobuf.ExtensionRangeOptions
)

// Field numbers for google.protobuf.DescriptorProto.ReservedRange.
const (
	DescriptorProto_ReservedRange_Start = 1 // optional int32
	DescriptorProto_ReservedRange_End   = 2 // optional int32
)

// Field numbers for google.protobuf.ExtensionRangeOptions.
const (
	ExtensionRangeOptions_UninterpretedOption = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.FieldDescriptorProto.
const (
	FieldDescriptorProto_Name           = 1  // optional string
	FieldDescriptorProto_Number         = 3  // optional int32
	FieldDescriptorProto_Label          = 4  // optional google.protobuf.FieldDescriptorProto.Label
	FieldDescriptorProto_Type           = 5  // optional google.protobuf.FieldDescriptorProto.Type
	FieldDescriptorProto_TypeName       = 6  // optional string
	FieldDescriptorProto_Extendee       = 2  // optional string
	FieldDescriptorProto_DefaultValue   = 7  // optional string
	FieldDescriptorProto_OneofIndex     = 9  // optional int32
	FieldDescriptorProto_JsonName       = 10 // optional string
	FieldDescriptorProto_Options        = 8  // optional google.protobuf.FieldOptions
	FieldDescriptorProto_Proto3Optional = 17 // optional bool
)

// Field numbers for google.protobuf.OneofDescriptorProto.
const (
	OneofDescriptorProto_Name    = 1 // optional string
	OneofDescriptorProto_Options = 2 // optional google.protobuf.OneofOptions
)

// Field numbers for google.protobuf.EnumDescriptorProto.
const (
	EnumDescriptorProto_Name          = 1 // optional string
	EnumDescriptorProto_Value         = 2 // repeated google.protobuf.EnumValueDescriptorProto
	EnumDescriptorProto_Options       = 3 // optional google.protobuf.EnumOptions
	EnumDescriptorProto_ReservedRange = 4 // repeated google.protobuf.EnumDescriptorProto.EnumReservedRange
	EnumDescriptorProto_ReservedName  = 5 // repeated string
)

// Field numbers for google.protobuf.EnumDescriptorProto.EnumReservedRange.
const (
	EnumDescriptorProto_EnumReservedRange_Start = 1 // optional int32
	EnumDescriptorProto_EnumReservedRange_End   = 2 // optional int32
)

// Field numbers for google.protobuf.EnumValueDescriptorProto.
const (
	EnumValueDescriptorProto_Name    = 1 // optional string
	EnumValueDescriptorProto_Number  = 2 // optional int32
	EnumValueDescriptorProto_Options = 3 // optional google.protobuf.EnumValueOptions
)

// Field numbers for google.protobuf.ServiceDescriptorProto.
const (
	ServiceDescriptorProto_Name    = 1 // optional string
	ServiceDescriptorProto_Method  = 2 // repeated google.protobuf.MethodDescriptorProto
	ServiceDescriptorProto_Options = 3 // optional google.protobuf.ServiceOptions
)

// Field numbers for google.protobuf.MethodDescriptorProto.
const (
	MethodDescriptorProto_Name            = 1 // optional string
	MethodDescriptorProto_InputType       = 2 // optional string
	MethodDescriptorProto_OutputType      = 3 // optional string
	MethodDescriptorProto_Options         = 4 // optional google.protobuf.MethodOptions
	MethodDescriptorProto_ClientStreaming = 5 // optional bool
	MethodDescriptorProto_ServerStreaming = 6 // optional bool
)

// Field numbers for google.protobuf.FileOptions.
const (
	FileOptions_JavaPackage               = 1   // optional string
	FileOptions_JavaOuterClassname        = 8   // optional string
	FileOptions_JavaMultipleFiles         = 10  // optional bool
	FileOptions_JavaGenerateEqualsAndHash = 20  // optional bool
	FileOptions_JavaStringCheckUtf8       = 27  // optional bool
	FileOptions_OptimizeFor               = 9   // optional google.protobuf.FileOptions.OptimizeMode
	FileOptions_GoPackage                 = 11  // optional string
	FileOptions_CcGenericServices         = 16  // optional bool
	FileOptions_JavaGenericServices       = 17  // optional bool
	FileOptions_PyGenericServices         = 18  // optional bool
	FileOptions_PhpGenericServices        = 42  // optional bool
	FileOptions_Deprecated                = 23  // optional bool
	FileOptions_CcEnableArenas            = 31  // optional bool
	FileOptions_ObjcClassPrefix           = 36  // optional string
	FileOptions_CsharpNamespace           = 37  // optional string
	FileOptions_SwiftPrefix               = 39  // optional string
	FileOptions_PhpClassPrefix            = 40  // optional string
	FileOptions_PhpNamespace              = 41  // optional string
	FileOptions_PhpMetadataNamespace      = 44  // optional string
	FileOptions_RubyPackage               = 45  // optional string
	FileOptions_UninterpretedOption       = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.MessageOptions.
const (
	MessageOptions_MessageSetWireFormat         = 1   // optional bool
	MessageOptions_NoStandardDescriptorAccessor = 2   // optional bool
	MessageOptions_Deprecated                   = 3   // optional bool
	MessageOptions_MapEntry                     = 7   // optional bool
	MessageOptions_UninterpretedOption          = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.FieldOptions.
const (
	FieldOptions_Ctype               = 1   // optional google.protobuf.FieldOptions.CType
	FieldOptions_Packed              = 2   // optional bool
	FieldOptions_Jstype              = 6   // optional google.protobuf.FieldOptions.JSType
	FieldOptions_Lazy                = 5   // optional bool
	FieldOptions_Deprecated          = 3   // optional bool
	FieldOptions_Weak                = 10  // optional bool
	FieldOptions_UninterpretedOption = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.OneofOptions.
const (
	OneofOptions_UninterpretedOption = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.EnumOptions.
const (
	EnumOptions_AllowAlias          = 2   // optional bool
	EnumOptions_Deprecated          = 3   // optional bool
	EnumOptions_UninterpretedOption = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.EnumValueOptions.
const (
	EnumValueOptions_Deprecated          = 1   // optional bool
	EnumValueOptions_UninterpretedOption = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.ServiceOptions.
const (
	ServiceOptions_Deprecated          = 33  // optional bool
	ServiceOptions_UninterpretedOption = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.MethodOptions.
const (
	MethodOptions_Deprecated          = 33  // optional bool
	MethodOptions_IdempotencyLevel    = 34  // optional google.protobuf.MethodOptions.IdempotencyLevel
	MethodOptions_UninterpretedOption = 999 // repeated google.protobuf.UninterpretedOption
)

// Field numbers for google.protobuf.UninterpretedOption.
const (
	UninterpretedOption_Name             = 2 // repeated google.protobuf.UninterpretedOption.NamePart
	UninterpretedOption_IdentifierValue  = 3 // optional string
	UninterpretedOption_PositiveIntValue = 4 // optional uint64
	UninterpretedOption_NegativeIntValue = 5 // optional int64
	UninterpretedOption_DoubleValue      = 6 // optional double
	UninterpretedOption_StringValue      = 7 // optional bytes
	UninterpretedOption_AggregateValue   = 8 // optional string
)

// Field numbers for google.protobuf.UninterpretedOption.NamePart.
const (
	UninterpretedOption_NamePart_NamePart    = 1 // required string
	UninterpretedOption_NamePart_IsExtension = 2 // required bool
)

// Field numbers for google.protobuf.SourceCodeInfo.
const (
	SourceCodeInfo_Location = 1 // repeated google.protobuf.SourceCodeInfo.Location
)

// Field numbers for google.protobuf.SourceCodeInfo.Location.
const (
	SourceCodeInfo_Location_Path                    = 1 // repeated int32
	SourceCodeInfo_Location_Span                    = 2 // repeated int32
	SourceCodeInfo_Location_LeadingComments         = 3 // optional string
	SourceCodeInfo_Location_TrailingComments        = 4 // optional string
	SourceCodeInfo_Location_LeadingDetachedComments = 6 // repeated string
)

// Field numbers for google.protobuf.GeneratedCodeInfo.
const (
	GeneratedCodeInfo_Annotation = 1 // repeated google.protobuf.GeneratedCodeInfo.Annotation
)

// Field numbers for google.protobuf.GeneratedCodeInfo.Annotation.
const (
	GeneratedCodeInfo_Annotation_Path       = 1 // repeated int32
	GeneratedCodeInfo_Annotation_SourceFile = 2 // optional string
	GeneratedCodeInfo_Annotation_Begin      = 3 // optional int32
	GeneratedCodeInfo_Annotation_End        = 4 // optional int32
)
