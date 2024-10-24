package main

import (
	"admin-demo/cmd"
)

// @title admin-demo
// @version v1.0.0
// @description 测试管理后台demo
func main() {
	defer cmd.Clean()
	cmd.Start()
}

//go get 只负责下载
//go install 安装
//从Go 1.17开始，go get不再安装，默认会加go get -d
