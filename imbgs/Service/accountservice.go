package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	Database "stvsljl.com/stvsl/imbgs/DBLogic"
	LsKy "stvsljl.com/stvsl/imbgs/lsky"
)

func Register(c *gin.Context) {
	// 直接打印json中的数据
	var user struct {
		Name   string
		Sex    string
		Avatar string
		Passwd string
		Tel    string
	}
	c.BindJSON(&user)
	url := LsKy.Upload(user.Avatar)
	db := Database.GetConn()
	db.Exec("INSERT INTO ic_user" +
		"(uuid, `type`, name, sex, avatar, passwd, tel, actnum)" +
		"VALUES(uuid(), 1, '" + user.Name + "', " + user.Sex + ", '" + url + "', '" + user.Passwd + "', '" + user.Tel + "', 0);")
	icuser := Database.IcUser{}
	// 查询数据库，返回信息
	db.Raw("SELECT * FROM ic_user WHERE name = '" + user.Name + "'AND passwd = '" + user.Passwd + "';").Scan(&icuser)

	c.JSON(200, gin.H{
		"status":   "success",
		"uuid":     icuser.Uuid,
		"id":       icuser.Uid,
		"username": icuser.Name,
		"avatar":   icuser.Avatar,
		"sex":      icuser.Sex,
		"tel":      icuser.Tel,
		"actnum":   icuser.Actnum,
	})
}

func Login(c *gin.Context) {
	var user struct {
		Id     string
		Passwd string
	}
	c.BindJSON(&user)
	db := Database.GetConn()
	icuser := Database.IcUser{}
	// 查询数据库，返回信息
	db.Raw("SELECT * FROM ic_user WHERE tel = '" + user.Id + "'AND passwd = '" + user.Passwd + "';").Scan(&icuser)
	if icuser.Uuid == "" {
		c.JSON(200, gin.H{
			"status": "fail",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":   "success",
		"uuid":     icuser.Uuid,
		"id":       icuser.Uid,
		"username": icuser.Name,
		"avatar":   icuser.Avatar,
		"sex":      icuser.Sex,
		"tel":      icuser.Tel,
		"actnum":   icuser.Actnum,
	})
}

func GetUserName(c *gin.Context) {
	var user struct {
		Uuid string
	}
	c.BindJSON(&user)
	db := Database.GetConn()
	icuser := Database.IcUser{}
	// 查询数据库，返回信息
	db.Raw("SELECT * FROM ic_user WHERE uuid = '" + user.Uuid + "';").Scan(&icuser)
	c.JSON(200, gin.H{
		"status": "success",
		"name":   icuser.Name,
	})
}

func GetUser(c *gin.Context) {
	uuid := c.Query("uuid")
	db := Database.GetConn()
	icuser := Database.IcUser{}
	// 查询数据库，返回信息
	db.Raw("SELECT * FROM ic_user WHERE uuid = '" + uuid + "';").Scan(&icuser)
	str, _ := json.Marshal(icuser)
	c.JSON(200, gin.H{
		"status": "success",
		"user":   string(str),
	})
}

func UpdateUser(c *gin.Context) {
	var user struct {
		Uuid   string
		Name   string
		Avatar string
		Sex    int
		Tel    string
	}
	c.BindJSON(&user)
	db := Database.GetConn()
	url := LsKy.Upload(user.Avatar)
	icuser := Database.IcUser{}
	// 查询数据库，返回信息
	db.Raw("UPDATE ic_user SET name = '" + user.Name + "', avatar = '" + url + "',Sex = " + strconv.Itoa(user.Sex) + ", tel = '" + user.Tel + "' WHERE uuid = '" + user.Uuid + "';").Scan(&icuser)
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func UpdatePasswd(c *gin.Context) {
	var user struct {
		Uuid   string
		Passwd string
	}
	c.BindJSON(&user)
	db := Database.GetConn()
	// 查询数据库，返回信息
	db.Exec("UPDATE ic_user SET passwd = '" + user.Passwd + "' WHERE uuid = '" + user.Uuid + "';")
	fmt.Println("UPDATE ic_user SET passwd = '" + user.Passwd + "' WHERE uuid = '" + user.Uuid + "';")
	c.JSON(200, gin.H{
		"status": "success",
	})
}
