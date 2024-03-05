package wrap

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"runtime/debug"
)

func WithErrorf(format string, args ...interface{}) error {
	debug.PrintStack()
	log.Errorf(format, args...)
	return fmt.Errorf(format, args...)
}

func IsDBError(db *gorm.DB) error {
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return fmt.Errorf("rows affected is zero")
	}

	return nil
}
