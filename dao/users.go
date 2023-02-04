package dao

import (
	"readSoftware/model"
)

func SearchUserByUserName(name string) (u model.UserInfo, err error) {
	row := DB.QueryRow("select * from user where username = ?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.PassWord, &u.Nickname, &u.Gender, &u.QQ, &u.Birthday, &u.Email, &u.Avatar, &u.Introduction, &u.Phone, &u.IsAdministrator)
	return
}

// 通过用户名查找

func InsertUser(u model.UserInfo) (err error) {
	_, err = DB.Exec("insert into user(id,username,password,nickname,gender,qq,birthday,email,avatar,introduction,phone,is_administrator) values (?,?,?,?,?,?,?,?,?,?,?,?)", u.Id, u.UserName, u.PassWord, u.Nickname, u.Gender, u.QQ, u.Birthday, u.Email, u.Avatar, u.Introduction, u.Phone, u.IsAdministrator)
	return err
}

// 注册，将用户信息录入数据库

func ChangePasswordByUsername(username string, newPassword string) (err error) {
	_, err = DB.Exec("update user set password=? where username=?", newPassword, username)
	return err
}

// 通过用户名改密码

func SearchUserByUserId(id int) (u model.UserInfo, err error) {
	row := DB.QueryRow("select * from user where id = ?", id)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.PassWord, &u.Nickname, &u.Gender, &u.QQ, &u.Birthday, &u.Email, &u.Avatar, &u.Introduction, &u.Phone, &u.IsAdministrator)
	return
}

// 根据id查找用户

func ChangeUserInfo(u model.UserInfo) (err error) {
	_, err = DB.Exec("update user set nickname=?,avatar=?,introduction=?,phone=?,qq=?,gender=?,email=?,birthday=? where username=?", u.Nickname, u.Avatar, u.Introduction, u.Phone, u.QQ, u.Gender, u.Email, u.Birthday, u.UserName)
	return err
}

//查询用户信息
