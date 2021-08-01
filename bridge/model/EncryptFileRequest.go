// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EncryptFileRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsEncryptFileRequest(buf []byte, offset flatbuffers.UOffsetT) *EncryptFileRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EncryptFileRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEncryptFileRequest(buf []byte, offset flatbuffers.UOffsetT) *EncryptFileRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EncryptFileRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EncryptFileRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EncryptFileRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *EncryptFileRequest) Input() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *EncryptFileRequest) Output() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *EncryptFileRequest) PublicKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *EncryptFileRequest) Options(obj *KeyOptions) *KeyOptions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KeyOptions)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *EncryptFileRequest) Signed(obj *Entity) *Entity {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Entity)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *EncryptFileRequest) FileHints(obj *FileHints) *FileHints {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FileHints)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func EncryptFileRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func EncryptFileRequestAddInput(builder *flatbuffers.Builder, input flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(input), 0)
}
func EncryptFileRequestAddOutput(builder *flatbuffers.Builder, output flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(output), 0)
}
func EncryptFileRequestAddPublicKey(builder *flatbuffers.Builder, publicKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(publicKey), 0)
}
func EncryptFileRequestAddOptions(builder *flatbuffers.Builder, options flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(options), 0)
}
func EncryptFileRequestAddSigned(builder *flatbuffers.Builder, signed flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(signed), 0)
}
func EncryptFileRequestAddFileHints(builder *flatbuffers.Builder, fileHints flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(fileHints), 0)
}
func EncryptFileRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
