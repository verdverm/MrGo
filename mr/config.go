package mr

import (
	"log"

	"github.com/robfig/config"
)

type MrGoConfig struct {
	// paths
	Base,
	Input,
	Temp,
	Output string

	// hosts
	HostFile string
	MaxNodes int
}

func (mgc *MrGoConfig) ReadConfig(filename string) {
	c, err := config.ReadDefault(filename)
	if err != nil {
		log.Fatalln(err)
	}

	mgc.Base, err = c.String("DEFAULT", "BaseDir")
	if err != nil {
		log.Fatalln(err)
	}

	mgc.Input, err = c.String("DEFAULT", "InputDir")
	if err != nil {
		log.Fatalln(err)
	}

	mgc.Temp, err = c.String("DEFAULT", "TempDir")
	if err != nil {
		log.Fatalln(err)
	}

	mgc.Output, err = c.String("DEFAULT", "OutputDir")
	if err != nil {
		log.Fatalln(err)
	}

	mgc.HostFile, err = c.String("DEFAULT", "HostFile")
	if err != nil {
		log.Fatalln(err)
	}

	mgc.MaxNodes, err = c.Int("DEFAULT", "MaxNodes")
	if err != nil {
		log.Fatalln(err)
	}

}
