package system

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
)

type SysRole struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	RoleCode                 string    `json:"roleCode" gorm:"comment:角色编码;not null;uniqueIndex;"`
	RoleName                 string    `json:"roleName" gorm:"comment:角色名称;not null;unique;"`
	DeptDataPermissionType   string    `json:"deptDataPermissionType" gorm:"comment:部门数据权限分类;not null;default:user"`
	DeptDataPermissionCustom string    `json:"deptDataPermissionCustom" gorm:"comment:自定义部门数据权限;type:text;"`
	User                     []SysUser `json:"user" gorm:"many2many:sys_user_role;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Username;joinReferences:SysUserUsername;"`
	Menu                     []SysMenu `json:"menu" gorm:"many2many:sys_role_menu;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Name;joinReferences:SysMenuName;"`
}

type RequestAddRole struct {
	Status   string `json:"status"`
	Sort     uint   `json:"sort"`
	Remark   string `json:"remark"`
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}

type RequestRoleList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddRole 中的字段
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysRole
}

type RequestRoleCode struct {
	RoleCode string `json:"roleCode"`
}

type RequestRoleUser struct {
	RoleCode string `json:"roleCode"`
	Username string `json:"username"`
}

type RequestRoleUserAdd struct {
	RoleCode string   `json:"roleCode"`
	Username []string `json:"username"`
}

type RequestRoleMenuEdit struct {
	RoleCode string            `json:"roleCode"`
	RoleMenu []RequestRoleMenu `json:"roleMenu"`
}

type RequestRoleMenu struct {
	RoleCode string `json:"roleCode"`
	MenuName string `json:"menuName"`
}

type RequestRoleApiEdit struct {
	RoleCode string      `json:"roleCode"`
	Policy   []SysCasbin `json:"policy"`
}

type RequestRoleDeptDataPermission struct {
	RoleCode                 string `json:"roleCode"`
	DeptDataPermissionType   string `json:"deptDataPermissionType"`
	DeptDataPermissionCustom string `json:"deptDataPermissionCustom"`
}
