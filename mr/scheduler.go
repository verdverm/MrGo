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
)

func (hs HostState) String() string {
	switch hs {
	case HOST_NULL:
		return "HOST_NULL"
	case HOST_LIVE:
		return "HOST_LIVE"
	case HOST_DEAD:
		return "HOST_DEAD"
	}
	return "HOST_NULL"
}

type Host struct {
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
	fmt.Println("Scheduler Initializing\n----------------------")
	s.readHostFile()
}

func (s *Scheduler) Run() {
	fmt.Println("Scheduler Starting\n------------------")

	fmt.Println("Scheduler Done")
}

func (s *Scheduler) readHostFile() {
	data, err := ioutil.ReadFile(s.config.HostFile)
	if err != nil {
		log.Fatalln(err)
	}

	hostnames := strings.Split(string(data), "\n")

	s.hosts = make([]*Host, len(hostnames))
	for _, h := range hostnames {
		h = strings.TrimSpace(h)
		if len(h) < 1 {
			continue
		}
		host := new(Host)
		host.name = h
		host.state = host.getHostState()

		fmt.Println(host)
		s.hosts = append(s.hosts, host)
	}
}

func (h *Host) getHostState() HostState {
	cmd := exec.Command("ssh", "-o StrictHostKeyChecking=no", "-o ConnectTimeout=6", "aworm1@"+h.name, "ls")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// log.Println(err)
		return HOST_DEAD
	}
	// fmt.Printf("host return: %q\n", out.String())

	return HOST_LIVE
}
