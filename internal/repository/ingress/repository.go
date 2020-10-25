package ingress

import (
	"github.com/getclasslabs/course/internal/repository"
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type Solicitation struct {
	db        db.Database
	traceName string
}

func NewSolicitation() *Solicitation {
	return &Solicitation{
		db:        repository.Db,
		traceName: "category repository",
	}
}

func (s *Solicitation) Create(i *tracer.Infos, courseID int, studentEmail, text, image string) error {
	i.TraceIt(s.traceName)
	defer i.Span.Finish()

	q := "INSERT INTO course_ingress_solicitation (student_id, course_id, text, image) " +
		"VALUES (?, ?, ?, ?);"

	_, err := s.db.Insert(i, q,
		studentEmail,
		courseID,
		text,
		image,
	)

	if err != nil {
		i.LogError(err)
		return err
	}

}