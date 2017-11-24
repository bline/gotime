
import (
	"github.com/bline/gotime/db"
	"github.com/bline/gotime/gotime.go"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	db := db.GetDB()

	db.AutoMigrate(&gotime.User{}, &gotime.TimeEntry{})
	db.Model(&gotime.User{}).AddIndex("idx_email", "Email")
}

