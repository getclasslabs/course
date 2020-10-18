package domain

type Course struct {
	Email                   string  `json:"email,omitempty"`
	ID                      int     `json:"id,omitempty"`
	Name                    string  `json:"name"`
	Description             string  `json:"description"`
	CategoryID              int     `json:"categoryID"`
	MaxStudents             int     `json:"maxStudents"`
	Classes                 int     `json:"classes"`
	Periods                 string  `json:"periods"`
	Price                   float64 `json:"price"`
	StartDay                string  `json:"startDay"`
	Type                    string  `json:"type"`
	Place                   string  `json:"place"`
	AllowStudentsAfterStart bool    `json:"allowStudentsAfterStart"`
	ClassOpen               bool    `json:"classOpen"`
	ClassesGiven			int 	`json:"classesGiven"`
	CreatedAt               string  `json:"createdAt"`
	Active                  bool    `json:"active"`
	ImagePath				string 	`json:"imagePath"`
}