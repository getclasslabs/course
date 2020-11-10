package ingress

import (
	"bytes"
	"github.com/getclasslabs/course/internal/domain"
	"github.com/getclasslabs/course/internal/repository/ingress"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func Request(i *tracer.Infos, solicitation *domain.IngressSolicitation) error {
	i.TraceIt("solicitation service")
	defer i.Span.Finish()

	var err error
	var imagePath string

	if solicitation.Image != nil{
		imagePath, err = saveReceipt(i, solicitation)
		if err != nil{
			return err
		}
	}

	solRepo := ingress.NewSolicitation()
	err = solRepo.Create(i, solicitation.CourseID, solicitation.StudentId, solicitation.Text, imagePath)
	if err != nil{
		return err
	}

	return nil
}

func saveReceipt(i *tracer.Infos, solicitation *domain.IngressSolicitation) (string, error) {
	now := time.Now()      // current local time
	cID := solicitation.CourseID
	name := strconv.Itoa(int(now.Unix())) + "_" + solicitation.StudentId + "_" + cID + ".png"

	receipt, err := os.Create("./receipt_photos/" + name)
	if err != nil {
		i.LogError(err)
		return "", err
	}
	defer receipt.Close()

	fileBytes, err := ioutil.ReadAll(solicitation.Image)
	if err != nil {
		i.LogError(err)
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		i.LogError(err)
		return "", err
	}

	enc := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	err = enc.Encode(receipt, img)
	if err != nil{
		i.LogError(err)
		return "", err
	}
	return name, nil
}

func ListRequests(i *tracer.Infos, courseID int, email string) ([]domain.IngressSolicitationResponse, error) {
	i.TraceIt("solicitation service")
	defer i.Span.Finish()

	solRepo := ingress.NewSolicitation()
	return solRepo.GetRequestsToCourse(i, courseID, email)
}
