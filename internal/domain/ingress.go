package domain

type IngressSolicitation struct {
	Email     string         `json:"email,omitempty"`
	CourseID  int            `json:"courseID"`
	Text      string         `json:"text"`
	//Image     multipart.File `json:"image"`
	CreatedAt string         `json:"createdAt,omitempty"`
}

type IngressAcceptance struct {
	Email                 string         `json:"email,omitempty"`
	CourseID              int            `json:"courseID"`
	IngressSolicitationID int            `json:"ingressSolicitationID"`
	CreatedAt             string         `json:"createdAt,omitempty"`
}
