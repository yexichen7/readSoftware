package dao

import (
	"readSoftware/model"
)

func GetBookLists() (u []model.BookInfo, err error) {
	row, err := DB.Query("select * from book")
	if err != nil {
		return
	}
	for row.Next() {
		var temp model.BookInfo
		err = row.Scan(&temp.BookId, &temp.Name, &temp.IsStar, &temp.Author, &temp.CommentNum, &temp.Score, &temp.Cover, &temp.PublishTime, &temp.Link, &temp.Label)
		if err != nil {
			return
		}
		u = append(u, temp)
	}
	return
}

func SearchBookInfo(bookName string) (book model.BookInfo, err error) {
	row := DB.QueryRow("select * from book where name = ?", bookName)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&book.BookId, &book.Name, &book.IsStar, &book.Author, &book.CommentNum, &book.Score, &book.Cover, &book.PublishTime, &book.Link, &book.Label)
	return
}

func GetBookMark(userID int, bookID int) (mark model.Mark, err error) {
	row := DB.QueryRow("SELECT * FROM mark WHERE  Id= ? AND bookId = ?", userID, bookID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&mark.BookID, &mark.Name, &mark.Page, &mark.Content, &mark.ID)
	return
}

func SearchUserStar(bookID int, userID int) (isStar bool, err error) {
	row := DB.QueryRow("SELECT * FROM star WHERE  Id= ? AND bookId = ?", userID, bookID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	var temp model.UserStar
	err = row.Scan(&temp.UserID, &temp.BookID, &temp.IsStar)
	return temp.IsStar, err
}

func StarBook(userID int, bookID int) (err error) {
	_, err = DB.Exec("insert into star(Id,bookId,is_star) values (?,?,?)", userID, bookID, true)
	return err
}

func GetBookByLabel(label string) (u []model.BookInfo, err error) {
	row, err := DB.Query("select * from book where lable like ?", "%"+label+"%")
	if err != nil {
		return
	}
	for row.Next() {
		var temp model.BookInfo
		err = row.Scan(&temp.BookId, &temp.Name, &temp.IsStar, &temp.Author, &temp.CommentNum, &temp.Score, &temp.Cover, &temp.PublishTime, &temp.Link, &temp.Label)
		if err != nil {
			return
		}
		u = append(u, temp)
	}
	return
}

func MarkBook(userID int, bookID int, page int, content string) (err error) {
	_, err = DB.Exec("insert into mark(Id,bookId,pages,content) values (?,?,?,?)", userID, bookID, page, content)
	return err
}
