package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/verdverm/MrGo/mrgo"
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

	hosts := make([]*mrgo.Host, 0)
	done := make(chan int, 128)

	for i, h := range hostnames {
		h = strings.TrimSpace(h)
		if len(h) < 1 {
			continue
		}
		host := new(mrgo.Host)
		host.Id = i
		host.Name = h
		host.State = mrgo.HOST_NULL
		if h[0] == '#' {
			host.Name = strings.TrimSpace(h[1:])
			host.State = mrgo.HOST_XXXX
		}
		hosts = append(hosts, host)

		go func() {
			if host.State != mrgo.HOST_XXXX {
				host.GetHostState()
			}
			done <- 1
		}()
	}

	for i := 0; i < len(hosts); i++ {
		<-done
	}

	for _, h := range hosts {
		fmt.Printf("%s:  %s\n", h.Name, h.State)
	}

}
