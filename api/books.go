package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"readSoftware/service"
	"readSoftware/tool"
	"readSoftware/util"
	"strconv"
)

func GetBookLists(c *gin.Context) {
	u, err := service.GetBookLists()
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.BookListRespSuccess(c, u)
}

func SearchBookInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	bookName := c.Query("book_name")
	uBook, err := service.SearchBook(bookName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			util.NormErr(c, 70001, "该书尚未收录")
			return
		}
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	for token != "" {
		isExist, username, err := tool.TokenExpired([]byte("77"), token)
		if err != nil {
			log.Printf("search user error:%v", err)
			util.NormErr(c, 60100, "token错误")
			break
		}
		if !isExist {
			util.NormErr(c, 60102, "token已过期")
			break
		}
		uUser, err := service.SearchUserByUserName(username)
		if err != nil {
			log.Printf("search user error:%v", err)
			util.RsepInternalErr(c)
			break
		}
		isStar, err := service.SearchUserStar(uUser.Id, uBook.BookId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				break
			}
			log.Printf("search user error:%v", err)
			util.RsepInternalErr(c)
			break
		}
		if isStar {
			uBook.IsStar = true
			break
		}
		break
	}
	util.BookRespSuccess(c, uBook)
}

func StarBook(c *gin.Context) {
	token := c.GetHeader("Authorization")
	bookIdString := c.Query("book_id")
	isExist, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 60102, "token过期")
		return
	}
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.NormErr(c, 70002, "book_id非法")
		return
	}
	err = service.StarBook(uUser.Id, bookId)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func GetBookByLabel(c *gin.Context) {
	label := c.Query("label")
	u, err := service.GetBookByLabel(label)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.BookListRespSuccess(c, u)
}

func MarkBook(c *gin.Context) {
	token := c.GetHeader("Authorization")
	bookIdString := c.Query("book_id")
	isExist, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 60102, "token过期")
		return
	}
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.NormErr(c, 70002, "book_id非法")
		return
	}
	pageString := c.Query("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	content := c.Query("content")
	err = service.MarkBook(uUser.Id, bookId, page, content)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}
func GetBookMark(c *gin.Context) {
	token := c.GetHeader("Authorization")
	bookName := c.Query("book_name")
	uBook, err := service.SearchBook(bookName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			util.NormErr(c, 70001, "该书尚未收录")
			return
		}
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}

	isExist, username, err := tool.TokenExpired([]byte("77"), token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 60100, "token错误")
		return
	}
	if !isExist {
		util.NormErr(c, 60102, "token已过期")
		return
	}
	uUser, err := service.SearchUserByUserName(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	u, err := service.GetBookMark(uUser.Id, uBook.BookId)
	if err != nil {
		log.Printf("search book error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.BookMarkRespSuccess(c, u)
}


