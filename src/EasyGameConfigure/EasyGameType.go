package EasyGameConfigure

type EasyGameType int64

// 枚举：游戏类型
const (
	EasyGameTypeUnkonw       EasyGameType = 1 << 0
	EasyGameTypeNiuniuCoin   EasyGameType = 1 << 1
	EasyGameTypeNiuniuPoints EasyGameType = 1 << 2
)

// Map：游戏类型 - 数据表表头
var easyGameTypeToTableNameString = map[EasyGameType]string{
	EasyGameTypeUnkonw:       "",
	EasyGameTypeNiuniuCoin:   "EasyGameNiuniuCoin",
	EasyGameTypeNiuniuPoints: "EasyGameTypeNiuniuPoints",
}

// EasyGameTypeTableString 根据EasyGameType获取数据库表名
func EasyGameTypeTableString(gameType EasyGameType) string {
	tableName, ok := easyGameTypeToTableNameString[gameType]

	if !ok {
		return ""
	}

	return tableName
}
