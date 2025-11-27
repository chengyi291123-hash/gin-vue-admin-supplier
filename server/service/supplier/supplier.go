package supplier

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/supplier"
)

type SupplierService struct{}

// ==================== 供应商基础操作 ====================

func (s *SupplierService) CreateSupplier(sup supplier.Supplier) (supplier.Supplier, error) {
	err := global.GVA_DB.Create(&sup).Error
	return sup, err
}

func (s *SupplierService) CreateSupplierWithCerts(sup supplier.Supplier, certs []supplier.SupplierCertificate) (supplier.Supplier, error) {
	tx := global.GVA_DB.Begin()
	if err := tx.Create(&sup).Error; err != nil {
		tx.Rollback()
		return sup, err
	}
	for i := range certs {
		certs[i].SupplierID = sup.ID
		if err := tx.Create(&certs[i]).Error; err != nil {
			tx.Rollback()
			return sup, err
		}
	}
	return sup, tx.Commit().Error
}

func (s *SupplierService) DeleteSupplier(sup supplier.Supplier) (err error) {
	err = global.GVA_DB.Delete(&sup).Error
	return err
}

func (s *SupplierService) UpdateSupplier(sup supplier.Supplier) (err error) {
	err = global.GVA_DB.Save(&sup).Error
	return err
}

func (s *SupplierService) GetSupplier(id uint) (sup supplier.Supplier, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sup).Error
	return
}

func (s *SupplierService) GetSupplierWithCerts(id uint) (sup supplier.Supplier, certs []supplier.SupplierCertificate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sup).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Where("supplier_id = ?", id).Find(&certs).Error
	return
}

func (s *SupplierService) GetSupplierInfoList(info request.PageInfo, search supplier.Supplier) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&supplier.Supplier{})
	var sups []supplier.Supplier

	if search.EnterpriseName != "" {
		db = db.Where("enterprise_name LIKE ?", "%"+search.EnterpriseName+"%")
	}
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	if search.ApprovalStatus != "" {
		db = db.Where("approval_status = ?", search.ApprovalStatus)
	}
	if search.IsBlacklist != "" {
		db = db.Where("is_blacklist = ?", search.IsBlacklist)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&sups).Error
	return sups, total, err
}

// ==================== 供应商证书操作 ====================

func (s *SupplierService) CreateCertificate(cert supplier.SupplierCertificate) (err error) {
	return global.GVA_DB.Create(&cert).Error
}

func (s *SupplierService) UpdateCertificate(cert supplier.SupplierCertificate) (err error) {
	return global.GVA_DB.Save(&cert).Error
}

func (s *SupplierService) DeleteCertificate(cert supplier.SupplierCertificate) (err error) {
	return global.GVA_DB.Delete(&cert).Error
}

func (s *SupplierService) GetCertificatesBySupplierID(supplierID uint) (certs []supplier.SupplierCertificate, err error) {
	err = global.GVA_DB.Where("supplier_id = ?", supplierID).Find(&certs).Error
	return
}

func (s *SupplierService) GetExpiringCertificates(months int) (certs []supplier.SupplierCertificate, err error) {
	futureDate := time.Now().AddDate(0, months, 0).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")
	err = global.GVA_DB.Where("valid_end_date <= ? AND valid_end_date >= ?", futureDate, today).Find(&certs).Error
	return
}

// ==================== 供应商审批操作 ====================

func (s *SupplierService) CreateApproval(approval supplier.SupplierApproval) (err error) {
	return global.GVA_DB.Create(&approval).Error
}

func (s *SupplierService) UpdateApproval(approval supplier.SupplierApproval) (err error) {
	return global.GVA_DB.Save(&approval).Error
}

func (s *SupplierService) GetApprovalsBySupplierID(supplierID uint) (approvals []supplier.SupplierApproval, err error) {
	err = global.GVA_DB.Where("supplier_id = ?", supplierID).Order("created_at asc").Find(&approvals).Error
	return
}

func (s *SupplierService) SubmitForApproval(supplierID uint, approver string) (err error) {
	tx := global.GVA_DB.Begin()

	// 更新供应商状态
	if err = tx.Model(&supplier.Supplier{}).Where("id = ?", supplierID).Updates(map[string]interface{}{
		"approval_status": "pending",
		"current_node":    "采购员审批",
	}).Error; err != nil {
		tx.Rollback()
		return
	}

	// 创建审批记录
	approval := supplier.SupplierApproval{
		SupplierID: supplierID,
		NodeName:   "采购员审批",
		Approver:   approver,
		Status:     "current",
	}
	if err = tx.Create(&approval).Error; err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}

