package common

type RolePrivilege struct {
	name     string
	id       string
	children []*RolePrivilege
	useful   bool
}

func newRolePrivilege(id string, name string) *RolePrivilege {
	return &RolePrivilege{
		name:   name,
		id:     id,
		useful: true,
	}
}

func (r *RolePrivilege) GetId() string {
	return r.id
}

func (r *RolePrivilege) GetName() string {
	return r.name
}

func (r *RolePrivilege) GetChildren() []*RolePrivilege {
	return r.children
}

func (r *RolePrivilege) setUseful(useful bool) *RolePrivilege {
	r.useful = useful
	return r
}

var (
	rolePrivilegeCustomer     = newRolePrivilege("ec-customer", "客户").setUseful(false)
	rolePrivilegeCustomerBase = newRolePrivilege("ec-customer-base", "客户基础").setUseful(false)
	rolePrivilegeCustomerGH   = newRolePrivilege("ec-customer-gh", "客户公海").setUseful(false)
	rolePrivilegeCustomerMore = newRolePrivilege("ec-customer-more", "客户更多").setUseful(false)
	rolePrivilegeSales        = newRolePrivilege("ec-sales", "营销").setUseful(false)

	rolePrivilegeList = make([]*RolePrivilege, 0)
)

// 客户权限点
// 基础
var (
	rolePrivilegeCustomerView         = newRolePrivilege("ec-customer-view", "查看客户")
	rolePrivilegeCustomerAbandon      = newRolePrivilege("ec-customer-abandon", "放弃客户")
	rolePrivilegeCustomerExport       = newRolePrivilege("ec-customer-export", "导出客户")
	rolePrivilegeCustomerShare        = newRolePrivilege("ec-customer-share", "共享客户")
	rolePrivilegeCustomerDelete       = newRolePrivilege("ec-customer-delete", "删除客户")
	rolePrivilegeCustomerMerge        = newRolePrivilege("ec-customer-merge", "合并客户")
	rolePrivilegeCustomerRollback     = newRolePrivilege("ec-customer-rollback", "回滚客户进展")
	rolePrivilegeCustomerImport       = newRolePrivilege("ec-customer-import", "导入客户给同事")
	rolePrivilegeCustomerTypeTransfer = newRolePrivilege("ec-customer-type-transfer", "转换客户类型")
	rolePrivilegeCustomerDealProgress = newRolePrivilege("ec-customer-deal-progress", "设定已成交客户进展")

	rolePrivilegeCustomerTransfer     = newRolePrivilege("ec-customer-transfer", "转让客户")
	rolePrivilegeCustomerTransferSolo = newRolePrivilege("ec-customer-transfer-solo", "仅个人")
	rolePrivilegeCustomerTransferAll  = newRolePrivilege("ec-customer-transfer-all", "全企业")
	rolePrivilegeCustomerTransferDept = newRolePrivilege("ec-customer-transfer-dept", "所在部门及下级部门")

	rolePrivilegeCustomerAuth     = newRolePrivilege("ec-customer-auth", "权限范围")
	rolePrivilegeCustomerAuthSolo = newRolePrivilege("ec-customer-auth-solo", "仅个人")
	rolePrivilegeCustomerAuthAll  = newRolePrivilege("ec-customer-auth-all", "全企业")
	rolePrivilegeCustomerAuthDept = newRolePrivilege("ec-customer-auth-dept", "所在部门及下级部门")
)

// 客户权限点
// 公海
var (
	rolePrivilegeCustomerGHObtain     = newRolePrivilege("ec-customer-gh-obtain", "领取客户")
	rolePrivilegeCustomerGHView       = newRolePrivilege("ec-customer-gh-view", "查看公司大公海")
	rolePrivilegeCustomerGHExport     = newRolePrivilege("ec-customer-gh-export", "导出公海客户")
	rolePrivilegeCustomerGHDelete     = newRolePrivilege("ec-customer-gh-delete", "删除公海客户")
	rolePrivilegeCustomerGHTransfer   = newRolePrivilege("ec-customer-gh-transfer", "转换客户类型")
	rolePrivilegeCustomerGHRepeatView = newRolePrivilege("ec-customer-gh-repeat-view", "查看所有重复客户")

	rolePrivilegeCustomerGHDistribution        = newRolePrivilege("ec-customer-gh-distribution", "分配公海客户")
	rolePrivilegeCustomerGHDistributionAll     = newRolePrivilege("ec-customer-gh-distribution-all", "全企业")
	rolePrivilegeCustomerGHDistributionDept    = newRolePrivilege("ec-customer-gh-distribution-dept", "所在部门及下级部门")
	rolePrivilegeCustomerGHDistributionChoosen = newRolePrivilege("ec-customer-gh-distribution-choosen", "指定部门")
)

