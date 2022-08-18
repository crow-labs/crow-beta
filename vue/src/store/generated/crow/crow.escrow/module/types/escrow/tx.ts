/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "crow.escrow";

export interface MsgItemReceived {
  userAddress: string;
  escrowId: number;
}

export interface MsgItemReceivedResponse {}

const baseMsgItemReceived: object = { userAddress: "", escrowId: 0 };

export const MsgItemReceived = {
  encode(message: MsgItemReceived, writer: Writer = Writer.create()): Writer {
    if (message.userAddress !== "") {
      writer.uint32(10).string(message.userAddress);
    }
    if (message.escrowId !== 0) {
      writer.uint32(16).uint64(message.escrowId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgItemReceived {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgItemReceived } as MsgItemReceived;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAddress = reader.string();
          break;
        case 2:
          message.escrowId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgItemReceived {
    const message = { ...baseMsgItemReceived } as MsgItemReceived;
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = String(object.userAddress);
    } else {
      message.userAddress = "";
    }
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = Number(object.escrowId);
    } else {
      message.escrowId = 0;
    }
    return message;
  },

  toJSON(message: MsgItemReceived): unknown {
    const obj: any = {};
    message.userAddress !== undefined &&
      (obj.userAddress = message.userAddress);
    message.escrowId !== undefined && (obj.escrowId = message.escrowId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgItemReceived>): MsgItemReceived {
    const message = { ...baseMsgItemReceived } as MsgItemReceived;
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = object.userAddress;
    } else {
      message.userAddress = "";
    }
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = object.escrowId;
    } else {
      message.escrowId = 0;
    }
    return message;
  },
};

const baseMsgItemReceivedResponse: object = {};

export const MsgItemReceivedResponse = {
  encode(_: MsgItemReceivedResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgItemReceivedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgItemReceivedResponse,
    } as MsgItemReceivedResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgItemReceivedResponse {
    const message = {
      ...baseMsgItemReceivedResponse,
    } as MsgItemReceivedResponse;
    return message;
  },

  toJSON(_: MsgItemReceivedResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgItemReceivedResponse>
  ): MsgItemReceivedResponse {
    const message = {
      ...baseMsgItemReceivedResponse,
    } as MsgItemReceivedResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ItemReceived(request: MsgItemReceived): Promise<MsgItemReceivedResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ItemReceived(request: MsgItemReceived): Promise<MsgItemReceivedResponse> {
    const data = MsgItemReceived.encode(request).finish();
    const promise = this.rpc.request("crow.escrow.Msg", "ItemReceived", data);
    return promise.then((data) =>
      MsgItemReceivedResponse.decode(new Reader(data))
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
