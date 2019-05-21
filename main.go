package main

import (
	_ "ranking/routers"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
}

func main() {
	//testRedis()
	beego.Run()

}
