import service from '@/utils/request'

// ==================== 供应商基础接口 ====================

// 创建供应商
export const createSupplier = (data) => {
  return service({
    url: '/suppliers',
    method: 'post',
    data
  })
}

// 创建供应商(含证书)
export const createSupplierWithCerts = (data) => {
  return service({
    url: '/suppliers/with-certs',
    method: 'post',
    data
  })
}

// 更新供应商
export const updateSupplier = (id, data) => {
  return service({
    url: `/suppliers/${id}`,
    method: 'put',
    data
  })
}

// 删除供应商
export const deleteSupplier = (data) => {
  return service({
    url: '/suppliers',
    method: 'delete',
    data
  })
}

// 获取供应商列表
export const getSupplierList = (params) => {
  return service({
    url: '/suppliers',
    method: 'get',
    params
  })
}

// 获取单个供应商
export const getSupplier = (id) => {
  return service({
    url: `/suppliers/${id}`,
    method: 'get'
  })
}

// 获取供应商(含证书)
export const getSupplierWithCerts = (id) => {
  return service({
    url: `/suppliers/${id}/with-certs`,
    method: 'get'
  })
}

// ==================== 供应商证书接口 ====================

// 创建证书
export const createCertificate = (data) => {
  return service({
    url: '/supplier-certificates',
    method: 'post',
    data
  })
}

// 更新证书
export const updateCertificate = (data) => {
  return service({
    url: '/supplier-certificates',
    method: 'put',
    data
  })
}

// 删除证书
export const deleteCertificate = (data) => {
  return service({
    url: '/supplier-certificates',
    method: 'delete',
    data
  })
}

// 获取供应商证书
export const getCertificatesBySupplierID = (supplierId) => {
  return service({
    url: `/supplier-certificates/supplier/${supplierId}`,
    method: 'get'
  })
}

// 获取即将过期证书
export const getExpiringCertificates = (months = 6) => {
  return service({
    url: '/supplier-certificates/expiring',
    method: 'get',
    params: { months }
  })
}

// ==================== 供应商审批接口 ====================

// 提交审批
export const submitForApproval = (data) => {
  return service({
    url: '/supplier-approvals/submit',
    method: 'post',
    data
  })
}

// 审批供应商
export const approveSupplier = (data) => {
  return service({
    url: '/supplier-approvals/approve',
    method: 'post',
    data
  })
}

// 获取审批记录
export const getApprovalsBySupplierID = (supplierId) => {
  return service({
    url: `/supplier-approvals/supplier/${supplierId}`,
    method: 'get'
  })
}

// ==================== 采购订单接口 ====================

// 创建采购订单
export const createPurchaseOrder = (data) => {
  return service({
    url: '/purchase-orders',
    method: 'post',
    data
  })
}

// 更新采购订单
export const updatePurchaseOrder = (data) => {
  return service({
    url: '/purchase-orders',
    method: 'put',
    data
  })
}

// 删除采购订单
export const deletePurchaseOrder = (data) => {
  return service({
    url: '/purchase-orders',
    method: 'delete',
    data
  })
}

// 获取采购订单列表
export const getPurchaseOrderList = (params) => {
  return service({
    url: '/purchase-orders',
    method: 'get',
    params
  })
}

// 获取采购订单详情
export const getPurchaseOrder = (id) => {
  return service({
    url: `/purchase-orders/${id}`,
    method: 'get'
  })
}

// 审批采购订单
export const approvePurchaseOrder = (data) => {
  return service({
    url: '/purchase-orders/approve',
    method: 'post',
    data
  })
}

// ==================== 协议价接口 ====================

// 创建协议价
export const createAgreementPrice = (data) => {
  return service({
    url: '/agreement-prices',
    method: 'post',
    data
  })
}

// 更新协议价
export const updateAgreementPrice = (data) => {
  return service({
    url: '/agreement-prices',
    method: 'put',
    data
  })
}

// 删除协议价
export const deleteAgreementPrice = (data) => {
  return service({
    url: '/agreement-prices',
    method: 'delete',
    data
  })
}

// 获取协议价列表
export const getAgreementPriceList = (params) => {
  return service({
    url: '/agreement-prices',
    method: 'get',
    params
  })
}

// ==================== 供应商变更记录接口 ====================

// 获取供应商变更记录
export const getSupplierChangeLogs = (supplierId) => {
  return service({
    url: `/suppliers/${supplierId}/change-logs`,
    method: 'get'
  })
}
