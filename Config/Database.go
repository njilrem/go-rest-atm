package config
import 
(
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

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
	 Host:     "localhost",
	 Port:     3306,
	 User:     goDotEnvVariable("MYSQL_USER"),
	 Password: goDotEnvVariable("MYSQL_PASSWORD"),
	 DBName:   goDotEnvVariable("MYSQL_DATABASE"),
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