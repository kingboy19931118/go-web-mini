package model

import (
	"encoding/json"
	"time"
)

type Student struct {
	ID            int64  `gorm:"type:bigint;not null;unique" json:"id"`
	Name          string `gorm:"type:varchar(50);not null;" json:"name"`
	Age           int    `gorm:"type:int(3);not null;" json:"age"`
	Gender        int    `gorm:"type:int(2);not null;" json:"gender"`
	ClassHour     int    `gorm:"type:int(10);" json:"classHour"`
	BaseClassHour int    `gorm:"type:int(10);" json:"baseClassHour"`
	FreeClassHour int    `gorm:"type:int(10);" json:"freeClassHour"`
	InClassHour   int    `gorm:"type:int(10);" json:"inClassHour"`
	LeftClassHour int    `gorm:"type:int(10);" json:"leftClassHour"`
	Course        string `gorm:"type:varchar(100)" json:"course"`

	SignAmount       float64 `gorm:"type:decimal(10,2);" json:"signAmount"`
	UnitPrice        float64 `gorm:"type:decimal(10,2);" json:"UnitPrice"`
	Balance          float64 `gorm:"type:decimal(10,2);" json:"balance"`
	SettlementAmount float64 `gorm:"type:decimal(10,2);" json:"settlementAmount"`

	OpenID *string `gorm:"type:varchar(50);" json:"openID,omitempty"`
	Extra  string  `gorm:"type:text;not null;" json:"extra"`
	Mobile string  `gorm:"type:varchar(15)" json:"mobile,omitempty"`

	ValidateTime time.Time `gorm:"type:datetime" json:"validateTime"`

	CreateAt time.Time `gorm:"type:datetime" json:"createAt"`
	ModifyAt time.Time `gorm:"type:datetime" json:"modifyAt"`
	Status   int       `gorm:"type:int(2)" json:"status"`

	HomeAddress string `gorm:"type:varchar(500)" json:"homeAddress"`
	RegistDate  string `gorm:"type:varchar(50)" json:"registDate"`
	Gift        string `gorm:"type:varchar(500)" json:"gift"`

	extraMap map[string]interface{}
}

func (s *Student) TableName() string {
	return "student"
}

func (s *Student) AddExtra(key string, value interface{}) {
	if s.extraMap == nil {
		extra := make(map[string]interface{})
		_ = json.Unmarshal([]byte(s.Extra), &extra)
		s.extraMap = extra
	}

	s.extraMap[key] = value
}

func (s *Student) MarshalExtra() {
	marshal, _ := json.Marshal(s.extraMap)
	s.Extra = string(marshal)
}

func (s *Student) InClassRecord(record string) {
	if s.extraMap == nil {
		extra := make(map[string]interface{})
		_ = json.Unmarshal([]byte(s.Extra), &extra)
		s.extraMap = extra
	}

	if datesString, ok := s.extraMap["in_class_record"].(string); ok {
		var dates []string
		_ = json.Unmarshal([]byte(datesString), &dates)
		dates = append(dates, record)
		records, _ := json.Marshal(dates)
		s.extraMap["in_class_record"] = string(records)
	} else {
		dates := []string{record}
		records, _ := json.Marshal(dates)
		s.extraMap["in_class_record"] = string(records)
	}
}

type InClassRecord struct {
	Date []string `json:"dates"`
}

var (
	NoActive = 1
	Active   = 10
	Finish   = 20
	Terminal = 90
	Delete   = 99
)
