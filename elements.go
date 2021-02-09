package larkcard

import (
	"encoding/json"
)

// Reference:
// https://open.larksuite.com/document/uMzMyEjLzMjMx4yMzITM/uETOxUjLxkTM14SM5ETN
// https://open.feishu.cn/document/ukTMukTMukTM/uAzNwUjLwcDM14CM3ATN

func NewEleImage(imgKey, alt string) *EImage {
	return &EImage{
		Tag:    "img",
		ImgKey: imgKey,
		Alt:    NewObjTextPlain(alt),
	}
}

func NewEleButton(text string, buttonType ...ButtonType) *EButton {
	b := NewEleButtonOrigin(NewObjTextPlain(text))
	if len(buttonType) > 0 {
		b.SetType(buttonType[0])
	}
	return b
}
func NewEleButtonOrigin(text *OText) *EButton {
	return &EButton{
		Tag:  "button",
		Text: text,
	}
}
func NewEleSelectMenu(placeholder string, options [][2]string) *ESelectMenu {
	return NewEleSelectMenuOrigin(placeholder, "", options, false)
}

func NewEleSelectMenuOrigin(placeholder, initialOption string, options [][2]string, selectPerson bool) *ESelectMenu {
	tag := "select_static"
	if selectPerson {
		tag = "select_person"
	}

	ops := make([]*OOption, 0, len(options))
	for _, v := range options {
		ops = append(ops, NewObjOption(v[0], v[1]))
	}

	s := &ESelectMenu{
		Tag:           tag,
		Placeholder:   NewObjTextPlain(placeholder),
		InitialOption: initialOption,
	}
	if len(ops) > 0 {
		s.Options = ops
	}
	return s
}

func NewEleOverflow(options ...*OOption) *EOverflow {
	return &EOverflow{
		Tag:     "overflow",
		Options: options,
	}
}

func NewEleDatePicker(datePickerType DatePickerType, placeholder string, initialTime string) *EDatePicker {
	d := &EDatePicker{
		Tag: datePickerType,
	}
	if initialTime != "" {
		if datePickerType == DatePickerDate {
			d.InitialDate = initialTime
		} else if datePickerType == DatePickerTime {
			d.InitialTime = initialTime
		} else if datePickerType == DatePickerDatetime {
			d.InitialDatetime = initialTime
		}
	}
	d.Placeholder = placeholder
	return d
}

type IElement interface{}

type EImage struct {
	Tag     string `json:"tag"`
	ImgKey  string `json:"img_key"`
	Alt     *OText `json:"alt"`
	Preview *bool  `json:"preview,omitempty"`
}

func (i *EImage) SetPreview(preview bool) *EImage {
	i.Preview = &preview
	return i
}

type EButton struct {
	Tag      string          `json:"tag"`
	Text     *OText          `json:"text"`
	URL      string          `json:"url"`
	MultiURL *OURL           `json:"multi_url,omitempty"`
	Type     *ButtonType     `json:"type,omitempty"`
	Value    json.RawMessage `json:"value,omitempty"`
	Confirm  *OConfirm       `json:"confirm,omitempty"`
}

func (b *EButton) SetURL(url string) *EButton {
	b.URL = url
	return b
}

func (b *EButton) SetMultiURL(url *OURL) *EButton {
	b.MultiURL = url
	return b
}

func (b *EButton) SetType(buttonType ButtonType) *EButton {
	b.Type = &buttonType
	return b
}

func (b *EButton) SetValue(val []byte) *EButton {
	b.Value = val
	return b
}

func (b *EButton) SetConfirm(confirm *OConfirm) *EButton {
	b.Confirm = confirm
	return b
}

type ESelectMenu struct {
	Tag           string          `json:"tag"`
	Placeholder   *OText          `json:"placeholder"`
	InitialOption string          `json:"initial_option,omitempty"`
	Options       []*OOption      `json:"options,omitempty"`
	Value         json.RawMessage `json:"value,omitempty"`
	Confirm       *OConfirm       `json:"confirm,omitempty"`
}

func (s *ESelectMenu) SetValue(val []byte) *ESelectMenu {
	s.Value = val
	return s
}

func (s *ESelectMenu) SetConfirm(confirm *OConfirm) *ESelectMenu {
	s.Confirm = confirm
	return s
}

type EOverflow struct {
	Tag     string          `json:"tag"`
	Options []*OOption      `json:"options"`
	Value   json.RawMessage `json:"value,omitempty"`
	Confirm *OConfirm       `json:"confirm,omitempty"`
}

func (o *EOverflow) SetValue(val []byte) *EOverflow {
	o.Value = val
	return o
}

func (o *EOverflow) SetConfirm(confirm *OConfirm) *EOverflow {
	o.Confirm = confirm
	return o
}

type EDatePicker struct {
	Tag             string          `json:"tag"`
	InitialDate     string          `json:"initial_date,omitempty"`
	InitialTime     string          `json:"initial_time,omitempty"`
	InitialDatetime string          `json:"initial_datetime,omitempty"`
	Placeholder     string          `json:"placeholder,omitempty"`
	Value           json.RawMessage `json:"value,omitempty"`
	Confirm         *OConfirm       `json:"confirm,omitempty"`
}

func (d *EDatePicker) SetValue(val []byte) *EDatePicker {
	d.Value = val
	return d
}

func (d *EDatePicker) SetConfirm(confirm *OConfirm) *EDatePicker {
	d.Confirm = confirm
	return d
}
