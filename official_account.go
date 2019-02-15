package wego

// OfficialAccount ...
type OfficialAccount struct {
	*OfficialAccountProperty
	client *Client
	config *Config
	option *OfficialAccountOption
}

// OfficialAccountOption ...
type OfficialAccountOption struct {
	RemoteAddress string
	LocalHost     string
	UseSandbox    bool
	Sandbox       *SandboxConfig
	NotifyURL     string
	RefundURL     string
}

// NewOfficialAccount ...
func NewOfficialAccount(config *Config, opts ...*OfficialAccountOption) *OfficialAccount {
	var opt *OfficialAccountOption
	if opts != nil {
		opt = opts[0]
	}
	bt := BodyTypeJSON
	return &OfficialAccount{
		client: NewClient(&ClientOption{
			//AccessToken: NewAccessToken(config.AccessToken.Credential()),
			BodyType: &bt,
		}),
		OfficialAccountProperty: config.OfficialAccount,
		config:                  config,
		option:                  opt,
	}
}
