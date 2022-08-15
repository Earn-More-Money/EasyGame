package EasyGamePlayerInfo

import (
	"EasyGameConfigure"
	"math"
)

type EasyGamePlayerGameInfo struct {
	// 用户统一id
	userid int64
	// 游戏类型
	gameType EasyGameConfigure.EasyGameType
	// 游戏积分
	points int64
}

// Init 初始化
func (object *EasyGamePlayerGameInfo) Init(gameType EasyGameConfigure.EasyGameType) bool {
	if object.userid <= 0 {
		return false
	}

	object.gameType = gameType

	// 1. 根据gameType获取表名
	tableName := EasyGameConfigure.EasyGameTypeTableString(object.gameType)

	// 2. 查表获取游戏积分
	print(tableName)

	return false
}

// GetPoints 获取当前账号的积分/金币
func (object *EasyGamePlayerGameInfo) GetPoints() int64 {
	if object.userid <= 0 {
		return math.MinInt64
	}

	return object.points
}

// UpdatePoints 改变积分/金币
func (object *EasyGamePlayerGameInfo) UpdatePoints(points int64) bool {
	if object.userid <= 0 {
		return false
	}

	var value = object.points + points
	// 溢出判断
	if points >= 0 && object.points > value {
		object.points = math.MaxInt64
		return false
	} else if points < 0 && value > object.points {
		object.points = math.MinInt64
		return false
	} else {
		object.points += points
		return true
	}
}
