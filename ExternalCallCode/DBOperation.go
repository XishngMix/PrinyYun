package ExternalCallCode

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

func CollBackDB(db *gorm.DB, err error, ctx iris.Context) (bool, int) {
	if err != nil && db == nil {
		db.Rollback()
		ctx.JSON(map[string]string {
			"message":"数据库操作错误,这里回调了",
		})
		return false, 1005
	}
	return true, 1000
}