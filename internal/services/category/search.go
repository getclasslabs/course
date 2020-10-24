package category

import (
	"github.com/getclasslabs/course/internal/config"
	"github.com/getclasslabs/course/internal/repository/category"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"strconv"
)

func Search(i *tracer.Infos, name, page string) (map[string]interface{}, error) {
	i.TraceIt("creating service")
	defer i.Span.Finish()
	limit := config.Config.SearchLimit

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 0 {
		pageNumber = 1
	}

	cRepo := category.NewCategory()

	categories, err := cRepo.Search(i, name, pageNumber)
	if err != nil {
		return nil, err
	}

	next, err := cRepo.GetNextPageCategory(i, name)
	if err != nil {
		return nil, err
	}

	hasNextCount := (pageNumber * limit) + 1
	var hasNext bool
	if len(next) > 0 && next["count"].(int64) >= int64(hasNextCount)  {
		hasNext = true
	} else {
		hasNext = false
	}

	return map[string]interface{}{
		"next": hasNext,
		"results": categories,
	}, err
}
