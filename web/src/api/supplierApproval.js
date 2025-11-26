import service from '@/utils/request'

// 获取供应商审批列表
export const getSupplierApprovalList = (params) => {
  return service({
    url: '/api/supplier/approval/list',
    method: 'get',
    params
  })
}

// 获取供应商详情
export const getSupplierDetail = (id) => {
  return service({
    url: `/api/supplier/approval/detail/${id}`,
    method: 'get'
  })
}

// 审批通过
export const approveSupplier = (data) => {
  return service({
    url: '/api/supplier/approval/approve',
    method: 'post',
    data
  })
}

// 审批驳回
export const rejectSupplier = (data) => {
  return service({
    url: '/api/supplier/approval/reject',
    method: 'post',
    data
  })
}

// 批量审批
export const batchApproval = (data) => {
  return service({
    url: '/api/supplier/approval/batch',
    method: 'post',
    data
  })
}

// 获取供应商附件
export const getSupplierAttachments = (supplierId) => {
  return service({
    url: `/api/supplier/approval/attachments/${supplierId}`,
    method: 'get'
  })
}

// 下载附件
export const downloadAttachment = (id) => {
  return service({
    url: `/api/supplier/approval/attachment/download/${id}`,
    method: 'get',
    responseType: 'blob'
  })
}

// 获取审批历史
export const getApprovalHistory = (supplierId) => {
  return service({
    url: `/api/supplier/approval/history/${supplierId}`,
    method: 'get'
  })
}