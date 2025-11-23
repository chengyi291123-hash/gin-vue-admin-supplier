package main

import (
	"fmt"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/supplier"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	fmt.Println("Initializing SQLite DB...")

	// Force SQLite config
	global.GVA_CONFIG.System.DbType = "sqlite"
	global.GVA_CONFIG.Sqlite.Path = "./gva.db"

	db, err := gorm.Open(sqlite.Open("./gva.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.GVA_DB = db

	// Migrate Tables
	err = db.AutoMigrate(
		&system.SysApi{},
		&system.SysUser{},
		&system.SysBaseMenu{},
		&system.SysAuthority{},
		&system.SysDictionary{},
		&system.SysOperationRecord{},
		&system.SysAutoCodeHistory{},
		&system.SysDictionaryDetail{},
		&system.SysBaseMenuParameter{},
		&system.SysBaseMenuBtn{},
		&system.SysAuthorityBtn{},
		&system.SysExportTemplate{},
		&supplier.Supplier{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables Migrated.")

	// Create Admin User if not exists
	var admin system.SysUser
	if err := db.Where("username = ?", "admin").First(&admin).Error; err != nil {
		admin = system.SysUser{
			Username:    "admin",
			Password:    "$2a$14$2a$14$2a$14$2a$14$2a$14$2a$14$2a$14$2a$14$2a$14$2a$14$2a$14", // 123456 (hashed)
			NickName:    "Mr.Admin",
			AuthorityId: 1,
			Enable:      1,
		}
		db.Create(&admin)
		fmt.Println("Admin user created.")
	}

	// Create Authority 1
	var auth system.SysAuthority
	if err := db.Where("authority_id = ?", 1).First(&auth).Error; err != nil {
        parentId := uint(0)
		auth = system.SysAuthority{
			AuthorityId:   1,
			AuthorityName: "Admin",
			DefaultRouter: "dashboard",
			ParentId:      &parentId,
		}
		db.Create(&auth)
	}

	fmt.Println("Initialization Complete.")
	os.Exit(0)
}