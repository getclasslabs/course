package course

import (
	"encoding/json"
	"errors"
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository"
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"strconv"
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

func (c Course) Create(i *tracer.Infos, course *domain.Course) (int, error) {
	q := "INSERT INTO course (teacher_id, name, description, category_id, max_students, classes, periods, price, " +
		"start_day, type)" +
		"VALUES ((SELECT t.id" +
		"         FROM teacher t" +
		"                  INNER JOIN users u on u.id = t.user_id" +
		"         where u.email = ?), ?, ?, ?, ?, ?, ?, ?, FROM_UNIXTIME(?), ?);"

	_, err := c.db.Insert(i, q,
		course.Email,
		course.Name,
		course.Description,
		course.CategoryID,
		course.MaxStudents,
		course.Classes,
		course.Periods,
		course.Price,
		course.StartDay,
		course.Type,
		)

	if err != nil {
		i.LogError(err)
		return 0, err
	}

	q2 := "SELECT LAST_INSERT_ID() AS id;"

	result, err := c.db.Get(i, q2)
	if err != nil {
		i.LogError(err)
		return 0, err
	}

	id, ok := result["id"].(int64)
	if !ok {
		err := errors.New("unable to get LID")
		i.LogError(err)
		return 0, err
	}

	return int(id), nil
}

func (c Course) Get(i *tracer.Infos, id int, email string) (*domain.Course, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	query := "SELECT " +
		"	id," +
		"	name," +
		"	description," +
		"	category_id as categoryID," +
		"	max_students as maxStudents," +
		"	classes," +
		"	periods," +
		"	price," +
		"	start_day as startDay," +
		"	type," +
		"	place," +
		"	class_open," +
		"	classes_given," +
		"	created_at as createdAt " +
		"FROM course " +
		"WHERE " +
		"	id = ? AND" +
		"	teacher_id = (SELECT t.id FROM teacher t INNER JOIN users u on u.id = t.user_id where u.email = ?)"

	result, err := c.db.Get(i, query, id, email)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	result["price"], err = strconv.ParseFloat(result["price"].(string), 64)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	course := domain.Course{}
	err = json.Unmarshal(jsonResult, &course)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return &course, nil
}
