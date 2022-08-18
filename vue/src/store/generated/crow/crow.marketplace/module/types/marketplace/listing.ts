/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crow.marketplace";

export interface Order {
  orderId: number;
  listingId: number;
  userId: number;
  maxPrice: string;
  status: string;
}

export interface Listing {
  listingId: number;
  producerId: number;
  escrowId: number;
  title: string;
  description: string;
  minPrice: string;
  status: string;
  ordersRecieved: Order[];
}

const baseOrder: object = {
  orderId: 0,
  listingId: 0,
  userId: 0,
  maxPrice: "",
  status: "",
};

export const Order = {
  encode(message: Order, writer: Writer = Writer.create()): Writer {
    if (message.orderId !== 0) {
      writer.uint32(8).uint64(message.orderId);
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
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrder } as Order;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orderId = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): Order {
    const message = { ...baseOrder } as Order;
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.orderId !== undefined && (obj.orderId = message.orderId);
    message.listingId !== undefined && (obj.listingId = message.listingId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.maxPrice !== undefined && (obj.maxPrice = message.maxPrice);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<Order>): Order {
    const message = { ...baseOrder } as Order;
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    return message;
  },
};

const baseListing: object = {
  listingId: 0,
  producerId: 0,
  escrowId: 0,
  title: "",
  description: "",
  minPrice: "",
  status: "",
};

export const Listing = {
  encode(message: Listing, writer: Writer = Writer.create()): Writer {
    if (message.listingId !== 0) {
      writer.uint32(8).uint64(message.listingId);
    }
    if (message.producerId !== 0) {
      writer.uint32(16).uint64(message.producerId);
    }
    if (message.escrowId !== 0) {
      writer.uint32(24).uint64(message.escrowId);
    }
    if (message.title !== "") {
      writer.uint32(34).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.minPrice !== "") {
      writer.uint32(50).string(message.minPrice);
    }
    if (message.status !== "") {
      writer.uint32(58).string(message.status);
    }
    for (const v of message.ordersRecieved) {
      Order.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Listing {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseListing } as Listing;
    message.ordersRecieved = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.listingId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.producerId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.escrowId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.title = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.minPrice = reader.string();
          break;
        case 7:
          message.status = reader.string();
          break;
        case 8:
          message.ordersRecieved.push(Order.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Listing {
    const message = { ...baseListing } as Listing;
    message.ordersRecieved = [];
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = Number(object.listingId);
    } else {
      message.listingId = 0;
    }
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = Number(object.producerId);
    } else {
      message.producerId = 0;
    }
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = Number(object.escrowId);
    } else {
      message.escrowId = 0;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.ordersRecieved !== undefined && object.ordersRecieved !== null) {
      for (const e of object.ordersRecieved) {
        message.ordersRecieved.push(Order.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Listing): unknown {
    const obj: any = {};
    message.listingId !== undefined && (obj.listingId = message.listingId);
    message.producerId !== undefined && (obj.producerId = message.producerId);
    message.escrowId !== undefined && (obj.escrowId = message.escrowId);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.minPrice !== undefined && (obj.minPrice = message.minPrice);
    message.status !== undefined && (obj.status = message.status);
    if (message.ordersRecieved) {
      obj.ordersRecieved = message.ordersRecieved.map((e) =>
        e ? Order.toJSON(e) : undefined
      );
    } else {
      obj.ordersRecieved = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Listing>): Listing {
    const message = { ...baseListing } as Listing;
    message.ordersRecieved = [];
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = object.listingId;
    } else {
      message.listingId = 0;
    }
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = object.producerId;
    } else {
      message.producerId = 0;
    }
    if (object.escrowId !== undefined && object.escrowId !== null) {
      message.escrowId = object.escrowId;
    } else {
      message.escrowId = 0;
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
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.ordersRecieved !== undefined && object.ordersRecieved !== null) {
      for (const e of object.ordersRecieved) {
        message.ordersRecieved.push(Order.fromPartial(e));
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
