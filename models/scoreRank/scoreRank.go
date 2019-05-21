package scoreRank

import (
	"ranking/common/message"
	"ranking/models/redisCtrl"
	"strconv"

	"github.com/astaxie/beego"
)

var instance *ScoreRank

type ScoreRank struct {
	curMinScore int
}

func GetInstance() *ScoreRank {
	if instance == nil {
		instance = &ScoreRank{}
	}
	return instance
}
func (this *ScoreRank) GetRankingList(userScore *message.GetRankingStruct) (ranking *message.Ranking, err error) {
	userList, err := redisCtrl.GetInstance().GetRankListByType(userScore.RankType, 0, 100)
	if err != nil {
		beego.Error("GetRankingList--GetRankListByType:" + err.Error())
		return
	}
	for key, val := range userList {
		curUser, err := redisCtrl.GetInstance().GetUserById(val.UserId)
		if err != nil {
			beego.Error("GetRankingList--GetRankingList:" + err.Error())
		}
		userList[key].UserUrl = curUser.UserUrl
	}

	ranking = &message.Ranking{
		RankType: userScore.RankType,
		UserList: userList,
	}
	return
}

//更新用户信息
func (this *ScoreRank) UpdataPlayerData(userScore *message.UserScoreType) {

	// 判断是否进前100名
	score, err := strconv.Atoi(userScore.Score)
	if score < 1000 {
		return
	}

	// 判断用户是否存在否则添加
	_, err = redisCtrl.GetInstance().GetUserById(userScore.UserId)
	if err != nil {
		err = redisCtrl.GetInstance().UpdataById(&message.User{UserId: userScore.UserId, UserUrl: userScore.UserUrl})
		if err != nil {
			beego.Error("UpdataPlayerData--UpdataById:" + err.Error())
			return
		}
	}

	// 更新数据
	err = redisCtrl.GetInstance().AddUserToRankList(userScore)
	if err != nil {
		beego.Error("UpdataPlayerData--AddUserToRankList:" + err.Error())
	}

}
