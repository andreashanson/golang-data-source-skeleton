package hubspot

import (
	"fmt"
	"time"
)

type HubspotRepo struct {
	Name string
	Chan chan int
}

func NewHubspotRepo(name string) HubspotRepo {
	hubspotChan := make(chan int)
	return HubspotRepo{Name: name, Chan: hubspotChan}
}

func (hsr HubspotRepo) Run() chan int {
	go func() {
		fmt.Println("Started tap: ", hsr.Name)
		time.Sleep(time.Second * 6)
		hsr.Chan <- 555
	}()
	return hsr.Chan
}
