package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-web-mini/common"
	"go-web-mini/model"
	"go-web-mini/vo"
	"strings"
)

type IStudentRepository interface {
	MultiGetStudent(ids []int64) ([]*model.Student, error)
	CreateStudent(user *model.Student) error                                 // 创建用户
	GetStudents(req *vo.StudentListRequest) ([]*model.Student, int64, error) // 获取用户列表
	UpdateStudent(user *model.Student) error                                 // 更新用户
	BatchDeleteStudentByIds(ids []int64, operator string) error              // 批量删除
}

type StudentRepository struct {
}

// StudentRepository构造函数
func NewStudentRepository() IStudentRepository {
	return StudentRepository{}
}

// 获取用户
func (ur StudentRepository) MultiGetStudent(ids []int64) ([]*model.Student, error) {
	fmt.Println("MultiGetStudent---")
	user := make([]*model.Student, 0, len(ids))
	err := common.DB.Where("id in ?", ids).Where("status != ?", model.Delete).Limit(10).Find(&user).Error
	return user, err
}

// 获取用户列表
func (ur StudentRepository) GetStudents(req *vo.StudentListRequest) ([]*model.Student, int64, error) {
	var list []*model.Student
	db := common.DB.Model(&model.Student{}).Order("create_at DESC")

	name := strings.TrimSpace(req.Name)
	db = db.Where("status != ?", model.Delete)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	if req.Gender != 0 {
		db = db.Where("gender = ?", req.Gender)
	}

	// 当pageNum > 0 且 pageSize > 0 才分页
	//记录总条数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := int(req.PageNum)
	pageSize := int(req.PageSize)
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}

// 创建用户
func (ur StudentRepository) CreateStudent(user *model.Student) error {
	err := common.DB.Create(user).Error
	return err
}

// 更新用户
func (ur StudentRepository) UpdateStudent(user *model.Student) error {
	return common.DB.Model(user).Updates(user).Error
}

// 批量删除
func (ur StudentRepository) BatchDeleteStudentByIds(ids []int64, operator string) error {
	students, err := ur.MultiGetStudent(ids)
	if err != nil {
		return errors.New("查询用户失败")
	}

	for _, student := range students {
		extraVal := make(map[string]interface{})
		_ = json.Unmarshal([]byte(student.Extra), &extraVal)
		extraVal["before_status"] = student.Status
		extraVal["operator"] = operator
		marshal, _ := json.Marshal(extraVal)
		student.Extra = string(marshal)
		student.Status = model.Delete
		if err = ur.UpdateStudent(student); err != nil {
			return err
		}
	}

	return nil
}
