package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	Database "stvsljl.com/stvsl/imbgs/DBLogic"
	LsKy "stvsljl.com/stvsl/imbgs/lsky"
)

func ArtUpLoad(c *gin.Context) {
	var art struct {
		Uuid     string
		Name     string
		Picture  string
		Describe string
		Tag      string
	}
	c.BindJSON(&art)
	url := LsKy.Upload(art.Picture)
	db := Database.GetConn()
	db.Exec("INSERT INTO ic_art" +
		"(img_id,`describe`, name, uri, upload_time, release_time, status, owner_uuid, views, likes, collects, label)" +
		"VALUES(NULL,'" + art.Describe + "', '" + art.Name + "','" + url + "', '" + time.Now().Format("2006-01-02 15:04:05") +
		"', '" + time.Now().Format("2006-01-02 15:04:05") + "', 1, '" + art.Uuid + "', 0, 0, 0, '" + art.Tag + "');")
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func ArtManage(c *gin.Context) {
	var user struct {
		Uuid string
	}
	c.BindJSON(&user)
	db := Database.GetConn()
	icarts := []Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE owner_uuid = '" + user.Uuid + "';").Scan(&icarts)
	// icarts转换为json
	str, _ := json.Marshal(icarts)
	c.JSON(200, gin.H{
		"status": "success",
		"arts":   string(str),
	})
}

func ArtUpdate(c *gin.Context) {
	var art struct {
		Id       int
		Name     string
		Describe string
	}
	c.BindJSON(&art)
	db := Database.GetConn()
	db.Exec("UPDATE ic_art SET name = '" + art.Name + "', `describe` = '" + art.Describe + "' WHERE img_id = " + strconv.Itoa(art.Id) + ";")
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func ArtDelete(c *gin.Context) {
	var art struct {
		Id int
	}
	c.BindJSON(&art)
	db := Database.GetConn()
	db.Exec("DELETE FROM ic_art WHERE img_id = " + strconv.Itoa(art.Id) + ";")
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func FindUserArt(c *gin.Context) {
	var art struct {
		Uuid string
	}
	c.BindJSON(&art)
	db := Database.GetConn()
	icarts := []Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE owner_uuid = '" + art.Uuid + "';").Scan(&icarts)
	str, _ := json.Marshal(icarts)
	c.JSON(200, gin.H{
		"status": "success",
		"arts":   string(str),
	})
}

func SearchArt(c *gin.Context) {
	var art struct {
		Tag string
	}
	c.BindJSON(&art)
	db := Database.GetConn()
	icarts := []Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE label LIKE '%" + art.Tag + "%' OR name LIKE '%" + art.Tag + "%';").Scan(&icarts)
	str, _ := json.Marshal(icarts)
	c.JSON(200, gin.H{
		"status": "success",
		"arts":   string(str),
	})
}

func GetArt(c *gin.Context) {
	id := c.Query("id")
	db := Database.GetConn()
	icarts := Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE img_id = " + id + ";").Scan(&icarts)
	// Art浏览量+1
	icarts.Views++
	db.Exec("UPDATE ic_art SET views = " + fmt.Sprintf("%d", icarts.Views) + " WHERE img_id = " + id + ";")
	str, _ := json.Marshal(icarts)
	c.JSON(200, gin.H{
		"status": "success",
		"arts":   string(str),
	})
}

func Like(c *gin.Context) {
	id := c.Query("id")
	db := Database.GetConn()
	icarts := Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE img_id = " + id + ";").Scan(&icarts)
	icarts.Likes++
	db.Exec("UPDATE ic_art SET likes = " + fmt.Sprintf("%d", icarts.Likes) + " WHERE img_id = " + id + ";")
	c.JSON(200, gin.H{
		"status": "success",
	})
}
