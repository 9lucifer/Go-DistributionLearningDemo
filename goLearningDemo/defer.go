package main

import "fmt"

func main ()  {
	defer fmt.Println("close.....")

	fmt.Println("hello1")
	fmt.Println("hello2")
}