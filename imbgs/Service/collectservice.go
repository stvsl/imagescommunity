package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	Database "stvsljl.com/stvsl/imbgs/DBLogic"
)

func CollectIt(c *gin.Context) {
	iccollect := Database.IcLikes{}
	c.BindJSON(&iccollect)
	db := Database.GetConn()
	db.Exec("INSERT INTO ic_likes (img_id, uuid) VALUES (" + fmt.Sprintf("%v", iccollect.ImgId) + ", '" + fmt.Sprintf("%v", iccollect.Uuid) + "');")
	icarts := Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE img_id = " + fmt.Sprintf("%v", iccollect.ImgId) + ";").Scan(&icarts)
	icarts.Collects++
	db.Exec("UPDATE ic_art SET collects = " + fmt.Sprintf("%v", icarts.Collects) + " WHERE img_id = " + fmt.Sprintf("%v", iccollect.ImgId) + ";")
	c.JSON(200, gin.H{
		"status": "success",
	})
}
