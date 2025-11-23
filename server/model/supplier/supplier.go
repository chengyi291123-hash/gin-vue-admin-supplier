package supplier

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Supplier struct {
	global.GVA_MODEL
	EnterpriseName string `json:"enterprise_name" form:"enterprise_name" gorm:"comment:企业名称"`
	CreditCode     string `json:"credit_code" form:"credit_code" gorm:"comment:信用代码"`
	EntryType      string `json:"entry_type" form:"entry_type" gorm:"comment:类型"`
	Category       string `json:"category" form:"category" gorm:"comment:品类"`
	Region         string `json:"region" form:"region" gorm:"comment:区域"`
	Industry       string `json:"industry" form:"industry" gorm:"comment:行业"`
	Brand          string `json:"brand" form:"brand" gorm:"comment:品牌"`
	ContactPerson  string `json:"contact_person" form:"contact_person" gorm:"comment:联系人"`
	Mobile         string `json:"mobile" form:"mobile" gorm:"comment:电话"`
	Purchaser      string `json:"purchaser" form:"purchaser" gorm:"comment:采购员"`
	BankName       string `json:"bank_name" form:"bank_name" gorm:"comment:开户行"`
	BranchName     string `json:"branch_name" form:"branch_name" gorm:"comment:支行"`
	BankAccount    string `json:"bank_account" form:"bank_account" gorm:"comment:账号"`
	Settlement     string `json:"settlement" form:"settlement" gorm:"comment:结算"`
	Status         string `json:"status" form:"status" gorm:"comment:状态"` // temp, qualified, blacklist
}

func (Supplier) TableName() string {
	return "sys_suppliers"
}
