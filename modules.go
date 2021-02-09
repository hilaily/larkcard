package larkcard

// Reference:
// https://open.larksuite.com/document/uMzMyEjLzMjMx4yMzITM/uYDOxUjL2gTM14iN4ETN
// https://open.feishu.cn/document/ukTMukTMukTM/uMjNwUjLzYDM14yM2ATN

func NewModContent(content string, isMD ...bool) *MContent {
	f := NewObjTextPlain
	if len(isMD) > 0 && isMD[0] {
		f = NewObjTextMD
	}
	return NewModContentOrigin(f(content), nil)
}

func NewModContentMulti(contents []string, isMD, isShort bool) *MContent {
	f := NewObjTextPlain
	if isMD {
		f = NewObjTextMD
	}
	var fields []*OField
	for _, v := range contents {
		fields = append(fields, &OField{
			IsShort: isShort,
			Text:    f(v),
		})
	}
	return NewModContentOrigin(nil, fields)
}

func NewModContentOrigin(text *OText, fields []*OField) *MContent {
	return &MContent{
		Tag:    "div",
		Text:   text,
		Fields: fields,
		Extra:  nil,
	}
}

func NewModDividerLine() *MDividerLine {
	return &MDividerLine{
		Tag: "hr",
	}
}

func NewModImage(imgKey, alt string) *MImage {
	return &MImage{
		Tag:    "img",
		ImgKey: imgKey,
		Alt:    NewObjTextPlain(alt),
	}
}

func NewModInteraction(actions ...IElement) *MInteraction {
	return &MInteraction{
		Tag:     "action",
		Actions: actions,
	}
}

func NewModNotes(elements []IElement) *MNote {
	return &MNote{
		Tag:      "note",
		Elements: elements,
	}
}

type MContent struct {
	Tag    string    `json:"tag"`
	Text   *OText    `json:"text,omitempty"`
	Fields []*OField `json:"fields,omitempty"`
	Extra  IElement  `json:"extra,omitempty"`
}

func (c *MContent) SetExtra(e IElement) *MContent {
	c.Extra = e
	return c
}

type MDividerLine struct {
	Tag string `json:"tag"`
}

type MImage struct {
	Tag     string        `json:"tag"`
	ImgKey  string        `json:"img_key"`
	Alt     *OText        `json:"alt"`
	Title   *OText        `json:"title,omitempty"`
	Mode    ImageModeType `json:"mode,omitempty"`
	Preview *bool         `json:"preview,omitempty"`
}

func (i *MImage) SetTitle(title string) *MImage {
	i.Title = NewObjTextPlain(title)
	return i
}

func (i *MImage) SetMode(mode ImageModeType) *MImage {
	i.Mode = mode
	return i
}

func (i *MImage) SetPreview(preview bool) *MImage {
	i.Preview = &preview
	return i
}

type MInteraction struct {
	Tag     string      `json:"tag"`
	Actions []IElement  `json:"actions"`
	Layout  *LayoutType `json:"layout,omitempty"`
}

func (i *MInteraction) SetLayout(layout LayoutType) *MInteraction {
	i.Layout = &layout
	return i
}

type MNote struct {
	Tag      string     `json:"tag"`
	Elements []IElement `json:"elements"`
}

type MMD struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
	Href    *OURL  `json:"href,omitempty"`
}

func newMODMD(content string, href *OURL) *MMD {
	return &MMD{
		Tag:  "markdown",
		Href: href,
	}
}
