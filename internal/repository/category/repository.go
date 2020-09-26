package category

import (
	"github.com/getclasslabs/course/internal/repository"
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type Category struct {
	db        db.Database
	traceName string
}

func NewCategory() *Category {
	return &Category{
		db:        repository.Db,
		traceName: "category repository",
	}
}

func (c Category) GetAll(i *tracer.Infos) ([]map[string]interface{}, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	query := "SELECT * FROM category"

	result, err := c.db.Fetch(i, query)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return result, nil
}
