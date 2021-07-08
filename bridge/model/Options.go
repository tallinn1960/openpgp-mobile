// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Options struct {
	_tab flatbuffers.Table
}

func GetRootAsOptions(buf []byte, offset flatbuffers.UOffsetT) *Options {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Options{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsOptions(buf []byte, offset flatbuffers.UOffsetT) *Options {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Options{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Options) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Options) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Options) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Options) Comment() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Options) Email() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Options) Passphrase() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Options) KeyOptions(obj *KeyOptions) *KeyOptions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
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

func OptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func OptionsAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func OptionsAddComment(builder *flatbuffers.Builder, comment flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(comment), 0)
}
func OptionsAddEmail(builder *flatbuffers.Builder, email flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(email), 0)
}
func OptionsAddPassphrase(builder *flatbuffers.Builder, passphrase flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(passphrase), 0)
}
func OptionsAddKeyOptions(builder *flatbuffers.Builder, keyOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(keyOptions), 0)
}
func OptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
