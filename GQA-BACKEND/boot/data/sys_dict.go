package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysDict = new(sysDict)

type sysDict struct{}

var sysDictData = []system.SysDict{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1001, Remark: "这是性别字典", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "gender", DictLabel: "性别",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1002, Remark: "这是启用状态", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "statusOnOff", DictLabel: "启用状态",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1003, Remark: "这是是否状态", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "statusYesNo", DictLabel: "是否状态",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1004, Remark: "这是部门数据权限分类", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "deptDataPermissionType", DictLabel: "部门数据权限分类",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1005, Remark: "这是消息类型", CreatedAt: time.Now(), CreatedBy: "admin"},
		DictCode: "noticeType", DictLabel: "消息类型",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是男", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "gender", DictCode: "m", DictLabel: "男",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是女", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "gender", DictCode: "f", DictLabel: "女",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "这是保密", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "gender", DictCode: "u", DictLabel: "保密",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是启用", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "statusOnOff", DictCode: "on", DictLabel: "启用",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是禁用", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "statusOnOff", DictCode: "off", DictLabel: "禁用",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是是", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "statusYesNo", DictCode: "yes", DictLabel: "是",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是否", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "statusYesNo", DictCode: "no", DictLabel: "否",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是全部数据权限", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "deptDataPermissionType", DictCode: "all", DictLabel: "全部部门数据权限",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是本人数据权限", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "deptDataPermissionType", DictCode: "user", DictLabel: "用户本人数据权限",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "这是用户所在部门数据权限", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "deptDataPermissionType", DictCode: "dept", DictLabel: "用户所在部门数据权限",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, Remark: "这是用户所在部门及下级部门数据权限", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "deptDataPermissionType", DictCode: "deptWithChildren", DictLabel: "用户所在部门及下级部门数据权限",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, Remark: "这是自定义数据权限", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "deptDataPermissionType", DictCode: "custom", DictLabel: "自定义部门数据权限",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是系统消息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "noticeType", DictCode: "system", DictLabel: "系统消息",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是消息提示", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "noticeType", DictCode: "message", DictLabel: "消息提示",
	},
}

func (s *sysDict) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysDict{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_dict 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[Gin-Quasar-Admin] --> sys_dict 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_dict 表初始数据成功！")
		global.GqaLog.Info("[Gin-Quasar-Admin] --> sys_dict 表初始数据成功！")
		return nil
	})
}
