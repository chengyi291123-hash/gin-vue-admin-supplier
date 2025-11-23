package supplier

import (
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

func (s *SupplierApi) CreateSupplier(c *gin.Context) {
	var sup supplier.Supplier
	_ = c.ShouldBindJSON(&sup)
	if err := supService.CreateSupplier(sup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
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
	if err := supService.UpdateSupplier(sup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (s *SupplierApi) FindSupplier(c *gin.Context) {
	var sup supplier.Supplier
	_ = c.ShouldBindQuery(&sup)
	if resSup, err := supService.GetSupplier(sup.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"ressup": resSup}, c)
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
