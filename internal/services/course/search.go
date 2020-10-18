package course

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository/course"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func Search(i *tracer.Infos, name string) ([]domain.Course, error) {
	i.TraceIt("creating service")
	defer i.Span.Finish()

	cRepo := course.NewCourse()
	return cRepo.Search(i, name)
}
