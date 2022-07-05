// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PrivateKeyMetadataResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsPrivateKeyMetadataResponse(buf []byte, offset flatbuffers.UOffsetT) *PrivateKeyMetadataResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PrivateKeyMetadataResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPrivateKeyMetadataResponse(buf []byte, offset flatbuffers.UOffsetT) *PrivateKeyMetadataResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PrivateKeyMetadataResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PrivateKeyMetadataResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PrivateKeyMetadataResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PrivateKeyMetadataResponse) Output(obj *PrivateKeyMetadata) *PrivateKeyMetadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PrivateKeyMetadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PrivateKeyMetadataResponse) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PrivateKeyMetadataResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PrivateKeyMetadataResponseAddOutput(builder *flatbuffers.Builder, output flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(output), 0)
}
func PrivateKeyMetadataResponseAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(error), 0)
}
func PrivateKeyMetadataResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
