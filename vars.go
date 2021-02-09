package larkcard

type HeaderColor = string

var (
	HeaderBlue      HeaderColor = "blue"
	HeaderWathet    HeaderColor = "wathet"
	HeaderTurquoise HeaderColor = "turquoise"
	HeaderGreen     HeaderColor = "green"
	HeaderYellow    HeaderColor = "Yellow"
	HeaderOrange    HeaderColor = "orange"
	HeaderRed       HeaderColor = "red"
	HeaderCarmine   HeaderColor = "carmine"
	HeaderViolet    HeaderColor = "violet"
	HeaderPurple    HeaderColor = "purple"
	HeaderIndigo    HeaderColor = "indigo"
	HeaderGrey      HeaderColor = "grey"
)

type TextTag = string

var (
	TextPlain  TextTag = "plain_text"
	TextLarkMD TextTag = "lark_md"
)

type ImageModeType = string

var (
	ImageModeFitHorizontal ImageModeType = "fit_horizontal"
	ImageModeOPCenter      ImageModeType = "op_center"
)

type LayoutType = string

var (
	LayoutBisected   LayoutType = "bisected"
	LayoutTrisection LayoutType = "trisection"
	LayoutFlow       LayoutType = "flow"
)

type ButtonType = string

var (
	ButtonDefault ButtonType = "default"
	ButtonPrimary ButtonType = "primary"
	ButtonDanger  ButtonType = "danger"
)

type DatePickerType = string

var (
	DatePickerDate     DatePickerType = "date_picker"
	DatePickerTime     DatePickerType = "picker_time"
	DatePickerDatetime DatePickerType = "picker_datetime"
)
