package handler

import (
	"encoding/json"
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/services/ingress"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Ingress(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"msg": "The image sent is bigger than 10mb"}`))
	}

	iDomain := domain.IngressSolicitation{}

	receipt, _, err := r.FormFile("receipt")
	if err != nil && err.Error() != "EOF" && err.Error() != "http: no such file" {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if receipt != nil {
		iDomain.Image = receipt
	}

	iDomain.StudentId = r.FormValue("studentID")
	iDomain.CourseID = r.FormValue("courseID")
	iDomain.Text = r.FormValue("text")

	err = ingress.Request(i, &iDomain)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ListSolicitations(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	courseID, err := strconv.Atoi(mux.Vars(r)["courseID"])
	email := r.Header.Get("X-Consumer-Username")

	solicitations, err := ingress.ListRequests(i, courseID, email)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret, _ := json.Marshal(solicitations)
	_, _ = w.Write(ret)
}

func AcceptSolicitation(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	acceptance := domain.IngressAcceptance{}
	err := json.NewDecoder(r.Body).Decode(&acceptance)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ingress.Accept(i, &acceptance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetCourseStudents(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	courseID, err := strconv.Atoi(mux.Vars(r)["courseID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	students, err := ingress.GetStudents(i, courseID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret, _ := json.Marshal(students)

	_, _ = w.Write(ret)
}
