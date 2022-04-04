package main

import (
	"fmt"

	"github.com/andreashanson/data-sources/pkg/external/plugins/hubspot"
	"github.com/andreashanson/data-sources/pkg/external/plugins/monday"
	"github.com/andreashanson/data-sources/pkg/taprunner"
)

func main() {
	hubspotRepo := hubspot.NewHubspotRepo("hubspot")
	mondayRepo := monday.NewMondyRepoa("monday")

	taps := []taprunner.Plugin{
		hubspotRepo,
		mondayRepo,
	}

	tap := taprunner.NewSerice(taps)
	tapChans := tap.RunTaps()

	chan0 := tapChans[0]
	chan1 := tapChans[1]
	defer close(chan0)
	defer close(chan1)

	for i := 0; i < 2; i++ {
		select {
		case first := <-chan0:
			fmt.Println(first)
		case second := <-chan1:
			fmt.Println(second)
		}
	}

}
