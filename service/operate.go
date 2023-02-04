package service

import (
	"readSoftware/dao"
	"readSoftware/model"
)

//其他操作

func PraiseComment(commentID int) (err error) {
	err = dao.PraiseComment(commentID)
	return
}

func PraiseDiscuss(discussID int) (err error) {
	err = dao.PraiseDiscuss(discussID)
	return
}

func GetCollectList(userID int) (u []model.BookInfo, err error) {
	u, err = dao.GetCollectList(userID)
	return
}

func Focus(followerID int, followeeID int) (err error) {
	err = dao.Focus(followerID, followeeID)
	return
}
