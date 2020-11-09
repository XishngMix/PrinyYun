package DBSource

import (
	"PrintYun/Config"
	"PrintYun/DBSource/Modules"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var DRIVER string
var PO *gormadapter.Adapter
var Enforcer *casbin.Enforcer
var err error

func DBConnect()  {
	DRIVER = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Config.MyUser,Config.Password, Config.Host, Config.Port, Config.DbName)
	var err error
	db,err = gorm.Open("mysql", DRIVER)
	if err !=nil{
		panic(err)
		fmt.Sprintf("gorm open 错误: %v", err)
	}
	db.SingularTable(true)
	CheckTableExist()				// 检查数据表存在情况
	//CreateAuthority()   			// casbin增加权限表
}

func GetDB() *gorm.DB {
	return db
}

func CheckTableExist()  {
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

// 权限验证
func CreateAuthority()  {
	// 将数据库连接同步给插件， 插件用来操作数据库
	PO = gormadapter.NewAdapterByDB(db)
	// 这里也可以使用原生字符串方式
	///home/xishng/文档/PrintYun/DBSource/rbac_model.conf
	Enforcer = casbin.NewEnforcer("DBSource/rbac_model.conf", PO)
	if err != nil {
		fmt.Sprintf("Happen a error: %v", err)
	}
	// 开启权限认证日志
	Enforcer.EnableLog(true)
	// 加载数据库中的策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		fmt.Println("loadPolicy error")
		panic(err)
	}
}