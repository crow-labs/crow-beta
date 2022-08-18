/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crow.whitelist";

export interface Producer {
  producerId: number;
  address: string;
  status: string;
}

export interface User {
  userId: number;
  address: string;
  status: string;
}

export interface Whitelist {
  whitelistId: number;
  governor: string;
  users: User | undefined;
  producers: Producer | undefined;
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

const baseUser: object = { userId: 0, address: "", status: "" };

export const User = {
  encode(message: User, writer: Writer = Writer.create()): Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): User {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUser } as User;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): User {
    const message = { ...baseUser } as User;
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = Number(object.userId);
    } else {
      message.userId = 0;
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

  toJSON(message: User): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.address !== undefined && (obj.address = message.address);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<User>): User {
    const message = { ...baseUser } as User;
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = object.userId;
    } else {
      message.userId = 0;
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

const baseWhitelist: object = { whitelistId: 0, governor: "" };

export const Whitelist = {
  encode(message: Whitelist, writer: Writer = Writer.create()): Writer {
    if (message.whitelistId !== 0) {
      writer.uint32(8).uint64(message.whitelistId);
    }
    if (message.governor !== "") {
      writer.uint32(18).string(message.governor);
    }
    if (message.users !== undefined) {
      User.encode(message.users, writer.uint32(26).fork()).ldelim();
    }
    if (message.producers !== undefined) {
      Producer.encode(message.producers, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Whitelist {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhitelist } as Whitelist;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.whitelistId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.governor = reader.string();
          break;
        case 3:
          message.users = User.decode(reader, reader.uint32());
          break;
        case 4:
          message.producers = Producer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Whitelist {
    const message = { ...baseWhitelist } as Whitelist;
    if (object.whitelistId !== undefined && object.whitelistId !== null) {
      message.whitelistId = Number(object.whitelistId);
    } else {
      message.whitelistId = 0;
    }
    if (object.governor !== undefined && object.governor !== null) {
      message.governor = String(object.governor);
    } else {
      message.governor = "";
    }
    if (object.users !== undefined && object.users !== null) {
      message.users = User.fromJSON(object.users);
    } else {
      message.users = undefined;
    }
    if (object.producers !== undefined && object.producers !== null) {
      message.producers = Producer.fromJSON(object.producers);
    } else {
      message.producers = undefined;
    }
    return message;
  },

  toJSON(message: Whitelist): unknown {
    const obj: any = {};
    message.whitelistId !== undefined &&
      (obj.whitelistId = message.whitelistId);
    message.governor !== undefined && (obj.governor = message.governor);
    message.users !== undefined &&
      (obj.users = message.users ? User.toJSON(message.users) : undefined);
    message.producers !== undefined &&
      (obj.producers = message.producers
        ? Producer.toJSON(message.producers)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Whitelist>): Whitelist {
    const message = { ...baseWhitelist } as Whitelist;
    if (object.whitelistId !== undefined && object.whitelistId !== null) {
      message.whitelistId = object.whitelistId;
    } else {
      message.whitelistId = 0;
    }
    if (object.governor !== undefined && object.governor !== null) {
      message.governor = object.governor;
    } else {
      message.governor = "";
    }
    if (object.users !== undefined && object.users !== null) {
      message.users = User.fromPartial(object.users);
    } else {
      message.users = undefined;
    }
    if (object.producers !== undefined && object.producers !== null) {
      message.producers = Producer.fromPartial(object.producers);
    } else {
      message.producers = undefined;
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
