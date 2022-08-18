/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "crow.whitelist";

export interface MsgCreateUser {
  creator: string;
  name: string;
}

export interface MsgCreateUserResponse {
  userId: string;
}

export interface MsgCreateProducer {
  creator: string;
  name: string;
}

export interface MsgCreateProducerResponse {
  producerId: number;
}

const baseMsgCreateUser: object = { creator: "", name: "" };

export const MsgCreateUser = {
  encode(message: MsgCreateUser, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateUser {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateUser } as MsgCreateUser;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateUser {
    const message = { ...baseMsgCreateUser } as MsgCreateUser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgCreateUser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateUser>): MsgCreateUser {
    const message = { ...baseMsgCreateUser } as MsgCreateUser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgCreateUserResponse: object = { userId: "" };

export const MsgCreateUserResponse = {
  encode(
    message: MsgCreateUserResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateUserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateUserResponse } as MsgCreateUserResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateUserResponse {
    const message = { ...baseMsgCreateUserResponse } as MsgCreateUserResponse;
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = String(object.userId);
    } else {
      message.userId = "";
    }
    return message;
  },

  toJSON(message: MsgCreateUserResponse): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateUserResponse>
  ): MsgCreateUserResponse {
    const message = { ...baseMsgCreateUserResponse } as MsgCreateUserResponse;
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = object.userId;
    } else {
      message.userId = "";
    }
    return message;
  },
};

const baseMsgCreateProducer: object = { creator: "", name: "" };

export const MsgCreateProducer = {
  encode(message: MsgCreateProducer, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateProducer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateProducer } as MsgCreateProducer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateProducer {
    const message = { ...baseMsgCreateProducer } as MsgCreateProducer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgCreateProducer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateProducer>): MsgCreateProducer {
    const message = { ...baseMsgCreateProducer } as MsgCreateProducer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgCreateProducerResponse: object = { producerId: 0 };

export const MsgCreateProducerResponse = {
  encode(
    message: MsgCreateProducerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.producerId !== 0) {
      writer.uint32(8).uint64(message.producerId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateProducerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateProducerResponse,
    } as MsgCreateProducerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.producerId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateProducerResponse {
    const message = {
      ...baseMsgCreateProducerResponse,
    } as MsgCreateProducerResponse;
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = Number(object.producerId);
    } else {
      message.producerId = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateProducerResponse): unknown {
    const obj: any = {};
    message.producerId !== undefined && (obj.producerId = message.producerId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateProducerResponse>
  ): MsgCreateProducerResponse {
    const message = {
      ...baseMsgCreateProducerResponse,
    } as MsgCreateProducerResponse;
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = object.producerId;
    } else {
      message.producerId = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateUser(request: MsgCreateUser): Promise<MsgCreateUserResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateProducer(
    request: MsgCreateProducer
  ): Promise<MsgCreateProducerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateUser(request: MsgCreateUser): Promise<MsgCreateUserResponse> {
    const data = MsgCreateUser.encode(request).finish();
    const promise = this.rpc.request("crow.whitelist.Msg", "CreateUser", data);
    return promise.then((data) =>
      MsgCreateUserResponse.decode(new Reader(data))
    );
  }

  CreateProducer(
    request: MsgCreateProducer
  ): Promise<MsgCreateProducerResponse> {
    const data = MsgCreateProducer.encode(request).finish();
    const promise = this.rpc.request(
      "crow.whitelist.Msg",
      "CreateProducer",
      data
    );
    return promise.then((data) =>
      MsgCreateProducerResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
