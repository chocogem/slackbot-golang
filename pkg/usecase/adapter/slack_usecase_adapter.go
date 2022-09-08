package adapter

import (
	repository "github.com/chocogem/slackbot-golang/pkg/repository"
	usecase "github.com/chocogem/slackbot-golang/pkg/usecase"
)

type slackUseCase struct {
	slackRepo repository.SlackRepository
}

func NewSlackUseCase(sr repository.SlackRepository) usecase.SlackUseCase {
	return &slackUseCase{
		slackRepo: sr,
	}
}

func (s *slackUseCase) SendHello()  error {
	 err := s.slackRepo.SendHello()
	 return  err
}
func (s *slackUseCase) SendCustomMessage(payload []byte)([]byte,error) {
	respBody,err := s.slackRepo.SendCustomMessage(payload)
    return  respBody,err
}