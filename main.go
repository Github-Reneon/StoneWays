package main

import "github.com/Github-Reneon/dicetools"

func main() {
	player := new(PC)

	//base stats
	player.Max_STR = dicetools.RollNotation("3d6")
	player.Max_DEX = dicetools.RollNotation("3d6")
	player.Max_Will = dicetools.RollNotation("3d6")
	player.STR = player.Max_STR
	player.DEX = player.Max_DEX
	player.Will = player.Max_Will

	//health
	player.Max_HP = dicetools.RollNotation("1d6")
	player.HP = player.Max_HP

	//gold
	player.Gold = dicetools.RollNotation("3d6")

	println(player)

}
