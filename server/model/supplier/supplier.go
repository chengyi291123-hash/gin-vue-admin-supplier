package supplier

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Supplier 供应商基础信息表
type Supplier struct {
	global.GVA_MODEL
	EnterpriseName   string `json:"enterprise_name" form:"enterprise_name" gorm:"comment:企业名称"`
	CreditCode       string `json:"credit_code" form:"credit_code" gorm:"comment:统一社会信用代码"`
	EntryType        string `json:"entry_type" form:"entry_type" gorm:"comment:供应商类型"` // manufacturing, trading
	Category         string `json:"category" form:"category" gorm:"comment:供应品类"`     // raw, electronic, office
	Region           string `json:"region" form:"region" gorm:"comment:合作区域"`
	Industry         string `json:"industry" form:"industry" gorm:"comment:合作行业"` // normal, nuclear, military, petrochemical
	Brand            string `json:"brand" form:"brand" gorm:"comment:合作品牌"`
	ContactPerson    string `json:"contact_person" form:"contact_person" gorm:"comment:联系人"`
	Mobile           string `json:"mobile" form:"mobile" gorm:"comment:联系电话"`
	Email            string `json:"email" form:"email" gorm:"comment:邮箱"`
	Purchaser        string `json:"purchaser" form:"purchaser" gorm:"comment:采购员"`
	BankName         string `json:"bank_name" form:"bank_name" gorm:"comment:开户行"`
	BranchName       string `json:"branch_name" form:"branch_name" gorm:"comment:开户行名称"`
	BankAccount      string `json:"bank_account" form:"bank_account" gorm:"comment:银行账号"`
	Settlement       string `json:"settlement" form:"settlement" gorm:"comment:结算方式"` // m30, m45, m60
	Status           string `json:"status" form:"status" gorm:"default:temp;comment:状态"` // temp(潜在), qualified(合格), blacklist(黑名单)
	IsBlacklist      string `json:"is_blacklist" form:"is_blacklist" gorm:"default:否;comment:是否黑名单"`
	BlacklistReason  string `json:"blacklist_reason" form:"blacklist_reason" gorm:"comment:黑名单原因"`
	ApprovalStatus   string `json:"approval_status" form:"approval_status" gorm:"comment:审批状态"` // pending, approved, rejected
	CurrentNode      string `json:"current_node" form:"current_node" gorm:"comment:当前审批节点"`
	RejectType       string `json:"reject_type" form:"reject_type" gorm:"comment:驳回类型"` // 填写问题, 资质不合格
	RejectReason     string `json:"reject_reason" form:"reject_reason" gorm:"comment:驳回原因"`
}

func (Supplier) TableName() string {
	return "sys_suppliers"
}

// SupplierCertificate 供应商资质证书表
type SupplierCertificate struct {
	global.GVA_MODEL
	SupplierID     uint   `json:"supplier_id" form:"supplier_id" gorm:"comment:供应商ID"`
	CertName       string `json:"cert_name" form:"cert_name" gorm:"comment:证书名称"`
	FilePath       string `json:"file_path" form:"file_path" gorm:"comment:文件路径"`
	FileName       string `json:"file_name" form:"file_name" gorm:"comment:文件名"`
	ValidStartDate string `json:"valid_start_date" form:"valid_start_date" gorm:"comment:有效期开始"`
	ValidEndDate   string `json:"valid_end_date" form:"valid_end_date" gorm:"comment:有效期结束"`
}

func (SupplierCertificate) TableName() string {
	return "sys_supplier_certificates"
}

// SupplierApproval 供应商审批记录表
type SupplierApproval struct {
	global.GVA_MODEL
	SupplierID   uint   `json:"supplier_id" form:"supplier_id" gorm:"comment:供应商ID"`
	NodeName     string `json:"node_name" form:"node_name" gorm:"comment:审批节点"`
	Approver     string `json:"approver" form:"approver" gorm:"comment:审批人"`
	Status       string `json:"status" form:"status" gorm:"comment:状态"` // pending, completed, rejected, current
	Comment      string `json:"comment" form:"comment" gorm:"comment:审批意见"`
	ApprovalTime string `json:"approval_time" form:"approval_time" gorm:"comment:审批时间"`
}

func (SupplierApproval) TableName() string {
	return "sys_supplier_approvals"
}

// PurchaseOrder 采购订单表
type PurchaseOrder struct {
	global.GVA_MODEL
	OrderNo        string  `json:"order_no" form:"order_no" gorm:"uniqueIndex;comment:订单编号"`
	SupplierID     uint    `json:"supplier_id" form:"supplier_id" gorm:"comment:供应商ID"`
	SupplierName   string  `json:"supplier_name" form:"supplier_name" gorm:"comment:供应商名称"`
	ProjectName    string  `json:"project_name" form:"project_name" gorm:"comment:关联项目"`
	TotalAmount    float64 `json:"total_amount" form:"total_amount" gorm:"comment:订单总额"`
	OverLimitCount int     `json:"over_limit_count" form:"over_limit_count" gorm:"comment:超限项数"`
	Applicant      string  `json:"applicant" form:"applicant" gorm:"comment:申请人"`
	ApplyTime      string  `json:"apply_time" form:"apply_time" gorm:"comment:申请时间"`
	Status         string  `json:"status" form:"status" gorm:"default:pending;comment:状态"` // pending, approved, rejected
	ApprovalRemark string  `json:"approval_remark" form:"approval_remark" gorm:"comment:审批意见"`
}

func (PurchaseOrder) TableName() string {
	return "sys_purchase_orders"
}

// PurchaseOrderItem 采购订单明细表
type PurchaseOrderItem struct {
	global.GVA_MODEL
	OrderID        uint    `json:"order_id" form:"order_id" gorm:"comment:订单ID"`
	MaterialName   string  `json:"material_name" form:"material_name" gorm:"comment:物料名称"`
	Specification  string  `json:"specification" form:"specification" gorm:"comment:规格型号"`
	Quantity       int     `json:"quantity" form:"quantity" gorm:"comment:数量"`
	Unit           string  `json:"unit" form:"unit" gorm:"comment:单位"`
	UnitPrice      float64 `json:"unit_price" form:"unit_price" gorm:"comment:采购单价"`
	AgreementPrice float64 `json:"agreement_price" form:"agreement_price" gorm:"comment:协议价"`
	YearLowestPrice float64 `json:"year_lowest_price" form:"year_lowest_price" gorm:"comment:本年度最低价"`
	LastYearPrice  float64 `json:"last_year_price" form:"last_year_price" gorm:"comment:上年度最后价"`
	OverLimit      bool    `json:"over_limit" form:"over_limit" gorm:"comment:是否超限"`
}

func (PurchaseOrderItem) TableName() string {
	return "sys_purchase_order_items"
}

// AgreementPrice 协议价表
type AgreementPrice struct {
	global.GVA_MODEL
	MaterialCode string  `json:"material_code" form:"material_code" gorm:"comment:物料编码"`
	MaterialName string  `json:"material_name" form:"material_name" gorm:"comment:物料名称"`
	SupplierID   uint    `json:"supplier_id" form:"supplier_id" gorm:"comment:供应商ID"`
	SupplierName string  `json:"supplier_name" form:"supplier_name" gorm:"comment:供应商名称"`
	Price        float64 `json:"price" form:"price" gorm:"comment:协议价"`
	ValidDate    string  `json:"valid_date" form:"valid_date" gorm:"comment:有效期"`
}

func (AgreementPrice) TableName() string {
	return "sys_agreement_prices"
}
