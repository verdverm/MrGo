package mrgo

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strings"
// )

// type Scheduler struct {
// 	config MrGoConfig

// 	// internal data
// 	files []string

// 	hosts []*Host
// 	tasks []*Task

// 	// stats
// 	done_cnt,
// 	fail_cnt,
// 	live_cnt,
// 	dead_cnt int
// }

// func (s *Scheduler) SetConfig(c MrGoConfig) {
// 	s.config = c
// }

// func (s *Scheduler) Init() {
// 	fmt.Println("Scheduler Initializing\n----------------------\n")
// 	s.initHostInfo()
// 	fmt.Println("\n\n")
// }

// func (s *Scheduler) Run() {
// 	fmt.Println("Scheduler Starting\n------------------\n")

// 	fmt.Println("Scheduler Done")
// }

// func runner(host *Host, tasks chan *Task, done chan int) {
// 	for t := range tasks {

// 		t.ssh_args = make([]string, 0)
// 		t.ssh_args = append(t.ssh_args, "ssh")
// 		t.ssh_args = append(t.ssh_args, "-o StrictHostKeyChecking=no")
// 		t.ssh_args = append(t.ssh_args, "-o ConnectTimeout=6")
// 		t.ssh_args = append(t.ssh_args, host.name)
// 		wkr_args := strings.Join(t.wkr_args, " ")
// 		t.ssh_args = append(t.ssh_args, wkr_args)

// 		t.Run()

// 		done <- 1
// 	}
// }

// func (s *Scheduler) initHostInfo() {
// 	fmt.Println("Getting host states")

// 	// read file
// 	data, err := ioutil.ReadFile(s.config.HostFile)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	hostnames := strings.Split(string(data), "\n")

// 	s.hosts = make([]*Host, 0)
// 	done := make(chan int, 128)

// 	for i, h := range hostnames {
// 		h = strings.TrimSpace(h)
// 		if len(h) < 1 {
// 			continue
// 		}
// 		host := new(Host)
// 		host.id = i
// 		host.name = h
// 		host.state = HOST_NULL
// 		if h[0] == '#' {
// 			host.name = strings.TrimSpace(h[1:])
// 			host.state = HOST_XXXX
// 		}
// 		s.hosts = append(s.hosts, host)

// 		go func() {
// 			if host.state != HOST_XXXX {
// 				host.getHostState()
// 			}
// 			done <- 1
// 		}()
// 	}

// 	for i := 0; i < len(s.hosts); i++ {
// 		<-done
// 	}

// 	// for _, h := range s.hosts {
// 	// 	fmt.Printf("%s:  %s\n", h.name, h.state)
// 	// }

// }
