package taprunner

import (
	"fmt"
)

type Plugin interface {
	Run(chan []byte)
}

type Service struct {
	plugins []Plugin
}

func NewSerice(repos []Plugin) *Service {
	return &Service{plugins: repos}
}

func (s *Service) RunPlugins() (chan Plugin, int) {
	numJobs := len(s.plugins)
	tapResponseChan := make(chan []byte, numJobs)

	completedJobs := make(chan Plugin, numJobs)
	//defer close(completedJobs)
	tapJobs := make(chan Plugin, numJobs)
	defer close(tapJobs)

	for w := 1; w <= 2; w++ {
		go worker(w, tapJobs, completedJobs, tapResponseChan)
	}

	for _, tap := range s.plugins {
		tapJobs <- tap
	}

	//for j := 1; j <= numJobs; j++ {
	//	jobs <- j
	//}
	//for a := 1; a <= numJobs; a++ {
	//	fmt.Println(<-completedJobs)
	//}

	//for _, tap := range s.plugins {
	//	go tap.Run(tapResponseChan)
	//}

	return completedJobs, numJobs
}

func worker(id int, jobs chan Plugin, completedJobs chan Plugin, tapChan chan []byte) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j, "with", len(jobs), "left")
		j.Run(tapChan)
		completedJobs <- j
	}
}
