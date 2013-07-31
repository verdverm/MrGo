package main

import (
	"flag"
	"fmt"

	MR "mapreduce/MrGo/mr"
)

var (
	arg_config = flag.String("conf", "conf/default.conf", "config file for MrGo")
)

func main() {
	flag.Parse()

	fmt.Println("Hello, World!")

	var mgc MR.MrGoConfig
	mgc.ReadConfig(*arg_config)

	var sched MR.Scheduler
	sched.SetConfig(mgc)

	sched.Init()
	sched.Run()
}
