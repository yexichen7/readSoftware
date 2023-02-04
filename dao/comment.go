package dao

import (
	"fmt"
	"readSoftware/model"
)

func CreateDiscuss(u model.CommentInfo) (discussID int, err error) {
	res, err := DB.Exec("insert into discuss(discuss_id,post_id,replay_id,comment,user_id,praise_count,replay_id) values (?,?,?,?,?,?,?)", u.DiscussID, u.PostID, u.ReplayID, u.Comment, u.UserID, u.PraiseCount, u.ReplayUid)
	if err != nil {
		return
	}
	discussID64, err := res.LastInsertId()
	discussID = int(discussID64)
	return
}

func GetDiscussList(postID int) (u []model.CommentInfo, err error) {
	row, err := DB.Query("select * from discuss where post_id=?", postID)
	if err != nil {
		return
	}
	for row.Next() {
		var temp model.CommentInfo
		err = row.Scan(&temp.DiscussID, &temp.PostID, &temp.ReplayID, &temp.Comment, &temp.UserID, &temp.PraiseCount, &temp.ReplayUid)
		if err != nil {
			return
		}
		u = append(u, temp)
	}
	return
}

func DeleteDiscuss(discussID int, userID int, isAdministrator bool) (err error) {
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM discuss WHERE discuss_id=? AND user_id=?", discussID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count != 1 && !isAdministrator {
		return fmt.Errorf("discuss_id and user_id not match")
	}
	_, err = DB.Exec("delete from discuss where discuss_id=?", discussID)
	return
}

func ReplayDiscuss(u model.CommentInfo) (discussID int, err error) {
	res, err := DB.Exec("insert into discuss(discuss_id,post_id,replay_id,comment,user_id,praise_count,replay_uid) values (?,?,?,?,?,?,?)", u.DiscussID, u.PostID, u.ReplayID, u.Comment, u.UserID, u.PraiseCount, u.ReplayUid)
	if err != nil {
		return
	}
	discussID64, err := res.LastInsertId()
	discussID = int(discussID64)
	return
}

func SearchPostAndUserByDiscussID(discussID int) (postID int, userID int, err error) { //根据回复查找postID和uid
	row := DB.QueryRow("select * from discuss where discuss_id = ?", discussID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	var temp model.CommentInfo
	err = row.Scan(&temp.DiscussID, &temp.PostID, &temp.ReplayID, &temp.Comment, &temp.UserID, &temp.PraiseCount, &temp.ReplayUid)
	return temp.PostID, temp.UserID, err
}

func CheckReplay(userID int) (u []model.CommentInfo, err error) { //查看回复
	row, err := DB.Query("select * from discuss where replay_uid=?", userID)
	if err != nil {
		return
	}
	for row.Next() {
		var temp model.CommentInfo
		err = row.Scan(&temp.DiscussID, &temp.PostID, &temp.ReplayID, &temp.Comment, &temp.UserID, &temp.PraiseCount, &temp.ReplayUid)
		if err != nil {
			return
		}
		u = append(u, temp)
	}
	return
}

func GetCommentLists(bookID int) (u []model.CommentInfo, err error) {
	row, err := DB.Query("select * from comment where book_id=?", bookID)
	if err != nil {
		return
	}
	for row.Next() {
		var temp model.CommentInfo
		err = row.Scan(&temp.PostID, &temp.BookID, &temp.PublishTime, &temp.Content, &temp.UserID, &temp.Avatar, &temp.Nickname, &temp.PraiseCount, &temp.IsPraised, &temp.IsFocus)
		if err != nil {
			return
		}
		u = append(u, temp)
	}
	return
}

func CreatComment(u model.CommentInfo) (commentID int, err error) {
	res, err := DB.Exec("insert into comment(post_id,book_id,publish_time,content,user_id,avatar,nickname,praise_count,is_praised,is_focus) values (?,?,?,?,?,?,?,?,?,?)", u.PostID, u.BookID, u.PublishTime, u.Content, u.UserID, u.Avatar, u.Nickname, u.PraiseCount, u.IsPraised, u.IsFocus)
	if err != nil {
		return
	}
	commentID64, err := res.LastInsertId()
	commentID = int(commentID64)
	return
}

func RefreshComment(userID int, commentID int, content string) (err error) {
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM comment WHERE post_id=? AND user_id=?", commentID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("post_id and user_id not match")
	}
	_, err = DB.Exec("update comment set content=? where post_id=? and user_id=?", content, commentID, userID)
	return
}

func DeleteComment(userID int, commentID int, isAdministrator bool) (err error) {
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM comment WHERE post_id=? AND user_id=?", commentID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count != 1 && !isAdministrator {
		return fmt.Errorf("post_id and user_id not match")
	}
	_, err = DB.Exec("delete from comment where post_id=?", commentID)
	return
}
