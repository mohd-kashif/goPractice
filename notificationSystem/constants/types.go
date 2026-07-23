package constants

type Status string

const (
	StatusPending Status = "Pending"
	StatusFailed  Status = "Failed"
	StatusSent    Status = "Sent"
)

type Channel string

const (
	ChannelEmail            Channel = "Email"
	ChannelSMS              Channel = "SMS"
	ChannelPushNotification Channel = "Push Notification"
)

type Priority string

const (
	PriorityLow    Priority = "Low"
	PriorityMedium Priority = "Medium"
	PriorityHigh   Priority = "High"
)
