package Modules

import (
	_ "github.com/jinzhu/gorm"
)


type User struct {
	Id int `gorm:"AUTO_INCREMENT;primary_key"`
	OpenId string `gorm:"type:varchar(64)"`  				// 用户唯一标识
	UnionId string `gorm:"type:varchar(64)"`				// 用户在开放平台的唯一标识符
	InviteCode string `gorm:"type:varchar(16)"`				// 邀请码
	SignTime int64 `gorm:"autoCreateTime"`					// 注册时间
	LashSignTime int64 `gorm:"autoUpdateTime:milli"`		// 用户上次登陆时间
	StatisticsTimes int `gorm:"type:int(16);default:0"`     // 使用次数
	Shop []Shop `gorm:"foreignKey:UserId"`					// 关联Shop表，建立外键
	Order []Order `gorm:"foreignKey:UserId"`				// 关联Order表，建立外键
}

type Admin struct {
		Id int `gorm:"AUTO_INCREMENT;primary_key"`
		UserName string `gorm:"type:varchar(32);not null;unique"`   	// 用户名, 注册使用手机号
		PassWord string `gorm:"type:varchar(64);not null"` 				// 密码
		Root int
		SignTime int64 `gorm:"autoCreateTime"`
		SystemAuthority int `gorm:"not null"`   						// 商户权限
		InviteCode string `gorm:"type:varchar(16);not null;unique"`
}