package monday

import (
	"fmt"
	"time"
)

type MondyRepoa struct {
	Name string
	Chan chan int
}

func NewMondyRepoa(name string, mondayChan chan int) MondyRepoa {
	return MondyRepoa{Name: name, Chan: mondayChan}
}

func (mr MondyRepoa) Run() {
	fmt.Println(mr.Name)
	defer close(mr.Chan)
	for {
		time.Sleep(time.Second * 2)
		mr.Chan <- 666
	}

}
