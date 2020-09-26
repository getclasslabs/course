package course

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository/classdate"
	"github.com/getclasslabs/course/internal/repository/course"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type Course struct {
	Domain *domain.Course
}

func (c *Course) Create(i *tracer.Infos) error {
	i.TraceIt("creating service")
	defer i.Span.Finish()

	cRepo := course.NewCourse()
	courseID, err := cRepo.Create(i, c.Domain)
	if err != nil {
		return err
	}

	classRepo := classdate.NewClassDate()
	err = classRepo.Create(i, courseID, c.Domain.Periods)
	if err != nil {
		return err
	}

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
