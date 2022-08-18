/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "crow.marketplace";

export interface MsgCreateListing {
  creator: string;
  title: string;
  description: string;
  minPrice: string;
  producerId: number;
}

export interface MsgCreateListingResponse {
  listingId: number;
}

export interface MsgCreateOrder {
  creator: string;
  listingId: number;
  userId: number;
  maxPrice: string;
}

export interface MsgCreateOrderResponse {
  orderId: number;
}

const baseMsgCreateListing: object = {
  creator: "",
  title: "",
  description: "",
  minPrice: "",
  producerId: 0,
};

export const MsgCreateListing = {
  encode(message: MsgCreateListing, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.minPrice !== "") {
      writer.uint32(34).string(message.minPrice);
    }
    if (message.producerId !== 0) {
      writer.uint32(40).uint64(message.producerId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateListing {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateListing } as MsgCreateListing;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.minPrice = reader.string();
          break;
        case 5:
          message.producerId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateListing {
    const message = { ...baseMsgCreateListing } as MsgCreateListing;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.minPrice !== undefined && object.minPrice !== null) {
      message.minPrice = String(object.minPrice);
    } else {
      message.minPrice = "";
    }
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = Number(object.producerId);
    } else {
      message.producerId = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateListing): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.minPrice !== undefined && (obj.minPrice = message.minPrice);
    message.producerId !== undefined && (obj.producerId = message.producerId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateListing>): MsgCreateListing {
    const message = { ...baseMsgCreateListing } as MsgCreateListing;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.minPrice !== undefined && object.minPrice !== null) {
      message.minPrice = object.minPrice;
    } else {
      message.minPrice = "";
    }
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = object.producerId;
    } else {
      message.producerId = 0;
    }
    return message;
  },
};

const baseMsgCreateListingResponse: object = { listingId: 0 };

export const MsgCreateListingResponse = {
  encode(
    message: MsgCreateListingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.listingId !== 0) {
      writer.uint32(8).uint64(message.listingId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateListingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateListingResponse,
    } as MsgCreateListingResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.listingId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateListingResponse {
    const message = {
      ...baseMsgCreateListingResponse,
    } as MsgCreateListingResponse;
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = Number(object.listingId);
    } else {
      message.listingId = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateListingResponse): unknown {
    const obj: any = {};
    message.listingId !== undefined && (obj.listingId = message.listingId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateListingResponse>
  ): MsgCreateListingResponse {
    const message = {
      ...baseMsgCreateListingResponse,
    } as MsgCreateListingResponse;
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = object.listingId;
    } else {
      message.listingId = 0;
    }
    return message;
  },
};

const baseMsgCreateOrder: object = {
  creator: "",
  listingId: 0,
  userId: 0,
  maxPrice: "",
};

export const MsgCreateOrder = {
  encode(message: MsgCreateOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.listingId !== 0) {
      writer.uint32(16).uint64(message.listingId);
    }
    if (message.userId !== 0) {
      writer.uint32(24).uint64(message.userId);
    }
    if (message.maxPrice !== "") {
      writer.uint32(34).string(message.maxPrice);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.listingId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.userId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.maxPrice = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateOrder {
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = Number(object.listingId);
    } else {
      message.listingId = 0;
    }
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = Number(object.userId);
    } else {
      message.userId = 0;
    }
    if (object.maxPrice !== undefined && object.maxPrice !== null) {
      message.maxPrice = String(object.maxPrice);
    } else {
      message.maxPrice = "";
    }
    return message;
  },

  toJSON(message: MsgCreateOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.listingId !== undefined && (obj.listingId = message.listingId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.maxPrice !== undefined && (obj.maxPrice = message.maxPrice);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateOrder>): MsgCreateOrder {
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = object.listingId;
    } else {
      message.listingId = 0;
    }
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = object.userId;
    } else {
      message.userId = 0;
    }
    if (object.maxPrice !== undefined && object.maxPrice !== null) {
      message.maxPrice = object.maxPrice;
    } else {
      message.maxPrice = "";
    }
    return message;
  },
};

const baseMsgCreateOrderResponse: object = { orderId: 0 };

export const MsgCreateOrderResponse = {
  encode(
    message: MsgCreateOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.orderId !== 0) {
      writer.uint32(8).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateOrderResponse {
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateOrderResponse): unknown {
    const obj: any = {};
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateOrderResponse>
  ): MsgCreateOrderResponse {
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateListing(request: MsgCreateListing): Promise<MsgCreateListingResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateOrder(request: MsgCreateOrder): Promise<MsgCreateOrderResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateListing(request: MsgCreateListing): Promise<MsgCreateListingResponse> {
    const data = MsgCreateListing.encode(request).finish();
    const promise = this.rpc.request(
      "crow.marketplace.Msg",
      "CreateListing",
      data
    );
    return promise.then((data) =>
      MsgCreateListingResponse.decode(new Reader(data))
    );
  }

  CreateOrder(request: MsgCreateOrder): Promise<MsgCreateOrderResponse> {
    const data = MsgCreateOrder.encode(request).finish();
    const promise = this.rpc.request(
      "crow.marketplace.Msg",
      "CreateOrder",
      data
    );
    return promise.then((data) =>
      MsgCreateOrderResponse.decode(new Reader(data))
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
