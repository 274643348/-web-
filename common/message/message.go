package message

//这里定义几个用户状态变量
const (
	UserOnline = iota
	AddUserScore
	GetRanking
)

var (
	Rank_ScoreType = "ScoreType"
	Rank_StarsType = "StarsType"
)

//定义一个用户的结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，
	// 必须保证用户信息的json字符串的key和结构体中字段对应的tag名字一样；
	UserId  string `json:"userId"`  //用户id
	UserUrl string `json:"userUrl"` //用户头像Url
}

type UserScore struct {
	User
	Score string `json:"Score"` //用户分数
}

type UserScoreType struct {
	UserScore
	RankType string `json:"rankType"` //排名类型{"ScoreType","StarsType"}
}

//定义一个用户的结构体
type Ranking struct {
	RankType string      `json:"rankType"` //数据类型
	UserList []UserScore `json:"userList"` //用户排行
}

/////////////////////////排行榜相关
type GetRankingStruct struct {
	UserId   string `json:"userId"`   //用户id
	RankType string `json:"rankType"` //排行榜类型
}
