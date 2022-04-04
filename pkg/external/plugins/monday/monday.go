package monday

import (
	"fmt"
	"time"
)

type MondyRepoa struct {
	Name string
	Chan chan int
}

func NewMondyRepoa(name string) MondyRepoa {
	mondayChan := make(chan int)
	return MondyRepoa{Name: name, Chan: mondayChan}
}

func (mr MondyRepoa) Run() chan int {
	go func() {
		fmt.Println("Started tap: ", mr.Name)
		time.Sleep(time.Second * 1)
		mr.Chan <- 666
	}()
	return mr.Chan
}
