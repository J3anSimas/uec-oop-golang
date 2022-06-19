package main

import (
	"fmt"
	"github.com/J3anSimas/uec-golang/pkg/fighter"
)

func main() {
	var fighter fighter.FighterInterface = &fighter.Fighter{}

	fighter.SetName("Jean Simas")
	fighter.SetNationality("Brazilian")
	fighter.SetAge(24)
	fighter.SetHeight(1.74)
	fighter.SetWeight(63.5)
	fighter.SetCategory("Leve")
	fmt.Println(fighter.Present())
	fmt.Println("=============================================")
	fighter.WinFight()
	fighter.WinFight()
	fighter.WinFight()
	fighter.DrawFight()
	fmt.Println(fighter.Status())

}
