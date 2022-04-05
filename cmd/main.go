package main

import (
	"fmt"

	"github.com/andreashanson/data-sources/pkg/external/plugins/hubspot"
	"github.com/andreashanson/data-sources/pkg/external/plugins/monday"
	"github.com/andreashanson/data-sources/pkg/taprunner"
)

func main() {
	hubspotRepo := hubspot.NewHubspotRepo()
	mondayRepo := monday.NewMondyRepo()

	taps := []taprunner.Plugin{
		hubspotRepo,
		mondayRepo,
	}

	type JSONResponse struct {
		Body string `json:"tap"`
	}

	tap := taprunner.NewSerice(taps)
	tapChan, numJobs := tap.RunPlugins()
	defer close(tapChan)
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-tapChan)
	}

	//for i := 0; i < len(taps); i++ {
	//	select {
	//	case first := <-tapChan:
	//		var jsonR JSONResponse
	//
	//		err := json.Unmarshal(first, &jsonR)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(jsonR.Body, "finished...")
	//	}
	//}
}
