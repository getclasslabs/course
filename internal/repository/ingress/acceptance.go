package ingress

import (
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository"
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type Acceptance struct {
	db        db.Database
	traceName string
}

func NewAcceptance() *Acceptance {
	return &Acceptance{
		db:        repository.Db,
		traceName: "acceptance repository",
	}
}

func (a *Acceptance) Create(i *tracer.Infos, acceptance *domain.IngressAcceptance) error {
	i.TraceIt(a.traceName)
	defer i.Span.Finish()

	q := "INSERT INTO course_registration (student_id,course_id,course_ingress_solicitation_id) VALUES " +
		"	((SELECT s.id FROM students s INNER JOIN users u ON u.id = s.user_id WHERE u.email = ?), ?, ?);"

	_, err := a.db.Insert(i, q,
		acceptance.Email,
		acceptance.CourseID,
		acceptance.IngressSolicitationID,
	)

	if err != nil {
		i.LogError(err)
		return err
	}
	return nil
}
