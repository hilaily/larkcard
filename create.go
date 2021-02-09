package larkcard

func New(config *Config, header *Header, modules ...IModule) *Card {
	c := &Card{
		Header: header,
		Config: config,
	}
	c.Modules = modules
	return c
}

// NewI18n
// Reference https://open.feishu.cn/document/ukTMukTMukTM/uEjNwUjLxYDM14SM2ATN
func NewI18n(config *Config, header *Header, modules map[string][]IModule) *Card {
	c := &Card{
		Header: header,
		Config: config,
	}
	c.I18nModules = modules
	return c
}

func NewByMD(config *Config, header *Header, mdContent string, href *OURL) *Card {
	c := &Card{
		Header: header,
		Config: config,
	}
	c.Modules = []IModule{newMODMD(mdContent, href)}
	return c
}

func NewHeader(header string) *Header {
	t := NewObjTextPlain(header)
	return &Header{
		Title: t,
	}
}
