package monday

import (
	"fmt"
	"time"
)

type MondyRepoa struct {
	Name string
}

func NewMondyRepoa(name string) MondyRepoa {
	return MondyRepoa{Name: name}
}

func (mr MondyRepoa) Run(c chan []byte) {
	fmt.Println("Started tap: ", mr.Name)
	time.Sleep(time.Second * 10)
	c <- []byte(`{"tap": "monday"}`)
}
