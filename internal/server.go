package internal

import (
	"github.com/getclasslabs/course/internal/handler"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	s := Server{r}

	s.serve()

	return &s
}

func (s *Server) serve() {
	s.Router.Path("/heartbeat").HandlerFunc(request.PreRequest(handler.Heartbeat)).Methods(http.MethodGet)
	s.Router.Path("/categories").HandlerFunc(request.PreRequest(handler.GetCategory)).Methods(http.MethodGet)

	//Course CRUD
	s.Router.Path("/create").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodPost)
	s.Router.Path("/{id}").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodGet)
	s.Router.Path("/edit/{id}").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodPut)
	s.Router.Path("/{id}").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodDelete)


	//Student Course
	s.Router.Path("/s/{id}").HandlerFunc(request.PreRequest(handler.CourseStudent)).Methods(http.MethodGet)

}
