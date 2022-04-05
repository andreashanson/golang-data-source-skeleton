package monday

import (
	"fmt"
	"time"
)

type MondyRepo struct {
	Name string
}

func NewMondyRepo() MondyRepo {
	return MondyRepo{Name: "Monday"}
}

func (mr MondyRepo) Run(c chan []byte) {
	fmt.Println("Started tap:", mr.Name)
	time.Sleep(time.Second * 10)
	c <- []byte(`{"tap": "monday"}`)
}
