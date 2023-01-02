package Database

import "time"

/**
CREATE TABLE `ic_user` (
		`uuid` uuid NOT NULL COMMENT 'UUID',
		`uid` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
		`type` int(11) NOT NULL DEFAULT 1 COMMENT '用户类型',
		`name` varchar(25) NOT NULL COMMENT '姓名',
		`sex` int(11) DEFAULT 2 COMMENT '性别',
		`avatar` blob DEFAULT NULL COMMENT '头像',
		`passwd` varchar(50) NOT NULL COMMENT '密码',
		`tel` int(12) DEFAULT NULL COMMENT '电话号码',
		`actnum` int(8) DEFAULT 0 COMMENT '作品数',
		PRIMARY KEY (`uuid`),
		UNIQUE KEY `ic_user_UN` (`uid`)
	  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
*/
// @Author: stvsl
type IcUser struct {
	Uuid   string `gorm:"column:uuid;primary_key;NOT NULL;comment:'UUID'"`
	Uid    int32  `gorm:"column:uid;AUTO_INCREMENT;NOT NULL;comment:'ID'"`
	Type   int32  `gorm:"column:type;default:1;NOT NULL;comment:'用户类型'"`
	Name   string `gorm:"column:name;NOT NULL;comment:'姓名'"`
	Sex    int32  `gorm:"column:sex;default:2;comment:'性别'"`
	Avatar string `gorm:"column:avatar;default:NULL;comment:'头像'"`
	Passwd string `gorm:"column:passwd;NOT NULL;comment:'密码'"`
	Tel    string `gorm:"column:tel;default:NULL;comment:'电话号码'"`
	Actnum int32  `gorm:"column:actnum;default:0;comment:'作品数'"`
}

func (i *IcUser) TableName() string {
	return "ic_user"
}

type IcLikes struct {
	Id        uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Uuid      string `gorm:"column:uuid;default:NULL;comment:'用户UUID'"`
	ImgId     uint64 `gorm:"column:img_id;default:NULL;comment:'图片ID'"`
	Name      string
	Uri       string
	OwnerUuid string
}

func (i *IcLikes) TableName() string {
	return "ic_likes"
}

type IcComment struct {
	Imgid    uint64    `gorm:"column:imgid;NOT NULL;comment:'作品编号'"`
	Comments string    `gorm:"column:comments;NOT NULL;comment:'评论内容'"`
	Uuid     string    `gorm:"column:uuid;NOT NULL;comment:'来自用户'"`
	Id       uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Time     time.Time `gorm:"column:time;NOT NULL"`
	Name     string
	Avatar   string
}

func (i *IcComment) TableName() string {
	return "ic_comment"
}

type IcArt struct {
	ImgId       uint64    `gorm:"column:img_id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'图片ID'"`
	Describe    string    `gorm:"column:describe;default:NULL;comment:'图片描述'"`
	Name        string    `gorm:"column:name;default:NULL;comment:'图片名称'"`
	Uri         string    `gorm:"column:uri;default:NULL;comment:'所在图床位置'"`
	UploadTime  time.Time `gorm:"column:upload_time;NOT NULL;comment:'上传时间'"`
	ReleaseTime time.Time `gorm:"column:release_time;default:NULL;comment:'发布时间'"`
	Status      int32     `gorm:"column:status;default:0;NOT NULL;comment:'状态'"`
	OwnerUuid   string    `gorm:"column:owner_uuid;default:NULL;comment:'所属用户'"`
	Views       uint64    `gorm:"column:views;default:0;NOT NULL;comment:'浏览量'"`
	Likes       uint64    `gorm:"column:likes;default:0;NOT NULL;comment:'点赞量'"`
	Collects    uint64    `gorm:"column:collects;default:0;NOT NULL;comment:'收藏量'"`
	Label       string    `gorm:"column:label;default:NULL;comment:'标签'"`
}

func (i *IcArt) TableName() string {
	return "ic_art"
}

type IcAdmin struct {
	Uuid   string `gorm:"column:uuid;primary_key;NOT NULL;comment:'uuid'"`
	Id     string `gorm:"column:id;NOT NULL;comment:'id'"`
	Name   string `gorm:"column:name;NOT NULL;comment:'姓名'"`
	Passwd string `gorm:"column:passwd;NOT NULL;comment:'密码'"`
}

func (i *IcAdmin) TableName() string {
	return "ic_admin"
}
