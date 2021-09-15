package global

import (
	"blog-api/blog-web/config"
	"blog-api/blog-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	// NacosConfig *config.NacosConfig = &config.NacosConfig{}

	BlogSrvClient proto.BlogClient
)