// 客户权限点
// 更多
var (
	rolePrivilegeCustomerMorePersonToEnt = newRolePrivilege("ec-customer-more-person-to-ent", "个人转企业客户")
)

// 营销权限点
var (
	rolePrivilegeSalesMaterial       = newRolePrivilege("ec-sales-material", "营销素材")
	rolePrivilegeSalesMaterialManage = newRolePrivilege("ec-sales-material-manage", "管理营销素材")
	rolePrivilegeSalesYKT            = newRolePrivilege("ec-sales-ykt", "营客通")
	rolePrivilegeSalesYKTManage      = newRolePrivilege("ec-sales-ykt-manage", "管理营客通")
)

// 装配权限点
func init() {
	// 客户权限点
	customerChildren := make([]*RolePrivilege, 0)
	customerChildren = append(customerChildren, rolePrivilegeCustomerBase, rolePrivilegeCustomerGH, rolePrivilegeCustomerMore)
	rolePrivilegeCustomer.children = customerChildren

	// 客户基础权限点
	customerBaseChildren := make([]*RolePrivilege, 0)
	customerBaseChildren = append(customerBaseChildren, rolePrivilegeCustomerView, rolePrivilegeCustomerAbandon, rolePrivilegeCustomerExport,
		rolePrivilegeCustomerShare, rolePrivilegeCustomerDelete, rolePrivilegeCustomerMerge, rolePrivilegeCustomerRollback, rolePrivilegeCustomerImport,
		rolePrivilegeCustomerTypeTransfer, rolePrivilegeCustomerDealProgress, rolePrivilegeCustomerTransfer, rolePrivilegeCustomerAuth)
	rolePrivilegeCustomerBase.children = customerBaseChildren

	// 客户-转让客户权限点
	customerTransferChildren := make([]*RolePrivilege, 0)
	customerTransferChildren = append(customerTransferChildren, rolePrivilegeCustomerTransferSolo, rolePrivilegeCustomerTransferAll, rolePrivilegeCustomerTransferDept)
	rolePrivilegeCustomerTransfer.children = customerTransferChildren

	// 客户-权限范围权限点
	customerAuthChildren := make([]*RolePrivilege, 0)
	customerAuthChildren = append(customerAuthChildren, rolePrivilegeCustomerAuthSolo, rolePrivilegeCustomerAuthAll, rolePrivilegeCustomerAuthDept)
	rolePrivilegeCustomerAuth.children = customerAuthChildren

	// 客户公海权限点
	customerGHChildren := make([]*RolePrivilege, 0)
	customerGHChildren = append(customerGHChildren, rolePrivilegeCustomerGHObtain, rolePrivilegeCustomerGHView, rolePrivilegeCustomerGHExport,
		rolePrivilegeCustomerGHDelete, rolePrivilegeCustomerGHDistribution, rolePrivilegeCustomerGHTransfer, rolePrivilegeCustomerGHRepeatView)
	rolePrivilegeCustomerGH.children = customerGHChildren

	// 客户-分配公海客户
	customerGHDistributionChildren := make([]*RolePrivilege, 0)
	customerGHDistributionChildren = append(customerGHDistributionChildren, rolePrivilegeCustomerGHDistributionAll, rolePrivilegeCustomerGHDistributionDept, rolePrivilegeCustomerGHDistributionChoosen)
	rolePrivilegeCustomerGHDistribution.children = customerGHDistributionChildren

	// 客户更多权限点
	customerMoreChildren := make([]*RolePrivilege, 0)
	customerMoreChildren = append(customerMoreChildren, rolePrivilegeCustomerMorePersonToEnt)
	rolePrivilegeCustomerMore.children = customerMoreChildren

	// 营销权限点
	salesChildren := make([]*RolePrivilege, 0)
	salesChildren = append(salesChildren, rolePrivilegeSalesMaterial, rolePrivilegeSalesYKT)
	rolePrivilegeSales.children = salesChildren

	// 营销素材权限点
	salesMaterialChildren := make([]*RolePrivilege, 0)
	salesMaterialChildren = append(salesMaterialChildren, rolePrivilegeSalesMaterialManage)
	rolePrivilegeSalesMaterial.children = salesMaterialChildren

	// 营销营客通权限点
	salesYKTChildren := make([]*RolePrivilege, 0)
	salesYKTChildren = append(salesYKTChildren, rolePrivilegeSalesYKTManage)
	rolePrivilegeSalesYKT.children = salesYKTChildren

	rolePrivilegeList = append(rolePrivilegeList, rolePrivilegeCustomer, rolePrivilegeSales)
}

func GetRolePrivilegeList() []*RolePrivilege {
	return rolePrivilegeList
}
