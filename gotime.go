package gotime

import (
	"time"
)

type UserID int
type TimeEntryID int

type User struct {
	ID         UserID     `gorm:"primary_key", json:"id"`
	GoogleID    string    `sql:"type:char(40)"`
	Email       string    `sql:"type:char(255)"`
	LastLogin   time.Time `sql:"type:bigint"`
	IsDisabled  bool
	IsAdmin     bool
	DisplayName string    `sql:"type: char(150)"`
	GivenName   string    `sql:"type: char(100)"`
	FamilyName  string    `sql:"type: char(150)"`
	Picture     string    `sql:"type: char(200)"`
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
