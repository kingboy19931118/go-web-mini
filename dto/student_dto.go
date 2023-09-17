package dto

import (
	"go-web-mini/model"
)

type StudentDTO struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Age           int     `json:"age"`
	Gender        string  `json:"gender"`
	ClassHour     int     `json:"classHour"`
	LeftClassHour int     `json:"leftClassHour"`
	SignAmount    float64 `json:"signAmount"`
	Extra         string  `json:"extra"`
	Mobile        string  `json:"mobile"`
	CreateAt      string  `json:"createAt"`
	ModifyAt      string  `json:"modifyAt"`
	Status        string  `json:"status"`
}

var (
	genderMap = map[int]string{
		1: "男", 2: "女", 3: "-",
	}

	statusMap = map[int]string{
		1:  "未激活",
		10: "在学",
		20: "毕业",
		90: "退学",
		99: "删除",
	}
)

func ToStudentsDto(students []*model.Student) []StudentDTO {
	var studentDTOS []StudentDTO
	for _, student := range students {
		studentDTO := StudentDTO{
			ID:            student.ID,
			Name:          student.Name,
			Age:           student.Age,
			Gender:        genderMap[student.Gender],
			ClassHour:     student.ClassHour,
			LeftClassHour: student.LeftClassHour,
			SignAmount:    student.SignAmount,
			Extra:         student.Extra,
			Mobile:        student.Mobile,
			CreateAt:      student.CreateAt.Format("2006-01-02 15:04:05"),
			ModifyAt:      student.ModifyAt.Format("2006-01-02 15:04:05"),
			Status:        statusMap[student.Status],
		}

		studentDTOS = append(studentDTOS, studentDTO)
	}

	return studentDTOS
}
