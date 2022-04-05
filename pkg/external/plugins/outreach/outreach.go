package outreach

import (
	"fmt"
	"time"
)

type OutreachRepo struct {
	Name string
}

func NewMondyRepo() OutreachRepo {
	return OutreachRepo{Name: "Outreach"}
}

func (mr OutreachRepo) Run(c chan []byte) {
	fmt.Println("Started tap:", mr.Name)
	time.Sleep(time.Second * 4)
	c <- []byte(`{"tap": "outreach"}`)
}
