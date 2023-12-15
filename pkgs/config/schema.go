package config

import (
	"fmt"
	"net/url"

	"github.com/go-webauthn/webauthn/protocol"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Auth  AuthConfig  `koanf:"auth_config"`
	DB    DBConfig    `koanf:"db_config"`
	HTTP  HTTPConfig  `koanf:"http_config"`
	Audit AuditConfig `koanf:"audit_config"`

	Log Log `koanf:"log"`

	ListenAddress string  `koanf:"listen_address"`
	ExternalURL   url.URL `koanf:"external_url"`
	Session       Session `koanf:"session"`
	DisplayName   string  `koanf:"display_name"`

	UserVerification        protocol.UserVerificationRequirement `koanf:"user_verification_requirement"`
	AuthenticatorAttachment protocol.AuthenticatorAttachment     `koanf:"authenticator_attachment"`
	ConveyancePreference    protocol.ConveyancePreference        `koanf:"conveyance_preference"`
}

func (c Config) AuthenticatorSelection(requirement protocol.ResidentKeyRequirement) (selection protocol.AuthenticatorSelection) {
	fmt.Println("AuthenticatorSelection", requirement)

	selection = protocol.AuthenticatorSelection{
		AuthenticatorAttachment: c.AuthenticatorAttachment,
		UserVerification:        c.UserVerification,
		ResidentKey:             requirement,
	}

	if selection.ResidentKey == "" {
		selection.ResidentKey = protocol.ResidentKeyRequirementDiscouraged
	}

	switch selection.ResidentKey {
	case protocol.ResidentKeyRequirementRequired:
		selection.RequireResidentKey = protocol.ResidentKeyRequired()
	case protocol.ResidentKeyRequirementDiscouraged:
		selection.RequireResidentKey = protocol.ResidentKeyNotRequired()
	}

	if selection.AuthenticatorAttachment == "" {
		selection.AuthenticatorAttachment = protocol.CrossPlatform
	}

	if selection.UserVerification == "" {
		selection.UserVerification = protocol.VerificationPreferred
	}

	return selection
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

type Session struct {
	CookieName string `koanf:"cookie_name"`
	Secure     bool   `koanf:"secure"`
	Domain     string `koanf:"domain"`
}

// ////////////////////////////////////////////////
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
