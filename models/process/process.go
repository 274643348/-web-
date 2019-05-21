package process

import (
	"ranking/common/message"
	"ranking/models/scoreRank"
)

func ProcessUpdataUser(userScore *message.UserScoreType) {
	scoreRank.GetInstance().UpdataPlayerData(userScore)
}
func ProcessGetRanking(userScore *message.GetRankingStruct) (ranking *message.Ranking, err error) {
	return scoreRank.GetInstance().GetRankingList(userScore)
}