func (s *SupplierService) ApproveSupplier(supplierID uint, nodeName string, approver string, status string, comment string, rejectType string) (err error) {
	tx := global.GVA_DB.Begin()

	// 更新当前审批节点
	now := time.Now().Format("2006-01-02 15:04:05")
	if err = tx.Model(&supplier.SupplierApproval{}).Where("supplier_id = ? AND node_name = ? AND status = ?", supplierID, nodeName, "current").Updates(map[string]interface{}{
		"status":        status,
		"comment":       comment,
		"approval_time": now,
	}).Error; err != nil {
		tx.Rollback()
		return
	}

	// 根据审批结果更新供应商状态
	updates := map[string]interface{}{}
	if status == "completed" {
		// 审批通过，进入下一节点或完成
		nextNode := getNextApprovalNode(nodeName)
		if nextNode == "" {
			// 所有节点审批完成
			updates["approval_status"] = "approved"
			updates["status"] = "qualified"
			updates["current_node"] = ""
		} else {
			updates["current_node"] = nextNode
			// 创建下一节点审批记录
			nextApproval := supplier.SupplierApproval{
				SupplierID: supplierID,
				NodeName:   nextNode,
				Status:     "current",
			}
			if err = tx.Create(&nextApproval).Error; err != nil {
				tx.Rollback()
				return
			}
		}
	} else if status == "rejected" {
		// 审批不通过
		updates["approval_status"] = "rejected"
		updates["reject_type"] = rejectType
		updates["reject_reason"] = comment
		if rejectType == "填写问题" {
			// 退回修改，重置状态
			updates["current_node"] = ""
		} else {
			// 资质不合格，终止
			updates["status"] = "blacklist"
			updates["is_blacklist"] = "是"
			updates["blacklist_reason"] = comment
		}
	}

	if err = tx.Model(&supplier.Supplier{}).Where("id = ?", supplierID).Updates(updates).Error; err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}

func getNextApprovalNode(currentNode string) string {
	nodes := []string{"采购员审批", "采购主管审批", "采购经理审批"}
	for i, node := range nodes {
		if node == currentNode && i < len(nodes)-1 {
			return nodes[i+1]
		}
	}
	return ""
}

// ==================== 采购订单操作 ====================

func (s *SupplierService) CreatePurchaseOrder(order supplier.PurchaseOrder, items []supplier.PurchaseOrderItem) (err error) {
	tx := global.GVA_DB.Begin()
	if err = tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return
	}
	for i := range items {
		items[i].OrderID = order.ID
		if err = tx.Create(&items[i]).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	return tx.Commit().Error
}

func (s *SupplierService) UpdatePurchaseOrder(order supplier.PurchaseOrder) (err error) {
	return global.GVA_DB.Save(&order).Error
}

func (s *SupplierService) DeletePurchaseOrder(order supplier.PurchaseOrder) (err error) {
	tx := global.GVA_DB.Begin()
	if err = tx.Where("order_id = ?", order.ID).Delete(&supplier.PurchaseOrderItem{}).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Delete(&order).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

func (s *SupplierService) GetPurchaseOrder(id uint) (order supplier.PurchaseOrder, items []supplier.PurchaseOrderItem, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Where("order_id = ?", id).Find(&items).Error
	return
}

func (s *SupplierService) GetPurchaseOrderList(info request.PageInfo, search supplier.PurchaseOrder) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&supplier.PurchaseOrder{})
	var orders []supplier.PurchaseOrder

	if search.SupplierName != "" {
		db = db.Where("supplier_name LIKE ?", "%"+search.SupplierName+"%")
	}
	if search.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+search.OrderNo+"%")
	}
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&orders).Error
	return orders, total, err
}

func (s *SupplierService) ApprovePurchaseOrder(id uint, status string, remark string) (err error) {
	return global.GVA_DB.Model(&supplier.PurchaseOrder{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":          status,
		"approval_remark": remark,
	}).Error
}

// ==================== 协议价操作 ====================

func (s *SupplierService) CreateAgreementPrice(price supplier.AgreementPrice) (err error) {
	return global.GVA_DB.Create(&price).Error
}

func (s *SupplierService) UpdateAgreementPrice(price supplier.AgreementPrice) (err error) {
	return global.GVA_DB.Save(&price).Error
}

func (s *SupplierService) DeleteAgreementPrice(price supplier.AgreementPrice) (err error) {
	return global.GVA_DB.Delete(&price).Error
}

func (s *SupplierService) GetAgreementPriceList(info request.PageInfo, search supplier.AgreementPrice) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&supplier.AgreementPrice{})
	var prices []supplier.AgreementPrice

	if search.MaterialName != "" {
		db = db.Where("material_name LIKE ?", "%"+search.MaterialName+"%")
	}
	if search.SupplierName != "" {
		db = db.Where("supplier_name LIKE ?", "%"+search.SupplierName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&prices).Error
	return prices, total, err
}
