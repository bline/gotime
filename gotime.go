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
  StateClockedOut = iota
  StateClockedIn  = iota
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
  TimeEntryService() TimeEntryService
}

type TimeEntryService interface {
  TimeEntries(id UserID, startDate time.Time, endDate time.Time) ([]TimeEntry, error)
  CreateTimeEntry(timesheet *TimeEntry) error
}

type UserService interface {
  User(id UserID) (*User, error)
  Users() ([]User, error)
  CreateUser(user *User) error

  SetToken(id UserID, token string) error
  SetIsDisabled(id UserID, isDisabled bool) error
}

type TimesheetService interface {
  Timesheet(id UserID) (*TimeEntry, error)
}

