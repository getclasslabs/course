package handler

import (
	"encoding/json"
	"errors"
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/services/course"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	emptyCourseID = "empty course id"
	badRequest    = "badRequest"
	notFound      = "no course found"
)

//CourseCRUD Handles course's CRUD
func CourseCRUD(w http.ResponseWriter, r *http.Request) {
	var status = 0

	i := r.Context().Value(request.ContextKey).(*tracer.Infos)

	i.TraceIt(spanName)
	defer i.Span.Finish()

	var err error
	email := r.Header.Get("X-Consumer-Username")

	c := domain.Course{}
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil && err.Error() != "EOF" {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c.Email = email
	service := course.Course{Domain: &c}

	switch r.Method {
	case http.MethodPost:
		err = service.Create(i)
		if err != nil {
			break
		}
		status = http.StatusCreated
		ret, _ := json.Marshal(c)
		_, _ = w.Write(ret)
		break
	case http.MethodPut:
		err = getID(r, &c)
		if err != nil {
			break
		}
		err = service.Edit(i)
		break
	case http.MethodDelete:
		err = getID(r, &c)
		if err != nil {
			break
		}
		err = service.Delete(i)
		break
	case http.MethodGet:
		err = getID(r, &c)
		if err != nil {
			break
		}
		response, err := service.Get(i)
		if err != nil{
			break
		}
		ret, _ := json.Marshal(response)
		_, _ = w.Write(ret)
		break
	}

	if err != nil {
		switch err.Error() {
		case notFound:
			w.WriteHeader(http.StatusNotFound)
			return
		case emptyCourseID:
		case badRequest:
			w.WriteHeader(http.StatusBadRequest)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if status > 0 {
		w.WriteHeader(status)
	}

}

func getID(r *http.Request, c *domain.Course) error {
	courseID := mux.Vars(r)["id"]
	if len(courseID) == 0 {
		err := errors.New(emptyCourseID)
		return err
	}
	id, err := strconv.Atoi(courseID)
	if err != nil {
		err = errors.New(badRequest)
		return err
	}
	c.ID = id
	return nil
}


func Search(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	name := r.URL.Query().Get("name")

	courses, err := course.Search(i, name)
	if err != nil{
		if err.Error() == notFound{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret, _ := json.Marshal(courses)
	_, _ = w.Write(ret)
}

