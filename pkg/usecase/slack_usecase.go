package usecase

type SlackUseCase interface {
	SendHello() error
	SendCustomMessage(payload []byte) ([]byte,error)
}
