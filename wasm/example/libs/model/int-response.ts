// automatically generated by the FlatBuffers compiler, do not modify

import * as flatbuffers from 'flatbuffers';

export class IntResponse {
  bb: flatbuffers.ByteBuffer|null = null;
  bb_pos = 0;
__init(i:number, bb:flatbuffers.ByteBuffer):IntResponse {
  this.bb_pos = i;
  this.bb = bb;
  return this;
}

static getRootAsIntResponse(bb:flatbuffers.ByteBuffer, obj?:IntResponse):IntResponse {
  return (obj || new IntResponse()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

static getSizePrefixedRootAsIntResponse(bb:flatbuffers.ByteBuffer, obj?:IntResponse):IntResponse {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new IntResponse()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

output():flatbuffers.Long {
  const offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? this.bb!.readInt64(this.bb_pos + offset) : this.bb!.createLong(0, 0);
}

mutate_output(value:flatbuffers.Long):boolean {
  const offset = this.bb!.__offset(this.bb_pos, 4);

  if (offset === 0) {
    return false;
  }

  this.bb!.writeInt64(this.bb_pos + offset, value);
  return true;
}

error():string|null
error(optionalEncoding:flatbuffers.Encoding):string|Uint8Array|null
error(optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.__string(this.bb_pos + offset, optionalEncoding) : null;
}

static startIntResponse(builder:flatbuffers.Builder) {
  builder.startObject(2);
}

static addOutput(builder:flatbuffers.Builder, output:flatbuffers.Long) {
  builder.addFieldInt64(0, output, builder.createLong(0, 0));
}

static addError(builder:flatbuffers.Builder, errorOffset:flatbuffers.Offset) {
  builder.addFieldOffset(1, errorOffset, 0);
}

static endIntResponse(builder:flatbuffers.Builder):flatbuffers.Offset {
  const offset = builder.endObject();
  return offset;
}

static createIntResponse(builder:flatbuffers.Builder, output:flatbuffers.Long, errorOffset:flatbuffers.Offset):flatbuffers.Offset {
  IntResponse.startIntResponse(builder);
  IntResponse.addOutput(builder, output);
  IntResponse.addError(builder, errorOffset);
  return IntResponse.endIntResponse(builder);
}
}
