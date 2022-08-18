/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crow.escrow";

export interface UserDispute {
  escrowId: number;
  userId: number;
  description: string;
}

const baseUserDispute: object = { escrowId: 0, userId: 0, description: "" };

export const UserDispute = {
  encode(message: UserDispute, writer: Writer = Writer.create()): Writer {
    if (message.escrowId !== 0) {
      writer.uint32(8).uint64(message.escrowId);
    }
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UserDispute {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUserDispute } as UserDispute;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.escrowId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.userId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserDispute {
    const message = { ...baseUserDispute } as UserDispute;
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = Number(object.escrowId);
    } else {
      message.escrowId = 0;
    }
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = Number(object.userId);
    } else {
      message.userId = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    return message;
  },

  toJSON(message: UserDispute): unknown {
    const obj: any = {};
    message.escrowId !== undefined && (obj.escrowId = message.escrowId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.description !== undefined &&
      (obj.description = message.description);
    return obj;
  },

  fromPartial(object: DeepPartial<UserDispute>): UserDispute {
    const message = { ...baseUserDispute } as UserDispute;
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = object.escrowId;
    } else {
      message.escrowId = 0;
    }
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = object.userId;
    } else {
      message.userId = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
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
