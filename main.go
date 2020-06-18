package main

import (
	"go-admin/cmd"
)

// @title go-admin API
// @version 0.0.1
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 添加qq群: 74520518 进入技术交流群 请备注，谢谢！
// @license.name MIT
// @license.url https://github.com/wenjianzhang/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
