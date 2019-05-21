package scoreRank

type rankMethod interface {
	//获取排行榜
	GetRankingList()

	//更新玩家数据
	UpdatPlayerData()
}
