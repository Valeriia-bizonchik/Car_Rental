package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var JwtKey = []byte("my_secret_key")

type Config struct {

	// Database settings
	DbDNS      string
	DbHost     string `env:"DB_HOST,required"`
	DbPort     string `env:"DB_PORT,required"`
	DbUser     string `env:"DB_USER,required"`
	DbName     string `env:"DB_NAME,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbSSLMode  string `env:"DB_SSL_MODE,required"`

	ServiceName string `env:"SERVICE_NAME" envDefault:"car-rental"`
	ServiceHost string `env:"SERVICE_HOST" envDefault:""`
	ServicePort string `env:"SERVICE_PORT" envDefault:""`

	LogFilePath string   `env:"LOG_FILE_PATH" envDefault:"./service.log"`
	LogFile     *os.File `env:"-"`

	DebugMode bool `env:"DEBUG_MODE"`
}

func InitEnvConfig() (Config, error) {
	readLocalEnvFile()
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return cfg, err
	}

	f, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err != nil {
		fmt.Printf("can't open log file by path'%v' err: %+v\n", cfg.LogFilePath, err)
		return cfg, err
	}
	cfg.LogFile = f

	cfg.buildPgURL()

	printConfig(cfg)

	return cfg, nil
}

func readLocalEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Info: can't load .env file")
	}
}

func printConfig(cfg Config) {
	fmt.Println("config:")
	fmt.Println(PrettyPrint(cfg))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func (c *Config) buildPgURL() {
	//c.DNS = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	//	c.User, c.Password, c.Host, c.Port, c.DBName)
	c.DbDNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DbHost, c.DbUser, c.DbPassword, c.DbName, c.DbPort)
}
