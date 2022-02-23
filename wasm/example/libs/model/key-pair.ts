// automatically generated by the FlatBuffers compiler, do not modify

import * as flatbuffers from 'flatbuffers';

export class KeyPair {
  bb: flatbuffers.ByteBuffer|null = null;
  bb_pos = 0;
__init(i:number, bb:flatbuffers.ByteBuffer):KeyPair {
  this.bb_pos = i;
  this.bb = bb;
  return this;
}

static getRootAsKeyPair(bb:flatbuffers.ByteBuffer, obj?:KeyPair):KeyPair {
  return (obj || new KeyPair()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

static getSizePrefixedRootAsKeyPair(bb:flatbuffers.ByteBuffer, obj?:KeyPair):KeyPair {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new KeyPair()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

publicKey():string|null
publicKey(optionalEncoding:flatbuffers.Encoding):string|Uint8Array|null
publicKey(optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? this.bb!.__string(this.bb_pos + offset, optionalEncoding) : null;
}

privateKey():string|null
privateKey(optionalEncoding:flatbuffers.Encoding):string|Uint8Array|null
privateKey(optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.__string(this.bb_pos + offset, optionalEncoding) : null;
}

static startKeyPair(builder:flatbuffers.Builder) {
  builder.startObject(2);
}

static addPublicKey(builder:flatbuffers.Builder, publicKeyOffset:flatbuffers.Offset) {
  builder.addFieldOffset(0, publicKeyOffset, 0);
}

static addPrivateKey(builder:flatbuffers.Builder, privateKeyOffset:flatbuffers.Offset) {
  builder.addFieldOffset(1, privateKeyOffset, 0);
}

static endKeyPair(builder:flatbuffers.Builder):flatbuffers.Offset {
  const offset = builder.endObject();
  return offset;
}

static createKeyPair(builder:flatbuffers.Builder, publicKeyOffset:flatbuffers.Offset, privateKeyOffset:flatbuffers.Offset):flatbuffers.Offset {
  KeyPair.startKeyPair(builder);
  KeyPair.addPublicKey(builder, publicKeyOffset);
  KeyPair.addPrivateKey(builder, privateKeyOffset);
  return KeyPair.endKeyPair(builder);
}
}