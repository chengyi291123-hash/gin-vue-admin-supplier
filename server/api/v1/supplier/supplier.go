package supplier

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/supplier"
	supplierService "github.com/flipped-aurora/gin-vue-admin/server/service/supplier"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SupplierApi struct{}

var supService = supplierService.SupplierService{}

// ==================== 供应商基础接口 ====================

func (s *SupplierApi) CreateSupplier(c *gin.Context) {
	var sup supplier.Supplier
	_ = c.ShouldBindJSON(&sup)
	if resSup, err := supService.CreateSupplier(sup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(resSup, c)
	}
}

type SupplierWithCertsRequest struct {
	Supplier     supplier.Supplier              `json:"supplier"`
	Certificates []supplier.SupplierCertificate `json:"certificates"`
}

func (s *SupplierApi) CreateSupplierWithCerts(c *gin.Context) {
	var req SupplierWithCertsRequest
	_ = c.ShouldBindJSON(&req)
	if resSup, err := supService.CreateSupplierWithCerts(req.Supplier, req.Certificates); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(gin.H{"supplier": resSup}, c)
	}
}

func (s *SupplierApi) DeleteSupplier(c *gin.Context) {
	var sup supplier.Supplier
	_ = c.ShouldBindJSON(&sup)
	if err := supService.DeleteSupplier(sup); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (s *SupplierApi) UpdateSupplier(c *gin.Context) {
	var sup supplier.Supplier
	_ = c.ShouldBindJSON(&sup)
	// 从 URL 路径获取 ID
	idStr := c.Param("id")
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		sup.ID = uint(id)
	}
	if sup.ID == 0 {
		global.GVA_LOG.Error("更新失败: ID 为空")
		response.FailWithMessage("更新失败: 缺少供应商ID", c)
		return
	}
	// 获取变更人（从请求中获取或使用默认值）
	changeBy := c.GetHeader("X-Change-By")
	if changeBy == "" {
		changeBy = "系统"
	}
	if err := supService.UpdateSupplier(sup, changeBy); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (s *SupplierApi) FindSupplier(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if resSup, err := supService.GetSupplier(uint(id)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"supplier": resSup}, c)
	}
}

func (s *SupplierApi) FindSupplierWithCerts(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if sup, certs, err := supService.GetSupplierWithCerts(uint(id)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"supplier": sup, "certificates": certs}, c)
	}
}

func (s *SupplierApi) GetSupplierList(c *gin.Context) {
	var pageInfo request.PageInfo
	var search supplier.Supplier
	_ = c.ShouldBindQuery(&pageInfo)
	_ = c.ShouldBindQuery(&search)

	if list, total, err := supService.GetSupplierInfoList(pageInfo, search); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// ==================== 供应商证书接口 ====================

func (s *SupplierApi) CreateCertificate(c *gin.Context) {
	var cert supplier.SupplierCertificate
	_ = c.ShouldBindJSON(&cert)
	if err := supService.CreateCertificate(cert); err != nil {
		global.GVA_LOG.Error("创建证书失败!", zap.Error(err))
		response.FailWithMessage("创建证书失败", c)
	} else {
		response.OkWithMessage("创建证书成功", c)
	}
}

func (s *SupplierApi) UpdateCertificate(c *gin.Context) {
	var cert supplier.SupplierCertificate
	_ = c.ShouldBindJSON(&cert)
	if err := supService.UpdateCertificate(cert); err != nil {
		global.GVA_LOG.Error("更新证书失败!", zap.Error(err))
		response.FailWithMessage("更新证书失败", c)
	} else {
		response.OkWithMessage("更新证书成功", c)
	}
}

func (s *SupplierApi) DeleteCertificate(c *gin.Context) {
	var cert supplier.SupplierCertificate
	_ = c.ShouldBindJSON(&cert)
	if err := supService.DeleteCertificate(cert); err != nil {
		global.GVA_LOG.Error("删除证书失败!", zap.Error(err))
		response.FailWithMessage("删除证书失败", c)
	} else {
		response.OkWithMessage("删除证书成功", c)
	}
}

func (s *SupplierApi) GetCertificatesBySupplierID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if certs, err := supService.GetCertificatesBySupplierID(uint(id)); err != nil {
		global.GVA_LOG.Error("获取证书失败!", zap.Error(err))
		response.FailWithMessage("获取证书失败", c)
	} else {
		response.OkWithData(gin.H{"certificates": certs}, c)
	}
}

