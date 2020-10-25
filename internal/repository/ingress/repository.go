package ingress

import (
	"encoding/json"
	"github.com/getclasslabs/course/internal/domain"
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
	return nil
}

func (s *Solicitation) GetRequestsToCourse(i *tracer.Infos, courseID int, email string) ([]domain.IngressSolicitation, error) {
	i.TraceIt(s.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"	cs.student_id, " +
		"	cs.course_id as courseID, " +
		"	cs.text, " +
		"	cs.image, " +
		"	cs.created_at as createdAt " +
		"FROM course_ingress_solicitation cs " +
		"INNER JOIN course c ON cs.course_id = c.id " +
		"INNER JOIN teacher t ON c.teacher_id = t.id " +
		"INNER JOIN user u ON t.user_id = u.id " +
		"WHERE " +
		"	u.email = ? AND " +
		"	c.id = ?"

	result, err := s.db.Fetch(i, q, email, courseID)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	var solicitations []domain.IngressSolicitation
	err = mapper(result, solicitations)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	return solicitations, nil
}

func mapper(data interface{}, to interface{}) error {
	jsonResult, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonResult, to)
	if err != nil {
		return err
	}
	return nil
}
