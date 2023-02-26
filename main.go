package main

import (
	"fmt"

	go_say_hello "github.com/rizkyAlfikri/go-say-hello/v2"
)

// use "go get" command to import module / library  in go mod

func main() {
	fmt.Println(go_say_hello.SayHello("Budi", "Utomo"))
}
