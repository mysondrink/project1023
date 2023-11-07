// 数据操作处理-用户
package controller

import (
	"assemblyline/project1023/common"
	"assemblyline/project1023/dto"
	"assemblyline/project1023/model"
	"assemblyline/project1023/response"
	"assemblyline/project1023/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 注册事务
func Register(ctx *gin.Context) {
	db := common.GetDB()
	// gin的绑定获取参数
	var requestUser = model.UserTable{}
	ctx.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 获取参数
	// name := ctx.PostForm("name")
	// telephone := ctx.PostForm("telephone")
	// password := ctx.PostForm("password")

	fmt.Println(name, telephone, password)

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	// 如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户存在")
		return
	}

	// 创建用户
	// 加密用户密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newUser := model.UserTable{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	db.Create(&newUser)

	// 发放token到前端
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

// 登录事务
func Login(ctx *gin.Context) {
	db := common.GetDB()
	var requestUser = model.UserTable{}
	ctx.Bind(&requestUser)

	// 获取参数
	// name := ctx.PostForm("name")
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	// if len(telephone) != 11 {
	// 	response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
	// 	return
	// }
	// if len(password) < 6 {
	// 	response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码大于6位")
	// 	return
	// }

	//判断手机号是否存在
	var usertable model.UserTable
	db.Unscoped().Where("telephone = ?", telephone).First(&usertable)
	if usertable.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(usertable.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(usertable)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

// 登录获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.UserTable))}, "ok")
}

// 验证电话是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var UserTable model.UserTable
	db.Unscoped().Where("telephone = ?", telephone).First(&UserTable)

	return UserTable.ID != 0
}

// 获取轨道信息业务
func GetInfo(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var request = model.Controller{}
	ctx.Bind(&request)
	controller_type := request.Controller_type

	// 获取id
	var controllerTable []model.Controller
	if result := db.Where("controller_type = ?", controller_type).Find(&controllerTable); result.Error != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}
	// 返回信息
	// fmt.Println(result.RowsAffected)
	response.Success(ctx, gin.H{"data": dto.ToControllerDto(controllerTable)}, "查询成功")
}

func GetCarInfo(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var request = model.Car{}
	ctx.Bind(&request)
	// car_id := request.Car_id

	// 获取id
	var carTable []model.Car
	if result := db.Find(&carTable); result.Error != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}
	// if result := db.Where("car_id = ?", car_id).Find(&controllerTable); result.Error != nil {
	// 	response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
	// 	return
	// }

	// 返回信息
	// fmt.Println(result.RowsAffected)
	response.Success(ctx, gin.H{"data": dto.ToCarDto(carTable)}, "查询成功")
}
