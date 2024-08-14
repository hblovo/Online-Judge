package main

import (
	"web_backend/core"
	"web_backend/global"
	"web_backend/initialize"
)

func main() {
	global.DB = initialize.Gorm()
	core.RunWindowsServer()
}
