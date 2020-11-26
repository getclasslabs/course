package domain

import "mime/multipart"

type IngressSolicitation struct {
	Email     string         `json:"email,omitempty"`
	StudentId string		 `json:"studentID"`
	CourseID  string         `json:"courseID"`
	Text      string         `json:"text"`
	Image     multipart.File `json:"image"`
	CreatedAt string         `json:"createdAt,omitempty"`
}

type IngressSolicitationResponse struct {
	Email     string         `json:"email,omitempty"`
	ID 		  int64			 `json:"id"`
	StudentId int64		 	 `json:"studentID"`
	CourseID  int64          `json:"courseID"`
	Text      string         `json:"text"`
	Image     string		 `json:"image"`
	Nickname  string		 `json:"nickname"`
	FirstName string		 `json:"firstName"`
	LastName  string		 `json:"lastName"`
	CreatedAt string         `json:"createdAt,omitempty"`
}

type IngressAcceptance struct {
	Email                 string         `json:"email,omitempty"`
	StudentID 			  int		 	 `json:"studentID"`
	CourseID              int            `json:"courseID"`
	IngressSolicitationID int            `json:"ingressSolicitationID"`
	CreatedAt             string         `json:"createdAt,omitempty"`
}
