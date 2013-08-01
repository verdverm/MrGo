package main

import (
	"flag"
	"fmt"
	// "io/ioutil"

	MR "github.com/verdverm/MrGo/mr"
)

var (
	arg_config = flag.String("conf", "conf/default.conf", "config file for MrGo")
	arg_split  = flag.Bool("split", false, "split files in input dir into chunks")
)

func main() {
	flag.Parse()

	fmt.Println("Hello, World!")

	var mgc MR.MrGoConfig
	mgc.ReadConfig(*arg_config)

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
