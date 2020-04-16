package radio

// Context defines command context
type Context interface {
	GetArgs() []string
	GetUser() User

	CommandInProgress()
	SendMessage(msg interface{})
}
