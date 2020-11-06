package Modules

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

type Shop struct {
	Id int `gorm:"AUTO_INCREMENT;primary_key"`
	Admin Admin `gorm:"ForeignKey:Id;AssociationForeignKey:AdminId"`
	AdminId int
	ShopName string `gorm:"type:varchar(64);not null"`   // 店铺名称
}

type Machine struct {
	Id int `gorm:"AUTO_INCREMENT;primary_key"`
	Province Province `gorm:"ForeignKey:Id;AssociationForeignKey:ProvinceId"`
	ProvinceId int 										// 用户选择打印的省份
	City City `gorm:"ForeignKey:Id;AssociationForeignKey:CityId"`
	CityId int											// 用户选择打印的城市
	Area Area `gorm:"ForeignKey:Id;AssociationForeignKey:AreaId"`
	AreaId int 											// 用户选择打印的区域
	Code string `gorm:"type:varchar(64);not null;unique"`
	Shop Shop `gorm:"ForeignKey:ID;AssociationForeignKey:ShopID"`
	ShopID int											// 用户选择打印的商铺ID
	SignTime time.Time
	WorkTimes int `gorm:"default:0"`
}

type Order struct {
	Id int `gorm:"AUTO_INCREMENT;primary_key"`
	User User `gorm:"ForeignKey:Id;AssociationForeignKey:UserId"`
	UserId int
	FileName string `gorm:"type:varchar(128);not null"`	// '上传时的文件名'
	OrderName string `gorm:"type:varchar(128);not null;unique"`	// '上传时的文件名'
	Page int											// 打印文件的页数
	Color int `gorm:"not null"`							// 打印的颜色, 1为黑白，2为彩色
	Direction int `gorm:"not null"`						// 打印的方向, 1为竖向，2为横向
	Num int `gorm:"not null"`							// 打印的份数
	Remarks string `gorm:"type:text(512);not null"`		// 用户备注
	Province Province `gorm:"ForeignKey:Id;AssociationForeignKey:ProvinceId"`
	ProvinceId int 										// 用户选择打印的省份
	City City `gorm:"ForeignKey:Id;AssociationForeignKey:CityId"`
	CityId int											// 用户选择打印的城市
	Area Area `gorm:"ForeignKey:Id;AssociationForeignKey:AreaId"`
	AreaId int 											// 用户选择打印的区域
	Machine Machine `gorm:"ForeignKey:Id;AssociationForeignKey:MachineId"`
	MachineId int										// 用户选择打印的机器
	Shop Shop `gorm:"ForeignKey:Id;AssociationForeignKey:ShopID"`
	ShopID int											// 用户选择打印的商铺ID
}