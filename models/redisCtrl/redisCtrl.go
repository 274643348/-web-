package redisCtrl

import (
	"encoding/json"
	"errors"
	"ranking/common/message"
	"ranking/common/myerror"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/gomodule/redigo/redis"
)

var instance *RedisCtrl

type RedisCtrl struct {
	pool *redis.Pool
}

func GetInstance() *RedisCtrl {
	if instance == nil {
		instance = &RedisCtrl{}
		instance.initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
		beego.Debug("redis init pool")
	}
	return instance
}
func (this *RedisCtrl) initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	this.pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (conn redis.Conn, e error) {
			conn, e = redis.Dial("tcp", address)
			return
		},
	}
}

func init() {
	//初始化redis数据库
	GetInstance()
}

/////////////////////////////////////////具体事务处理

// 获取用户
func (this *RedisCtrl) GetUserById(userId string) (user *message.User, err error) {
	beego.Debug("redisCtrl GetUserById")
	conn := this.pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("get", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = myerror.ERROR_USER_NOTEXISTS
		}
		beego.Error("GetUserById:conn.Do err = ", err)
		return
	}

	user = &message.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		beego.Error("GetUserById:josn.Unmarshal err = ", err)
		return
	}
	return
}

// 更新用户
func (this *RedisCtrl) UpdataById(user *message.User) (err error) {
	beego.Debug("redisCtrl UpdataById")
	conn := this.pool.Get()
	defer conn.Close()

	//格式化数据
	data, err := json.Marshal(user)
	if err != nil {
		beego.Error("UpdataById:json.Marshal err = ", err)
		return
	}

	// 入库
	_, err = conn.Do("set", user.UserId, string(data))
	if err != nil {
		beego.Error("保存注册用户数据错误，err=", err)
		return
	}
	return
}

// 添加到排行榜
func (this *RedisCtrl) AddUserToRankList(userScoreType *message.UserScoreType) (err error) {
	beego.Debug("redisCtrl AddUserList")
	conn := this.pool.Get()
	defer conn.Close()

	score, err := strconv.Atoi(userScoreType.Score)
	_, err = conn.Do("zadd", userScoreType.RankType, score, userScoreType.UserId)
	if err != nil {
		err = errors.New("用户添加失败1")
	}
	return
}

//获取排行榜
func (this *RedisCtrl) GetRankListByType(rankType string, start, end int) (userList []message.UserScore, err error) {
	beego.Debug("redisCtrl GetRankListByType")
	conn := this.pool.Get()
	defer conn.Close()

	res, err := redis.Strings(conn.Do("zrevrange", rankType, start, end, "withscores"))
	if err != nil {
		if err == redis.ErrNil {
			err = errors.New("获取排行榜失败1--ranktype:" + rankType + "--start:" + strconv.Itoa(start) + "--end:" + strconv.Itoa(end) + "--err:" + err.Error())
		}
		err = errors.New("获取排行榜失败2--ranktype:" + rankType + "--start:" + strconv.Itoa(start) + "--end:" + strconv.Itoa(end) + "--err:" + err.Error())
		return
	}

	userList = make([]message.UserScore, 0)
	for i := 0; i < len(res)-1; i += 2 {
		userScore := message.UserScore{
			Score: res[i+1],
		}
		userScore.UserId = res[i]

		userList = append(userList, userScore)
		beego.Debug("players ", i, "---name:", res[i])
	}
	return
}
