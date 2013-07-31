package mr

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Collector struct {
	file *os.File
	buf  *bufio.Writer
}

func (c *Collector) Open(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("couldn't create errs log", err)
	}

	buf := bufio.NewWriter(file)
	buf.Flush()

	c.file = file
	c.buf = buf
}

func (c *Collector) Collect(obj interface{}) {
	fmt.Fprint(c.buf, obj)
}

func (c *Collector) Collectln(obj interface{}) {
	fmt.Fprintln(c.buf, obj)
}

func (c *Collector) Flush() {
	c.buf.Flush()
}

func (c *Collector) Close() {
	c.buf.Flush()
	c.file.Close()
}
