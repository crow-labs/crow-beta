/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crow.escrow";

export interface BasicEscrow {
  escrowId: number;
  address: string;
  orderId: number;
  price: string;
  status: string;
}

const baseBasicEscrow: object = {
  escrowId: 0,
  address: "",
  orderId: 0,
  price: "",
  status: "",
};

export const BasicEscrow = {
  encode(message: BasicEscrow, writer: Writer = Writer.create()): Writer {
    if (message.escrowId !== 0) {
      writer.uint32(8).uint64(message.escrowId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.orderId !== 0) {
      writer.uint32(24).uint64(message.orderId);
    }
    if (message.price !== "") {
      writer.uint32(34).string(message.price);
    }
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BasicEscrow {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBasicEscrow } as BasicEscrow;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.escrowId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.price = reader.string();
          break;
        case 5:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BasicEscrow {
    const message = { ...baseBasicEscrow } as BasicEscrow;
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = Number(object.escrowId);
    } else {
      message.escrowId = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: BasicEscrow): unknown {
    const obj: any = {};
    message.escrowId !== undefined && (obj.escrowId = message.escrowId);
    message.address !== undefined && (obj.address = message.address);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    message.price !== undefined && (obj.price = message.price);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<BasicEscrow>): BasicEscrow {
    const message = { ...baseBasicEscrow } as BasicEscrow;
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = object.escrowId;
    } else {
      message.escrowId = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
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
