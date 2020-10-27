package domain

type Course struct {
	Email                   string  `json:"email,omitempty"`
	ID                      int     `json:"id,omitempty"`
	Name                    string  `json:"name"`
	Description             string  `json:"description"`
	CategoryID              int     `json:"categoryID"`
	CategoryName            string  `json:"categoryName,omitempty"`
	TeacherName				string  `json:"teacherName,omitempty"`
	TeacherId				int		`json:"teacher_id,omitempty"`
	MaxStudents             int     `json:"maxStudents,omitempty"`
	Classes                 int     `json:"classes,omitempty"`
	Periods                 string  `json:"periods,omitempty"`
	Price                   float64 `json:"price"`
	Payment					string	`json:"payment"`
	StartDay                string  `json:"startDay"`
	Type                    string  `json:"type"`
	Place                   string  `json:"place,omitempty"`
	AllowStudentsAfterStart bool    `json:"allowStudentsAfterStart,omitempty"`
	ClassOpen               bool    `json:"classOpen,omitempty"`
	ClassesGiven			int 	`json:"classesGiven,omitempty"`
	CreatedAt               string  `json:"createdAt"`
	Active                  bool    `json:"active,omitempty"`
	ImagePath				string 	`json:"imagePath"`
	Registered				*bool    `json:"registered,omitempty"`
	Solicitation			*bool    `json:"solicitation,omitempty"`
	StudentsRegistered		int		`json:"studentsRegistered"`
}