package mr

import (
	"bufio"
	"log"
	"os"
)

type MrLog struct {
	*log.Logger
	file *os.File
	buf  *bufio.Writer
}

func newLogger(filename string) (mr *MrLog) {

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("couldn't create errs log", err)
	}

	buf := bufio.NewWriter(file)
	buf.Flush()
	mr = new(MrLog)
	mr.Logger = log.New(buf, "", log.LstdFlags)
	mr.file = file
	mr.buf = buf

	return mr
}

func (Mr *MrLog) Flush() {
	Mr.buf.Flush()
}

func (Mr *MrLog) Close() {
	Mr.buf.Flush()
	Mr.file.Close()
}
