package main

import (
	"fmt"

	"github.com/andreashanson/data-sources/pkg/external/plugins/hubspot"
	"github.com/andreashanson/data-sources/pkg/external/plugins/linkedin"
	"github.com/andreashanson/data-sources/pkg/external/plugins/monday"
	"github.com/andreashanson/data-sources/pkg/external/plugins/outreach"
	"github.com/andreashanson/data-sources/pkg/external/plugins/todilo"
	"github.com/andreashanson/data-sources/pkg/taprunner"
)

func main() {
	hubspotRepo := hubspot.NewHubspotRepo()
	mondayRepo := monday.NewMondyRepo()
	outreachRepo := outreach.NewMondyRepo()
	todiloRepo := todilo.NewTodiloRepo()
	linkedinRepo := linkedin.NewLinkedinRepo()

	taps := []taprunner.Plugin{
		hubspotRepo,
		mondayRepo,
		outreachRepo,
		todiloRepo,
		linkedinRepo,
	}

	type JSONResponse struct {
		Body string `json:"tap"`
	}

	tap := taprunner.NewSerice(taps)
	tapChan, numJobs := tap.RunPlugins()
	defer close(tapChan)
	for a := 1; a <= numJobs; a++ {
		response := <-tapChan
		fmt.Println(string(response))

	}
}
