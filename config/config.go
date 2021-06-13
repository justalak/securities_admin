package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/url"
)

type ApplicationConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	Host     string
	GRPCPort string
	HTTPPort string
}

func (c DatabaseConfig) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Username,
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.Database)
}

func (c DatabaseConfig) String() string {
	return fmt.Sprintf("mysql://%s", c.DNS())
}

func LoadConfig() *ApplicationConfig {
	yamlFile, err := ioutil.ReadFile("default.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
		return DefaultConfig()
	}

	conf := &ApplicationConfig{}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return DefaultConfig()
	}

	return conf
}

func DefaultConfig() *ApplicationConfig {
	return &ApplicationConfig{
		Database: DatabaseConfig{
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "mysql",
			Database: "stocks",
		},
		Server: ServerConfig{
			Host:     "localhost",
			GRPCPort: "10443",
			HTTPPort: "8080",
		},
	}
}
