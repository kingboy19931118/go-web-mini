package model

import (
	"time"
)

type Student struct {
	ID            int64     `gorm:"type:bigint;not null;unique" json:"id"`
	Name          string    `gorm:"type:varchar(50);not null;" json:"name"`
	Age           int       `gorm:"type:int(3);not null;" json:"age"`
	Gender        int       `gorm:"type:int(2);not null;" json:"gender"`
	ClassHour     int       `gorm:"type:int(10);not null;" json:"classHour"`
	LeftClassHour int       `gorm:"type:int(10);not null;" json:"leftClassHour"`
	SignAmount    float64   `gorm:"type:decimal(10,2);not null;" json:"signAmount"`
	OpenID        *string   `gorm:"type:varchar(50);" json:"openID,omitempty"`
	Extra         string    `gorm:"type:text;not null;" json:"extra"`
	Mobile        string    `gorm:"type:varchar(15)" json:"mobile,omitempty"`
	CreateAt      time.Time `gorm:"type:datetime" json:"createAt"`
	ModifyAt      time.Time `gorm:"type:datetime" json:"modifyAt"`
	Status        int       `gorm:"type:int(2)" json:"status"`
}

func (s Student) TableName() string {
	return "student"
}

var (
	NoActive = 1
	Active   = 10
	Finish   = 20
	Terminal = 90
	Delete   = 99
)
