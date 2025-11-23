package supplier

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/supplier"
)

type SupplierService struct{}

func (s *SupplierService) CreateSupplier(sup supplier.Supplier) (err error) {
	err = global.GVA_DB.Create(&sup).Error
	return err
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

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sups).Error
	return sups, total, err
}
