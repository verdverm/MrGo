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
