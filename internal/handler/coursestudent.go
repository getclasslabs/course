package handler

import (
	"encoding/json"
	"github.com/getclasslabs/course/internal/services/course"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CourseStudent(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	email := r.Header.Get("X-Consumer-Username")
	courseID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := course.GetCourse(i, email, courseID)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret, _ := json.Marshal(c)

	_, _ = w.Write(ret)
}