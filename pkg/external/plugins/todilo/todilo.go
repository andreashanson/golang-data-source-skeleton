package todilo

import (
	"fmt"
	"time"
)

type TodiloRepo struct {
	Name string
}

func NewTodiloRepo() TodiloRepo {
	return TodiloRepo{Name: "Todilo"}
}

func (mr TodiloRepo) Run(c chan []byte) {
	fmt.Println("Started tap:", mr.Name)
	time.Sleep(time.Second * 1)
	c <- []byte(`{"tap": "todilo"}`)
}