func (s *SupplierApi) GetExpiringCertificates(c *gin.Context) {
	monthsStr := c.DefaultQuery("months", "6")
	months, _ := strconv.Atoi(monthsStr)
	if certs, err := supService.GetExpiringCertificates(months); err != nil {
		global.GVA_LOG.Error("获取即将过期证书失败!", zap.Error(err))
		response.FailWithMessage("获取即将过期证书失败", c)
	} else {
		response.OkWithData(gin.H{"certificates": certs}, c)
	}
}

// ==================== 供应商审批接口 ====================

func (s *SupplierApi) GetApprovalsBySupplierID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if approvals, err := supService.GetApprovalsBySupplierID(uint(id)); err != nil {
		global.GVA_LOG.Error("获取审批记录失败!", zap.Error(err))
		response.FailWithMessage("获取审批记录失败", c)
	} else {
		response.OkWithData(gin.H{"approvals": approvals}, c)
	}
}

type SubmitApprovalRequest struct {
	SupplierID uint   `json:"supplier_id"`
	Approver   string `json:"approver"`
}

func (s *SupplierApi) SubmitForApproval(c *gin.Context) {
	var req SubmitApprovalRequest
	_ = c.ShouldBindJSON(&req)
	if err := supService.SubmitForApproval(req.SupplierID, req.Approver); err != nil {
		global.GVA_LOG.Error("提交审批失败!", zap.Error(err))
		response.FailWithMessage("提交审批失败", c)
	} else {
		response.OkWithMessage("提交审批成功", c)
	}
}

type ApproveRequest struct {
	SupplierID uint   `json:"supplier_id"`
	NodeName   string `json:"node_name"`
	Approver   string `json:"approver"`
	Status     string `json:"status"` // completed, rejected
	Comment    string `json:"comment"`
	RejectType string `json:"reject_type"` // 填写问题, 资质不合格
}

func (s *SupplierApi) ApproveSupplier(c *gin.Context) {
	var req ApproveRequest
	_ = c.ShouldBindJSON(&req)
	if err := supService.ApproveSupplier(req.SupplierID, req.NodeName, req.Approver, req.Status, req.Comment, req.RejectType); err != nil {
		global.GVA_LOG.Error("审批失败!", zap.Error(err))
		response.FailWithMessage("审批失败", c)
	} else {
		response.OkWithMessage("审批成功", c)
	}
}

// ==================== 采购订单接口 ====================

type PurchaseOrderRequest struct {
	Order supplier.PurchaseOrder       `json:"order"`
	Items []supplier.PurchaseOrderItem `json:"items"`
}

func (s *SupplierApi) CreatePurchaseOrder(c *gin.Context) {
	var req PurchaseOrderRequest
	_ = c.ShouldBindJSON(&req)
	if err := supService.CreatePurchaseOrder(req.Order, req.Items); err != nil {
		global.GVA_LOG.Error("创建采购订单失败!", zap.Error(err))
		response.FailWithMessage("创建采购订单失败", c)
	} else {
		response.OkWithMessage("创建采购订单成功", c)
	}
}

func (s *SupplierApi) UpdatePurchaseOrder(c *gin.Context) {
	var order supplier.PurchaseOrder
	_ = c.ShouldBindJSON(&order)
	if err := supService.UpdatePurchaseOrder(order); err != nil {
		global.GVA_LOG.Error("更新采购订单失败!", zap.Error(err))
		response.FailWithMessage("更新采购订单失败", c)
	} else {
		response.OkWithMessage("更新采购订单成功", c)
	}
}

