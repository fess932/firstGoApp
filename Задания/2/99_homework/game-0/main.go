package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type player struct {
	name     string
	position room
}

type room struct {
	actions map[string]string
}

func (p *player) watch() {
	fmt.Println(p.position.actions["ABOUT"])
}

func (p *player) move(s string) {
	p.position = rooms[s]
}

func (p *player) use(o, s string) {
	p.position = rooms[s]
}

var rooms = map[string]room{
	"KITCHEN": {
		actions: map[string]string{
			"USE":   "ANOE",
			"ABOUT": "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		},
	},
	"HALL": {
		actions: map[string]string{
			"ABOUT": "ничего интересного. можно пройти - кухня, комната, улица",
		},
	},
}

var mainPlayer = player{name: "HERO", position: rooms["KITCHEN"]}

// var rooms = []room {
// 	room{name: "main", answers: ["осмотреться"]}
// }

func main() {
	initGame()
}

func initGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Добро пожаловать в игру!")
	text, _, _ := reader.ReadLine()

	// if arr[0] == "осмотреться" {
	// 	fmt.Println("ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор")
	// }
	mainPlayer.watch()
	mainPlayer.move("HALL")
	mainPlayer.watch()
	handleCommand(string(text))
}

func handleCommand(input string) string {
	arr := strings.Split(input, " ")
	command := arr[0]
	// param1 := arr[1]
	// param2 := arr[2]
	switch {
	case command == "осмотреться":
		mainPlayer.watch()
		// case command == "идти":
		// 	mainPlayer.move(param1)
	}
	return command
}
