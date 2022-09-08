package repository

type SlackRepository interface {
	SendHello()  error
	SendCustomMessage(payload[]byte)([]byte,error)
}
