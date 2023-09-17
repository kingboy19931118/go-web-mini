package controller

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/pkg/idutil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-mini/common"
	"go-web-mini/dto"
	"go-web-mini/model"
	"go-web-mini/repository"
	"go-web-mini/response"
	"go-web-mini/vo"
	"time"
)

var (
	idGen = idutil.NewGenerator(0, time.Now())
)

type IStudentController interface {
	GetStudents(c *gin.Context)             // 获取学生列表
	CreateStudent(c *gin.Context)           // 创建用户
	UpdateStudent(c *gin.Context)           // 更新用户
	BatchDeleteStudentByIds(c *gin.Context) // 批量删除用户
}

type StudentController struct {
	StudentRepository repository.IStudentRepository
}

// 构造函数
func NewStudentController() IStudentController {
	studentRepository := repository.NewStudentRepository()
	StudentController := StudentController{StudentRepository: studentRepository}
	return StudentController
}

// 获取用户列表
func (uc StudentController) GetStudents(c *gin.Context) {
	var req vo.StudentListRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取
	users, total, err := uc.StudentRepository.GetStudents(&req)
	if err != nil {
		response.Fail(c, nil, "获取学生列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"users": dto.ToStudentsDto(users), "total": total}, "获取学生列表成功")
}

// 创建用户
func (uc StudentController) CreateStudent(c *gin.Context) {
	var req vo.CreateStudentRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取当前用户
	ur := repository.NewUserRepository()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	student := &model.Student{
		ID:            int64(idGen.Next()),
		Name:          req.Name,
		Age:           req.Age,
		Gender:        req.Gender,
		ClassHour:     req.ClassHour,
		LeftClassHour: req.LeftClassHour,
		SignAmount:    req.SignAmount,
		Mobile:        req.Mobile,
		CreateAt:      req.InDate,
		ModifyAt:      req.InDate,
		Status:        model.NoActive,
	}

	extra := make(map[string]interface{})
	extra["operator"] = ctxUser.Username
	marshal, _ := json.Marshal(extra)
	student.Extra = string(marshal)

	err = uc.StudentRepository.CreateStudent(student)
	if err != nil {
		response.Fail(c, nil, "创建学生失败: "+err.Error())
		return
	}
	response.Success(c, nil, "创建学生成功")
}

// 更新用户
func (uc StudentController) UpdateStudent(c *gin.Context) {
	var req vo.UpdateStudentRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取当前用户
	ur := repository.NewUserRepository()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	student, err := uc.StudentRepository.MultiGetStudent([]int64{req.ID})
	if err != nil {
		response.Fail(c, nil, "获取需要更新的用户信息失败: "+err.Error())
		return
	}

	if len(student) == 0 {
		response.Fail(c, nil, "需要更新的用户不存在！ ")
		return
	}

	u := student[0]
	extra := make(map[string]interface{})
	_ = json.Unmarshal([]byte(u.Extra), &extra)

	if req.Name != u.Name {
		extra["before_name"] = u.Name
		u.Name = req.Name
	}
	if req.Age != u.Age {
		extra["before_age"] = u.Age
		u.Age = req.Age
	}
	if req.Gender != u.Gender {
		extra["before_gender"] = u.Gender
		u.Gender = req.Gender
	}
	if req.ClassHour != u.ClassHour {
		extra["before_ch"] = u.ClassHour
		u.ClassHour = req.ClassHour
	}
	if req.LeftClassHour != u.LeftClassHour {
		extra["before_lch"] = u.LeftClassHour
		u.LeftClassHour = req.LeftClassHour
	}
	if req.SignAmount != u.SignAmount {
		extra["before_sa"] = u.SignAmount
		u.SignAmount = req.SignAmount
	}
	if req.OpenID != "" {
		extra["before_openID"] = u.OpenID
		u.OpenID = &req.OpenID
	}
	if req.Mobile != u.Mobile {
		extra["before_mobile"] = u.Mobile
		u.Mobile = req.Mobile
	}
	if req.Status != u.Status {
		extra["before_status"] = u.Status
		u.Status = req.Status
	}
	extra["operator"] = ctxUser.Username
	marshal, _ := json.Marshal(extra)
	u.Extra = string(marshal)

	// 更新用户
	err = uc.StudentRepository.UpdateStudent(u)
	if err != nil {
		response.Fail(c, nil, "更新用户失败: "+err.Error())
		return
	}
	response.Success(c, nil, "更新用户成功")

}

// 批量删除用户
func (uc StudentController) BatchDeleteStudentByIds(c *gin.Context) {
	var req vo.DeleteStudentRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取当前用户
	ur := repository.NewUserRepository()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户信息失败")
		return
	}
	err = uc.StudentRepository.BatchDeleteStudentByIds(req.StudentIds, ctxUser.Username)
	if err != nil {
		fmt.Println(err.Error())
		response.Fail(c, nil, "删除用户失败: "+err.Error())
		return
	}

	response.Success(c, nil, "删除用户成功")

}
