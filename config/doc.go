package config

import "github.com/gofiber/swagger"

func NewDoc() swagger.Config {
	return swagger.Config{
		Title:                 "订座系统",
		Layout:                "StandaloneLayout", // 最简洁的布局
		DocExpansion:          "none",             // 不展开文档
		DefaultModelRendering: "example",
		DeepLinking:           false,
		QueryConfigEnabled:    false,
		CustomStyle:           "body { margin: 0 !important; }  /* 去除页面body的默认边距 */",
	}
}
