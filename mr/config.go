package mr

import (
	"log"
	"math"

	"github.com/robfig/config"
)

type MrGoConfig struct {
	// paths
	Input,
	Temp,
	Output string

	// hosts
	HostFile string
	MaxNodes int

	// M/R params
	NumMaps,
	NumReduces,
	NumPhases int
}

func (mgc *MrGoConfig) ReadConfig(filename string) {
	c, err := config.ReadDefault(filename)
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

	mgc.NumMaps, err = c.Int("DEFAULT", "NumMaps")
	if err != nil {
		log.Fatalln(err)
	}

	mgc.NumReduces, err = c.Int("DEFAULT", "NumReduces")
	if err != nil {
		log.Fatalln(err)
	}

	mgc.NumPhases, err = c.Int("DEFAULT", "NumPhases")
	if err != nil {
		log.Fatalln(err)
	}

	if math.Pow(float64(mgc.NumReduces), float64(mgc.NumPhases)) > float64(mgc.NumMaps) {
		log.Fatalln("Error:  Reduces/Phases is greater than the number of Maps")
	}

}
