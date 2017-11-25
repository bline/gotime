package gotime

import (
  "time"
)

type UserID int
type TimeEntryID int

type User struct {
	ID         UserID     `gorm:"primary_key"`
	Email      string     `sql:"type:varchar(255)"`
	Token      string     `sql:"type:varchar(48)"`
	LastLogin  time.Time  `sql:"type:bigint"`
	IsDisabled bool
	IsAdmin    bool
}

const (
  StateClockedOut = 0
  StateClockedIn  = 1
  StateOnBreak    = 3
)

type TimeEntry struct {
	ID         TimeEntryID     `gorm:"primary_key"`
	UserID     UserID
	Timestamp  time.Time       `sql:"type:bigint"`
	State      int
}

type Client interface {
  OAuthService() OAuthService
  UserService() UserService
  TimeSheetService() TimeSheetService
}

type TimeEntryService interface {
  TimeEntries(id UserID, startDate time.Time, endDate time.Time) ([]TimeEntry, error)
  CreateTimeEntry(timesheet *TimeEntry) error
}


type UserService interface {
	GetUser(context.Context, *api.GetUserRequest) (*api.User, error)
	GetUsers(*api.GetUsersRequest) error
	DeleteUser(*api.DeleteUserRequest) error
	DisableUser(*api.DisableUserRequest) error
	LockUser(*api.LockUserRequest) error
}

type TimeSheetService interface {
	ClockIn () error
	ClockOut () error
	GetCurrentStatus () api.TSStatusResponse
	GetEntries (r api.TimeSheetRequest) error
}

