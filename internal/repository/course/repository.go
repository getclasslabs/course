package course

import (
	"encoding/json"
	"errors"
	"github.com/getclasslabs/course/internal/config"
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

func (c *Course) Create(i *tracer.Infos, course *domain.Course) (int, error) {
	q := "INSERT INTO course (teacher_id, name, description, category_id, max_students, classes, periods, price, " +
		"start_day, type, place, payment, allow_students_after_start)" +
		"VALUES ((SELECT t.id" +
		"         FROM teacher t" +
		"                  INNER JOIN users u on u.id = t.user_id" +
		"         where u.email = ?), ?, ?, ?, ?, ?, ?, ?, FROM_UNIXTIME(?), ?, ?, ?, ?);"

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
		course.Place,
		course.Payment,
		course.AllowStudentsAfterStart,
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

func (c *Course) Get(i *tracer.Infos, id int, email string) (*domain.Course, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	query := "SELECT " +
		"	id," +
		"	name," +
		"	teacher_id," +
		"	description," +
		"	category_id as categoryID," +
		"	max_students as maxStudents," +
		"	classes," +
		"	periods," +
		"	price," +
		"	start_day as startDay," +
		"	type," +
		"	place," +
		"	allow_students_after_start," +
		"	payment," +
		"	class_open as classOpen," +
		"	classes_given as classesGiven," +
		"	created_at as createdAt, " +
		"	active, " +
		"	image " +
		"FROM course " +
		"WHERE " +
		"	id = ? AND " +
		"	teacher_id = (SELECT t.id FROM teacher t INNER JOIN users u on u.id = t.user_id where u.email = ?) AND " +
		"	active = true"

	result, err := c.db.Get(i, query, id, email)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	if len(result) == 0{
		err = errors.New("no course found")
		i.LogError(err)
		return nil, err
	}

	result["classOpen"] = result["classOpen"].(int64) != 0
	result["active"] = result["active"].(int64) != 0
	result["price"], err = strconv.ParseFloat(result["price"].(string), 64)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	if result["allow_students_after_start"].(int64) == 1 {
		result["allow_students_after_start"] = true
	} else {
		result["allow_students_after_start"] = false
	}

	course := domain.Course{}
	err = mapper(result, &course)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return &course, nil
}

func (c *Course) Update(i *tracer.Infos, course *domain.Course) error {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	query := "UPDATE course SET " +
		"	name = ?," +
		"	description = ?," +
		"	category_id = ?," +
		"	max_students = ?," +
		"	classes = ?," +
		"	periods = ?," +
		"	price = ?," +
		"	start_day = FROM_UNIXTIME(?)," +
		"	type = ?," +
		"	place = ?," +
		"	payment = ?," +
		"	allow_students_after_start = ? " +
		"WHERE " +
		"	id = ? AND" +
		"	teacher_id = (SELECT t.id FROM teacher t INNER JOIN users u on u.id = t.user_id where u.email = ?)"

	_, err := c.db.Update(i, query,
		course.Name,
		course.Description,
		course.CategoryID,
		course.MaxStudents,
		course.Classes,
		course.Periods,
		course.Price,
		course.StartDay,
		course.Type,
		course.Place,
		course.Payment,
		course.AllowStudentsAfterStart,
		course.ID,
		course.Email)

	if err != nil {
		i.LogError(err)
		return err
	}

	return nil
}

func (c *Course) Delete(i *tracer.Infos,  id int, email string) error {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	query := "UPDATE course SET " +
		"	active = false " +
		"WHERE " +
		"	id = ? AND " +
		"	teacher_id = (SELECT t.id FROM teacher t INNER JOIN users u on u.id = t.user_id where u.email = ?)"

	_, err := c.db.Update(i, query, id, email)

	if err != nil {
		i.LogError(err)
		return err
	}

	return nil
}

func (c *Course) Search(i *tracer.Infos, name string, page int) ([]domain.Course, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	limit := config.Config.SearchLimit
	offset := (page - 1) * limit

	query := "SELECT " +
		"	c.id as id," +
		"	c.name as name," +
		"	c.description as description," +
		"	ca.name as categoryName," +
		"	ca.id as categoryID," +
		"	c.start_day as startDay," +
		"	c.price," +
		"	c.type," +
		" 	c.image," +
		"	c.created_at as createdAt " +
		"FROM course c " +
		"INNER JOIN category ca ON ca.id = c.category_id " +
		"WHERE " +
		"	soundex(c.name) = soundex(?) AND " +
		"	active is true " +
		"LIMIT ? " +
		"OFFSET ?"

	result, err := c.db.Fetch(i, query, name, limit, offset)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	for _, c := range result {
		c["price"], err = strconv.ParseFloat(c["price"].(string), 64)
		if err != nil {
			i.LogError(err)
			return nil, err
		}
	}
	var courses []domain.Course
	err = mapper(result, &courses)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return courses, nil
}

func (c *Course) GetNextPageCourse(i *tracer.Infos, name string) (map[string]interface{}, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"	count(c.id) as count " +
		"FROM course c " +
		"INNER JOIN category ca ON ca.id = c.category_id " +
		"INNER JOIN teacher te ON te.id = c.teacher_id " +
		"INNER JOIN users u ON u.id = te.user_id " +
		"WHERE " +
		"	soundex(c.name) = soundex(?) AND " +
		"	active is true "

	result, err := c.db.Get(i, q, name)

	if err != nil {
		i.LogError(err)
		return nil, err
	}
	return result, nil
}

func (c *Course) GetToRegistered(i *tracer.Infos, courseID int) (*domain.Course, error){
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	fields := "c.class_open, " +
		"	c.place, " +
		"	" +
		"	c.name as name, " +
		"	c.description as description," +
		"	ca.name as categoryName," +
		"	ca.id as categoryID," +
		"	c.start_day as startDay," +
		"	c.price," +
		"	c.type, " +
		"	c.image, " +
		"	c.created_at as createdAt"

	return c.getToStudent(i, courseID, fields)
}

func (c *Course) GetToNotRegistered(i *tracer.Infos, courseID int) (*domain.Course, error){
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	fields := "" +
		"	c.name as name," +
		"	c.description as description," +
		"	ca.name as categoryName," +
		"	ca.id as categoryID," +
		"	c.start_day as startDay," +
		"	c.price," +
		"	c.type, " +
		"	c.image, " +
		"	c.created_at as createdAt, " +
		"	c.periods "

	return c.getToStudent(i, courseID, fields)

}

func (c *Course) getToStudent(i *tracer.Infos, courseID int, fields string) (*domain.Course, error){
	q := "SELECT " +
		fields +
		"FROM course c " +
		"INNER JOIN category ca ON ca.id = c.category_id " +
		"WHERE " +
		"	c.id = ?"

	result, err := c.db.Get(i, q, courseID)

	if err != nil {
		i.LogError(err)
		return nil, err
	}

	ret := &domain.Course{}
	err = mapper(result, ret)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return ret, nil
}

func (c *Course) GetCourseFromUser(i *tracer.Infos, email, id string) (map[string]interface{}, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"	c.id " +
		"FROM course c " +
		"INNER JOIN teacher te ON te.id = c.teacher_id " +
		"INNER JOIN users u ON u.id = te.user_id " +
		"WHERE " +
		"	c.id = ? AND " +
		"	u.email = ? "

	result, err := c.db.Get(i, q, id, email)

	if err != nil {
		i.LogError(err)
		return nil, err
	}
	return result, nil
}

func (c *Course) UpdateImage(i *tracer.Infos, id string, name string) error {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	q := "UPDATE course SET " +
		"	image = ? " +
		"WHERE " +
		"	id = ?"

	_, err := c.db.Update(i, q, name, id)
	if err != nil {
		i.LogError(err)
		return err
	}
	return nil
}

func (c *Course) GetCourseById(i *tracer.Infos, id string) (map[string]interface{}, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"id, " +
		"image " +
		"FROM course " +
		"WHERE " +
		"id = ?"

	result, err := c.db.Get(i, q, id)

	if err != nil {
		i.LogError(err)
		return nil, err
	}
	return result, nil
}

func (c *Course) UpdatePhoto(i *tracer.Infos, id string, path string) error {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	q := "UPDATE course SET " +
		"	image = ? " +
		"WHERE " +
		"	id = ?"

	_, err := c.db.Update(i, q, path, id)
	if err != nil {
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

