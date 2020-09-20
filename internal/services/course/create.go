package course

import "github.com/getclasslabs/go-tools/pkg/tracer"

type Course struct {
	Email string
	ID int
	Name string `json:"name"`
	CategoryID int `json:"categoryID"`
	MaxStudents int `json:"maxStudents"`
	Classes int `json:"numClasses"`
	Periods []Period `json:"periods"`
}

type Period struct {
	Day int `json:"weekDay"`
	Hour int `json:"hour"`
}


func (c *Course) Create(i *tracer.Infos) error {
	
	return nil
}


func (c *Course) Edit(i *tracer.Infos) error {
	return nil
}

func (c *Course) Delete(i *tracer.Infos) error {
	return nil
}

func (c *Course) Get(i *tracer.Infos) error {
	return nil
}
