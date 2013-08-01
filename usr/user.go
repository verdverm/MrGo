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

// optional function
func (mr *MyMapReduce) Init() {
	// always call the MR.Init()
	mr.MapReduceBase.Init()

	fmt.Println("MyInit")
}

func (mr *MyMapReduce) Map() {
	fmt.Println("MyMap")
}

func (mr *MyMapReduce) Reduce() {
	fmt.Println("MyReduce")
}

// optinonal function
func (mr *MyMapReduce) Finish() {
	// always call the MR.Init()
	mr.MapReduceBase.Finish()

}
