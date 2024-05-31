package db

type Quest struct {
	Id       int
	Message  string
	Status   string
    Objectives []string
	Rewards   []string
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

