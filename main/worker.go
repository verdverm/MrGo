package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"runtime"

	USER "github.com/verdverm/MrGo/usr"
)

// ONLY EDIT THIS LINE if you change the name of MyMapReduce
type MRTYPE USER.MyMapReduce

var (
	arg_task   = flag.String("task", "", "[map or reduce], you choose")
	arg_taskid = flag.Int("id", -1, "id of the current task")

	arg_reduces = flag.Int("reduces", 1, "number of reducer processes")
	arg_phases  = flag.Int("phases", 1, "current phase of reduce stages")
	arg_phase   = flag.Int("phase", 1, "current phase of reduce stages")

	arg_tmpdir = flag.String("tmpdir", "tmp", "directory for intermediate files")
)

func init() {
	flag.Parse()
}

func main() {

	fmt.Println("Mr. Go - Starting")

	// printFlags(os.Stdout)
	checkFlags()

	// start mapping process
	if *arg_task == "map" {
		runMap()
		return
	}

	// start reduce process
	if *arg_task == "reduce" {
		runReduce()
		return
	}

	panic("Should Not Get Here")
}

func runMap() {
	// for parallel execution on a single node
	numCPU := runtime.NumCPU()
	// open & partition file for each Map goroutine
	inFn := fmt.Sprintf("file%4d", *arg_taskid)
	data, err := ioutil.ReadFile(inFn)
	if err != nil {
		log.Fatalln(err)
	}
	dl := len(data)
	part_size := dl / numCPU

	done := make(chan string, numCPU)

	for i := 0; i < numCPU; i++ {
		// calc partion end points
		s, e := i*part_size, (i+1)*part_size
		if e > dl {
			e = dl
		}
		go func() {
			S, E := s, e

			mr := new(MRTYPE)
			done <- mr.Map(string(data[S:E]))
		}()
	}

	for i := 0; i < numCPU; i++ {
		result := <-done
		outFn := fmt.Sprintf("%s/temp_t%4d_p%2d_i%2d", *arg_tmpdir, *arg_taskid, *arg_phases, i)
		ioutil.WriteFile(outFn, []byte(result), 0644)
	}
}

func runReduce() {

	numCPU := runtime.NumCPU()
	inputs := make(chan string, numCPU)

	// determine filenames
	fns := make([]string, 0)

	go func() {
		for _, fn := range fns {
			data, err := ioutil.ReadFile(fn)
			if err != nil {
				panic(err)
			}
			inputs <- string(data)
		}
		close(inputs)
	}()

	// mr := new(MRTYPE)
	// result := mr.Reduce(inputs)
	// outFn := fmt.Sprintf("%s/temp_t%4d_p%2d_i%2d", *arg_tmpdir, *arg_taskid, *arg_phase-1, i)

	// ioutil.WriteFile(filename, data, perm)

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
