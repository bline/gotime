// package: api
// file: user.proto

import * as user_pb from "./user_pb";
import * as common_pb from "./common_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
export class Accounts {
  static serviceName = "api.Accounts";
}
export namespace Accounts {
  export class GetUser {
    static readonly methodName = "GetUser";
    static readonly service = Accounts;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = user_pb.GetUserRequest;
    static readonly responseType = user_pb.User;
  }
  export class GetUsers {
    static readonly methodName = "GetUsers";
    static readonly service = Accounts;
    static readonly requestStream = false;
    static readonly responseStream = true;
    static readonly requestType = user_pb.GetUsersRequest;
    static readonly responseType = user_pb.User;
  }
  export class DisableUser {
    static readonly methodName = "DisableUser";
    static readonly service = Accounts;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = user_pb.DisableUserRequest;
    static readonly responseType = common_pb.SimpleResponse;
  }
  export class DeleteUser {
    static readonly methodName = "DeleteUser";
    static readonly service = Accounts;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = user_pb.DeleteUserRequest;
    static readonly responseType = common_pb.SimpleResponse;
  }
  export class LockUser {
    static readonly methodName = "LockUser";
    static readonly service = Accounts;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = user_pb.LockUserRequest;
    static readonly responseType = common_pb.SimpleResponse;
  }
}
