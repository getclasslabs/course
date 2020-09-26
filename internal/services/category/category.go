package category

import (
	"github.com/getclasslabs/course/internal/repository/category"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

func GetCategories(i *tracer.Infos) ([]map[string]interface{}, error) {
	i.TraceIt("getting category")
	defer i.Span.Finish()

	r := category.NewCategory()
	categories, err := r.GetAll(i)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
