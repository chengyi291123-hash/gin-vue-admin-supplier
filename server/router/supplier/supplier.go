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
	{
		supplierRouter.POST("", supApi.CreateSupplier)   // 新建
		supplierRouter.PUT("/:id", supApi.UpdateSupplier) // 更新
		supplierRouter.DELETE("", supApi.DeleteSupplier) // 删除
	}
	{
		supplierRouterWithoutRecord.GET("", supApi.GetSupplierList) // 获取列表
		supplierRouterWithoutRecord.GET("/:id", supApi.FindSupplier) // 获取单条
	}
}
