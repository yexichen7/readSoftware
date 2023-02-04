package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"readSoftware/model"
	"readSoftware/service"
	"readSoftware/tool"
	"readSoftware/util"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	u, err := service.SearchUserByUserName(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)

		return
	}
	if u.UserName != "" {
		util.NormErr(c, 60001, "用户名已注册")
		return
	}
	//用户名查重
	err = service.InsertUser(model.UserInfo{
		UserName: username,
		PassWord: password,
	})
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)

		return
	}
	//检查连接
	util.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	u, err := service.SearchUserByUserName(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName != username {
		util.NormErr(c, 60002, "用户未注册或用户名输入错误")
		return
	}
	if u.PassWord != password {
		util.NormErr(c, 60003, "密码错误")
	}
	//登录信息检查
	token, err := tool.GenerateToken([]byte("77"), 3600*time.Second, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60010, "登陆错误")
		return
	}
	refreshToken, err := tool.GenerateToken([]byte("777"), 86400*time.Second, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60010, "登陆错误")
		return
	}
	util.RespSuccess(c, token, refreshToken)
}

func RefreshToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	refreshToken := c.Query("refresh_token")
	_, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	}
	isTure, _, err := tool.TokenExpired([]byte("777"), refreshToken)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	} else if err == nil && isTure == false {
		util.NormErr(c, 60103, "refresh_token过期，请重新登陆")
		return
	}
	token, err = tool.GenerateToken([]byte("77"), 3600*time.Second, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600104, "token申请错误")
		return
	}
	refreshToken, err = tool.GenerateToken([]byte("777"), 86400*time.Second, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60104, "token申请错误")
		return
	}
	util.RespSuccess(c, token, refreshToken)
}

func ChangePassword(c *gin.Context) {

	token := c.GetHeader("Authorization")
	oldPassword := c.Query("old_password")
	newPassword := c.Query("new_password")
	//获取用户信息

	isExpired, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExpired {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600102, "token已过期")
		return
	}
	//token解密

	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.PassWord != oldPassword {
		util.NormErr(c, 60003, "密码错误")
		return
	}
	err = service.ChangePasswordByUsername(username, newPassword)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	//查找用户信息
	util.RespOK(c)
}

func GetUserInfo(c *gin.Context) {
	idString := c.Param("user_id")
	fmt.Println(idString)
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60004, "UID非法")
		return
	}
	u, err := service.SearchUserByUserId(id)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespUserInfoSuccess(c, u)
}

func ChangeUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	nickname := c.PostForm("nickname")
	avatar := c.PostForm("avatar")
	introduction := c.PostForm("introduction")
	telephoneString := c.PostForm("telephone")
	telephone, err := strconv.Atoi(telephoneString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60011, "数值非法")
		return
	}
	qqString := c.PostForm("qq")
	qq, err := strconv.Atoi(qqString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60011, "数值非法")
		return
	}
	gender := c.PostForm("gender")
	email := c.PostForm("email")
	birthday := c.PostForm("birthday")
	isExist, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExist {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600102, "token已过期")
		return
	}
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	u.Nickname = nickname
	u.Avatar = avatar
	u.Introduction = introduction
	u.Phone = telephone
	u.QQ = qq
	u.Gender = gender
	u.Email = email
	u.Birthday = birthday
	//各类用户信息获取，尝试ShouldBind失败
	err = service.ChangeUserInfo(u)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}
