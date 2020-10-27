package course

import (
	"bytes"
	"errors"
	"github.com/disintegration/imaging"
	"github.com/getclasslabs/course/internal/repository/course"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"image"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func UpdateImage(i *tracer.Infos, email, courseId string, file multipart.File) (string, error) {
	i.TraceIt("updating photo")
	defer i.Span.Finish()

	now := time.Now()      // current local time
	name := strconv.Itoa(int(now.Unix())) + ".png"

	isValid , err := validUpdate(i, email, courseId)
	if !isValid || err != nil {
		i.LogError(err)
		return "", err
	}
	if !isValid {
		err = errors.New("user not allowed")
		i.LogError(err)
		return "", err
	}

	photoFile, err := os.Create("./course_photos/" + name)
	if err != nil {
		i.LogError(err)
		return "", err
	}
	defer photoFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		i.LogError(err)
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		i.LogError(err)
		return "", err
	}

	resized := imaging.Resize(img, 200, 200, imaging.Lanczos)

	enc := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	err = enc.Encode(photoFile, resized)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	cRepo := course.NewCourse()
	err = cRepo.UpdateImage(i, courseId, name)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	return name, nil
}

func ErasePhoto(i *tracer.Infos, courseId string) error {
	cRepo := course.NewCourse()
	resp, err := cRepo.GetCourseById(i, courseId)
	if err != nil{
		i.LogError(err)
		return err
	}
	filename, ok := resp["image"].(string)

	if !ok || len(filename) == 0 {
		err := errors.New("could not get course image")
		i.LogError(err)
		return err
	}

	err = cRepo.UpdatePhoto(i, courseId, "")
	if err != nil{
		i.LogError(err)
		return err
	}

	//The file could not be removed but the register was updated, must remove manually
	err = os.Remove("./course_photos/" + filename)
	if err != nil{
		i.LogError(err)
	}

	return nil
}

func validUpdate(i *tracer.Infos, email, id string) (bool, error) {
	cRepo := course.NewCourse()
	courses, err := cRepo.GetCourseFromUser(i, email, id)

	if err != nil || len(courses) == 0 {
		return false, err
	}

	return true, nil
}