package larkcard

// Reference:
// https://open.larksuite.com/document/uMzMyEjLzMjMx4yMzITM/uYTOxUjL2kTM14iN5ETN
// https://open.feishu.cn/document/ukTMukTMukTM/uUzNwUjL1cDM14SN3ATN

func NewObjTextPlain(text string) *OText {
	return NewObjText(text, TextPlain)
}
func NewObjTextMD(text string) *OText {
	return NewObjText(text, TextLarkMD)
}

func NewObjText(content string, tag TextTag) *OText {
	return &OText{
		Tag:     tag,
		Content: content,
	}
}

func NewObjField(isShort bool, text string) *OField {
	return &OField{
		IsShort: isShort,
		Text:    NewObjTextPlain(text),
	}
}

func NewObjURL(url, androidURL, iOSURL, pcURL string) *OURL {
	return &OURL{
		URL:        url,
		AndroidURL: androidURL,
		IOSURL:     iOSURL,
		PCURL:      pcURL,
	}
}

func NewObjOption(text string, val string) *OOption {
	return &OOption{
		Text:  NewObjTextPlain(text),
		Value: val,
	}
}

func NewObjConfirm(title, text string) *OConfirm {
	return &OConfirm{
		Title: NewObjTextPlain(title),
		Text:  NewObjTextPlain(text),
	}
}

type OText struct {
	Tag     TextTag `json:"tag"`
	Content string  `json:"content"`
	Lines   int     `json:"lines,omitempty"`
}

func (t *OText) SetLines(l int) *OText {
	t.Lines = l
	return t
}

type OField struct {
	IsShort bool   `json:"is_short"`
	Text    *OText `json:"text"`
}

type OURL struct {
	URL        string `json:"url"`
	AndroidURL string `json:"android_url"`
	IOSURL     string `json:"ios_url"`
	PCURL      string `json:"pc_url"`
}

type OOption struct {
	Text     *OText `json:"text,omitempty"`
	Value    string `json:"value"`
	URL      string `json:"url"`
	MultiURL *OURL  `json:"multi_url,omitempty"`
}

func (o *OOption) SetURL(url string) *OOption {
	o.URL = url
	return o
}

func (o *OOption) SetMultiURL(url *OURL) *OOption {
	o.MultiURL = url
	return o
}

type OConfirm struct {
	Title *OText `json:"title"`
	Text  *OText `json:"text"`
}
