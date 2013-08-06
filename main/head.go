package main

import (
	"fmt"

	"github.com/verdverm/MrGo/mrgo"
)

var HEAD *mrgo.HeadNode

func init() {
	fmt.Println("Head Node Starting...")
	HEAD = new(mrgo.HeadNode)
	HEAD.InitHostInfo()
}

func main() {
	fmt.Println("Head Node Running...")
}
