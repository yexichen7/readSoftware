package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	use := r.Group("/user")
	{
		use.POST("/register", Register)
		use.GET("/token", Login)
		use.GET("/token/refresh", RefreshToken)
		use.PUT("/password", ChangePassword)
		use.GET("/info/:user_id", GetUserInfo)
		use.PUT("/info", ChangeUserInfo)
	}
	book := r.Group("/book")
	{
		book.GET("/list", GetBookLists)
		book.GET("/search", SearchBookInfo)
		book.PUT("/star", StarBook)
		book.GET("/label", GetBookByLabel)
	}
	comment := r.Group("/comment")
	{
		comment.GET("/:book_id", GetCommentLists)
		comment.POST("/:book_id", CreatComment)
		comment.DELETE("/:comment_id", DeleteComment)
		comment.PUT("/:comment_id", RefreshComment)
		comment.POST("/:post_id", CreateDiscuss)
		comment.GET("/:post_id", GetDiscussList)
		comment.DELETE("/:discuss_id", DeleteDiscuss)
		comment.POST("/replay/:discuss_id", ReplayDiscuss)
		comment.GET("/check", CheckReplay)
	}
	operate := r.Group("/operate")
	{
		operate.PUT("/praise/:target_id/model", Praise)
		operate.GET("/collect/list", GetCollectList)
		operate.PUT("/focus/:user_id", FocusUser)
	}
	r.Run()
}
