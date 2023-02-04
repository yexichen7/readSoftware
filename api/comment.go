package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"readSoftware/model"
	"readSoftware/service"
	"readSoftware/tool"
	"readSoftware/util"

	"strconv"
)

//

func GetCommentLists(c *gin.Context) {
	bookIDString := c.Param("book_id")
	bookID, err := strconv.Atoi(bookIDString)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.NormErr(c, 70002, "book_id非法")
		return
	}
	u, err := service.GetCommentList(bookID)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CommentRespSuccess(c, u)
}

func CreatComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	bookIDString := c.Param("book_id")
	content := c.PostForm("content")
	bookID, err := strconv.Atoi(bookIDString)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.NormErr(c, 70002, "book_id非法")
		return
	}
	isExist, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 600102, "token已过期")
		return
	}
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	var uComment model.CommentInfo
	uComment.BookID = bookID
	uComment.PublishTime = tool.FormatTime()
	uComment.Content = content
	uComment.UserID = uUser.Id
	uComment.Avatar = uUser.Avatar
	uComment.Nickname = uUser.Nickname
	uComment.PostID, err = service.CreatComment(uComment)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CreatCommentRespSuccess(c, uComment.PostID)
}

func DeleteComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	commentIDString := c.Param("comment_id")
	isExist, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 600102, "token已过期")
		return
	}
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	commentID, err := strconv.Atoi(commentIDString)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.NormErr(c, 80002, "comment_id非法")
		return
	}
	err = service.DeleteComment(u.Id, commentID, u.IsAdministrator)
	if err != nil {
		if err.Error() == "post_id and user_id not match" {
			util.NormErr(c, 80004, "用户无权限删除此书评")
			return
		}
		log.Printf("search comment error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func RefreshComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	commentIDString := c.Param("comment_id")
	content := c.PostForm("content")
	commentID, err := strconv.Atoi(commentIDString)
	if err != nil {
		log.Printf("search comment error:%v", err)
		util.NormErr(c, 80002, "comment_id非法")
		return
	}
	isExist, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 600100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 600102, "token已过期")
		return
	}
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	err = service.RefreshComment(u.Id, commentID, content)
	if err != nil {
		if err.Error() == "post_id and user_id not match" {
			util.NormErr(c, 80003, "权限不足")
			return
		}
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func CreateDiscuss(c *gin.Context) {
	token := c.GetHeader("Authorization")
	postIDString := c.Param("post_id")
	comment := c.PostForm("comment")
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
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70012, "post_id非法")
		return
	}
	var uDiscuss model.CommentInfo
	uDiscuss.PostID = postID
	uDiscuss.Comment = comment
	uDiscuss.UserID = uUser.Id
	discussID, err := service.CreateDiscuss(uDiscuss)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CreatDiscussRespSuccess(c, discussID)
}

func GetDiscussList(c *gin.Context) {
	postIDString := c.Param("post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70012, "post_id非法")
		return
	}
	uCommentInfo, err := service.GetDiscussList(postID)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.GetCommentInfoSuccess(c, uCommentInfo)
}

//获取一个帖子下全部的回复信息

func DeleteDiscuss(c *gin.Context) {
	token := c.GetHeader("Authorization")
	discussIDString := c.Param("discuss_id")
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
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	discussID, err := strconv.Atoi(discussIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70013, "discuss_id非法")
		return
	}
	err = service.DeleteDiscuss(discussID, uUser.Id, uUser.IsAdministrator)
	if err != nil {
		if err.Error() == "discuss_id and user_id not match" {
			util.NormErr(c, 70014, "权限不足")
			return
		}
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func ReplayDiscuss(c *gin.Context) {
	token := c.GetHeader("Authorization")
	discussIDString := c.Param("discuss_id")
	comment := c.PostForm("comment")
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
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	discussID, err := strconv.Atoi(discussIDString)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.NormErr(c, 70013, "discuss_id非法")
		return
	}
	postID, userID, err := service.SearchPostAndUserByDiscussID(discussID)
	//得到帖子ID，被回复用户UID
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	var uDiscuss model.CommentInfo
	uDiscuss.PostID = postID
	uDiscuss.Comment = comment
	uDiscuss.UserID = uUser.Id
	uDiscuss.ReplayID = discussID
	uDiscuss.ReplayUid = userID
	discussID, err = service.ReplayDiscuss(uDiscuss)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.CreatDiscussRespSuccess(c, discussID)
}

func CheckReplay(c *gin.Context) {
	token := c.GetHeader("Authorization")
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
	uUser, err := service.SearchUserByUserName(username)
	u, err := service.CheckReplay(uUser.Id)
	if err != nil {
		log.Printf("search discuss error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.GetCommentInfoSuccess(c, u)
}
