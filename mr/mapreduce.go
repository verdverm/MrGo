package mr

import (
	"fmt"
)

type MapReduce interface {
	Setup(task string, task_id, reduces, phases int, tmpdir string)
	Init()
	Map()
	Reduce()
}

type MapReduceBase struct {
	// job args
	Task    string
	TaskId  int
	Reduces int
	Phases  int
	Tmpdir  string
}

func (mr *MapReduceBase) Setup(task string, task_id, reduces, phases int, tmpdir string) {
	mr.Task = task
	mr.TaskId = task_id
	mr.Reduces = reduces
	mr.Phases = phases
	mr.Tmpdir = tmpdir
}

func (mr *MapReduceBase) Init() {
	fmt.Println("Base Init", mr.TaskId)
}

func (mr *MapReduceBase) Map() {
	fmt.Println("Base Mapper", mr.TaskId)
}

func (mr *MapReduceBase) Reduce() {
	fmt.Println("Base Reducer", mr.TaskId)
}
