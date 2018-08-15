package main

import (
	"fmt"
)

type iStuff interface {
	DoStuff()
}

type stuff struct {
	iStuff
	Name string
}

type fakeStuff int

func (f fakeStuff) DoStuff() {
	fmt.Println("its a trap")
}

type realStuff string

func (r realStuff) DoStuff() {
	fmt.Println(r)
}

func (s stuff) SomeComplex() {
	s.DoStuff()
}

// //func main() {
// 	r := realStuff("hey")
// 	f := fakeStuff(0)

// 	rS := stuff{r, "stuff"}

// 	rS.SomeComplex()

// 	fS := stuff{f, "fake"}
// 	fS.DoStuff()
// }
