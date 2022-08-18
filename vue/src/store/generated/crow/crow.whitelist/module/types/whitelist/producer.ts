/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crow.whitelist";

export interface Producer {
  producerId: number;
  address: string;
  status: string;
}

const baseProducer: object = { producerId: 0, address: "", status: "" };

export const Producer = {
  encode(message: Producer, writer: Writer = Writer.create()): Writer {
    if (message.producerId !== 0) {
      writer.uint32(8).uint64(message.producerId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Producer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProducer } as Producer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.producerId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Producer {
    const message = { ...baseProducer } as Producer;
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = Number(object.producerId);
    } else {
      message.producerId = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: Producer): unknown {
    const obj: any = {};
    message.producerId !== undefined && (obj.producerId = message.producerId);
    message.address !== undefined && (obj.address = message.address);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<Producer>): Producer {
    const message = { ...baseProducer } as Producer;
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = object.producerId;
    } else {
      message.producerId = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
