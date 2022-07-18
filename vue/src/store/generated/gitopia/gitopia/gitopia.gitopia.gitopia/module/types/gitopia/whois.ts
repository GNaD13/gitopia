/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "gitopia.gitopia.gitopia";

export interface Whois {
  creator: string;
  name: string;
  address: string;
  ownerType: Whois_OwnerType;
}

export enum Whois_OwnerType {
  USER = 0,
  ORGANIZATION = 1,
  UNRECOGNIZED = -1,
}

export function whois_OwnerTypeFromJSON(object: any): Whois_OwnerType {
  switch (object) {
    case 0:
    case "USER":
      return Whois_OwnerType.USER;
    case 1:
    case "ORGANIZATION":
      return Whois_OwnerType.ORGANIZATION;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Whois_OwnerType.UNRECOGNIZED;
  }
}

export function whois_OwnerTypeToJSON(object: Whois_OwnerType): string {
  switch (object) {
    case Whois_OwnerType.USER:
      return "USER";
    case Whois_OwnerType.ORGANIZATION:
      return "ORGANIZATION";
    default:
      return "UNKNOWN";
  }
}

const baseWhois: object = { creator: "", name: "", address: "", ownerType: 0 };

export const Whois = {
  encode(message: Whois, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.ownerType !== 0) {
      writer.uint32(32).int32(message.ownerType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Whois {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhois } as Whois;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.ownerType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Whois {
    const message = { ...baseWhois } as Whois;
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
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.ownerType !== undefined && object.ownerType !== null) {
      message.ownerType = whois_OwnerTypeFromJSON(object.ownerType);
    } else {
      message.ownerType = 0;
    }
    return message;
  },

  toJSON(message: Whois): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.address !== undefined && (obj.address = message.address);
    message.ownerType !== undefined &&
      (obj.ownerType = whois_OwnerTypeToJSON(message.ownerType));
    return obj;
  },

  fromPartial(object: DeepPartial<Whois>): Whois {
    const message = { ...baseWhois } as Whois;
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
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.ownerType !== undefined && object.ownerType !== null) {
      message.ownerType = object.ownerType;
    } else {
      message.ownerType = 0;
    }
    return message;
  },
};

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
