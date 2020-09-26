package classdate

import (
	"github.com/getclasslabs/course/internal/domain"
	"reflect"
	"testing"
)

func Test_buildClassQuery(t *testing.T) {
	type args struct {
		courseId int
		classes  []domain.Period
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
	}{
		{
			"test1",
			args{
				1,
				[]domain.Period{
					{
						2,
						12,
					},
				},
			},
			"INSERT INTO class_date(course_id, day, hour) VALUES (?, ?, ?)",
			[]interface{}{1, 2, 12},
		},
		{
			"test2",
			args{
				1,
				[]domain.Period{
					{
						2,
						12,
					},
					{
						4,
						12,
					},
					{
						6,
						12,
					},
				},
			},
			"INSERT INTO class_date(course_id, day, hour) VALUES (?, ?, ?),(?, ?, ?),(?, ?, ?)",
			[]interface{}{1, 2, 12, 1, 4, 12, 1, 6, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := buildClassQuery(tt.args.courseId, tt.args.classes)
			if got != tt.want {
				t.Errorf("buildClassQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("buildClassQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}