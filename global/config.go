package global

import (
	"os"
	"path/filepath"
	"time"

	_ "github.com/dxvgef/filter"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// 全局配置
var Config struct {
	IP              string        `yaml:"ip"`
	Debug           bool          `yaml:"debug"`
	QuitWaitTimeout time.Duration `yaml:"quitWaitTimeout"`
	HTTP            struct {
		Port              uint          `yaml:"port"`
		ReadTimeout       time.Duration `yaml:"readTimeout"`
		ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
		WriteTimeout      time.Duration `yaml:"writeTimeout"`
		IdleTimeout       time.Duration `yaml:"idleTimeout"`
	} `yaml:"http"`
	HTTPS struct {
		Port              uint          `yaml:"port"`
		HTTP2             bool          `yaml:"http2"`
		CertFile          string        `yaml:"certFile"`
		KeyFile           string        `yaml:"keyFile"`
		ReadTimeout       time.Duration `yaml:"readTimeout"`
		ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
		WriteTimeout      time.Duration `yaml:"writeTimeout"`
		IdleTimeout       time.Duration `yaml:"idleTimeout"`
	} `yaml:"https"`
	Logger struct {
		Level      string      `yaml:"level"`
		FilePath   string      `yaml:"filePath"`
		FileMode   os.FileMode `yaml:"fileMode"`
		Encode     string      `yaml:"encode"`
		TimeFormat string      `yaml:"timeFormat"`
	} `yaml:"logger"`
	Source struct {
		Name   string `json:"name"`
		Config string `json:"config"`
	} `yaml:"source"`
}

// 加载配置文件
func LoadConfigFile(configPath string) error {
	file, err := os.Open(filepath.Clean(configPath))
	if err != nil {
		return err
	}
	err = yaml.NewDecoder(file).Decode(&Config)
	if err != nil {
		return err
	}

	if Config.HTTP.Port == 0 {
		Config.HTTP.Port = 80
	}
	if Config.HTTPS.CertFile != "" &&
		Config.HTTPS.KeyFile != "" &&
		Config.HTTPS.Port == 0 {
		Config.HTTPS.Port = 443
	}
	if Config.Logger.Level != "empty" &&
		Config.Logger.Level != "debug" &&
		Config.Logger.Level != "info" &&
		Config.Logger.Level != "warn" &&
		Config.Logger.Level != "error" {
		Config.Logger.Level = "debug"
	}
	if Config.Logger.Encode != "console" &&
		Config.Logger.Encode != "json" {
		Config.Logger.Encode = "console"
	}
	if Config.Logger.TimeFormat == "" {
		Config.Logger.TimeFormat = "y-m-d h:i:s"
	}
	log.Debug().Interface("config", Config.Source.Config)
	// var endpoints []string
	// for k := range Config.Etcd.Endpoints {
	// 	if _, err = url.Parse(Config.Etcd.Endpoints[k]); err == nil {
	// 		endpoints = append(endpoints, Config.Etcd.Endpoints[k])
	// 	}
	// }
	// if len(endpoints) == 0 {
	// 	return errors.New("至少要配置一个有效的etcd的端点")
	// }
	// Config.Etcd.Endpoints = endpoints
	return nil
}
