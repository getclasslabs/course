package course

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository/course"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func isRegistered(email string, courseID int) (bool, error){
	return true, nil
}

func GetCourse(i *tracer.Infos, email string, courseID int) (*domain.Course, error) {
	i.TraceIt("getting course")
	defer i.Span.Finish()

	registered, err := isRegistered(email, courseID)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	cRepo := course.NewCourse()
	if registered {
		return cRepo.GetToRegistered(i, courseID)
	}
	return cRepo.GetToNotRegistered(i, courseID)
}