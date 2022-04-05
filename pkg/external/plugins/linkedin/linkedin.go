package linkedin

import (
	"fmt"
	"time"
)

type LinkedinRepo struct {
	Name string
}

func NewLinkedinRepo() LinkedinRepo {
	return LinkedinRepo{Name: "Linkedin"}
}

func (mr LinkedinRepo) Run(c chan []byte) {
	fmt.Println("Started tap:", mr.Name)
	time.Sleep(time.Second * 7)
	c <- []byte(`{"tap": "linkedin"}`)
}
