package mrgo

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type HeadNode struct {
	hosts      []*Host
	live_hosts []*Host
}

func (head *HeadNode) InitHostInfo() {
	// fmt.Println("Getting host states")

	// read file
	data, err := ioutil.ReadFile("conf/nodes.hosts")
	if err != nil {
		log.Fatalln(err)
	}

	hostnames := strings.Split(string(data), "\n")

	head.hosts = make([]*Host, 0)
	done := make(chan int, 128)

	for i, h := range hostnames {
		h = strings.TrimSpace(h)
		if len(h) < 1 {
			continue
		}
		host := new(Host)
		host.Id = i
		host.Name = h
		host.State = HOST_NULL
		if h[0] == '#' {
			host.Name = strings.TrimSpace(h[1:])
			host.State = HOST_XXXX
		}
		head.hosts = append(head.hosts, host)

		go func() {
			if host.State != HOST_XXXX {
				host.GetHostState()
			}
			done <- 1
		}()
	}

	for i := 0; i < len(head.hosts); i++ {
		<-done
	}

	var L, D, X int

	head.live_hosts = make([]*Host, 0)
	for _, h := range head.hosts {
		// fmt.Printf("%s:  %s\n", h.Name, h.State)
		switch h.State {
		case HOST_LIVE:
			L++
			head.live_hosts = append(head.live_hosts, h)

		case HOST_DEAD:
			D++
		case HOST_XXXX:
			X++
		}
	}

	fmt.Printf("There are %d live nodes\n", len(head.live_hosts))

}
