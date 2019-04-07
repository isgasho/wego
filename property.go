package wego

import (
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

// SandboxOption ...
type SandboxOption struct {
	SubMchID string `xml:"sub_mch_id"`
	SubAppID string `xml:"sub_app_id"`
}

// SandboxProperty ...
type SandboxProperty struct {
	AppID     string
	AppSecret string
	MchID     string
	Key       string
}

// SafeCertProperty ...
type SafeCertProperty struct {
	Cert   []byte
	Key    []byte
	RootCA []byte
}

// PaymentOption ...
type PaymentOption struct {
	BodyType   *BodyType        `xml:"body_type"`
	SubMchID   string           `xml:"sub_mch_id"`
	SubAppID   string           `xml:"sub_app_id"`
	PublicKey  string           `xml:"public_key"`
	PrivateKey string           `xml:"private_key"`
	RemoteHost string           `xml:"remote_host"`
	LocalHost  string           `xml:"local_host"`
	UseSandbox bool             `xml:"use_sandbox"`
	Sandbox    *SandboxProperty `xml:"sandbox"`

	NotifyURL   string `xml:"notify_url"`
	RefundedURL string `xml:"refunded_url"`
	ScannedURL  string `xml:"scanned_url"`
}

// PaymentProperty ...
type PaymentProperty struct {
	AppID     string            `xml:"app_id"`
	AppSecret string            `xml:"app_secret"`
	MchID     string            `xml:"mch_id"`
	Key       string            `xml:"key"`
	SafeCert  *SafeCertProperty `xml:"safe_cert"`
}

// OAuthProperty ...
type OAuthProperty struct {
	Scopes      []string
	RedirectURI string
}

// OpenPlatformProperty ...
type OpenPlatformProperty struct {
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
}

// OfficialAccountProperty ...
type OfficialAccountProperty struct {
	AppID       string
	AppSecret   string
	Token       string
	AesKey      string
	AccessToken *AccessTokenProperty
	OAuth       *OAuthProperty
}

// MiniProgramProperty ...
type MiniProgramProperty struct {
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
}

// GrantTypeClient ...
const GrantTypeClient string = "client_credential"

// AccessTokenProperty ...
type AccessTokenProperty struct {
	GrantType string `toml:"grant_type"`
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

// ToMap ...
func (obj *AccessTokenProperty) ToMap() util.Map {
	return util.Map{
		"grant_type": obj.GrantType,
		"appid":      obj.AppID,
		"secret":     obj.AppSecret,
	}
}

// ToJSON ...
func (obj *AccessTokenProperty) ToJSON() []byte {
	bytes, err := jsoniter.Marshal(obj)
	if err != nil {
		return nil
	}
	return bytes
}

// JSSDKProperty ...
type JSSDKProperty struct {
	AppID       string
	MchID       string
	Key         string
	AccessToken *AccessTokenProperty
}

// JSSDKOption ...
type JSSDKOption struct {
	SubAppID string
	URL      string
}

// Property 属性配置，各个接口用到的参数
type Property struct {
	JSSDK                 *JSSDKProperty
	JSSDKOption           *JSSDKOption
	AccessToken           *AccessTokenProperty
	AccessTokenOption     *AccessTokenOption
	OAuth                 *OAuthProperty
	OpenPlatform          *OpenPlatformProperty
	OfficialAccount       *OfficialAccountProperty
	OfficialAccountOption *OfficialAccountOption
	MiniProgram           *MiniProgramProperty
	Payment               *PaymentProperty
	PaymentOption         *PaymentOption
	SafeCert              *SafeCertProperty
}

// ParseConfig ...
func ParseProperty(config *Config, v interface{}) (e error) {
	if config == nil {
		return xerrors.New("nil config")
	}
	switch t := v.(type) {
	case *JSSDKProperty:
		e = parseJSSDKProperty(config, t)
	case *AccessTokenProperty:
		e = parseAccessTokenProperty(config, t)
	default:
		e = xerrors.Errorf("some error with %T", t)
	}
	return
}

func parseAccessTokenProperty(config *Config, property *AccessTokenProperty) error {
	if property == nil {
		return xerrors.New("cannot point to nil")
	}
	property = &AccessTokenProperty{
		GrantType: GrantTypeClient,
		AppID:     config.AppID,
		AppSecret: config.AppSecret,
	}
	return nil
}

func parseJSSDKProperty(config *Config, property *JSSDKProperty) error {
	if property == nil {
		return xerrors.New("cannot point to nil")
	}

	return nil
}
