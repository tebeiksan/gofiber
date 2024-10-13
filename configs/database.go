package configs

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

type Database struct {
	*gorm.DB
}

func DatabaseNew(config *DatabaseConfig) (*Database, error) {
	var db *gorm.DB
	var err error
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &Database{db}, err
}
