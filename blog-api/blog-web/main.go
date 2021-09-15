package main

import (
	"fmt"

	"blog-api/blog-web/global"
	"blog-api/blog-web/initialize"

	"go.uber.org/zap"
)

func main() {
	// 1. 初始化 logger
	initialize.InitLogger()

	// 2.初始化配置文件
	initialize.InitConfig()

	// 3. 初始化 routers
	Router := initialize.Routers()

	// 4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	zap.S().Infof("启动服务器，端口: %d", global.ServerConfig.Port)

	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
