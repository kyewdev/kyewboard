package db

type Quest struct {
	Id       int
	Message  string
	Status   string
	Reward   string
	Assignee string
    Questtype string
    Category string
}
