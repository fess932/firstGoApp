package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type player struct {
	name     string
	position *room
}

type room struct {
	name  string
	about string
	route []string
	// actions map[string]func()
}

func (p *player) watch() {
	fmt.Println(p.position.about)
}

func (p *player) move(s string) {
	if p.position.name == s {

		fmt.Printf("Вы уже в комнате %v \n", s)
		return
	}
	if p.position.name == "" || rooms[s] == nil {
		fmt.Printf("нет пути в %v", s)
		return
	}

	p.position = rooms[s]
	p.watch()
	return

}

func (p *player) use(o, s string) {
	p.position = rooms[s]
}

var rooms = map[string]*room{
	"кухня": {
		name:  "кухня",
		about: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		// actions: map[string]func(){
		// 	"USE": func() {},
		// },
		route: []string{"коридор"},
	},
	"коридор": {
		name:  "коридор",
		about: "ничего интересного. можно пройти - кухня, комната, улица",
		// actions: map[string]func(){
		// 	"USE": func() {},
		// },
		route: []string{"кухня"},
	},
}

var mainPlayer player

var reader = bufio.NewReader(os.Stdin)

// var rooms = []room {
// 	room{name: "main", answers: ["осмотреться"]}
// }

func main() {
	initGame()
}

func initGame() {
	mainPlayer.name = "YOU"
	mainPlayer.position = rooms["КУХНЯ"]
	// if arr[0] == "осмотреться" {
	// 	fmt.Println("ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор")
	// }
	// mainPlayer.move("коридор")
	mainPlayer.watch()
	// for {
	// 	text, _, _ := reader.ReadLine()
	// 	handleCommand(string(text))
	// }

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
		fmt.Print(input)
		if param1 == "" {
			fmt.Print("Укажите куда идти")
			break
		}

		mainPlayer.move(param1)

	case command == "применить":
		mainPlayer.use(param1, param2)
	default:
		fmt.Print("неизвестная команда")
	}

	return "LEL"

}
