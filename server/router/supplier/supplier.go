package supplier

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/supplier"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SupplierRouter struct{}

func (s *SupplierRouter) InitSupplierRouter(Router *gin.RouterGroup) {
	supplierRouter := Router.Group("suppliers").Use(middleware.OperationRecord())
	supplierRouterWithoutRecord := Router.Group("suppliers")
	var supApi = supplier.SupplierApi{}

	// 供应商基础操作
	{
		supplierRouter.POST("", supApi.CreateSupplier)                      // 新建供应商
		supplierRouter.POST("/with-certs", supApi.CreateSupplierWithCerts)  // 新建供应商(含证书)
		supplierRouter.PUT("/:id", supApi.UpdateSupplier)                   // 更新供应商
		supplierRouter.DELETE("", supApi.DeleteSupplier)                    // 删除供应商
	}
	{
		supplierRouterWithoutRecord.GET("", supApi.GetSupplierList)                // 获取供应商列表
		supplierRouterWithoutRecord.GET("/:id", supApi.FindSupplier)               // 获取单个供应商
		supplierRouterWithoutRecord.GET("/:id/with-certs", supApi.FindSupplierWithCerts) // 获取供应商(含证书)
		supplierRouterWithoutRecord.GET("/:id/change-logs", supApi.GetChangeLogsBySupplierID) // 获取供应商变更记录
	}

	// 供应商证书操作
	certRouter := Router.Group("supplier-certificates").Use(middleware.OperationRecord())
	certRouterWithoutRecord := Router.Group("supplier-certificates")
	{
		certRouter.POST("", supApi.CreateCertificate)    // 创建证书
		certRouter.PUT("", supApi.UpdateCertificate)     // 更新证书
		certRouter.DELETE("", supApi.DeleteCertificate)  // 删除证书
	}
	{
		certRouterWithoutRecord.GET("/supplier/:id", supApi.GetCertificatesBySupplierID) // 获取供应商证书
		certRouterWithoutRecord.GET("/expiring", supApi.GetExpiringCertificates)         // 获取即将过期证书
	}

	// 供应商审批操作
	approvalRouter := Router.Group("supplier-approvals").Use(middleware.OperationRecord())
	approvalRouterWithoutRecord := Router.Group("supplier-approvals")
	{
		approvalRouter.POST("/submit", supApi.SubmitForApproval)  // 提交审批
		approvalRouter.POST("/approve", supApi.ApproveSupplier)   // 审批
	}
	{
		approvalRouterWithoutRecord.GET("/supplier/:id", supApi.GetApprovalsBySupplierID) // 获取审批记录
	}

	// 采购订单操作
	orderRouter := Router.Group("purchase-orders").Use(middleware.OperationRecord())
	orderRouterWithoutRecord := Router.Group("purchase-orders")
	{
		orderRouter.POST("", supApi.CreatePurchaseOrder)          // 创建采购订单
		orderRouter.PUT("", supApi.UpdatePurchaseOrder)           // 更新采购订单
		orderRouter.DELETE("", supApi.DeletePurchaseOrder)        // 删除采购订单
		orderRouter.POST("/approve", supApi.ApprovePurchaseOrder) // 审批采购订单
	}
	{
		orderRouterWithoutRecord.GET("", supApi.GetPurchaseOrderList)   // 获取采购订单列表
		orderRouterWithoutRecord.GET("/:id", supApi.GetPurchaseOrder)   // 获取采购订单详情
	}

	// 协议价操作
	priceRouter := Router.Group("agreement-prices").Use(middleware.OperationRecord())
	priceRouterWithoutRecord := Router.Group("agreement-prices")
	{
		priceRouter.POST("", supApi.CreateAgreementPrice)    // 创建协议价
		priceRouter.PUT("", supApi.UpdateAgreementPrice)     // 更新协议价
		priceRouter.DELETE("", supApi.DeleteAgreementPrice)  // 删除协议价
	}
	{
		priceRouterWithoutRecord.GET("", supApi.GetAgreementPriceList) // 获取协议价列表
	}
}
