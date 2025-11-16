package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConfig holds database connection info
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// AppConfig holds all config for the application
type AppConfig struct {
	ServerPort int
	Database   DBConfig
	RedisAddr  string
}

var DB *gorm.DB

// LoadConfig reads config.yaml and returns AppConfig
func LoadConfig() *AppConfig {
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return &AppConfig{
		ServerPort: viper.GetInt("server.port"),
		Database: DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
		RedisAddr: viper.GetString("redis.addr"),
	}
}

func ConnectDB() {
	cfg := LoadConfig()

	// Build DSN for Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = db
	fmt.Println("Database connected successfully")
}
