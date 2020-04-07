package schema

import (
	"errors"
	"github.com/Nathan-CY/shortlink/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ErrDataExist = errors.New("data exist")

func IsDataExistError(err error) bool {
	return err == ErrDataExist
}

func InitDB(config *config.DBConf) *gorm.DB {
	db, err := gorm.Open(config.Type, config.DriveAddr)
	if err != nil {
		panic("Database connection failed")
	}
	return db
}

func TxAction(db *gorm.DB, txContext func(tx *gorm.DB) error) error {
	tx := db.Begin()
	err := txContext(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
