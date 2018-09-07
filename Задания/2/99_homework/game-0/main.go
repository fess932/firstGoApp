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
	about   string
	route   map[string]room
	actions map[string]func()
}

func (p *player) watch() {
	fmt.Println(p.position.about)
}

func (p *player) move(s string) {
	If p.position = rooms[s]
}

func (p *player) use(o, s string) {
	p.position = rooms[s]
}

var rooms = map[string]room{
	"кухня": {
		about: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		actions: map[string]func(){
			"USE": func() {},
		},
		route: {
			"коридор": rooms["коридор"],
		},
	},
	"коридор": {
		about: "ничего интересного. можно пройти - кухня, комната, улица",
		actions: map[string]func(){
			"USE": func() {},
		},
		route: {
			"кухня": rooms["кухня"],
		},
	},
}

var mainPlayer = player{name: "ты", position: rooms["кухня"]}

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
	mainPlayer.move("коридор")
	mainPlayer.watch()
	handleCommand(string(text))
}

func handleCommand(input string) string {
	arr := strings.Split(input, " ")
	command := arr[0]
	var param1, param2 string

	if len(arr) == 2 {
		param1 = arr[1]
	}
	if len(arr) == 3 {
		param2 = arr[2]
	}

	switch {
	case command == "осмотреться":
		mainPlayer.watch()
	case command == "идти":
		mainPlayer.move(param1)
	}
	return command
}
