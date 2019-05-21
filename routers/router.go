package routers

import (
	"ranking/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/rank", &controllers.RankController{}, "get:GetRanking")
	beego.Router("/api/update", &controllers.RankController{}, "get:UpdataUserData")
}
