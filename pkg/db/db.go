package db

type Quest struct {
	Id       int
	Message  string
	Status   string
    Objectives []Objective
	Rewards   []string
	Assignee string
    Questtype string
    Category string
}

type Objective struct {
    Done bool
    Text string
}

type Player struct {
    Name string
    Id int
    Level int
    Experience int
    Skills map[string]Skill
    Stats map[string]int
}

type Skill struct {
    Category string
    Level int
    Experience int
    Title string
}
