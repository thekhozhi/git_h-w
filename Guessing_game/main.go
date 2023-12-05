package main

import (
	"fmt"
	"math/rand"
)

func main() {
	player := Player{}
	game := Game{}
	name, age, favnum := getPlayerInfo()
	newPlayer := player.NewPlayer(name, age, favnum, player.Chances)
	newGame := game.NewGame(newPlayer)
	newGame.StartGame()

}

var counter = 0

type Player struct {
	Name        string
	Age         int
	Chances     int
	FavoriteNum int
}

type Game struct {
	RandomNumber int
	Player       Player
}

func getPlayerInfo() (string, int, int) {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	var favnum int
	fmt.Print("Enter your favorite number: ")
	fmt.Scan(&favnum)

	return name, age, favnum
}

func (p Player) NewPlayer(name string, chances int, age int, favnum int) Player {
	return Player{
		Name:        name,
		Age:         age,
		Chances:     chances,
		FavoriteNum: favnum,
	}
}

func (g Game) NewGame(player Player) Game {
	return Game{
		Player: player,
	}
}

func (g Game) StartGame() {

	var chances int
	fmt.Printf("Welcome %s\n", g.Player.Name)
	fmt.Println("This is guessing game!")
	fmt.Println("Before starting our game you can choose chances.\nYou can have 1 chance, 3 chances, 5 chances or 10 chances.")
	fmt.Print("Choose one and write the number: ")
	fmt.Scan(&chances)
	if chances == 1 {
		g.RandomNumber = rand.Intn(5)
	} else if chances == 3 {
		g.RandomNumber = rand.Intn(10)
	} else if chances == 5 {
		g.RandomNumber = rand.Intn(20)
	} else if chances == 10 {
		g.RandomNumber = rand.Intn(50)
	} else {
		fmt.Println("Error!!!")
		return
	}

	for i := 2; i < g.RandomNumber/2; i++ {
		if g.RandomNumber%i == 0 {
			counter++
		}
		if counter == 0 {
			fmt.Println("The random number is prime!")
		} else {
		}
	}

	for i := 0; i < chances; i++ {
		var n int
		fmt.Print(i+1, " chance, enter number: ")
		fmt.Scan(&n)

		if n == g.RandomNumber && g.RandomNumber == g.Player.FavoriteNum {
			fmt.Printf("You won! You found the random number in %d try, And the random number was your favorite number", i+1)
			fmt.Println()
			return
		}
		if n == g.RandomNumber {
			fmt.Println("You won! you found the random number in", i+1, "try")
			return
		} else {
			fmt.Println("Incorrect!")
			if n < g.RandomNumber {
				fmt.Println("The random number is greater than this number!")
			} else {
				fmt.Println("The random number is lower than this number!")
			}
		}
	}
	fmt.Println("You lost the game, the random number was:", g.RandomNumber)
}
