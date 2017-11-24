package gotime

import (
  "time"
)

type UserID int
type TimeEntryID int

type User struct {
  ID         UserID
  Email      string
  Token      string
  LastLogin  time.Time
  IsDisabled bool
}

const (
  StateClockedOut = iota
  StateClockedIn  = iota
)

type TimeEntry struct {
  ID         TimeEntryID
  UserID     UserID
  Timestamp  time.Time
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

