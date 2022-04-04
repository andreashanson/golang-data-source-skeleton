package taprunner

type Plugin interface {
	Run() chan int
}

type Service struct {
	plugins []Plugin
}

func NewSerice(repos []Plugin) *Service {
	return &Service{plugins: repos}
}

func (s *Service) RunTaps() []chan int {
	var chanSlice []chan int
	for _, tap := range s.plugins {
		tapchan := tap.Run()
		chanSlice = append(chanSlice, tapchan)
	}
	return chanSlice
}
