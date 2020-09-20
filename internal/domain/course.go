package domain

type Course struct {
	Email       string
	ID          int
	Name        string   `json:"name"`
	CategoryID  int      `json:"categoryID"`
	MaxStudents int      `json:"maxStudents"`
	Classes     int      `json:"numClasses"`
	Periods     []Period `json:"periods"`
}

type Period struct {
	Day  int `json:"weekDay"`
	Hour int `json:"hour"`
}
