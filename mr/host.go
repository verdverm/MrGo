package mr

import (
	"bytes"
	"fmt"
	"os/exec"
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

var (
	ssh_args = []string{
		"ssh",
		"-o StrictHostKeyChecking=no",
		"-o ConnectTimeout=6",
	}
)

func sshRunStrings(user, host, command string) []string {
	ssh_cmd := make([]string, 5)
	ssh_cmd[0] = ssh_args[0]
	ssh_cmd[1] = ssh_args[1]
	ssh_cmd[2] = ssh_args[2]
	ssh_cmd[3] = user + "@" + host
	ssh_cmd[4] = command
	return ssh_cmd
}

func sshStartStrings(user, host, command string) []string {
	ssh_cmd := make([]string, 5)
	ssh_cmd[0] = ssh_args[0]
	ssh_cmd[1] = ssh_args[1]
	ssh_cmd[2] = ssh_args[2]
	ssh_cmd[3] = user + "@" + host
	ssh_cmd[4] = "nohup bash " + command
	return ssh_cmd
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
	err := h.RunCommand("aworm1", "ls")
	if err != nil {
		// log.Println(err)
		h.state = HOST_DEAD
		return
	}
	// fmt.Printf("host return: %q\n", out.String())

	h.state = HOST_LIVE
}

func (h *Host) RunCommand(user, command string) error {
	ssh_cmd := sshRunStrings(user, h.name, command)
	cmd := exec.Command(ssh_cmd[0], ssh_cmd[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	return cmd.Run()
}

func (h *Host) StartCommand(user, command string) error {
	ssh_cmd := sshStartStrings(user, h.name, command)
	cmd := exec.Command(ssh_cmd[0], ssh_cmd[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	return cmd.Run()
}
