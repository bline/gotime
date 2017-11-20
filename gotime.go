package gotime

import (
  "log"
  "time"
)

type UserID int
type TimeEntryID int

type User struct {
  ID         UserID    `json:"ID"`
  Token      string    `json:"token"`
  Email      string    `json:"email"`
  LastLogin  time.Time `json:"lastLogin"`
  IsDisabled bool      `json:"isDisabled"`
}

const (
  ClockedOut = iota
  ClockedIn  = iota
)

type TimeEntry struct {
  ID         TimeEntryID `json:"ID"`
  UserID     UserID      `json:"userID"`
  Timestamp  time.Time   `json:"timestamp"`
  State      int         `json:"state"`
}

type Client interface {
  UserService() UserService
  TimeEntryService() TimeEntryService
}

type TimeEntryService interface {
  TimeEntries(id UserID, startDate time.Time, endDate time.Time) ([]Timesheet, error)
  CreateTimeEntry(timesheet *Timesheet) error
}

type UserService interface {
  User(id UserID) (*User, error)
  CreateUser(user *User) error

  SetToken(id UserID, token string) error
  SetIsDisabled(id UserID, isDisabled bool) error
}

type TimesheetService interface {
  Timesheet(id UserID) (*Timesheet, error)
}

