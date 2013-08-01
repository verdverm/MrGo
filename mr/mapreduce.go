package mr

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type MapReduce interface {
	Setup(task string, task_id, reduces, phases int, tmpdir string)
	Init()
	Map()
	Reduce()
	Finish()
}

type MapReduceBase struct {
	// job args
	Task    string
	TaskId  int
	Reduces int
	Phases  int
	Tmpdir  string

	inputFn,
	outputFn string
	input    *bufio.Reader
	output   *bufio.Writer
	inputFd  *os.File
	outputFd *os.File
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

	// open input file for reading
	var err error
	if mr.Task == "map" {
		dataID := mr.TaskId
		mr.inputFn = fmt.Sprintf("file%04d", dataID)

		tempID := mr.TaskId
		mr.outputFn = fmt.Sprintln("temp%04d", tempID)
	} else if mr.Task == "reduce" {

	} else {
		log.Fatalln("Uknown task error in MapReduceBase.Init()")
	}

	mr.inputFd, err = os.Open(mr.inputFn)
	if err != nil {
		log.Fatalln(err)
	}
	mr.input = bufio.NewReader(mr.inputFd)

	mr.outputFd, err = os.Create(mr.outputFn)
	if err != nil {
		log.Fatalln(err)
	}
	mr.output = bufio.NewWriter(mr.outputFd)
}

func (mr *MapReduceBase) Map() {
	fmt.Println("Base Mapper", mr.TaskId)
}

func (mr *MapReduceBase) Reduce() {
	fmt.Println("Base Reducer", mr.TaskId)
}

func (mr *MapReduceBase) Finish() {
	fmt.Println("Base Finish")
	var err error

	err = mr.inputFd.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = mr.output.Flush()
	if err != nil {
		log.Fatalln(err)
	}

	err = mr.outputFd.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
