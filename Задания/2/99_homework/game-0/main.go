package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Добро пожаловать в игру!")
	text, _, _ := reader.ReadLine()
	arr := strings.Split(string(text[:]), " ")
	if arr[0] == "осмотреться" {
		fmt.Println("ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор")
	}
}
func initGame() {

}

func handleCommand(command string) string {

	return command
}
