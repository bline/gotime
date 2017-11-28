package gotime

import (
	"context"

	"github.com/bline/gotime"
	"github.com/bline/gotime/api/proto"
	"github.com/bline/gotime/db"
	"github.com/bline/gotime/config"
	"github.com/grpc/grpc-go"
)


type TimeSheetService struct {
}

func New() *TimeSheetService {
	return &TimeSheetService{}
}

func (tss *TimeSheetService) ClockIn(ctx context.Context, _ *api.ClockRequest) (api.SimpleResponse, error) {
	timeEntry := &TimeEntry{}
	timeEntry.UserID
}
