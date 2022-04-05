package hubspot

import (
	"fmt"
	"time"
)

type HubspotRepo struct {
	Name string
}

func NewHubspotRepo() HubspotRepo {
	return HubspotRepo{Name: "Hubspot"}
}

func (hsr HubspotRepo) Run(c chan []byte) {
	fmt.Println("Started tap:", hsr.Name)
	time.Sleep(time.Second * 5)
	c <- []byte(`{"tap": "hubspot"}`)
}
