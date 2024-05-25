package db

type Quest struct {
	Id       int
	Message  string
	Accepted bool
	Reward   string
	Assignee string
}
