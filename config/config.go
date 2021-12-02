package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	MySqlDNS      string
	MySqlHost     string `env:"MYSQL_HOST,required"`
	MySqlPort     string `env:"MYSQL_PORT,required"`
	MySqlUser     string `env:"MYSQL_USER,required"`
	MySqlDBName   string `env:"MYSQL_DB_NAME,required"`
	MySqlPassword string `env:"MYSQL_PASSWORD,required"`
	MySqlSSLMode  string `env:"MYSQL_SSL_MODE,required"`

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
	fmt.Println(prettyPrint(cfg))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func (c *Config) buildPgURL() {
	c.MySqlDNS = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.MySqlUser, c.MySqlPassword, c.MySqlHost, c.MySqlPort, c.MySqlDBName)
}
