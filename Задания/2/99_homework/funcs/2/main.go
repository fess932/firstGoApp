package main

import (
	"reflect"
)

func showMeTheType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {

}
