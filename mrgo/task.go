package mrgo

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

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
