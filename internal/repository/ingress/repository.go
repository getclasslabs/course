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
		traceName: "solicitation repository",
	}
}

func (s *Solicitation) Create(i *tracer.Infos, courseID, studentId, text, image string) error {
	i.TraceIt(s.traceName)
	defer i.Span.Finish()

	q := "INSERT INTO course_ingress_solicitation (student_id, course_id, text, image) " +
		"VALUES (?, ?, ?, ?);"

	_, err := s.db.Insert(i, q,
		studentId,
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

func (s *Solicitation) GetRequestsToCourse(i *tracer.Infos, courseID int, email string) ([]domain.IngressSolicitationResponse, error) {
	i.TraceIt(s.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"	cs.id, " +
		"	cs.student_id as studentID, " +
		"	cs.course_id as courseID, " +
		"	cs.text, " +
		"	cs.image, " +
		"	cs.created_at as createdAt, " +
		"	us.nickname, " +
		"	us.first_name as firstName, " +
		"	us.last_name as lastName " +
		"FROM course_ingress_solicitation cs " +
		"INNER JOIN course c ON cs.course_id = c.id " +
		"INNER JOIN teacher t ON c.teacher_id = t.id " +
		"INNER JOIN users u ON t.user_id = u.id " +
		"INNER JOIN students s ON cs.student_id = s.id " +
		"INNER JOIN users us ON s.user_id = us.id " +
		"LEFT JOIN course_registration cr ON cs.id = cr.course_ingress_solicitation_id " +
		"WHERE " +
		"	u.email = ? AND " +
		"	c.id = ? AND " +
		"	cr.id IS NULL"

	result, err := s.db.Fetch(i, q, email, courseID)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	var solicitations []domain.IngressSolicitationResponse
	err = mapper(result, &solicitations)
	if err != nil{
		i.LogError(err)
		return nil, err
	}

	return solicitations, nil
}

func (s *Solicitation) DelRequestsToCourse(i *tracer.Infos, solicitationID int) error {
	i.TraceIt(s.traceName)
	defer i.Span.Finish()

	q := "DELETE " +
		"FROM course_ingress_solicitation " +
		"WHERE " +
		"	id = ?"

	_, err := s.db.Fetch(i, q, solicitationID)
	if err != nil{
		i.LogError(err)
		return err
	}

	return nil
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
