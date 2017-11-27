// package: api
// file: timesheet.proto

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";

export class ClockRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClockRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ClockRequest): ClockRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClockRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClockRequest;
  static deserializeBinaryFromReader(message: ClockRequest, reader: jspb.BinaryReader): ClockRequest;
}

export namespace ClockRequest {
  export type AsObject = {
  }
}

export class TimeSheetRequest extends jspb.Message {
  hasOptions(): boolean;
  clearOptions(): void;
  getOptions(): common_pb.SearchOptions | undefined;
  setOptions(value?: common_pb.SearchOptions): void;

  hasStarttime(): boolean;
  clearStarttime(): void;
  getStarttime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStarttime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasEndtime(): boolean;
  clearEndtime(): void;
  getEndtime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEndtime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasEmail(): boolean;
  clearEmail(): void;
  getEmail(): string;
  setEmail(value: string): void;

  hasUserid(): boolean;
  clearUserid(): void;
  getUserid(): number;
  setUserid(value: number): void;

  getUserIdentCase(): TimeSheetRequest.UserIdentCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TimeSheetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TimeSheetRequest): TimeSheetRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TimeSheetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TimeSheetRequest;
  static deserializeBinaryFromReader(message: TimeSheetRequest, reader: jspb.BinaryReader): TimeSheetRequest;
}

export namespace TimeSheetRequest {
  export type AsObject = {
    options?: common_pb.SearchOptions.AsObject,
    starttime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    endtime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    email: string,
    userid: number,
  }

  export enum UserIdentCase {
    USERIDENT_NOT_SET = 0,
    EMAIL = 4,
    USERID = 5,
  }
}

export class TSStatusResponse extends jspb.Message {
  getState(): TSStatusResponse.Status;
  setState(value: TSStatusResponse.Status): void;

  hasTimestamp(): boolean;
  clearTimestamp(): void;
  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasLastworked(): boolean;
  clearLastworked(): void;
  getLastworked(): google_protobuf_duration_pb.Duration | undefined;
  setLastworked(value?: google_protobuf_duration_pb.Duration): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TSStatusResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TSStatusResponse): TSStatusResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TSStatusResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TSStatusResponse;
  static deserializeBinaryFromReader(message: TSStatusResponse, reader: jspb.BinaryReader): TSStatusResponse;
}

export namespace TSStatusResponse {
  export type AsObject = {
    state: TSStatusResponse.Status,
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastworked?: google_protobuf_duration_pb.Duration.AsObject,
  }

  export enum Status {
    ClockedOut = 0,
    ClockedIn = 1,
  }
}

export class TimeEntry extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getUserid(): string;
  setUserid(value: string): void;

  hasTimestamp(): boolean;
  clearTimestamp(): void;
  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getState(): TimeEntry.Status;
  setState(value: TimeEntry.Status): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TimeEntry.AsObject;
  static toObject(includeInstance: boolean, msg: TimeEntry): TimeEntry.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TimeEntry, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TimeEntry;
  static deserializeBinaryFromReader(message: TimeEntry, reader: jspb.BinaryReader): TimeEntry;
}

export namespace TimeEntry {
  export type AsObject = {
    id: number,
    userid: string,
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    state: TimeEntry.Status,
  }

  export enum Status {
    ClockedOut = 0,
    ClockedIn = 1,
  }
}

