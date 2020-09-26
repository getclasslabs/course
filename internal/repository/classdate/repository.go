package classdate

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository"
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type ClassDate struct {
	db        db.Database
	traceName string
}

func NewClassDate() *ClassDate {
	return &ClassDate{
		db:        repository.Db,
		traceName: "classdate repository",
	}
}

func (c *ClassDate) Create(i *tracer.Infos, courseId int, classes []domain.Period) error {
	q, valuesUnpack := buildClassQuery(courseId, classes)

	_, err := c.db.Insert(i, q, valuesUnpack...)

	if err != nil {
		i.LogError(err)
		return err
	}
	return nil
}

func buildClassQuery(courseId int, classes []domain.Period) (string, []interface{}){
	var values string
	var valuesUnpack []interface{}
	for _, period := range classes {
		if len(values) > 0 {
			values += ","
		}
		values += "(?, ?, ?)"

		valuesUnpack = append(valuesUnpack, courseId, period.Day, period.Hour)
	}

	q := "INSERT INTO class_date(course_id, day, hour) VALUES " + values

	return q, valuesUnpack
}
