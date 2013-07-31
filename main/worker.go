package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	MR "github.com/verdverm/MrGo/mr"
	USER "github.com/verdverm/MrGo/usr"
)

var (
	arg_task   = flag.String("task", "", "[map or reduce], you choose")
	arg_taskid = flag.Int("id", -1, "id of the current task")

	arg_reduces = flag.Int("reduces", 1, "number of reducer processes")
	arg_phases  = flag.Int("phases", 1, "number of reduce stages")

	arg_tmpdir = flag.String("tmpdir", "tmp", "directory for intermediate files")
)

func init() {
	flag.Parse()
}

func main() {

	fmt.Println("MrGo - Starting")

	// printFlags(os.Stdout)
	checkFlags()

	// ONLY EDIT THIS LINE vvv
	/**********************/
	mr := new(USER.MyMapReduce)
	/**********************/

	setupMR(mr)

	// start mapping process
	if *arg_task == "map" {
		mr.Map()
		return
	}

	// start reduce process
	if *arg_task == "reduce" {
		mr.Reduce()
		return
	}

	panic("Should Not Get Here")
}

func setupMR(mr MR.MapReduce) {
	mr.Setup(*arg_task, *arg_taskid, *arg_reduces, *arg_phases, *arg_tmpdir)
}

func printFlags(w io.Writer) {
	v := func(f *flag.Flag) {
		fmt.Fprintf(w, "%s : %v  [%s]\n", f.Name, f.Value, f.DefValue)
	}
	flag.VisitAll(v)
	fmt.Fprintln(w)
}

func checkFlags() {
	if *arg_task != "map" && *arg_task != "reduce" {
		log.Fatalln("Task not set, expected [map or reduce]")
	}
	if *arg_taskid < 0 {
		log.Fatalln("Task ID not set, must be a non-negative integer")
	}
}
