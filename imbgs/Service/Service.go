package service

import (
	"github.com/gin-gonic/gin"
	Database "stvsljl.com/stvsl/imbgs/DBLogic"
)

func Run() {
	Database.Connect()
	r := gin.Default()
	// 允许跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// 路由
	r.POST("/api/user/register", Register)
	r.POST("/api/user/login", Login)
	r.POST("/api/user/getname", GetUserName)
	r.GET("/api/user/getuser", GetUser)
	r.POST("/api/user/updateuser", UpdateUser)
	r.POST("/api/user/updatepasswd", UpdatePasswd)
	r.POST("/api/art/upload", ArtUpLoad)
	r.POST("/api/art/manage", ArtManage)
	r.POST("/api/art/update", ArtUpdate)
	r.POST("/api/art/delete", ArtDelete)
	r.POST("/api/art/findUserArt", FindUserArt)
	r.POST("/api/art/searchArt", SearchArt)
	r.GET("/api/art/getArt", GetArt)
	r.GET("/api/art/like", Like)
	r.POST("/api/art/recommend", Recommend)
	r.POST("/api/comment/upload", CommentUpLoad)
	r.POST("/api/comment/get", CommentGet)
	r.POST("/api/collect/collectit", CollectIt)
	r.POST("/api/collect/get", GetCollect)
	r.Run(":8521")
}
