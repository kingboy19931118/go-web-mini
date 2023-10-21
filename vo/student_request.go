package vo

import "time"

// CreateStudentRequest 创建学生结构体
type CreateStudentRequest struct {
	Name          string    `form:"name" json:"name" validate:"required,min=1,max=10"`
	Age           int       `form:"age" json:"age,omitempty"`
	Gender        int       `form:"gender" json:"gender,omitempty" validate:"oneof=1 2 3"`
	ClassHour     int       `form:"classHour" json:"classHour,omitempty"`
	BaseClassHour int       `form:"baseClassHour" json:"baseClassHour,omitempty"`
	FreeClassHour int       `form:"freeClassHour" json:"freeClassHour,omitempty"`
	LeftClassHour int       `form:"leftClassHour" json:"leftClassHour,omitempty"`
	SignAmount    float64   `form:"signAmount" json:"signAmount,omitempty"`
	Course        string    `form:"course" json:"course"`
	HomeAddress   string    `form:"homeAddress" json:"homeAddress"`
	Gift          string    `form:"gift" json:"gift"`
	Status        int       `form:"status" json:"status" validate:"oneof=1 10 20 90 99"`
	InDate        time.Time `form:"inDate" json:"inDate,omitempty"`
	ValidateDate  time.Time `form:"validateDate" json:"validateDate,omitempty"`
	Mobile        string    `form:"mobile" json:"mobile,omitempty" validate:"checkMobile"`
}

type UpdateStudentRequest struct {
	ID            int64     `form:"id" json:"ID" validate:"required"`
	Name          string    `form:"name" json:"name" validate:"min=1,max=10"`
	Age           int       `form:"age" json:"age"`
	Gender        int       `form:"gender" json:"gender" validate:"oneof=1 2 3"`
	ClassHour     int       `form:"classHour" json:"classHour"`
	InClassHOur   int       `form:"inClassHour" json:"inClassHour"`
	LeftClassHour int       `form:"leftClassHour" json:"leftClassHour"`
	Course        int       `form:"course" json:"course"`
	SignAmount    float64   `form:"signAmount" json:"signAmount"`
	InDate        time.Time `form:"inDate" json:"inDate"`
	Mobile        string    `form:"mobile" json:"mobile" validate:"checkMobile"`
	// Status        int       `form:"status" json:"status" validate:"oneof=0 1 10 20 90 99"`
	OpenID string `form:"openID" json:"openID"`
}

// 获取用户列表结构体
type StudentListRequest struct {
	Name     string `json:"name" form:"name"`
	Gender   int    `json:"gender" form:"gender"`
	Status   int    `form:"status" json:"status" validate:"oneof=0 1 10 20 90 99"`
	Mobile   string `form:"mobile" json:"mobile" validate:"checkMobile"`
	Course   string `form:"course" json:"course"`
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 批量删除用户结构体
type DeleteStudentRequest struct {
	StudentIds []int64 `json:"studentIDs" form:"studentIDs"`
}

// 批量销课时
type BatchInClassRequest struct {
	StudentIds []int64 `json:"studentIDs" form:"studentIDs"`
}
