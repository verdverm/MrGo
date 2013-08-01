package mr

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

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
	user  string
	name  string
	state HostState

	task *Task
}

func (h *Host) String() string {
	return fmt.Sprintf("%s:  %s", h.name, h.state)
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
	host  *Host
	state TaskState

	task string
	id   int
}

type Scheduler struct {
	config MrGoConfig

	// internal data
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

	fmt.Println("Mapping...\n")

	fmt.Println("Reducing...\n")

	fmt.Println("Scheduler Done")
}

func (s *Scheduler) initHostInfo() {

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

	for _, h := range s.hosts {
		fmt.Printf("%s:  %s\n", h.name, h.state)
	}

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
