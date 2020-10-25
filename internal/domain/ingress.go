package domain

import "mime/multipart"

type IngressSolicitation struct {
	Email 		string  `json:"email,omitempty"`
	CourseID 	int		`json:"courseID"`
	Text		string  `json:"text"`
	Image		multipart.File `json:"image"`
	CreatedAt   string  `json:"createdAt,omitempty"`
}
