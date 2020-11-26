package ingress

import (
	"github.com/getclasslabs/course/internal/repository/ingress"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func GetStudents(i *tracer.Infos, courseID int) ([]map[string]interface{}, error) {
	i.TraceIt("get students service")
	defer i.Span.Finish()

	sRepo := ingress.NewAcceptance()
	return sRepo.GetStudents(i, courseID)
}
