package main

import (
	"fmt"

	"github.com/andreashanson/data-sources/pkg/external/plugins/hubspot"
	"github.com/andreashanson/data-sources/pkg/external/plugins/monday"
	"github.com/andreashanson/data-sources/pkg/taprunner"
)

func main() {
	hubspotChan := make(chan int)
	hubspotRepo := hubspot.NewHubspotRepo("hubspot", hubspotChan)

	mondayChan := make(chan int)
	mondayRepo := monday.NewMondyRepoa("monday", mondayChan)

	taps := []taprunner.Plugin{
		hubspotRepo,
		mondayRepo,
	}

	tap := taprunner.NewSerice(taps)
	tap.RunTaps()

	for {
		select {

		case hubspotData := <-hubspotChan:
			fmt.Println(hubspotData)

		case mondayData := <-mondayChan:
			fmt.Println(mondayData)
		}
	}

}
