package category

import (
	"github.com/getclasslabs/course/internal/repository/category"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func Search(i *tracer.Infos, name string) ([]map[string]interface{}, error) {
	i.TraceIt("creating service")
	defer i.Span.Finish()

	cRepo := category.NewCategory()
	return cRepo.Search(i, name)
}
