package hubspot

import (
	"fmt"
	"time"
)

type HubspotRepo struct {
	Name string
	Chan chan int
}

func NewHubspotRepo(name string, hubspotChan chan int) HubspotRepo {
	return HubspotRepo{Name: name, Chan: hubspotChan}
}

func (hsr HubspotRepo) Run() {
	fmt.Println(hsr.Name)
	defer close(hsr.Chan)
	for {
		time.Sleep(time.Second * 2)
		hsr.Chan <- 555
	}
}
