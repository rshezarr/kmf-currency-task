package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"os"
)

type Configs struct {
	Port           string       `json:"port"`
	Host           string       `json:"host"`
	MaxHeaderBytes int          `json:"max_header_bytes"`
	ReadTimeout    int          `json:"read_timeout"`
	WriteTimeout   int          `json:"write_timeout"`
	IdleTimeout    int          `json:"idle_timeout"`
	DB             DBConfig     `json:"db"`
	NationalBank   NationalBank `json:"national_bank"`
}

type DBConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	Name       string `json:"name"`
	DriverName string `json:"driver"`
	OpenConns  int    `json:"open_conns"`
	IdleConns  int    `json:"idle_conns"`
}

func (dbc *DBConfig) GetConnectionString() string {
	// Если будет решение перейти на другую БД, добавить новые кейсы
	switch dbc.DriverName {
	case "sqlserver":
		user := url.QueryEscape(dbc.User)
		pass := url.QueryEscape(dbc.Pass)
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable", user, pass, dbc.Host, dbc.Port, dbc.Name)
	default:
		return ""
	}
}

type NationalBank struct {
	BaseUrl     string `json:"base_url"`
	RatesMethod string `json:"rates_method"`
}

func (n *NationalBank) GetFullURL() string {
	return n.BaseUrl + n.RatesMethod
}

var Conf = &Configs{}

func LoadLocalConfig(config interface{}) error {
	var path string
	flag.StringVar(&path, "config", "./config/conf.json", "Config path")
	flag.Parse()

	zap.S().Infof("config: %s", path)

	return LoadConfig(path, config)
}

// LoadConfig needed for external things, like tests, when we want to specify the config path manually.
func LoadConfig(path string, config interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		zap.S().Fatalf("Can not open the config file: %s(%v)", path, err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zap.S().Error("error closing body: " + err.Error())
		}
	}(file)

	return nil
}
