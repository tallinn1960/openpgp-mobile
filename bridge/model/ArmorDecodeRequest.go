// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArmorDecodeRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsArmorDecodeRequest(buf []byte, offset flatbuffers.UOffsetT) *ArmorDecodeRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArmorDecodeRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArmorDecodeRequest(buf []byte, offset flatbuffers.UOffsetT) *ArmorDecodeRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArmorDecodeRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArmorDecodeRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArmorDecodeRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArmorDecodeRequest) Packet() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ArmorDecodeRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArmorDecodeRequestAddPacket(builder *flatbuffers.Builder, packet flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(packet), 0)
}
func ArmorDecodeRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
