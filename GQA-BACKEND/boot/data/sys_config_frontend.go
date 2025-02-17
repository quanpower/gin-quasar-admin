package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysConfigFrontend = new(sysConfigFrontend)

type sysConfigFrontend struct{}

var sysConfigFrontendData = []system.SysConfigFrontend{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1001, Remark: "网站主标题", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "gqaMainTitle", Default: "Gin-Quasar-Admin",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1002, Remark: "网站次标题", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "gqaSubTitle", Default: "Gin-Quasar-Admin",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1003, Remark: "网站描述", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "gqaDescribe", Default: "Lorem ipsum dolor sit amet consectetur adipisicing elit",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1004, Remark: "登录页插件", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "gqaPluginLoginLayout", Default: "",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1005, Remark: "网站Logo", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "gqaWebLogo", Default: "",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1006, Remark: "标签页Logo", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "gqaHeaderLogo", Default: "",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1007, Remark: "显示仓库入口", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "gqaShowGit", Default: "yes",
	},
}

func (s *sysConfigFrontend) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysConfigFrontend{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_config_frontend 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[Gin-Quasar-Admin] --> sys_config_frontend 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysConfigFrontendData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_config_frontend 表初始数据成功！")
		global.GqaLog.Info("[Gin-Quasar-Admin] --> sys_config_frontend 表初始数据成功！")
		return nil
	})
}
