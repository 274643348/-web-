package controllers

import (
	"ranking/common/message"
	"ranking/models/process"

	"github.com/astaxie/beego"
)

type RankController struct {
	beego.Controller
}

func (c *RankController) GetRanking() {
	// 解析数据
	userId := c.GetString("userId", "111")
	rankType := c.GetString("rankType", "111")

	userScore := message.GetRankingStruct{
		UserId:   userId,
		RankType: rankType,
	}

	// models中传递数据（进行事物处理）
	ranking, err := process.ProcessGetRanking(&userScore)
	if err != nil {
		c.Data["json"] = ""
		c.ServeJSON()
		return
	}
	c.Data["json"] = ranking
	c.ServeJSON()
}

func (c *RankController) UpdataUserData() {
	// 解析数据
	userId := c.GetString("userId", "111")
	userUrl := c.GetString("userUrl", "111")
	score := c.GetString("score", "111")
	rankType := c.GetString("rankType", "111")

	userScoreType := message.UserScoreType{}
	userScoreType.UserId = userId
	userScoreType.UserUrl = userUrl
	userScoreType.Score = score
	userScoreType.RankType = rankType

	// models中传递数据（进行事物处理）
	process.ProcessUpdataUser(&userScoreType)

	c.Ctx.WriteString("")
}
