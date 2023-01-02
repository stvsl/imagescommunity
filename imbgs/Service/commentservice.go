package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	Database "stvsljl.com/stvsl/imbgs/DBLogic"
)

func CommentUpLoad(c *gin.Context) {
	iccomment := Database.IcComment{}
	c.BindJSON(&iccomment)
	iccomment.Time = time.Now()
	db := Database.GetConn()
	db.Exec("INSERT INTO ic_comment" +
		"(imgid, comments, `uuid`, `time`)" +
		"VALUES(" + fmt.Sprintf("%d", iccomment.Imgid) + ", '" + iccomment.Comments + "', '" + iccomment.Uuid + "', '" + iccomment.Time.Format("2006-01-02 15:04:05") + "');")
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func CommentGet(c *gin.Context) {
	var art struct {
		Id int
	}
	c.BindJSON(&art)
	db := Database.GetConn()
	iccomments := []Database.IcComment{}
	db.Raw("SELECT * FROM ic_comment WHERE imgid = " + fmt.Sprintf("%d", art.Id) + ";").Scan(&iccomments)
	// 在ic_user表中查询uuid对应的name和avatar
	for i := 0; i < len(iccomments); i++ {
		icuser := Database.IcUser{}
		db.Raw("SELECT * FROM ic_user WHERE uuid = '" + iccomments[i].Uuid + "';").Scan(&icuser)
		iccomments[i].Name = icuser.Name
		iccomments[i].Avatar = icuser.Avatar
	}
	str, _ := json.Marshal(iccomments)
	c.JSON(200, gin.H{
		"status":   "success",
		"comments": string(str),
	})
}

func GetCollect(c *gin.Context) {
	var user struct {
		Uuid string
	}
	c.BindJSON(&user)
	db := Database.GetConn()
	iccollects := []Database.IcLikes{}
	db.Raw("SELECT * FROM ic_likes WHERE uuid = '" + user.Uuid + "';").Scan(&iccollects)
	// 在ic_art表中查询imgid对应的name和uri,owneruuid,并将其添加到iccollects中
	for i := 0; i < len(iccollects); i++ {
		icart := Database.IcArt{}
		db.Raw("SELECT * FROM ic_art WHERE img_id = " + fmt.Sprintf("%d", iccollects[i].ImgId) + ";").Scan(&icart)
		iccollects[i].Name = icart.Name
		iccollects[i].Uri = icart.Uri
		iccollects[i].OwnerUuid = icart.OwnerUuid
	}
	str, _ := json.Marshal(iccollects)
	c.JSON(200, gin.H{
		"status":   "success",
		"collects": string(str),
	})
}
