package gotime

import (
	"time"
	"github.com/jinzhu/gorm"
)

type UserID uint

type User struct {
	gorm.Model
	GoogleID    string      `sql:"type:char(40)" json:"google_id"`
	Email       string      `sql:"type:char(255)" json:"email"`
	DisplayName string      `sql:"type: char(150)" json:"display_name"`
	GivenName   string      `sql:"type: char(100)" json:"given_name"`
	FamilyName  string      `sql:"type: char(150)" json:"family_name"`
	Picture     string      `sql:"type: char(200)" json:"picture"`
	LastLogin   time.Time   `sql:"type:bigint" json:"last_login"`
	IsAdmin     bool        `json:"is_admin"`
	TimeEntries []TimeEntry `gorm:"ForeignKey:UserID" json:"-"`
}

const (
	StateClockedOut = 0
	StateClockedIn  = 1
	StateOnBreak    = 3
)

type TimeEntry struct {
	gorm.Model
	UserID     UserID          `json:"user_id"`
	Timestamp  time.Time       `sql:"type:bigint"`
	State      int
}
