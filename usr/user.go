package usr

import (
	"fmt"

	MR "github.com/verdverm/MrGo/mr"
)

type MyMapReduce struct {
	// do not remove this
	MR.MapReduceBase

	// add any additional data or variables you need here
}

func (mr *MyMapReduce) Init() {
	fmt.Println("MyInit")
}

func (mr *MyMapReduce) Map() {
	fmt.Println("MyMap")
}

func (mr *MyMapReduce) Reduce() {
	fmt.Println("MyReduce")
}
