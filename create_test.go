package larkcard

import (
	"fmt"
	"testing"
)

func TestCreateCard(t *testing.T) {
	c := Create(
		nil,
		NewHeader("test header").SetColor(HeaderGreen),
		NewModContent("[飞书](https://www.feishu.cn)整合即时沟通、日历、音视频会议、云文档、云盘、工作台等功能于一体，成就组织和个人，更高效、更愉悦。", true),
		NewModInteraction(
			NewEleButton("主按钮", ButtonPrimary),
			NewEleButton("次按钮"),
			NewEleButton("危险按钮", ButtonDanger),
		),
		NewModContent("深度整合使用率极高的办公工具，企业成员在一处即可实现高效沟通与协作。").SetExtra(
			NewEleImage("img_e344c476-1e58-4492-b40d-7dcffe9d6dfg", "hover"),
		),
		NewModContent("在移动端同样进行便捷的沟通、互动与协作，手机电脑随时随地保持同步。").SetExtra(
			NewEleSelectMenu("默认提示文本", [][2]string{
				{"选项1", "1"},
				{"选项2", "2"},
				{"选项3", "3"},
			}),
		),
		NewModContent("ISV产品接入及企业自主开发，更好地对接现有系统，满足不同组织的需求。").SetExtra(
			NewEleOverflow(
				NewObjOption("打开飞书应用目录", "https://app.feishu.cn"),
				NewObjOption("打开飞书开发文档", "https://open.feishu.cn"),
			),
		),
		NewModContent("国际权威安全认证与信息安全管理体系，为企业提供全生命周期安全保障。").SetExtra(
			NewEleDatePicker(
				DatePickerDate, "", "2020-09-20",
			),
		),
	)
	en := c.Encode()
	fmt.Println(string(en))
}
