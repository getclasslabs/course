package course

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository"
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type Course struct {
	db        db.Database
	traceName string
}

func NewCourse() *Course {
	return &Course{
		db:        repository.Db,
		traceName: "course repository",
	}
}

func (c Course) Create(i *tracer.Infos, course *domain.Course) error {
	return nil
}

func (c Course) Get(i *tracer.Infos, id int) (map[string]interface{}, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	query := "SELECT * FROM course"

	result, err := c.db.Get(i, query, id)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	return result, nil
}
