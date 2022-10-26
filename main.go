package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}


func main() {
	profile := makeProfile("Sasuke")
	fmt.Println("Name: ",profile.Name)
	fmt.Println("Health: ",profile.Health)
	fmt.Println("Power: ",profile.Power )
	fmt.Println("Exp: ",profile.Exp )
	fmt.Println("=====HEROES POWER UP=====")
	powerUp := PowerUp(profile,3)
	fmt.Println("Name: ",powerUp.Name)
	fmt.Println("Health: ",powerUp.Health)
	fmt.Println("Power: ",powerUp.Power )
	fmt.Println("Exp: ",powerUp.Exp )
	
}
func makeProfile(inputName string) Profile {
	var anime Profile
	anime.Name = inputName
	if anime.Name == "Sasuke" {
		anime.Health = 200
		anime.Power = 200
		anime.Exp = 0
	}
	if anime.Name == "Goku" {
		anime.Health = 400
		anime.Power = 300
		anime.Exp = 100
	}
	if anime.Name == "Naruto" {
		anime.Health = 150
		anime.Power = 200
		anime.Exp = 50
	}
	return anime
	
}
func PowerUp(profile Profile,multiplayer int) Profile {
	var animePowerUp Profile

	animePowerUp.Name = profile.Name
	animePowerUp.Health = profile.Health + profile.Health * multiplayer
	animePowerUp.Power = profile.Power + profile.Power * multiplayer
	animePowerUp.Exp = profile.Exp + profile.Exp * multiplayer

return animePowerUp
}



