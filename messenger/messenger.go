package messenger

// Send messages through some medium to the user
type Messenger interface {
	// Sends an alert message
	SendAlert() error
}
