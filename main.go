package main

import (
	"fmt"
	"llamadas/setup"
)

func init() {
	setup.SetDatabase()
}

func main() {
	fmt.Println("Hello world")
}
