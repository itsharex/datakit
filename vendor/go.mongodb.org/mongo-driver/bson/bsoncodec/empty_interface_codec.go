// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package bsoncodec

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsonoptions"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmptyInterfaceCodec is the Codec used for interface{} values.
type EmptyInterfaceCodec struct {
	DecodeBinaryAsSlice bool
}

var (
	defaultEmptyInterfaceCodec = NewEmptyInterfaceCodec()

	_ ValueCodec  = defaultEmptyInterfaceCodec
	_ typeDecoder = defaultEmptyInterfaceCodec
)

// NewEmptyInterfaceCodec returns a EmptyInterfaceCodec with options opts.
func NewEmptyInterfaceCodec(opts ...*bsonoptions.EmptyInterfaceCodecOptions) *EmptyInterfaceCodec {
	interfaceOpt := bsonoptions.MergeEmptyInterfaceCodecOptions(opts...)

	codec := EmptyInterfaceCodec{}
	if interfaceOpt.DecodeBinaryAsSlice != nil {
		codec.DecodeBinaryAsSlice = *interfaceOpt.DecodeBinaryAsSlice
	}
	return &codec
}

// EncodeValue is the ValueEncoderFunc for interface{}.
func (eic EmptyInterfaceCodec) EncodeValue(ec EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tEmpty {
		return ValueEncoderError{Name: "EmptyInterfaceEncodeValue", Types: []reflect.Type{tEmpty}, Received: val}
	}

	if val.IsNil() {
		return vw.WriteNull()
	}
	encoder, err := ec.LookupEncoder(val.Elem().Type())
	if err != nil {
		return err
	}

	return encoder.EncodeValue(ec, vw, val.Elem())
}

func (eic EmptyInterfaceCodec) getEmptyInterfaceDecodeType(dc DecodeContext, valueType bsontype.Type) (reflect.Type, error) {
	isDocument := valueType == bsontype.Type(0) || valueType == bsontype.EmbeddedDocument
	if isDocument {
		if dc.defaultDocumentType != nil {
			// If the bsontype is an embedded document and the DocumentType is set on the DecodeContext, then return
			// that type.
			return dc.defaultDocumentType, nil
		}
		if dc.Ancestor != nil {
			// Using ancestor information rather than looking up the type map entry forces consistent decoding.
			// If we're decoding into a bson.D, subdocuments should also be decoded as bson.D, even if a type map entry
			// has been registered.
			return dc.Ancestor, nil
		}
	}

	rtype, err := dc.LookupTypeMapEntry(valueType)
	if err == nil {
		return rtype, nil
	}

	if isDocument {
		// For documents, fallback to looking up a type map entry for bsontype.Type(0) or bsontype.EmbeddedDocument,
		// depending on the original valueType.
		var lookupType bsontype.Type
		switch valueType {
		case bsontype.Type(0):
			lookupType = bsontype.EmbeddedDocument
		case bsontype.EmbeddedDocument:
			lookupType = bsontype.Type(0)
		}

		rtype, err = dc.LookupTypeMapEntry(lookupType)
		if err == nil {
			return rtype, nil
		}
	}

	return nil, err
}

func (eic EmptyInterfaceCodec) decodeType(dc DecodeContext, vr bsonrw.ValueReader, t reflect.Type) (reflect.Value, error) {
	if t != tEmpty {
		return emptyValue, ValueDecoderError{Name: "EmptyInterfaceDecodeValue", Types: []reflect.Type{tEmpty}, Received: reflect.Zero(t)}
	}

	rtype, err := eic.getEmptyInterfaceDecodeType(dc, vr.Type())
	if err != nil {
		switch vr.Type() {
		case bsontype.Null:
			return reflect.Zero(t), vr.ReadNull()
		default:
			return emptyValue, err
		}
	}

	decoder, err := dc.LookupDecoder(rtype)
	if err != nil {
		return emptyValue, err
	}

	elem, err := decodeTypeOrValue(decoder, dc, vr, rtype)
	if err != nil {
		return emptyValue, err
	}

	if eic.DecodeBinaryAsSlice && rtype == tBinary {
		binElem := elem.Interface().(primitive.Binary)
		if binElem.Subtype == bsontype.BinaryGeneric || binElem.Subtype == bsontype.BinaryBinaryOld {
			elem = reflect.ValueOf(binElem.Data)
		}
	}

	return elem, nil
}

// DecodeValue is the ValueDecoderFunc for interface{}.
func (eic EmptyInterfaceCodec) DecodeValue(dc DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tEmpty {
		return ValueDecoderError{Name: "EmptyInterfaceDecodeValue", Types: []reflect.Type{tEmpty}, Received: val}
	}

	elem, err := eic.decodeType(dc, vr, val.Type())
	if err != nil {
		return err
	}

	val.Set(elem)
	return nil
}
