package service

import (
	"readSoftware/dao"
	"readSoftware/model"
)

func CreateDiscuss(u model.CommentInfo) (discussID int, err error) {
	discussID, err = dao.CreateDiscuss(u)
	return
}

func GetDiscussList(postID int) (u []model.CommentInfo, err error) {
	u, err = dao.GetDiscussList(postID)
	return
}

func DeleteDiscuss(discussID int, userID int, isAdministrator bool) (err error) {
	err = dao.DeleteDiscuss(discussID, userID, isAdministrator)
	return
}

func ReplayDiscuss(u model.CommentInfo) (discussID int, err error) {
	discussID, err = dao.ReplayDiscuss(u)
	return
}

func SearchPostAndUserByDiscussID(discussID int) (postID int, userID int, err error) {
	postID, userID, err = dao.SearchPostAndUserByDiscussID(discussID)
	return
}

func CheckReplay(userID int) (u []model.CommentInfo, err error) {
	u, err = dao.CheckReplay(userID)
	return
}

func GetCommentList(bookID int) (u []model.CommentInfo, err error) {
	u, err = dao.GetCommentLists(bookID)
	return
}

func CreatComment(u model.CommentInfo) (commentID int, err error) {
	commentID, err = dao.CreatComment(u)
	return
}

func RefreshComment(userID int, commentID int, content string) (err error) {
	err = dao.RefreshComment(userID, commentID, content)
	return err
}

func DeleteComment(userID int, commentID int, isAdministrator bool) (err error) {
	err = dao.DeleteComment(userID, commentID, isAdministrator)
	return
}
