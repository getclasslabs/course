package handler

import (
	"encoding/json"
	"github.com/getclasslabs/course/internal/services/category"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"net/http"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)

	i.TraceIt(spanName)
	defer i.Span.Finish()

	categories, err := category.GetCategories(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal server error"))
		return
	}
	ret, _ := json.Marshal(categories)

	_, _ = w.Write(ret)

}

func CategorySearch(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)

	i.TraceIt(spanName)
	defer i.Span.Finish()

	name := r.URL.Query().Get("name")
	page := r.URL.Query().Get("page")

	if page == "" {
		page = "1"
	}

	categories, err := category.Search(i, name, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ret, _ := json.Marshal(categories)

	_, _ = w.Write(ret)
}
