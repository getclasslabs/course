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

	//Course Ingress
	s.Router.Path("/ingress").HandlerFunc(request.PreRequest(handler.Ingress)).Methods(http.MethodPost)
	s.Router.Path("/solicitations/{courseID}").HandlerFunc(request.PreRequest(handler.ListSolicitations)).Methods(http.MethodGet)
	s.Router.Path("/accept").HandlerFunc(request.PreRequest(handler.AcceptSolicitation)).Methods(http.MethodPost)
	s.Router.Path("/students/{courseID}").HandlerFunc(request.PreRequest(handler.GetCourseStudents)).Methods(http.MethodGet)

	//Get
	s.Router.Path("/mine").HandlerFunc(request.PreRequest(handler.GetMyCourses)).Methods(http.MethodGet)
	s.Router.Path("/categories").HandlerFunc(request.PreRequest(handler.GetCategory)).Methods(http.MethodGet)

	//Searches
	s.Router.Path("/search").HandlerFunc(request.PreRequest(handler.Search)).Methods(http.MethodGet)
	s.Router.Path("/category/search").HandlerFunc(request.PreRequest(handler.CategorySearch)).Methods(http.MethodGet)

	//Course Image
	s.Router.Path("/image/{courseId}").HandlerFunc(request.PreRequest(handler.UpdatePhoto)).Methods(http.MethodPut)
	s.Router.Path("/image/{courseId}").HandlerFunc(request.PreRequest(handler.DeletePhoto)).Methods(http.MethodDelete)

	//Course CRUD
	s.Router.Path("/create").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodPost)
	s.Router.Path("/{id}").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodGet)
	s.Router.Path("/edit/{id}").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodPut)
	s.Router.Path("/{id}").HandlerFunc(request.PreRequest(handler.CourseCRUD)).Methods(http.MethodDelete)

	//Student Course
	s.Router.Path("/s/{id}").HandlerFunc(request.PreRequest(handler.CourseStudent)).Methods(http.MethodGet)

	s.Router.PathPrefix("/category/images/").Handler(http.StripPrefix("/category/images/",
		http.FileServer(http.Dir("./category_photos/"))))

	s.Router.PathPrefix("/course/images/").Handler(http.StripPrefix("/course/images/",
		http.FileServer(http.Dir("./course_photos/"))))



}
