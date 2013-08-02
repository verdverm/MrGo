package mr

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

type Scheduler struct {
	config MrGoConfig

	// internal data
	files []string

	hosts []*Host
	tasks []*Task

	// stats
	done_cnt,
	fail_cnt,
	live_cnt,
	dead_cnt int
}

func (s *Scheduler) SetConfig(c MrGoConfig) {
	s.config = c
}

func (s *Scheduler) Init() {
	fmt.Println("Scheduler Initializing\n----------------------\n")
	s.initHostInfo()
	fmt.Println("\n\n")
}

func (s *Scheduler) Run() {
	fmt.Println("Scheduler Starting\n------------------\n")

	fmt.Println("Splitting...\n")
	fmt.Println("assuming split already...")

	// ioutil.ReadDir(s.config.Input)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// etc...
	// ...

	fmt.Println("Mapping...\n")

	hosts := make([]*Host, 0)
	for i := 0; len(hosts) < s.config.MaxNodes; i++ {
		if i == len(s.hosts) {
			log.Fatalln("Not enough Live Hosts for number of requested nodes")
		}
		if s.hosts[i].state == HOST_LIVE {
			hosts = append(hosts, s.hosts[i])
		}
	}

	tasks := make(chan *Task, s.config.NumMaps)
	done := make(chan int)

	// start runner goroutines
	for i := 0; i < s.config.MaxNodes; i++ {
		go runner(hosts[i], tasks, done)
	}

	for i := 0; i < s.config.NumMaps; i++ {
		t := new(Task)
		t.wkr_args = []string{
			"gocode/bin/MrWorker",
			"-task=map",
			fmt.Sprintf("-id=%d", i),
			fmt.Sprintf("-conf=%s", s.config.Base+"/conf/default.conf"),
		}
		// setup t
		tasks <- t
	}

	fmt.Println("Waiting for map to finish")
	for i := 0; i < s.config.NumMaps; i++ {
		<-done
	}

	fmt.Println("Reducing...\n")

	fmt.Println("Scheduler Done")
}

func runner(host *Host, tasks chan *Task, done chan int) {
	for t := range tasks {

		t.ssh_args = make([]string, 0)
		t.ssh_args = append(t.ssh_args, "ssh")
		t.ssh_args = append(t.ssh_args, "-o StrictHostKeyChecking=no")
		t.ssh_args = append(t.ssh_args, "-o ConnectTimeout=6")
		t.ssh_args = append(t.ssh_args, host.name)
		wkr_args := strings.Join(t.wkr_args, " ")
		t.ssh_args = append(t.ssh_args, wkr_args)

		t.Run()

		done <- 1
	}
}

func (s *Scheduler) initHostInfo() {
	fmt.Println("Getting host states")

	// read file
	data, err := ioutil.ReadFile(s.config.HostFile)
	if err != nil {
		log.Fatalln(err)
	}

	hostnames := strings.Split(string(data), "\n")

	s.hosts = make([]*Host, 0)
	done := make(chan int, 128)

	for i, h := range hostnames {
		h = strings.TrimSpace(h)
		if len(h) < 1 {
			continue
		}
		host := new(Host)
		host.id = i
		host.name = h
		host.state = HOST_NULL
		if h[0] == '#' {
			host.name = strings.TrimSpace(h[1:])
			host.state = HOST_XXXX
		}
		s.hosts = append(s.hosts, host)

		go func() {
			if host.state != HOST_XXXX {
				host.getHostState()
			}
			done <- 1
		}()
	}

	for i := 0; i < len(s.hosts); i++ {
		<-done
	}

	// for _, h := range s.hosts {
	// 	fmt.Printf("%s:  %s\n", h.name, h.state)
	// }

}

type HostState int

const (
	HOST_NULL HostState = iota
	HOST_LIVE
	HOST_DEAD
	HOST_XXXX
)

func (hs HostState) String() string {
	switch hs {
	case HOST_NULL:
		return "HOST_NULL"
	case HOST_LIVE:
		return "HOST_LIVE"
	case HOST_DEAD:
		return "HOST_DEAD"
	case HOST_XXXX:
		return "HOST_XXXX"
	}
	return "HOST_NULL"
}

type Host struct {
	id    int
	name  string
	state HostState
}

func (h *Host) String() string {
	return fmt.Sprintf("%s:  %s", h.name, h.state)
}

func (h *Host) getHostState() {
	cmd := exec.Command("ssh", "-o StrictHostKeyChecking=no", "-o ConnectTimeout=6", h.name, "ls")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// log.Println(err)
		h.state = HOST_DEAD
		return
	}
	// fmt.Printf("host return: %q\n", out.String())

	h.state = HOST_LIVE
}

type TaskState int

const (
	TASK_NULL TaskState = iota
	TASK_RUNNING
	TASK_DONE
	TASK_FAIL
)

func (ts TaskState) String() string {
	switch ts {
	case TASK_NULL:
		return "TASK_NULL"
	case TASK_RUNNING:
		return "TASK_RUNNING"
	case TASK_DONE:
		return "TASK_DONE"
	case TASK_FAIL:
		return "TASK_FAIL"
	}
	return "TASK_NULL"
}

type Task struct {
	state TaskState

	wkr_args []string
	ssh_args []string
}

func (t *Task) Run() {

	// ssh_args := strings.Join(t.ssh_args, " ")

	fmt.Println(len(t.ssh_args), t.ssh_args)

	cmd := exec.Command(
		t.ssh_args[0],
		t.ssh_args[1],
		t.ssh_args[2],
		t.ssh_args[3],
		t.ssh_args[4],
		// "ls -l",
		// t.ssh_args[5],
		// t.ssh_args[6],
		// t.ssh_args[7],
		// t.ssh_args[8],
		// t.ssh_args[9],
	)
	var out bytes.Buffer
	cmd.Stdout = &out

	t.state = TASK_RUNNING
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		t.state = TASK_FAIL
		return
	}
	t.state = TASK_DONE
}
