// package: api
// file: user.proto

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class EmailOrId extends jspb.Message {
  hasEmail(): boolean;
  clearEmail(): void;
  getEmail(): string;
  setEmail(value: string): void;

  hasUserid(): boolean;
  clearUserid(): void;
  getUserid(): number;
  setUserid(value: number): void;

  getUserIdentCase(): EmailOrId.UserIdentCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmailOrId.AsObject;
  static toObject(includeInstance: boolean, msg: EmailOrId): EmailOrId.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EmailOrId, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmailOrId;
  static deserializeBinaryFromReader(message: EmailOrId, reader: jspb.BinaryReader): EmailOrId;
}

export namespace EmailOrId {
  export type AsObject = {
    email: string,
    userid: number,
  }

  export enum UserIdentCase {
    USERIDENT_NOT_SET = 0,
    EMAIL = 1,
    USERID = 2,
  }
}

export class LockUserRequest extends jspb.Message {
  hasIdent(): boolean;
  clearIdent(): void;
  getIdent(): EmailOrId | undefined;
  setIdent(value?: EmailOrId): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LockUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LockUserRequest): LockUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LockUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LockUserRequest;
  static deserializeBinaryFromReader(message: LockUserRequest, reader: jspb.BinaryReader): LockUserRequest;
}

export namespace LockUserRequest {
  export type AsObject = {
    ident?: EmailOrId.AsObject,
  }
}

export class DisableUserRequest extends jspb.Message {
  hasIdent(): boolean;
  clearIdent(): void;
  getIdent(): EmailOrId | undefined;
  setIdent(value?: EmailOrId): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DisableUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DisableUserRequest): DisableUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DisableUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DisableUserRequest;
  static deserializeBinaryFromReader(message: DisableUserRequest, reader: jspb.BinaryReader): DisableUserRequest;
}

export namespace DisableUserRequest {
  export type AsObject = {
    ident?: EmailOrId.AsObject,
  }
}

export class DeleteUserRequest extends jspb.Message {
  hasIdent(): boolean;
  clearIdent(): void;
  getIdent(): EmailOrId | undefined;
  setIdent(value?: EmailOrId): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserRequest): DeleteUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserRequest;
  static deserializeBinaryFromReader(message: DeleteUserRequest, reader: jspb.BinaryReader): DeleteUserRequest;
}

export namespace DeleteUserRequest {
  export type AsObject = {
    ident?: EmailOrId.AsObject,
  }
}

export class DeleteUserResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserResponse): DeleteUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserResponse;
  static deserializeBinaryFromReader(message: DeleteUserResponse, reader: jspb.BinaryReader): DeleteUserResponse;
}

export namespace DeleteUserResponse {
  export type AsObject = {
    success: boolean,
    message: string,
  }
}

export class GetUserRequest extends jspb.Message {
  hasIdent(): boolean;
  clearIdent(): void;
  getIdent(): EmailOrId | undefined;
  setIdent(value?: EmailOrId): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRequest;
  static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    ident?: EmailOrId.AsObject,
  }
}

export class GetUsersRequest extends jspb.Message {
  hasSearchoptions(): boolean;
  clearSearchoptions(): void;
  getSearchoptions(): common_pb.SearchOptions | undefined;
  setSearchoptions(value?: common_pb.SearchOptions): void;

  getIsdeleted(): boolean;
  setIsdeleted(value: boolean): void;

  getIsdisabled(): boolean;
  setIsdisabled(value: boolean): void;

  getIslockedout(): boolean;
  setIslockedout(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUsersRequest): GetUsersRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUsersRequest;
  static deserializeBinaryFromReader(message: GetUsersRequest, reader: jspb.BinaryReader): GetUsersRequest;
}

export namespace GetUsersRequest {
  export type AsObject = {
    searchoptions?: common_pb.SearchOptions.AsObject,
    isdeleted: boolean,
    isdisabled: boolean,
    islockedout: boolean,
  }
}

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getEmail(): string;
  setEmail(value: string): void;

  getToken(): string;
  setToken(value: string): void;

  hasLastlogin(): boolean;
  clearLastlogin(): void;
  getLastlogin(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastlogin(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasLockedout(): boolean;
  clearLockedout(): void;
  getLockedout(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLockedout(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getIsdisabled(): boolean;
  setIsdisabled(value: boolean): void;

  getIsdeleted(): boolean;
  setIsdeleted(value: boolean): void;

  getIslockedout(): boolean;
  setIslockedout(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    email: string,
    token: string,
    lastlogin?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lockedout?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    isdisabled: boolean,
    isdeleted: boolean,
    islockedout: boolean,
  }
}

