// 数据操作处理-用户
package controller

import (
	"assemblyline/project1023/common"
	"assemblyline/project1023/dto"
	"assemblyline/project1023/model"
	"assemblyline/project1023/response"
	"assemblyline/project1023/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 注册事务
func Register(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
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
	// 返回结果
	response.Success(ctx, nil, "注册成功")
}

// 登录事务
func Login(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	// name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码大于6位")
		return
	}

	//判断手机号是否存在
	var usertable model.UserTable
	db.Where("telephone = ?", telephone).First(&usertable)
	if usertable.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
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

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var UserTable model.UserTable
	db.Where("telephone = ?", telephone).First(&UserTable)

	return UserTable.ID != 0
}
