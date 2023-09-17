package vo

import "time"

// 创建用户结构体
type CreateStudentRequest struct {
	Name          string    `form:"name" json:"name" validate:"required,min=1,max=10"`
	Age           int       `form:"age" json:"age,omitempty"`
	Gender        int       `form:"gender" json:"gender,omitempty" validate:"oneof=1 2 3"`
	ClassHour     int       `form:"classHour" json:"classHour,omitempty"`
	LeftClassHour int       `form:"leftClassHour" json:"leftClassHour,omitempty"`
	SignAmount    float64   `form:"signAmount" json:"signAmount,omitempty"`
	Status        int       `form:"status" json:"status" validate:"oneof=1 10 20 90 99"`
	InDate        time.Time `form:"inDate" json:"inDate,omitempty"`
	Mobile        string    `form:"mobile" json:"mobile,omitempty" validate:"checkMobile"`
}

type UpdateStudentRequest struct {
	ID            int64     `form:"id" json:"ID" validate:"required"`
	Name          string    `form:"name" json:"name" validate:"min=1,max=10"`
	Age           int       `form:"age" json:"age"`
	Gender        int       `form:"gender" json:"gender" validate:"oneof=1 2 3"`
	ClassHour     int       `form:"classHour" json:"classHour"`
	LeftClassHour int       `json:"leftClassHour" json:"leftClassHour"`
	SignAmount    float64   `form:"signAmount" json:"signAmount"`
	InDate        time.Time `form:"inDate" json:"inDate"`
	Mobile        string    `form:"mobile" json:"mobile" validate:"checkMobile"`
	Status        int       `form:"status" json:"status" validate:"oneof=1 10 20 90 99"`
	OpenID        string    `form:"openID" json:"openID"`
}

// 获取用户列表结构体
type StudentListRequest struct {
	Name     string `json:"name" form:"name"`
	Gender   int    `json:"gender" form:"gender"`
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 批量删除用户结构体
type DeleteStudentRequest struct {
	StudentIds []int64 `json:"studentIDs" form:"studentIDs"`
}
