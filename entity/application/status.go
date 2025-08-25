package application

type Status string

const (
	StatusReceived  = "Received"
	StatusViewed    = "Viewed"
	StatusInProcess = "InProcess"
	StatusRejected  = "Rejected"
	StatusAccepted  = "Accepted"
	StatusCancelled = "Cancelled"
	StatusOnHold    = "OnHold"
)
