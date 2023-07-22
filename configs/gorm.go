package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetReaderGorm() *gorm.DB {
	reader, err := gorm.Open(mysql.Open(DB_CONNECTION))

	if err != nil {
		panic("failed to connect to database")
	}

	return reader
}

func GetWriterGorm() *gorm.DB {
	reader, err := gorm.Open(mysql.Open(DB_CONNECTION))

	if err != nil {
		panic("failed to connect to database")
	}

	return reader
}
