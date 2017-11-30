package gotime

import (
	"context"
	"github.com/bline/gotime/api/proto"
	"github.com/golang/protobuf/ptypes"
	"time"
	"log"
)


type TimeSheetService struct {
}

func (tss *TimeSheetService) ClockIn(ctx context.Context, _ *api.ClockRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "Clock-In Success"}, nil
}

func (tss *TimeSheetService) ClockOut(context.Context, *api.ClockRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "Clock-Out Success"}, nil

}
func (tss *TimeSheetService) GetCurrentStatus(context.Context, *api.ClockRequest) (*api.TSStatusResponse, error) {
	ts, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatal("Bad time")
	}
	return &api.TSStatusResponse{State: 1, Timestamp: ts}, nil
}

func (tss *TimeSheetService) GetEntries(*api.TimeSheetRequest, api.TimeSheet_GetEntriesServer) error {
	return nil
}
