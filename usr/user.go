package usr

import (
	"fmt"
)

type MyMapReduce struct {
	// add any additional data or variables you need here
}

// optional function
func (mr *MyMapReduce) Init() {
	fmt.Println("MyInit")
}

func (mr *MyMapReduce) Map(input string) (output string) {
	fmt.Println("MyMap")

	return "mymap"
}

func (mr *MyMapReduce) Reduce(input chan string) (output string) {
	fmt.Println("MyReduce")

	return "myreduce"
}

// optinonal function
func (mr *MyMapReduce) Finish() {
	fmt.Println("MyFinish")
}
