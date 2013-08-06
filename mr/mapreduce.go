package mr

type MapReduce interface {
	Init()
	Map()
	Reduce()
	Finish()
}
