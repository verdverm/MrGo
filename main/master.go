package main

import (
	"flag"
	"fmt"
	// "io/ioutil"

	MR "github.com/verdverm/MrGo/mr"
)

var (
	arg_dir   = flag.String("dir", ".", "base dir for MrGo to find files")
	arg_split = flag.Bool("split", false, "split files in input dir into chunks")
)

func main() {
	flag.Parse()

	fmt.Println("Hello, World!")

	var mgc MR.MrGoConfig
	mgc.ReadConfig(*arg_dir + "/conf/default.conf")

	var sched MR.Scheduler
	sched.SetConfig(mgc)

	if *arg_split {
		run_split(mgc)
	}

	sched.Init()
	sched.Run()
}

func run_split(config MR.MrGoConfig) {

}
