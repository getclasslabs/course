package ingress

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository/ingress"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func Accept(i *tracer.Infos, acceptance *domain.IngressAcceptance) error {
	i.TraceIt("acceptance service")
	defer i.Span.Finish()

	aRepo := ingress.NewAcceptance()
	return aRepo.Create(i, acceptance)
}

