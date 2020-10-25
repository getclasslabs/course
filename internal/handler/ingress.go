package handler

import (
	"encoding/json"
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/services/ingress"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"net/http"
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

	receipt, _, err := r.FormFile("receipt")

	iDomain := domain.IngressSolicitation{}
	err = json.NewDecoder(r.Body).Decode(&iDomain)
	if err != nil && err.Error() != "EOF" {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	iDomain.Image = receipt

	err = ingress.Request(i, &iDomain)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
