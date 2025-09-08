package application

type Status string

const (
	StatusReceived  Status = "Received"
	StatusViewed    Status = "Viewed"
	StatusInProcess Status = "InProcess"
	StatusRejected  Status = "Rejected"
	StatusAccepted  Status = "Accepted"
	StatusCancelled Status = "Cancelled"
	StatusOnHold    Status = "OnHold"
)
