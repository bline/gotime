// package: api
// file: common.proto

import * as jspb from "google-protobuf";

export class SimpleResponse extends jspb.Message {
  getIssuccess(): boolean;
  setIssuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SimpleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SimpleResponse): SimpleResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SimpleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SimpleResponse;
  static deserializeBinaryFromReader(message: SimpleResponse, reader: jspb.BinaryReader): SimpleResponse;
}

export namespace SimpleResponse {
  export type AsObject = {
    issuccess: boolean,
    message: string,
  }
}

export class SearchOptions extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): void;

  getPage(): number;
  setPage(value: number): void;

  getPerpage(): number;
  setPerpage(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchOptions.AsObject;
  static toObject(includeInstance: boolean, msg: SearchOptions): SearchOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SearchOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchOptions;
  static deserializeBinaryFromReader(message: SearchOptions, reader: jspb.BinaryReader): SearchOptions;
}

export namespace SearchOptions {
  export type AsObject = {
    query: string,
    page: number,
    perpage: number,
  }
}

