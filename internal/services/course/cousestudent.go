package course

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository/course"
	"github.com/getclasslabs/course/internal/repository/ingress"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func isRegistered(i *tracer.Infos, email string, courseID int) (bool, error){
	aRepo := ingress.NewAcceptance()
	registered, err := aRepo.Get(i, email, courseID)
	if err != nil{
		return false, err
	}
	return registered, nil
}

func isSolicitation(i *tracer.Infos, email string, courseID int) (bool, error){
	aRepo := ingress.NewAcceptance()
	solicitation, err := aRepo.GetSolicitation(i, email, courseID)
	if err != nil{
		return false, err
	}
	return solicitation, nil
}

func GetCourse(i *tracer.Infos, email string, courseID int) (*domain.Course, error) {
	i.TraceIt("getting course")
	defer i.Span.Finish()

	solicitation, err := isSolicitation(i, email, courseID)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	approved, err := isRegistered(i, email, courseID)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	cRepo := course.NewCourse()
	if solicitation {
		return cRepo.GetToRegistered(i, courseID, approved)
	}
	return cRepo.GetToNotRegistered(i, courseID)
}