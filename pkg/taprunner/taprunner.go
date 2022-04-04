package taprunner

type Plugin interface {
	Run(chan []byte)
}

type Service struct {
	plugins []Plugin
}

func NewSerice(repos []Plugin) *Service {
	return &Service{plugins: repos}
}

func (s *Service) RunPlugins() chan []byte {
	tapResponseChan := make(chan []byte)

	for _, tap := range s.plugins {
		go tap.Run(tapResponseChan)
	}
	return tapResponseChan
}
