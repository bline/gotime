
package client

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	"github.com/bline/gotime/config"

	"github.com/bline/gotime/api/proto/timesheet.pb"
)

func ClockIn() error {
}

func ClockOut() error {
}

func GetCurrentStatus() (*api.TSStatusResponse, error) {
}

func GetEntries(req *TimeSheetRequest) error (
}

