package service

import (
	"readSoftware/dao"
	"readSoftware/model"
)

func GetBookLists() (u []model.BookInfo, err error) {
	u, err = dao.GetBookLists()
	return
}

func SearchBook(bookName string) (book model.BookInfo, err error) {
	book, err = dao.SearchBookInfo(bookName)
	return
}

func SearchUserStar(userID int, bookID int) (isStar bool, err error) {
	isStar, err = dao.SearchUserStar(bookID, userID)
	return
}

func StarBook(userID int, bookID int) (err error) {
	err = dao.StarBook(userID, bookID)
	return err
}

func GetBookByLabel(label string) (u []model.BookInfo, err error) {
	u, err = dao.GetBookByLabel(label)
	return
}
