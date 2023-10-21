package dto

import (
	"go-web-mini/model"
	"time"
)

type StudentDTO struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Age              int       `json:"age"`
	Gender           string    `json:"gender"`
	ClassHour        int       `json:"classHour"`
	LeftClassHour    int       `json:"leftClassHour"`
	BaseClassHour    int       `json:"baseClassHour"`
	FreeClassHour    int       `json:"freeClassHour"`
	InClassHour      int       `json:"inClassHour"`
	RegistDate       string    `json:"registDate"`
	HomeAddress      string    `json:"homeAddress"`
	ValidateTime     time.Time `json:"validateTime"`
	SignAmount       float64   `json:"signAmount"`
	UnitPrice        float64   `json:"unitPrice"`
	Balance          float64   `json:"balance"`
	SettlementAmount float64   `json:"settlementAmount"`
	OpenID           string    `json:"openID"`
	Extra            string    `json:"extra"`
	Mobile           string    `json:"mobile"`
	CreateAt         string    `json:"createAt"`
	ModifyAt         string    `json:"modifyAt"`
	Status           string    `json:"status"`
	Course           string    `json:"course"`
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
			ID:               student.ID,
			Name:             student.Name,
			Age:              student.Age,
			Gender:           genderMap[student.Gender],
			ClassHour:        student.ClassHour,
			LeftClassHour:    student.LeftClassHour,
			BaseClassHour:    student.BaseClassHour,
			FreeClassHour:    student.FreeClassHour,
			InClassHour:      student.InClassHour,
			RegistDate:       student.RegistDate,
			HomeAddress:      student.HomeAddress,
			ValidateTime:     student.ValidateTime,
			SignAmount:       student.SignAmount,
			UnitPrice:        student.UnitPrice,
			Balance:          student.Balance,
			SettlementAmount: student.SettlementAmount,
			Extra:            student.Extra,
			Mobile:           student.Mobile,
			Status:           statusMap[student.Status],
			Course:           student.Course,
		}

		studentDTOS = append(studentDTOS, studentDTO)
	}

	return studentDTOS
}
