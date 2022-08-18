/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crow.whitelist";

export interface Producer {
  producerId: number;
  name: string;
  address: string;
  status: string;
}

export interface User {
  userId: number;
  name: string;
  address: string;
  status: string;
}

export interface Whitelist {
  whitelistId: number;
  governor: string;
  users: User[];
  producers: Producer[];
}

const baseProducer: object = {
  producerId: 0,
  name: "",
  address: "",
  status: "",
};

export const Producer = {
  encode(message: Producer, writer: Writer = Writer.create()): Writer {
    if (message.producerId !== 0) {
      writer.uint32(8).uint64(message.producerId);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
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
          message.name = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
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
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
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
    message.name !== undefined && (obj.name = message.name);
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
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
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

const baseUser: object = { userId: 0, name: "", address: "", status: "" };

export const User = {
  encode(message: User, writer: Writer = Writer.create()): Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
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
          message.name = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
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
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
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
    message.name !== undefined && (obj.name = message.name);
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
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
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
    for (const v of message.users) {
      User.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.producers) {
      Producer.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Whitelist {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhitelist } as Whitelist;
    message.users = [];
    message.producers = [];
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
          message.users.push(User.decode(reader, reader.uint32()));
          break;
        case 4:
          message.producers.push(Producer.decode(reader, reader.uint32()));
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
    message.users = [];
    message.producers = [];
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
      for (const e of object.users) {
        message.users.push(User.fromJSON(e));
      }
    }
    if (object.producers !== undefined && object.producers !== null) {
      for (const e of object.producers) {
        message.producers.push(Producer.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Whitelist): unknown {
    const obj: any = {};
    message.whitelistId !== undefined &&
      (obj.whitelistId = message.whitelistId);
    message.governor !== undefined && (obj.governor = message.governor);
    if (message.users) {
      obj.users = message.users.map((e) => (e ? User.toJSON(e) : undefined));
    } else {
      obj.users = [];
    }
    if (message.producers) {
      obj.producers = message.producers.map((e) =>
        e ? Producer.toJSON(e) : undefined
      );
    } else {
      obj.producers = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Whitelist>): Whitelist {
    const message = { ...baseWhitelist } as Whitelist;
    message.users = [];
    message.producers = [];
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
      for (const e of object.users) {
        message.users.push(User.fromPartial(e));
      }
    }
    if (object.producers !== undefined && object.producers !== null) {
      for (const e of object.producers) {
        message.producers.push(Producer.fromPartial(e));
      }
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
