package larkcard

func Create(config *Config, header *Header, modules ...IModule) *Card {
	c := &Card{
		Header: header,
		Config: config,
	}
	c.Modules = modules
	return c
}

func CreateI18n(config *Config, header *Header, modules map[string][]IModule) *Card {
	c := &Card{
		Header: header,
		Config: config,
	}
	c.I18nModules = modules
	return c
}

func CreateByMD(config *Config, header *Header, mdContent string, href *OURL) *Card {
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
