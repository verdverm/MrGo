package main

import (
	"fmt"
)

func init() {

}

func main() {
	fmt.Println("Head Node Starting...")
}

func initHostInfo() {
	fmt.Println("Getting host states")

	// read file
	data, err := ioutil.ReadFile("conf/nodes.hosts")
	if err != nil {
		log.Fatalln(err)
	}

	hostnames := strings.Split(string(data), "\n")

	hosts := make([]*Host, 0)
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
