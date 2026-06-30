package main

import "fmt"

func main() {
	process()
}

func process() {
	defer fmt.Println("This will be printed last.")
	defer fmt.Println("This will be printed second last.")
	fmt.Println("This will be printed first.")
}
