package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readSoftware/model"
)

// books相关--------------------------------
type bookListResp struct {
	Status int          `json:"status"`
	Info   string       `json:"info"`
	Data   bookInfoList `json:"data"`
}
type bookInfoList struct {
	Books []model.BookInfo `json:"books"`
}

type bookResp struct {
	Status int      `json:"status"`
	Info   string   `json:"info"`
	Data   bookInfo `json:"data"`
}

type bookInfo struct {
	Book model.BookInfo `json:"book"`
}

func BookListRespSuccess(c *gin.Context, u []model.BookInfo) {
	response := bookListResp{
		Status: 10000,
		Info:   "success",
		Data:   bookInfoList{u},
	}
	c.JSON(http.StatusOK, response)
}

func BookRespSuccess(c *gin.Context, u model.BookInfo) {
	response := bookResp{
		Status: 10000,
		Info:   "success",
		Data:   bookInfo{u},
	}
	c.JSON(http.StatusOK, response)
}

//comment相关-------------------------------

type commentRespSuccess struct {
	Status   int                 `json:"status"`
	Info     string              `json:"info"`
	Comments []model.CommentInfo `json:"comments"`
}

type creatCommentRespSuccess struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   int    `json:"data"`
}

func CommentRespSuccess(c *gin.Context, u []model.CommentInfo) {
	response := commentRespSuccess{
		Status:   10000,
		Info:     "success",
		Comments: u,
	}
	c.JSON(http.StatusOK, response)
}

func CreatCommentRespSuccess(c *gin.Context, postID int) {
	response := creatCommentRespSuccess{
		Status: 10000,
		Info:   "success",
		Data:   postID,
	}
	c.JSON(http.StatusOK, response)
}

// comment/discuss相关--
type discussRespSuccess struct {
	Status   int                 `json:"status"`
	Info     string              `json:"info"`
	Comments []model.CommentInfo `json:"comments"`
}

func GetCommentInfoSuccess(c *gin.Context, u []model.CommentInfo) {
	response := discussRespSuccess{
		Status:   10000,
		Info:     "success",
		Comments: u,
	}
	c.JSON(http.StatusOK, response)
}

func CreatDiscussRespSuccess(c *gin.Context, data int) {
	response := creatCommentRespSuccess{
		Status: 10000,
		Info:   "success",
		Data:   data,
	}
	c.JSON(http.StatusOK, response)
}

//users相关---------------------------------

type TokenResponse struct {
	Status int       `json:"status"`
	Info   string    `json:"info"`
	Data   TokenData `json:"data"`
}

type TokenData struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
}

type userInfo struct {
	ID           int    `json:"id"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Phone        int    `json:"phone"`
	QQ           int    `json:"qq"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
}

type Data struct {
	User userInfo `json:"user"`
}

type userInfoResponse struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Data   `json:"data"`
}

func RespSuccess(c *gin.Context, token string, refreshToken string) {
	response := TokenResponse{
		Status: 10000,
		Info:   "success",
		Data: TokenData{
			RefreshToken: refreshToken,
			Token:        token,
		},
	}

	c.JSON(http.StatusOK, response)
}

//token获取成功

func RespUserInfoSuccess(c *gin.Context, u model.UserInfo) {
	response := userInfoResponse{
		Status: 10000,
		Info:   "success",
		Data: Data{User: userInfo{
			ID:           u.Id,
			Avatar:       u.Avatar,
			Nickname:     u.Nickname,
			Introduction: u.Introduction,
			Phone:        u.Phone,
			QQ:           u.QQ,
			Gender:       u.Gender,
			Email:        u.Email,
			Birthday:     u.Birthday,
		}},
	}
	c.JSON(http.StatusOK, response)
}

//用户信息获取成功

//operate相关----------------------------------

type respCollectSuccess struct {
	Status int             `json:"status"`
	Info   string          `json:"info"`
	Data   collectInfoList `json:"data"`
}

type collectInfoList struct {
	Collections []collectBookInfoList `json:"collections"`
}

type collectBookInfoList struct {
	BookId      int    `json:"book_id"`
	Name        string `json:"name"`
	PublishTime string `json:"publish_time"`
	Link        string `json:"link"`
}

func RespCollectSuccess(c *gin.Context, u []model.BookInfo) {
	var temp []collectBookInfoList
	for _, book := range u {
		temp = append(temp, collectBookInfoList{
			BookId:      book.BookId,
			Name:        book.Name,
			PublishTime: book.PublishTime,
			Link:        book.Link,
		})
	}
	response := respCollectSuccess{
		Status: 10000,
		Info:   "success",
		Data:   collectInfoList{temp},
	}
	c.JSON(http.StatusOK, response)
}

//resp相关----------------------

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var ok = respTemplate{
	Status: 10000,
	Info:   "success",
}

//访问成功

var ParamError = respTemplate{
	Status: 30000,
	Info:   "params error",
}

var InternalError = respTemplate{
	Status: 50000,
	Info:   "internal error",
}

//访问错误

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, ok)
}
func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

func RsepInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
}

//连接错误

func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}

//其他错误