func (s *SupplierApi) DeletePurchaseOrder(c *gin.Context) {
	var order supplier.PurchaseOrder
	_ = c.ShouldBindJSON(&order)
	if err := supService.DeletePurchaseOrder(order); err != nil {
		global.GVA_LOG.Error("删除采购订单失败!", zap.Error(err))
		response.FailWithMessage("删除采购订单失败", c)
	} else {
		response.OkWithMessage("删除采购订单成功", c)
	}
}

func (s *SupplierApi) GetPurchaseOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if order, items, err := supService.GetPurchaseOrder(uint(id)); err != nil {
		global.GVA_LOG.Error("获取采购订单失败!", zap.Error(err))
		response.FailWithMessage("获取采购订单失败", c)
	} else {
		response.OkWithData(gin.H{"order": order, "items": items}, c)
	}
}

func (s *SupplierApi) GetPurchaseOrderList(c *gin.Context) {
	var pageInfo request.PageInfo
	var search supplier.PurchaseOrder
	_ = c.ShouldBindQuery(&pageInfo)
	_ = c.ShouldBindQuery(&search)

	if list, total, err := supService.GetPurchaseOrderList(pageInfo, search); err != nil {
		global.GVA_LOG.Error("获取采购订单列表失败!", zap.Error(err))
		response.FailWithMessage("获取采购订单列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

type ApprovePurchaseOrderRequest struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
	Remark string `json:"remark"`
}

func (s *SupplierApi) ApprovePurchaseOrder(c *gin.Context) {
	var req ApprovePurchaseOrderRequest
	_ = c.ShouldBindJSON(&req)
	if err := supService.ApprovePurchaseOrder(req.ID, req.Status, req.Remark); err != nil {
		global.GVA_LOG.Error("审批采购订单失败!", zap.Error(err))
		response.FailWithMessage("审批采购订单失败", c)
	} else {
		response.OkWithMessage("审批成功", c)
	}
}

// ==================== 协议价接口 ====================

func (s *SupplierApi) CreateAgreementPrice(c *gin.Context) {
	var price supplier.AgreementPrice
	_ = c.ShouldBindJSON(&price)
	if err := supService.CreateAgreementPrice(price); err != nil {
		global.GVA_LOG.Error("创建协议价失败!", zap.Error(err))
		response.FailWithMessage("创建协议价失败", c)
	} else {
		response.OkWithMessage("创建协议价成功", c)
	}
}

func (s *SupplierApi) UpdateAgreementPrice(c *gin.Context) {
	var price supplier.AgreementPrice
	_ = c.ShouldBindJSON(&price)
	if err := supService.UpdateAgreementPrice(price); err != nil {
		global.GVA_LOG.Error("更新协议价失败!", zap.Error(err))
		response.FailWithMessage("更新协议价失败", c)
	} else {
		response.OkWithMessage("更新协议价成功", c)
	}
}

func (s *SupplierApi) DeleteAgreementPrice(c *gin.Context) {
	var price supplier.AgreementPrice
	_ = c.ShouldBindJSON(&price)
	if err := supService.DeleteAgreementPrice(price); err != nil {
		global.GVA_LOG.Error("删除协议价失败!", zap.Error(err))
		response.FailWithMessage("删除协议价失败", c)
	} else {
		response.OkWithMessage("删除协议价成功", c)
	}
}

func (s *SupplierApi) GetAgreementPriceList(c *gin.Context) {
	var pageInfo request.PageInfo
	var search supplier.AgreementPrice
	_ = c.ShouldBindQuery(&pageInfo)
	_ = c.ShouldBindQuery(&search)

	if list, total, err := supService.GetAgreementPriceList(pageInfo, search); err != nil {
		global.GVA_LOG.Error("获取协议价列表失败!", zap.Error(err))
		response.FailWithMessage("获取协议价列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// ==================== 供应商变更记录接口 ====================

func (s *SupplierApi) GetChangeLogsBySupplierID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if logs, err := supService.GetChangeLogsBySupplierID(uint(id)); err != nil {
		global.GVA_LOG.Error("获取变更记录失败!", zap.Error(err))
		response.FailWithMessage("获取变更记录失败", c)
	} else {
		response.OkWithData(gin.H{"list": logs}, c)
	}
}
