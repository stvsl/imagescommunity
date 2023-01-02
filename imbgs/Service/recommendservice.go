package service

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	Database "stvsljl.com/stvsl/imbgs/DBLogic"
)

func Recommend(c *gin.Context) {
	ktarts := []Database.IcArt{}
	db := Database.GetConn()
	db.Raw("SELECT * FROM ic_art WHERE name LIKE '%动漫%' OR label LIKE '%动漫%' ORDER BY RAND() LIMIT 15;").Scan(&ktarts)
	ktartsJson, _ := json.Marshal(ktarts)
	// 第二次查询与特摄相关的图片
	tsarts2 := []Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE name LIKE '%特摄%' OR label LIKE '%特效%' ORDER BY RAND() LIMIT 15;").Scan(&tsarts2)
	tsarts2Json, _ := json.Marshal(tsarts2)
	// 第三次查询与照片相关的图片
	zparts3 := []Database.IcArt{}
	db.Raw("SELECT * FROM ic_art WHERE name LIKE '%照片%' OR label LIKE '%照片%' ORDER BY RAND() LIMIT 15;").Scan(&zparts3)
	zparts3Json, _ := json.Marshal(zparts3)
	// 随机查询30条数据
	part := []Database.IcArt{}
	db.Raw("SELECT * FROM ic_art ORDER BY RAND() LIMIT 30;").Scan(&part)
	partJson, _ := json.Marshal(part)

	c.JSON(200, gin.H{
		"status": "success",
		"ktarts": string(ktartsJson),
		"tsarts": string(tsarts2Json),
		"zparts": string(zparts3Json),
		"parts":  string(partJson),
	})
}
