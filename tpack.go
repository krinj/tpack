package tpack

import "fmt"

const UserLevel int = 20

func Execute(x int) int {
	fmt.Println("Execute X TPACK! %")
	return x + UserLevel
}