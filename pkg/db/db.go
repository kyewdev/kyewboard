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


type Player struct {
    Name string
    Id int
    Level int
    Experience int
    Skills map[string]int
}

