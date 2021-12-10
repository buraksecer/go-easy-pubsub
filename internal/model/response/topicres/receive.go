package topicres

// Receive return model
type Receive struct {
	AckId  string
	Data   []byte
	Result bool
	Error  string
}
