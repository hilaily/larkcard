package larkcard

import "encoding/json"

// Card represent card field in lark card message
// Reference
// https://open.larksuite.com/document/uMzMyEjLzMjMx4yMzITM/ukzNxUjL5cTM14SO3ETN
// https://open.feishu.cn/document/ukTMukTMukTM/ugTNwUjL4UDM14CO1ATN
type Card struct {
	Config      *Config              `json:"config,omitempty"`
	Header      *Header              `json:"header,omitempty"`
	Modules     []IModule            `json:"elements"`
	CardLink    *OURL                `json:"card_link,omitempty"`
	I18nModules map[string][]IModule `json:"i18n_elements,omitempty"`
}

func (c *Card) SetCardLink(link *OURL) *Card {
	c.CardLink = link
	return c
}

// Encode to marshal card structure to []byte
func (c *Card) Encode() []byte {
	en, _ := json.Marshal(c)
	return en
}

// Header represent header info in cark structure
type Header struct {
	Title    *OText            `json:"title"`
	Template *HeaderColor      `json:"template,omitempty"`
	I18n     map[string]string `json:"i18n,omitempty"`
}

// SetI18n
// zh_cn for chinese content, en_us for english, ja_jp for japanese
func (h *Header) SetI18n(i18nInfo map[string]string) *Header {
	h.I18n = i18nInfo
	return h
}

// SetColor set color for header
// Reference
// https://open.larksuite.com/document/uMzMyEjLzMjMx4yMzITM/uADOxUjLwgTM14CM4ETN
// https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN
func (h *Header) SetColor(color HeaderColor) *Header {
	h.Template = &color
	return h
}

// Config set custom config for card
// Reference
// https://open.larksuite.com/document/uMzMyEjLzMjMx4yMzITM/uEDOxUjLxgTM14SM4ETN
// https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
type Config struct {
	WideScreenMode *bool `json:"wide_screen_mode,omitempty"`
	EnableForward  *bool `json:"enable_forward,omitempty"`
}

type IModule interface{}
