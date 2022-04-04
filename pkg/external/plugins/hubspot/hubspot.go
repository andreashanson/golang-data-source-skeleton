package hubspot

import (
	"fmt"
	"time"
)

type HubspotRepo struct {
	Name string
}

func NewHubspotRepo(name string) HubspotRepo {
	return HubspotRepo{Name: name}
}

func (hsr HubspotRepo) Run(c chan []byte) {
	fmt.Println("Started tap: ", hsr.Name)
	time.Sleep(time.Second * 7)
	c <- []byte(`{"tap": "hubspot"}`)
}
