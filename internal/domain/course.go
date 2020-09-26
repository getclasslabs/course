package domain

type Course struct {
	Email       string
	ID          int
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CategoryID  int      `json:"categoryID"`
	StartDay    string   `json:"startDay"`
	MaxStudents int      `json:"maxStudents"`
	Classes     int      `json:"numClasses"`
	Periods     []Period `json:"periods"`
	Type        string   `json:"type"`
}

type Period struct {
	Day  int `json:"weekDay"`
	Hour int `json:"hour"`
}
