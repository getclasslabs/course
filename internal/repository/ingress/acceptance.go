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

func (a *Acceptance) Get(i *tracer.Infos, email string, courseID int) (bool, error){
	i.TraceIt(a.traceName)
	defer i.Span.Finish()

	q := "SELECT" +
		"	COUNT(*) as exist " +
		"FROM course_registration " +
		"WHERE" +
		"	course_id = ? AND" +
		"	student_id = (SELECT s.id FROM students s INNER JOIN users u ON u.id = s.user_id WHERE u.email = ?) AND" +
		"	valid = true"

	result, err := a.db.Get(i, q, courseID, email)

	if err != nil {
		i.LogError(err)
		return false, err
	}
	if len(result) > 0 {
		return result["exist"].(int64) > 0, nil
	}

	return false, nil
}

func (a *Acceptance) GetSolicitation(i *tracer.Infos, email string, courseID int) (bool, error){
	i.TraceIt(a.traceName)
	defer i.Span.Finish()

	q := "SELECT" +
		"	COUNT(*) as exist " +
		"FROM course_ingress_solicitation " +
		"WHERE" +
		"	course_id = ? AND" +
		"	student_id = (SELECT s.id FROM students s INNER JOIN users u ON u.id = s.user_id WHERE u.email = ?)"

	result, err := a.db.Get(i, q, courseID, email)

	if err != nil {
		i.LogError(err)
		return false, err
	}
	if len(result) > 0 {
		return result["exist"].(int64) > 0, nil
	}

	return false, nil
}

func (a *Acceptance) GetStudents(i *tracer.Infos, email string, courseID int) ([]map[string]interface{}, error) {
	i.TraceIt(a.traceName)
	defer i.Span.Finish()
	q := "SELECT" +
		"	r.student_id as studentID " +
		"FROM course_registration r " +
		"INNER JOIN course c ON c.id = r.course_id " +
		"INNER JOIN teacher t ON t.id = c.teacher_id " +
		"INNER JOIN users u ON u.id = t.user_id " +
		"WHERE" +
		"	course_id = ? AND " +
		"	valid = true"

	result, err := a.db.Fetch(i, q, courseID, email)

	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return result, nil
}