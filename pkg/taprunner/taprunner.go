package taprunner

type Plugin interface {
	Run()
}

type Service struct {
	plugins []Plugin
}

func NewSerice(repos []Plugin) *Service {
	return &Service{plugins: repos}
}

func (s *Service) RunTaps() {
	for _, tap := range s.plugins {
		go tap.Run()
	}
}
