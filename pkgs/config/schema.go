package config

import (
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Auth     AuthConfig     `koanf:"auth_config"`
	DB       DBConfig       `koanf:"db_config"`
	HTTP     HTTPConfig     `koanf:"http_config"`
	Audit    AuditConfig    `koanf:"audit_config"`
	Services ServicesConfig `koanf:"services_config"`

	Log Log `koanf:"log"`
}

type Log struct {
	Level zapcore.Level `koanf:"level"`

	File    FileLog    `koanf:"file"`
	Console ConsoleLog `koanf:"console"`
}

type FileLog struct {
	Level    *zapcore.Level `koanf:"level"`
	Encoding string         `koanf:"encoding"`
	Path     string         `koanf:"path"`
}

type ConsoleLog struct {
	Level    *zapcore.Level `koanf:"level"`
	Encoding string         `koanf:"encoding"`
	Disable  bool           `koanf:"disable"`
}

type AuthConfig struct {
	AccessSecret  string `koanf:"access_secret"`
	RefreshSecret string `koanf:"refresh_secret"`
}

type DBConfig struct {
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Driver   string `koanf:"driver"`
	Name     string `koanf:"name"`
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
}

type HTTPConfig struct {
	Host       string `koanf:"host"`
	Port       string `koanf:"port"`
	ExposePort string `koanf:"expose_port"`
}

type AuditConfig struct {
	Url string `koanf:"url"`
}

type ServicesConfig struct {
	CommonTool        string `koanf:"common_tool"`
	Audit             string `koanf:"audit"`
	TransactionServer string `koanf:"transaction_server"`
}
