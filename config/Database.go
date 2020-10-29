package config

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB 
// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
	 Host:     "host.docker.internal",
	 Port:     3306,
	 User:     os.Getenv("MYSQL_USER"),
	 Password: os.Getenv("MYSQL_PASSWORD"),
	 DBName:   os.Getenv("MYSQL_DATABASE"),
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
	 "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	 dbConfig.User,
	 dbConfig.Password,
	 dbConfig.Host,
	 dbConfig.Port,
	 dbConfig.DBName,
	)
}