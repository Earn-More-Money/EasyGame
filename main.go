package main

import (
	"EasyGameConfigure"
)

func main() {
	var gameType = EasyGameConfigure.EasyGameTypeNiuniuCoin

	print(EasyGameConfigure.EasyGameTypeTableString(gameType))
}
