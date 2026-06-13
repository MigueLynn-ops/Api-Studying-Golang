package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Port      string
	Cfg       mysql.Config
	SecretKey []byte
)

func LoadEnv() {
	// Carregar o arquivo .env localizado na mesma pasta do config.go
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not determine the path of the config file")
	}

	configDir := path.Dir(filePath)
	envPath := filepath.Join(configDir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Warning: .env not found at %s: %s", envPath, err)
	}

	log.Println("Environment variables loaded successfully from", envPath)

	Port = os.Getenv("API_PORT")
	Cfg = mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
	}
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
