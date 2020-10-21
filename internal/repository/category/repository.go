package category

import (
	"github.com/getclasslabs/course/internal/config"
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

func (c Category) Search(i *tracer.Infos, name string, page int) ([]map[string]interface{}, error){
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
		"	c.type, " +
		"	c.created_at as createdAt " +
		"FROM category ca " +
		"INNER JOIN course c ON ca.id = c.category_id " +
		"WHERE soundex(ca.name) = soundex(?) " +
		"LIMIT ? " +
		"OFFSET ?"
	result, err := c.db.Fetch(i, query, name, limit, offset)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return result, nil
}

func (c Category) GetNextPageCategory(i *tracer.Infos, name string) (map[string]interface{}, error) {
	i.TraceIt(c.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"	count(c.id) as count " +
		"FROM category ca " +
		"INNER JOIN course c ON ca.id = c.category_id " +
		"WHERE soundex(ca.name) = soundex(?) "
	result, err := c.db.Get(i, q, name)

	if err != nil {
		i.LogError(err)
		return nil, err
	}
	return result, nil
}