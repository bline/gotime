// package: api
// file: timesheet.proto

import * as timesheet_pb from "./timesheet_pb";
import * as common_pb from "./common_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
export class TimeSheet {
  static serviceName = "api.TimeSheet";
}
export namespace TimeSheet {
  export class ClockIn {
    static readonly methodName = "ClockIn";
    static readonly service = TimeSheet;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = timesheet_pb.ClockRequest;
    static readonly responseType = common_pb.SimpleResponse;
  }
  export class ClockOut {
    static readonly methodName = "ClockOut";
    static readonly service = TimeSheet;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = timesheet_pb.ClockRequest;
    static readonly responseType = common_pb.SimpleResponse;
  }
  export class GetCurrentStatus {
    static readonly methodName = "GetCurrentStatus";
    static readonly service = TimeSheet;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = timesheet_pb.ClockRequest;
    static readonly responseType = timesheet_pb.TSStatusResponse;
  }
  export class GetEntries {
    static readonly methodName = "GetEntries";
    static readonly service = TimeSheet;
    static readonly requestStream = false;
    static readonly responseStream = true;
    static readonly requestType = timesheet_pb.TimeSheetRequest;
    static readonly responseType = timesheet_pb.TimeEntry;
  }
}
