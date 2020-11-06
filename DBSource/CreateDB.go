package DBSource

import (
	"PrintYun/Config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"PrintYun/DBSource/Modules"
)

var db *gorm.DB

func DBConnect()  {
	DRIVER := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Config.MyUser,Config.Password, Config.Host, Config.Port, Config.DbName )
	var err error
	db,err = gorm.Open("mysql", DRIVER)
	if err !=nil{
		panic(err)
	}
	db.SingularTable(true)
	CheckTableExist(db)
}

func GetDB() *gorm.DB {
	return db
}

func CheckTableExist(db *gorm.DB)  {
	// Admin
	if db.HasTable(&Modules.Admin{}) == false{
		fmt.Println("Create Admin Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.Admin{})
	}else{
		fmt.Println("Admin Table has exist~~")
	}
	// User
	if db.HasTable(&Modules.User{}) == false{
		fmt.Println("Create Users Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.User{})
	}else{
		fmt.Println("Admin Users has exist~~")
	}
	// Province
	if db.HasTable(&Modules.Province{}) == false{
		fmt.Println("Create Province Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.Province{})
	}else{
		fmt.Println("Admin Province has exist~~")
	}
	// CIty
	if db.HasTable(&Modules.City{}) == false{
		fmt.Println("Create City Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.City{})
	}else{
		fmt.Println("Admin City has exist~~")
	}
	// Area
	if db.HasTable(&Modules.Area{}) == false{
		fmt.Println("Create Area Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.Area{})
	}else{
		fmt.Println("Admin Area has exist~~")
	}
	// Shop
	if db.HasTable(&Modules.Shop{}) == false{
		fmt.Println("Create Shop Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.Shop{})
		db.Model(&Modules.Shop{}).AddForeignKey("admin_id", "admin(id)", "RESTRICT", "RESTRICT")
	}else{
		fmt.Println("Admin Shop has exist~~")
	}
	// Machine
	if db.HasTable(&Modules.Machine{}) == false{
		fmt.Println("Create Machine Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.Machine{})
		db.Model(&Modules.Machine{}).AddForeignKey("shop_id", "shop(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Machine{}).AddForeignKey("province_id", "province(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Machine{}).AddForeignKey("city_id", "city(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Machine{}).AddForeignKey("area_id", "area(id)", "RESTRICT", "RESTRICT")
	}else{
		fmt.Println("Admin Machine has exist~~")
	}
	// Order
	if db.HasTable(&Modules.Order{}) == false{
		fmt.Println("Create Order Table~~")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Modules.Order{})
		db.Model(&Modules.Order{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Order{}).AddForeignKey("province_id", "province(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Order{}).AddForeignKey("city_id", "city(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Order{}).AddForeignKey("area_id", "area(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Order{}).AddForeignKey("machine_id", "machine(id)", "RESTRICT", "RESTRICT")
		db.Model(&Modules.Order{}).AddForeignKey("shop_id", "shop(id)", "RESTRICT", "RESTRICT")
	}else{
		fmt.Println("Admin Order has exist~~")
	}
}
