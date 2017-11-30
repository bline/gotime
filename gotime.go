package gotime

import (
	"time"

	"github.com/bline/gotime/api/proto"
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
