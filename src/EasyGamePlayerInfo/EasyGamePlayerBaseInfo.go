package EasyGamePlayerInfo

import "EasyGameConfigure"

type EasyGamePlayerBaseInfo struct {
	// 统一iD
	userid int64
	// 用户名
	username string
	// 密码
	password string
	// 昵称
	nickname string
	// 游戏类型(二进制) 游戏数量受限
	gameType int64
}

// CreateGamePlayer 创建用户
func (object *EasyGamePlayerBaseInfo) CreateGamePlayer(username string, password string, nickname string) bool {
	return false
}

// Init 初始化对象
// param username 用户名
// return userid
func (object *EasyGamePlayerBaseInfo) Init(username string) int64 {
	// 数据库查询
	return -1
}

// generalConditions 通用判断条件
func (object *EasyGamePlayerBaseInfo) generalConditions() bool {
	// 必须先完成初始化，成果获取userid
	if object.userid < 0 {
		return false
	}

	return true
}

// UpdateNickname 修改昵称
func (object *EasyGamePlayerBaseInfo) UpdateNickname(newNickname string) bool {
	// 通用条件不满足
	if object.generalConditions() {
		return false
	}

	if object.nickname == newNickname {
		return false
	}

	object.nickname = newNickname
	return true
}

// UpdatePassword 修改密码
func (object *EasyGamePlayerBaseInfo) UpdatePassword(oldPassword string, newPassword string) bool {
	// 通用条件不满足
	if object.generalConditions() {
		return false
	}
	// 密码不匹配
	if object.password != oldPassword {
		return false
	}
	// 密码未发生改变
	if object.password == newPassword {
		return false
	}

	object.password = newPassword

	// 写入数据库
	return true
}

// IsVaild 校验密码是否正确
func (object *EasyGamePlayerBaseInfo) IsVaild(password string) bool {
	// 通用条件不满足
	if object.generalConditions() {
		return false
	}

	return object.password == password
}

// AddGameType 添加游戏类型
func (object *EasyGamePlayerBaseInfo) AddGameType(gameType EasyGameConfigure.EasyGameType) bool {
	// 通用条件不满足
	if object.generalConditions() {
		return false
	}

	// 游戏类型已存在
	if object.gameType&int64(gameType) == int64(gameType) {
		return true
	}

	// 添加游戏类型
	object.gameType = object.gameType | int64(gameType)
	// 更新数据库

	return true
}
