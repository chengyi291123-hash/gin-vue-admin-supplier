<template>
  <div class="supplier-container">
    <!-- Sidebar -->
    <div class="aside">
      <div class="logo">
        <span>叭盛采购管理平台</span>
      </div>
      <el-menu
        :default-openeds="['1', '2', '3', '4']"
        :default-active="currentMenu"
        background-color="#2b3649"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
        @select="handleSelect"
      >
        <!-- 供应商管理 -->
        <el-sub-menu index="1">
          <template #title>
            <el-icon><Menu /></el-icon>
            <span>供应商管理</span>
          </template>
          <el-menu-item index="apply">供应商准入申请</el-menu-item>
          <el-menu-item index="temp">潜在供应商</el-menu-item>
          <el-menu-item index="qualified">合格供应商</el-menu-item>
          <el-menu-item index="change">供应商变更</el-menu-item>
          <el-menu-item index="supplier-approval">供应商审批</el-menu-item>
          <el-menu-item index="supplier-list">供应商档案库</el-menu-item>
        </el-sub-menu>
        <!-- 询价管理 -->
        <el-sub-menu index="2">
          <template #title>
            <el-icon><Document /></el-icon>
            <span>询价管理</span>
          </template>
          <el-menu-item index="my-inquiry">我的询价（销售）</el-menu-item>
          <el-menu-item index="pending-quote">待报价项目（采购）</el-menu-item>
          <el-menu-item index="inquiry-ledger">全部询报价台账</el-menu-item>
          <el-menu-item index="bid-projects">中标项目管理</el-menu-item>
        </el-sub-menu>
        <!-- 采购审批 -->
        <el-sub-menu index="3">
          <template #title>
            <el-icon><Check /></el-icon>
            <span>采购审批</span>
          </template>
          <el-menu-item index="payment-approval">吧盛支付申请</el-menu-item>
          <el-menu-item index="order-approval">采购订单审批</el-menu-item>
        </el-sub-menu>
        <!-- 协议价管理（仅管理员） -->
        <el-sub-menu index="4" v-if="isAdmin">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>协议价管理</span>
          </template>
          <el-menu-item index="agreement-price">协议价管理</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </div>

    <!-- Main Layout -->
    <div style="flex: 1; display: flex; flex-direction: column; overflow: hidden;">
      <!-- Header -->
      <div class="header">
        <div class="breadcrumb" style="display: flex; align-items: center; gap: 15px;">
          <span>首页 / 供应商管理 / {{ getCurrentTitle() }}</span>
          <el-tag v-if="isApprovalMode" type="warning" size="small">审批模式</el-tag>
          <!-- 潜在供应商/合格供应商切换按钮 -->
          <el-radio-group v-if="(currentMenu === 'apply' || currentMenu === 'supplier-apply') && !isApprovalMode" v-model="supplierApplyType" size="small">
            <el-radio-button value="temp">潜在供应商</el-radio-button>
            <el-radio-button value="qualified">合格供应商</el-radio-button>
          </el-radio-group>
        </div>
        <div style="display: flex; align-items: center;">
          <template v-if="isApprovalMode && (currentMenu === 'apply' || currentMenu === 'supplier-apply')">
            <el-button size="small" @click="backToApprovalList">返回列表</el-button>
            <el-button type="success" size="small" @click="approveCurrentSupplier">审批通过</el-button>
            <el-button type="danger" size="small" @click="rejectCurrentSupplier" style="margin-right: 15px;">审批驳回</el-button>
          </template>
          <template v-else-if="currentMenu === 'apply' || currentMenu === 'supplier-apply'">
            <el-button v-if="supplierApplyType === 'temp'" type="success" size="small" @click="submitTempSupplier" style="margin-right: 15px;">保存潜在供应商</el-button>
            <el-button v-else type="primary" size="small" @click="submitForApproval" style="margin-right: 15px;">提交审批</el-button>
          </template>
          <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
        </div>
      </div>

      <!-- Content -->
      <div class="main">
        <div class="content-card">
          <!-- VIEW 1: 供应商准入申请 -->
          <div v-if="currentMenu === 'apply' || currentMenu === 'supplier-apply'">
            <el-alert title="提示：潜在供方只需填写红色必填项，直接提交；申请合格供方需填写全部信息并审批。" type="info" show-icon style="margin-bottom: 20px;"></el-alert>

            <el-tabs v-model="activeTab">
              <el-tab-pane label="基本信息" name="basic">
                <div class="form-section-title">基本信息</div>
                <el-form :model="form" label-width="150px" style="max-width: 900px">
                  <el-form-item label="企业名称">
                    <el-input v-model="form.enterprise_name" placeholder="请输入企业全称"></el-input>
                  </el-form-item>
                  <el-form-item label="统一社会信用代码" v-if="!isTempAdd">
                    <el-input v-model="form.credit_code"></el-input>
                  </el-form-item>
                  <el-row>
                    <el-col :span="12">
                      <el-form-item label="供应商类型">
                        <template #label><span class="red-star">*</span>供应商类型</template>
                        <el-select v-model="form.entry_type" placeholder="请选择" style="width: 100%">
                          <el-option label="生产商" value="manufacturing"></el-option>
                          <el-option label="贸易商" value="trading"></el-option>
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="供应品类">
                        <template #label><span class="red-star">*</span>供应品类</template>
                        <el-select v-model="form.category" placeholder="请选择分类" style="width: 100%">
                          <el-option label="原材料" value="raw"></el-option>
                          <el-option label="电子元器件" value="electronic"></el-option>
                          <el-option label="办公用品" value="office"></el-option>
                        </el-select>
                      </el-form-item>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="12" v-if="!isTempAdd">
                      <el-form-item label="合作区域">
                        <el-input v-model="form.region"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="合作行业">
                        <template #label><span class="red-star">*</span>合作行业</template>
                        <el-select v-model="form.industry" placeholder="请选择" style="width: 100%" @change="handleIndustryChange">
                          <el-option label="普通行业" value="normal"></el-option>
                          <el-option label="核电" value="nuclear"></el-option>
                          <el-option label="军工" value="military"></el-option>
                          <el-option label="石化" value="petrochemical"></el-option>
                          <el-option label="其他" value="other"></el-option>
                        </el-select>
                      </el-form-item>
                    </el-col>
                  </el-row>
                  <el-form-item label="合作品牌">
                    <template #label><span class="red-star">*</span>合作品牌</template>
                    <el-input v-model="form.brand" placeholder="请输入品牌"></el-input>
                  </el-form-item>

                  <!-- 核电/军工特殊字段 -->
                  <div v-if="form.industry === 'nuclear' || form.industry === 'military'">
                    <el-divider content-position="left">
                      <el-tag type="warning" size="small">特殊行业要求</el-tag>
                    </el-divider>
                    <el-form-item label="质保体系认证" required>
                      <el-upload
                        class="upload-demo"
                        action="#"
                        :auto-upload="false"
                        :on-change="(file) => handleSpecialFile(file, 'quality_cert')"
                        :file-list="specialFiles.quality_cert"
                        :limit="5"
                        multiple
                      >
                        <el-button type="primary" size="small">点击上传</el-button>
                        <template #tip>
                          <div class="el-upload__tip">支持ISO9001等质量体系认证文件，最大5MB</div>
                        </template>
                      </el-upload>
                    </el-form-item>
                    <el-form-item label="安全保密资质" required>
                      <el-upload
                        class="upload-demo"
                        action="#"
                        :auto-upload="false"
                        :on-change="(file) => handleSpecialFile(file, 'security_cert')"
                        :file-list="specialFiles.security_cert"
                        :limit="5"
                        multiple
                      >
                        <el-button type="primary" size="small">点击上传</el-button>
                        <template #tip>
                          <div class="el-upload__tip">支持保密资质、安全许可等文件，最大5MB</div>
                        </template>
                      </el-upload>
                    </el-form-item>
                    <el-form-item label="其他资质">
                      <el-upload
                        class="upload-demo"
                        action="#"
                        :auto-upload="false"
                        :on-change="(file) => handleSpecialFile(file, 'other_cert')"
                        :file-list="specialFiles.other_cert"
                        :limit="5"
                        multiple
                      >
                        <el-button type="primary" size="small">点击上传</el-button>
                        <template #tip>
                          <div class="el-upload__tip">其他业务相关资质文件，最大5MB</div>
                        </template>
                      </el-upload>
                    </el-form-item>
                  </div>

                  <el-divider></el-divider>
                  <div class="form-section-title">联系人与采购</div>
                  <el-row>
                    <el-col :span="12">
                      <el-form-item label="联系人">
                        <template #label><span class="red-star">*</span>联系人</template>
                        <el-input v-model="form.contact_person"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="联系电话">
                        <template #label><span class="red-star">*</span>联系电话</template>
                        <el-input v-model="form.mobile"></el-input>
                      </el-form-item>
                    </el-col>
                  </el-row>
                  <el-form-item label="采购员" v-if="!isTempAdd">
                    <template #label><span class="red-star">*</span>采购员</template>
                    <el-input v-model="form.purchaser" placeholder="默认当前登录用户"></el-input>
                  </el-form-item>

                  <div v-if="!isTempAdd">
                    <el-divider></el-divider>
                    <div class="form-section-title">财务账户信息</div>
                    <el-form-item label="开户行">
                      <el-input v-model="form.bank_name"></el-input>
                    </el-form-item>
                    <el-form-item label="开户行名称 (支行)">
                      <el-input v-model="form.branch_name"></el-input>
                    </el-form-item>
                    <el-form-item label="银行账号">
                      <el-input v-model="form.bank_account"></el-input>
                    </el-form-item>
                    <el-form-item label="结算方式">
                      <el-select v-model="form.settlement" placeholder="请选择" style="width: 100%">
                        <el-option label="月结30天" value="m30"></el-option>
                        <el-option label="预付" value="prepay"></el-option>
                        <el-option label="货到付款" value="cod"></el-option>
                      </el-select>
                    </el-form-item>
                  </div>

                  <el-form-item>
                    <el-button>上一步</el-button>
                    <el-button type="primary" @click="nextTab('qual')">下一步：资质证书</el-button>
                  </el-form-item>
                </el-form>
              </el-tab-pane>

              <el-tab-pane label="资质证书" name="qual">
                <div class="form-section-title">资质文件上传</div>

                <el-form label-width="150px">
                  <el-form-item label="供应商类型">
                    <el-radio-group v-model="form.entry_type">
                      <el-radio-button label="manufacturing">生产商</el-radio-button>
                      <el-radio-button label="trading">贸易商</el-radio-button>
                    </el-radio-group>
                  </el-form-item>
                  
                  <div style="margin-left: 50px; margin-right: 50px;">
                    <el-table :data="certificateList" border style="width: 100%">
                        <el-table-column label="证书名称" min-width="200">
                            <template #default="scope">
                                <span v-if="isCertExpiringSoon(scope.row.validEndDate)" style="color: #f56c6c; font-weight: bold;">
                                    {{ scope.row.name }}
                                    <el-tooltip effect="dark" placement="top">
                                        <template #content>
                                            <div style="max-width: 200px;">
                                                <div style="font-weight: bold; margin-bottom: 5px;">证书即将到期</div>
                                                <div>该证书有效期不足6个月，请及时更新！</div>
                                                <div v-if="scope.row.validEndDate" style="margin-top: 5px;">
                                                    到期日期：{{ scope.row.validEndDate }}
                                                </div>
                                            </div>
                                        </template>
                                        <el-icon style="margin-left: 5px; cursor: pointer;"><WarningFilled /></el-icon>
                                    </el-tooltip>
                                </span>
                                <span v-else>{{ scope.row.name }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column label="已上传份数" width="100" align="center">
                            <template #default="scope">
                                {{ getFileCount(scope.row.name) }} 份
                            </template>
                        </el-table-column>
                        <el-table-column label="有效期" width="320" align="center">
                            <template #default="scope">
                                <div style="display: flex; align-items: center; gap: 5px;">
                                    <el-date-picker
                                        v-model="scope.row.validStartDate"
                                        type="date"
                                        placeholder="开始日期"
                                        format="YYYY-MM-DD"
                                        value-format="YYYY-MM-DD"
                                        size="small"
                                        style="width: 120px;"
                                    />
                                    <span>至</span>
                                    <el-date-picker
                                        v-model="scope.row.validEndDate"
                                        type="date"
                                        placeholder="结束日期"
                                        format="YYYY-MM-DD"
                                        value-format="YYYY-MM-DD"
                                        size="small"
                                        style="width: 120px;"
                                    />
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" width="150">
                            <template #default="scope">
                                <div v-if="getFileForCert(scope.row.name)" style="display: flex; align-items: center;">
                                    <el-icon style="margin-right: 5px;"><document /></el-icon>
                                    <span style="margin-right: 10px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; max-width: 80px;">{{ getFileForCert(scope.row.name).name }}</span>
                                    <el-button type="danger" size="small" link icon="Delete" @click="removeFile(scope.row.name)"></el-button>
                                </div>
                                <el-upload
                                    v-else
                                    class="upload-demo"
                                    action="#"
                                    :auto-upload="false"
                                    :show-file-list="false"
                                    :on-change="(file) => handleFileChange(file, scope.row.name)"
                                >
                                    <el-button type="primary" size="small" icon="Upload">点击上传</el-button>
                                </el-upload>
                            </template>
                        </el-table-column>
                    </el-table>
                  </div>

                  <el-form-item style="margin-top: 20px;">
                    <el-button @click="activeTab = 'basic'">上一步</el-button>
                    <el-button type="success" @click="saveData">提交申请</el-button>
                  </el-form-item>
                </el-form>
              </el-tab-pane>
            </el-tabs>
          </div>

          <!-- VIEW: 供应商审批列表 -->
          <div v-else-if="currentMenu === 'supplier-approval'">
            <div class="form-section-title">供应商审批列表</div>

            <!-- 搜索筛选 -->
            <div style="margin-bottom: 20px;">
              <el-form :inline="true" :model="searchForm" class="demo-form-inline">
                <el-form-item label="企业名称">
                  <el-input v-model="searchForm.name" placeholder="请输入名称" prefix-icon="Search" clearable></el-input>
                </el-form-item>
                <el-form-item label="审批状态">
                  <el-select v-model="searchForm.approvalStatus" placeholder="全部" style="width: 120px;" clearable>
                    <el-option label="待审批" value="待审批"></el-option>
                    <el-option label="已通过" value="已通过"></el-option>
                    <el-option label="已驳回" value="已驳回"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item label="行业">
                  <el-select v-model="searchForm.industry" placeholder="全部" style="width: 100px;" clearable>
                    <el-option label="核电" value="nuclear"></el-option>
                    <el-option label="军工" value="military"></el-option>
                    <el-option label="普通" value="normal"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="onSearch">筛选</el-button>
                  <el-button @click="resetApprovalSearch">重置</el-button>
                </el-form-item>
              </el-form>
            </div>

            <!-- 审批列表表格 -->
            <el-table :data="approvalData" border style="width: 100%">
              <el-table-column prop="code" label="供应商编号" width="150" />
              <el-table-column prop="name" label="供应商名称" min-width="180" />
              <el-table-column prop="type" label="供应商类型" width="120" />
              <el-table-column prop="industry" label="行业" width="80">
                <template #default="{ row }">
                  <el-tag :type="row.industry === 'nuclear' || row.industry === 'military' ? 'danger' : 'info'" size="small">
                    {{ row.industry === 'nuclear' ? '核电' : row.industry === 'military' ? '军工' : '普通' }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="contact_person" label="联系人" width="100" />
              <el-table-column prop="status" label="状态" width="90">
                <template #default="{ row }">
                  <el-tag :type="getApprovalStatusType(row.status)">{{ row.status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="current_node" label="当前节点" width="120" />
              <el-table-column prop="created_time" label="申请时间" width="160" />
              <el-table-column label="操作" width="180" fixed="right">
                <template #default="{ row }">
                  <el-button type="primary" size="small" @click="goToApprovalDetail(row)">查看详情</el-button>
                  <el-button v-if="row.status === '待审批'" type="success" size="small" @click="quickApprove(row)">快速审批</el-button>
                </template>
              </el-table-column>
            </el-table>

            <div style="margin-top: 20px; text-align: right;">
              <el-pagination background layout="prev, pager, next, sizes, total" :total="approvalData.length"></el-pagination>
            </div>

            <!-- 快速审批对话框 -->
            <el-dialog v-model="showQuickApprovalDialog" title="供应商审批" width="550px" :close-on-click-modal="false">
              <div v-if="quickApprovalRow" style="margin-bottom: 20px; padding: 15px; background: #f5f7fa; border-radius: 8px;">
                <div style="font-size: 16px; font-weight: bold; margin-bottom: 10px;">{{ quickApprovalRow.name }}</div>
                <div style="display: flex; gap: 20px; color: #606266; font-size: 14px;">
                  <span>类型：{{ quickApprovalRow.supplier_type }}</span>
                  <span>行业：{{ quickApprovalRow.industry }}</span>
                  <span>联系人：{{ quickApprovalRow.contact }}</span>
                </div>
              </div>

              <el-form label-width="100px">
                <el-form-item label="审批结果" required>
                  <el-radio-group v-model="quickApprovalResult" style="display: flex; flex-direction: column; gap: 15px;">
                    <el-radio value="approve" style="height: auto; padding: 12px; border: 1px solid #dcdfe6; border-radius: 6px; margin-right: 0;">
                      <div>
                        <div style="font-weight: bold; color: #67c23a;">通过</div>
                        <div style="font-size: 12px; color: #909399; margin-top: 4px;">审批通过，供应商将进入合格供应商名录</div>
                      </div>
                    </el-radio>
                    <el-radio value="reject_fill" style="height: auto; padding: 12px; border: 1px solid #dcdfe6; border-radius: 6px; margin-right: 0;">
                      <div>
                        <div style="font-weight: bold; color: #e6a23c;">不通过 - 填写问题</div>
                        <div style="font-size: 12px; color: #909399; margin-top: 4px;">资料填写有误或不完整，退回供应商补充修改</div>
                      </div>
                    </el-radio>
                    <el-radio value="reject_unqualified" style="height: auto; padding: 12px; border: 1px solid #dcdfe6; border-radius: 6px; margin-right: 0;">
                      <div>
                        <div style="font-weight: bold; color: #f56c6c;">不通过 - 资质不合格</div>
                        <div style="font-size: 12px; color: #909399; margin-top: 4px;">供应商本身资质不符合要求，审批终止</div>
                      </div>
                    </el-radio>
                  </el-radio-group>
                </el-form-item>

                <el-form-item v-if="quickApprovalResult === 'reject_fill'" label="补充说明" required>
                  <el-input v-model="quickApprovalRemark" type="textarea" :rows="3" placeholder="请说明需要补充或修改的内容..."></el-input>
                </el-form-item>

                <el-form-item v-if="quickApprovalResult === 'reject_unqualified'" label="不合格原因" required>
                  <el-input v-model="quickApprovalRemark" type="textarea" :rows="3" placeholder="请说明资质不合格的具体原因..."></el-input>
                </el-form-item>
              </el-form>

              <template #footer>
                <el-button @click="showQuickApprovalDialog = false">取消</el-button>
                <el-button
                  :type="quickApprovalResult === 'approve' ? 'success' : quickApprovalResult ? 'warning' : 'primary'"
                  @click="confirmQuickApproval"
                  :disabled="!quickApprovalResult"
                >
                  确认提交
                </el-button>
              </template>
            </el-dialog>
          </div>

          <!-- VIEW 2: 供应商变更 (NEW - Inline Edit) -->
          <div v-else-if="currentMenu === 'change'">
            <div class="form-section-title">供应商信息变更</div>
            <el-alert title="提示：直接在表格中修改信息，点击右侧“保存”按钮提交变更。" type="warning" style="margin-bottom: 20px;" :closable="false"></el-alert>

            <div style="margin-bottom: 20px;">
              <el-form :inline="true" :model="searchForm" class="demo-form-inline">
                <el-form-item label="企业名称">
                  <el-input v-model="searchForm.name" placeholder="请输入名称"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                  <el-select v-model="searchForm.status" placeholder="全部" style="width: 150px;">
                    <el-option label="合格" value="qualified"></el-option>
                    <el-option label="潜在" value="temp"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="onSearch">筛选</el-button>
                </el-form-item>
              </el-form>
            </div>

            <el-table :data="tableData" border style="width: 100%" height="500">
              <el-table-column prop="id" label="序号" width="60" fixed></el-table-column>
              
              <el-table-column label="企业名称" min-width="180">
                <template #default="scope">
                    <el-input v-model="scope.row.enterprise_name" size="small"></el-input>
                </template>
              </el-table-column>
              
              <el-table-column label="统一社会信用代码" min-width="180">
                <template #default="scope">
                    <el-input v-model="scope.row.credit_code" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="供应商类型" min-width="120">
                <template #default="scope">
                    <el-select v-model="scope.row.entry_type" size="small">
                        <el-option label="生产商" value="manufacturing"></el-option>
                        <el-option label="贸易商" value="trading"></el-option>
                    </el-select>
                </template>
              </el-table-column>

              <el-table-column label="供应品类" min-width="120">
                <template #default="scope">
                    <el-select v-model="scope.row.category" size="small">
                        <el-option label="原材料" value="raw"></el-option>
                        <el-option label="电子元器件" value="electronic"></el-option>
                        <el-option label="办公用品" value="office"></el-option>
                    </el-select>
                </template>
              </el-table-column>

              <el-table-column label="合作区域" min-width="120">
                <template #default="scope">
                    <el-input v-model="scope.row.region" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="合作行业" min-width="120">
                <template #default="scope">
                    <el-select v-model="scope.row.industry" size="small">
                        <el-option label="核电" value="nuclear"></el-option>
                        <el-option label="石化" value="petrochemical"></el-option>
                        <el-option label="军工" value="military"></el-option>
                        <el-option label="其他" value="other"></el-option>
                    </el-select>
                </template>
              </el-table-column>

              <el-table-column label="合作品牌" min-width="120">
                <template #default="scope">
                    <el-input v-model="scope.row.brand" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="联系人" min-width="100">
                <template #default="scope">
                    <el-input v-model="scope.row.contact_person" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="联系电话" min-width="130">
                <template #default="scope">
                    <el-input v-model="scope.row.mobile" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="开户行" min-width="150">
                <template #default="scope">
                    <el-input v-model="scope.row.bank_name" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="开户行名称" min-width="180">
                <template #default="scope">
                    <el-input v-model="scope.row.branch_name" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="结算方式" min-width="120">
                <template #default="scope">
                    <el-select v-model="scope.row.settlement" size="small">
                        <el-option label="月结30天" value="m30"></el-option>
                        <el-option label="预付" value="prepay"></el-option>
                        <el-option label="货到付款" value="cod"></el-option>
                    </el-select>
                </template>
              </el-table-column>

              <el-table-column label="采购员" min-width="100">
                <template #default="scope">
                    <el-input v-model="scope.row.purchaser" size="small"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="潜在或合格" min-width="120">
                <template #default="scope">
                    <el-select v-model="scope.row.status_base" size="small">
                        <el-option label="潜在" value="temp"></el-option>
                        <el-option label="合格" value="qualified"></el-option>
                    </el-select>
                </template>
              </el-table-column>

              <el-table-column label="是否黑名单" width="100">
                <template #default="scope">
                    <el-select v-model="scope.row.is_blacklist" size="small" style="width: 100%;">
                      <el-option label="否" value="否"></el-option>
                      <el-option label="是" value="是"></el-option>
                    </el-select>
                </template>
              </el-table-column>

              <el-table-column label="黑名单原因" min-width="150">
                <template #default="scope">
                    <el-input
                      v-model="scope.row.blacklist_reason"
                      size="small"
                      placeholder="请输入原因"
                      :disabled="scope.row.is_blacklist !== '是'"
                    ></el-input>
                </template>
              </el-table-column>

              <el-table-column label="操作" width="80" fixed="right">
                <template #default="scope">
                  <el-button type="primary" size="small" @click="submitRowChange(scope.row)">保存</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>

          <!-- VIEW: 供应商档案库 -->
          <div v-else-if="currentMenu === 'supplier-list'">
            <div class="form-section-title">供应商档案库</div>

            <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
              <div>
                <el-input
                  placeholder="搜索供应商名称/信用代码"
                  style="width: 250px; margin-right: 10px;"
                  prefix-icon="Search"
                  v-model="supplierSearch.keyword"
                  clearable
                ></el-input>
                <el-select v-model="supplierSearch.industry" placeholder="全部行业" style="width: 120px; margin-right: 10px;" clearable>
                  <el-option label="全部" value=""></el-option>
                  <el-option label="核电" value="nuclear"></el-option>
                  <el-option label="军工" value="military"></el-option>
                  <el-option label="石化" value="petrochemical"></el-option>
                  <el-option label="普通" value="normal"></el-option>
                </el-select>
                <el-select v-model="supplierSearch.status" placeholder="全部状态" style="width: 120px; margin-right: 10px;" clearable>
                  <el-option label="全部" value=""></el-option>
                  <el-option label="准入" value="approved"></el-option>
                  <el-option label="驳回" value="rejected"></el-option>
                </el-select>
                <el-select v-model="supplierSearch.supplierType" placeholder="供应商类型" style="width: 120px; margin-right: 10px;" clearable>
                  <el-option label="全部" value=""></el-option>
                  <el-option label="生产商" value="manufacturing"></el-option>
                  <el-option label="贸易商" value="trading"></el-option>
                </el-select>
                <el-button type="primary" @click="searchSuppliers">搜索</el-button>
                <el-button @click="resetSupplierSearch">重置</el-button>
              </div>
              <div>
                <el-button type="success" @click="exportSupplierList">导出全部 Excel</el-button>
              </div>
            </div>

            <!-- 完整字段表格 -->
            <el-table :data="supplierListData" border style="width: 100%" height="500" @row-click="viewSupplierDetail">
              <el-table-column type="index" label="序号" width="60" fixed />
              <!-- 准入申请字段 -->
              <el-table-column prop="name" label="企业名称" min-width="180" fixed />
              <el-table-column prop="credit_code" label="统一社会信用代码" width="180" />
              <el-table-column prop="type" label="供应商类型" width="100" />
              <el-table-column prop="category" label="供应品类" width="100" />
              <el-table-column prop="region" label="合作区域" width="100" />
              <el-table-column prop="industry" label="合作行业" width="100">
                <template #default="{ row }">
                  <el-tag v-if="row.industry === 'nuclear'" type="danger" size="small">核电</el-tag>
                  <el-tag v-else-if="row.industry === 'military'" type="warning" size="small">军工</el-tag>
                  <el-tag v-else-if="row.industry === 'petrochemical'" type="info" size="small">石化</el-tag>
                  <el-tag v-else type="info" size="small">普通</el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="brand" label="合作品牌" width="100" />
              <el-table-column prop="contact_person" label="联系人" width="90" />
              <el-table-column prop="mobile" label="联系电话" width="120" />
              <!-- 财务信息 -->
              <el-table-column prop="bank_name" label="开户行" width="120" />
              <el-table-column prop="branch_name" label="开户行名称" width="150" />
              <el-table-column prop="bank_account" label="银行账号" width="180" />
              <el-table-column prop="settlement" label="结算方式" width="100" />
              <!-- 采购信息 -->
              <el-table-column prop="purchaser" label="采购员" width="90" />
              <!-- 供应商变更字段 -->
              <el-table-column prop="supplier_status" label="潜在/合格" width="90">
                <template #default="{ row }">
                  <el-tag :type="row.supplier_status === 'qualified' ? 'success' : 'warning'" size="small">
                    {{ row.supplier_status === 'qualified' ? '合格' : '潜在' }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="is_blacklist" label="黑名单" width="80">
                <template #default="{ row }">
                  <el-tag v-if="row.is_blacklist" type="danger" size="small">是</el-tag>
                  <el-tag v-else type="success" size="small">否</el-tag>
                </template>
              </el-table-column>
              <!-- 审批信息 -->
              <el-table-column prop="approval_status" label="审批状态" width="90">
                <template #default="{ row }">
                  <el-tag :type="row.approval_status === '已通过' ? 'success' : row.approval_status === '已驳回' ? 'danger' : 'warning'" size="small">
                    {{ row.approval_status }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="current_node" label="当前节点" width="110" />
              <el-table-column prop="apply_time" label="申请时间" width="160" />
              <el-table-column prop="approval_time" label="审批完成时间" width="160" />
              <el-table-column prop="last_change_time" label="最后变更时间" width="160" />
              <el-table-column label="操作" width="100" fixed="right">
                <template #default="{ row }">
                  <el-button type="primary" size="small" link @click.stop="viewSupplierDetail(row)">详情</el-button>
                </template>
              </el-table-column>
            </el-table>

            <div style="margin-top: 20px; text-align: right;">
              <el-pagination
                background
                layout="prev, pager, next, sizes, total"
                :total="supplierList.total"
                :current-page="supplierList.page"
                :page-size="supplierList.pageSize"
                @size-change="handleSupplierPageSizeChange"
                @current-change="handleSupplierPageChange"
              >
              </el-pagination>
            </div>

            <!-- 供应商详情弹窗 -->
            <el-dialog v-model="supplierDetailVisible" title="供应商档案详情" width="90%" top="3vh">
              <div v-if="selectedSupplier">
                <el-tabs type="border-card">
                  <el-tab-pane label="基本信息">
                    <el-descriptions :column="3" border>
                      <el-descriptions-item label="企业名称">{{ selectedSupplier.name }}</el-descriptions-item>
                      <el-descriptions-item label="统一社会信用代码">{{ selectedSupplier.credit_code }}</el-descriptions-item>
                      <el-descriptions-item label="供应商类型">{{ selectedSupplier.type }}</el-descriptions-item>
                      <el-descriptions-item label="供应品类">{{ selectedSupplier.category }}</el-descriptions-item>
                      <el-descriptions-item label="合作区域">{{ selectedSupplier.region }}</el-descriptions-item>
                      <el-descriptions-item label="合作行业">
                        <el-tag v-if="selectedSupplier.industry === 'nuclear'" type="danger">核电</el-tag>
                        <el-tag v-else-if="selectedSupplier.industry === 'military'" type="warning">军工</el-tag>
                        <el-tag v-else>{{ selectedSupplier.industry_text || selectedSupplier.industry }}</el-tag>
                      </el-descriptions-item>
                      <el-descriptions-item label="合作品牌">{{ selectedSupplier.brand }}</el-descriptions-item>
                      <el-descriptions-item label="联系人">{{ selectedSupplier.contact_person }}</el-descriptions-item>
                      <el-descriptions-item label="联系电话">{{ selectedSupplier.mobile }}</el-descriptions-item>
                    </el-descriptions>
                  </el-tab-pane>

                  <el-tab-pane label="财务信息">
                    <el-descriptions :column="2" border>
                      <el-descriptions-item label="开户行">{{ selectedSupplier.bank_name }}</el-descriptions-item>
                      <el-descriptions-item label="开户行名称(支行)">{{ selectedSupplier.branch_name }}</el-descriptions-item>
                      <el-descriptions-item label="银行账号">{{ selectedSupplier.bank_account }}</el-descriptions-item>
                      <el-descriptions-item label="结算方式">{{ selectedSupplier.settlement }}</el-descriptions-item>
                    </el-descriptions>
                  </el-tab-pane>

                  <el-tab-pane label="采购信息">
                    <el-descriptions :column="2" border>
                      <el-descriptions-item label="采购员">{{ selectedSupplier.purchaser }}</el-descriptions-item>
                      <el-descriptions-item label="供应商状态">
                        <el-tag :type="selectedSupplier.supplier_status === 'qualified' ? 'success' : 'warning'">
                          {{ selectedSupplier.supplier_status === 'qualified' ? '合格供应商' : '潜在供应商' }}
                        </el-tag>
                      </el-descriptions-item>
                      <el-descriptions-item label="黑名单状态">
                        <el-tag :type="selectedSupplier.is_blacklist ? 'danger' : 'success'">
                          {{ selectedSupplier.is_blacklist ? '已列入黑名单' : '正常' }}
                        </el-tag>
                      </el-descriptions-item>
                    </el-descriptions>
                  </el-tab-pane>

                  <el-tab-pane label="审批信息">
                    <el-descriptions :column="2" border style="margin-bottom: 20px;">
                      <el-descriptions-item label="审批状态">
                        <el-tag :type="selectedSupplier.approval_status === '已通过' ? 'success' : selectedSupplier.approval_status === '已驳回' ? 'danger' : 'warning'">
                          {{ selectedSupplier.approval_status }}
                        </el-tag>
                      </el-descriptions-item>
                      <el-descriptions-item label="当前节点">{{ selectedSupplier.current_node }}</el-descriptions-item>
                      <el-descriptions-item label="申请时间">{{ selectedSupplier.apply_time }}</el-descriptions-item>
                      <el-descriptions-item label="审批完成时间">{{ selectedSupplier.approval_time }}</el-descriptions-item>
                    </el-descriptions>

                    <h4 style="margin-bottom: 15px;">审批流程</h4>
                    <el-timeline>
                      <el-timeline-item
                        v-for="(step, index) in selectedSupplier.approval_flow || approvalSteps"
                        :key="index"
                        :type="step.status === 'completed' ? 'success' : step.status === 'current' ? 'primary' : 'info'"
                        :hollow="step.status === 'pending'"
                        :timestamp="step.time"
                      >
                        <h4>{{ step.node }}</h4>
                        <p v-if="step.approver">审批人：{{ step.approver }}</p>
                        <p v-if="step.comment">意见：{{ step.comment }}</p>
                      </el-timeline-item>
                    </el-timeline>
                  </el-tab-pane>

                  <el-tab-pane label="变更记录">
                    <el-table :data="selectedSupplier.change_history || []" border style="width: 100%">
                      <el-table-column prop="change_time" label="变更时间" width="160" />
                      <el-table-column prop="change_field" label="变更字段" width="120" />
                      <el-table-column prop="old_value" label="原值" min-width="150" />
                      <el-table-column prop="new_value" label="新值" min-width="150" />
                      <el-table-column prop="change_by" label="变更人" width="100" />
                      <el-table-column prop="change_reason" label="变更原因" min-width="150" />
                    </el-table>
                    <el-empty v-if="!selectedSupplier.change_history?.length" description="暂无变更记录" />
                  </el-tab-pane>

                  <el-tab-pane label="附件资料">
                    <div class="attachment-list">
                      <el-tag
                        v-for="file in selectedSupplier.attachments"
                        :key="file.name"
                        style="margin: 5px; cursor: pointer;"
                        @click="previewAttachment(file)"
                      >
                        <el-icon style="margin-right: 5px;"><Document /></el-icon>
                        {{ file.name }}
                      </el-tag>
                      <el-empty v-if="!selectedSupplier.attachments?.length" description="暂无附件" />
                    </div>
                  </el-tab-pane>
                </el-tabs>

                <div style="text-align: center; margin-top: 20px;">
                  <el-button @click="supplierDetailVisible = false">关闭</el-button>
                  <el-button type="primary" @click="printApprovalForm">打印审批单</el-button>
                </div>
              </div>
            </el-dialog>
          </div>

          <!-- VIEW: 我的询价（销售） -->
          <div v-else-if="currentMenu === 'my-inquiry'">
            <!-- Tab 切换：询价录入 / 我的询价列表 -->
            <el-tabs v-model="inquiryTab" type="card">
              <el-tab-pane label="询价需求录入" name="entry">
                <div class="form-section-title">询价需求录入</div>
                <el-form :model="inquiryForm" label-width="120px" ref="inquiryFormRef">
                  <el-row :gutter="20">
                    <el-col :span="12">
                      <el-form-item label="项目名称" prop="project_name" :rules="[{required: true, message: '请输入项目名称'}]">
                        <el-input v-model="inquiryForm.project_name" placeholder="请输入项目名称"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="终端用户名称" prop="end_user" :rules="[{required: true, message: '请输入终端用户名称'}]">
                        <el-input v-model="inquiryForm.end_user" placeholder="请输入终端用户名称"></el-input>
                      </el-form-item>
                    </el-col>
                  </el-row>

                  <el-row :gutter="20">
                    <el-col :span="12">
                      <el-form-item label="要求报价时效" prop="quote_deadline" :rules="[{required: true, message: '请选择报价截止时间'}]">
                        <el-date-picker v-model="inquiryForm.quote_deadline" type="datetime" placeholder="选择报价截止时间" style="width: 100%"></el-date-picker>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="期望交期">
                        <el-date-picker v-model="inquiryForm.expected_delivery" type="date" placeholder="选择期望交货日期" style="width: 100%"></el-date-picker>
                      </el-form-item>
                    </el-col>
                  </el-row>

                  <el-form-item label="物料信息">
                    <div style="width: 100%; overflow-x: auto;">
                      <el-table :data="inquiryForm.items" border style="width: 100%" size="small">
                        <el-table-column label="序号" width="60" align="center">
                          <template #default="scope">{{ scope.$index + 1 }}</template>
                        </el-table-column>
                        <el-table-column label="品牌" width="100">
                          <template #default="scope">
                            <el-input v-model="scope.row.brand" placeholder="品牌" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="商品名称" width="140">
                          <template #default="scope">
                            <el-input v-model="scope.row.product_name" placeholder="商品名称" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="规格型号" width="120">
                          <template #default="scope">
                            <el-input v-model="scope.row.specification" placeholder="规格型号" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="数量" width="100">
                          <template #default="scope">
                            <el-input-number v-model="scope.row.quantity" :min="1" size="small" controls-position="right" style="width: 80px;"></el-input-number>
                          </template>
                        </el-table-column>
                        <el-table-column label="单位" width="80">
                          <template #default="scope">
                            <el-input v-model="scope.row.unit" placeholder="单位" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="介质" width="100">
                          <template #default="scope">
                            <el-input v-model="scope.row.medium" placeholder="介质" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="口径" width="80">
                          <template #default="scope">
                            <el-input v-model="scope.row.diameter" placeholder="口径" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="材质" width="100">
                          <template #default="scope">
                            <el-input v-model="scope.row.material" placeholder="材质" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="备注" width="150">
                          <template #default="scope">
                            <el-input v-model="scope.row.remark" placeholder="备注" size="small"></el-input>
                          </template>
                        </el-table-column>
                        <el-table-column label="操作" width="80" fixed="right">
                          <template #default="scope">
                            <el-button type="danger" size="small" link @click="removeInquiryItem(scope.$index)">删除</el-button>
                          </template>
                        </el-table-column>
                      </el-table>
                    </div>
                    <el-button type="primary" size="small" @click="addInquiryItem" style="margin-top: 10px;">+ 添加物料</el-button>
                  </el-form-item>

                  <el-form-item label="技术要求">
                    <el-input v-model="inquiryForm.technical_requirements" type="textarea" :rows="3" placeholder="请输入技术要求"></el-input>
                  </el-form-item>

                  <el-form-item label="技术附件">
                    <el-upload action="#" :auto-upload="false" :on-change="handleTechFile" :file-list="techFiles" multiple>
                      <el-button type="primary">上传附件</el-button>
                      <template #tip>
                        <div class="el-upload__tip">支持上传技术图纸、要求文档等</div>
                      </template>
                    </el-upload>
                  </el-form-item>

                  <el-row :gutter="20">
                    <el-col :span="12">
                      <el-form-item label="录入人">
                        <el-input :value="currentUser" disabled></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="录入时间">
                        <el-input :value="currentTime" disabled></el-input>
                      </el-form-item>
                    </el-col>
                  </el-row>

                  <el-form-item>
                    <el-button type="primary" @click="submitInquiry">单条提交</el-button>
                    <el-button type="success" @click="showImportDialog">批量导入</el-button>
                  </el-form-item>
                </el-form>
              </el-tab-pane>

              <el-tab-pane label="我的询价列表" name="list">
                <div class="form-section-title">我的询价列表</div>

                <!-- 搜索条件 -->
                <div style="margin-bottom: 20px; padding: 15px; background-color: #fafafa; border-radius: 4px;">
                  <el-form :inline="true" :model="inquirySearchForm">
                    <el-form-item label="项目名称">
                      <el-input v-model="inquirySearchForm.project_name" placeholder="项目名称" clearable style="width: 180px;"></el-input>
                    </el-form-item>
                    <el-form-item label="状态">
                      <el-select v-model="inquirySearchForm.status" placeholder="选择状态" clearable style="width: 150px;">
                        <el-option label="待采购处理" value="pending"></el-option>
                        <el-option label="澄清中" value="clarifying"></el-option>
                        <el-option label="已报价" value="quoted"></el-option>
                        <el-option label="待反馈中标" value="waiting_feedback"></el-option>
                        <el-option label="已中标" value="won"></el-option>
                        <el-option label="未中标" value="lost"></el-option>
                      </el-select>
                    </el-form-item>
                    <el-form-item>
                      <el-button type="primary" @click="searchMyInquiries">查询</el-button>
                      <el-button @click="resetInquirySearch">重置</el-button>
                    </el-form-item>
                  </el-form>
                </div>

                <!-- 询价列表 -->
                <el-table :data="myInquiryList" border style="width: 100%">
                  <el-table-column prop="inquiry_no" label="询价编号" width="150"></el-table-column>
                  <el-table-column prop="project_name" label="项目名称" width="180"></el-table-column>
                  <el-table-column prop="end_user" label="终端用户名称" width="150"></el-table-column>
                  <el-table-column prop="quote_deadline" label="要求报价时效" width="160"></el-table-column>
                  <el-table-column prop="material_summary" label="物料摘要" width="150"></el-table-column>
                  <el-table-column prop="status" label="状态" width="120">
                    <template #default="scope">
                      <el-tag :type="getInquiryStatusType(scope.row.status)">{{ getInquiryStatusText(scope.row.status) }}</el-tag>
                    </template>
                  </el-table-column>
                  <el-table-column prop="project_result" label="项目结果" width="100">
                    <template #default="scope">
                      <el-tag v-if="scope.row.project_result === 'won'" type="success">成交</el-tag>
                      <el-tag v-else-if="scope.row.project_result === 'lost'" type="danger">未成交</el-tag>
                      <el-tag v-else-if="scope.row.project_result === 'pending'" type="warning">待定</el-tag>
                      <el-tag v-else-if="scope.row.project_result === 'delayed'" type="info">延迟</el-tag>
                      <span v-else>-</span>
                    </template>
                  </el-table-column>
                  <el-table-column prop="created_at" label="创建时间" width="160"></el-table-column>
                  <el-table-column prop="feedback_deadline" label="反馈截止" width="160">
                    <template #default="scope">
                      <span v-if="scope.row.status === 'waiting_feedback'" :style="{color: isOverdue(scope.row.feedback_deadline) ? '#f56c6c' : ''}">
                        {{ scope.row.feedback_deadline }}
                        <el-tag v-if="isOverdue(scope.row.feedback_deadline)" type="danger" size="small">已逾期</el-tag>
                      </span>
                      <span v-else>-</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="200" fixed="right">
                    <template #default="scope">
                      <el-button type="primary" size="small" link @click="viewInquiryDetail(scope.row)">查看</el-button>
                      <el-button v-if="scope.row.status === 'clarifying'" type="warning" size="small" link @click="replyClarification(scope.row)">回复澄清</el-button>
                      <el-button v-if="scope.row.status === 'waiting_feedback'" type="success" size="small" link @click="showFeedbackDialog(scope.row)">反馈中标</el-button>
                    </template>
                  </el-table-column>
                </el-table>

                <el-pagination
                  style="margin-top: 20px; justify-content: flex-end;"
                  background
                  layout="total, sizes, prev, pager, next"
                  :total="myInquiryTotal"
                  :page-sizes="[10, 20, 50]"
                  v-model:current-page="myInquiryPage"
                  v-model:page-size="myInquiryPageSize"
                ></el-pagination>
              </el-tab-pane>
            </el-tabs>

            <!-- 批量导入对话框 -->
            <el-dialog v-model="importDialogVisible" title="批量导入询价" width="500px">
              <el-upload class="upload-demo" drag action="#" :auto-upload="false" :on-change="handleExcelFile" accept=".xlsx,.xls">
                <el-icon class="el-icon--upload"><UploadFilled /></el-icon>
                <div class="el-upload__text">将Excel文件拖到此处，或<em>点击上传</em></div>
                <template #tip>
                  <div class="el-upload__tip">请按照模板格式上传Excel文件 <el-link type="primary" @click="downloadTemplate">下载模板</el-link></div>
                </template>
              </el-upload>
              <template #footer>
                <el-button @click="importDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="importExcel">确认导入</el-button>
              </template>
            </el-dialog>

            <!-- 报价结果反馈对话框（7天规则） -->
            <el-dialog v-model="feedbackDialogVisible" title="报价结果反馈" width="650px">
              <el-alert title="温馨提示：采购报价后第7天，系统自动提醒销售反馈结果。请及时反馈！" type="warning" show-icon style="margin-bottom: 20px;"></el-alert>
              <el-form :model="feedbackForm" label-width="120px" ref="feedbackFormRef">
                <el-form-item label="询价编号">
                  <el-input v-model="feedbackForm.inquiry_no" disabled></el-input>
                </el-form-item>
                <el-form-item label="项目名称">
                  <el-input v-model="feedbackForm.project_name" disabled></el-input>
                </el-form-item>
                <el-form-item label="采购报价日期">
                  <el-input v-model="feedbackForm.quote_date" disabled></el-input>
                </el-form-item>
                <el-form-item label="反馈结果" prop="result" :rules="[{required: true, message: '请选择反馈结果'}]">
                  <el-radio-group v-model="feedbackForm.result">
                    <el-radio label="won">成交</el-radio>
                    <el-radio label="lost">未成交</el-radio>
                    <el-radio label="pending">结果待定</el-radio>
                    <el-radio label="delayed">项目延迟</el-radio>
                  </el-radio-group>
                </el-form-item>

                <!-- 成交时必填字段 -->
                <template v-if="feedbackForm.result === 'won'">
                  <el-form-item label="上传报价计算书" prop="calculation_file" :rules="[{required: true, message: '成交必须上传报价计算书'}]">
                    <el-upload action="#" :auto-upload="false" :on-change="handleCalculationFile" :file-list="calculationFiles" :limit="1">
                      <el-button type="primary">上传计算书</el-button>
                      <template #tip>
                        <div class="el-upload__tip" style="color: #f56c6c;">成交必须上传报价计算书</div>
                      </template>
                    </el-upload>
                  </el-form-item>
                  <el-form-item label="销售订单号/生产单号" prop="order_no" :rules="[{required: true, message: '成交必须填写订单号'}]">
                    <el-input v-model="feedbackForm.order_no" placeholder="请输入销售订单号或生产单号"></el-input>
                  </el-form-item>
                  <el-form-item label="成交金额">
                    <el-input-number v-model="feedbackForm.deal_amount" :min="0" :precision="2" style="width: 200px;"></el-input-number>
                    <span style="margin-left: 10px;">元</span>
                  </el-form-item>
                </template>

                <!-- 未成交原因 -->
                <el-form-item v-if="feedbackForm.result === 'lost'" label="未成交原因" :rules="[{required: true, message: '请填写未成交原因'}]">
                  <el-input v-model="feedbackForm.no_deal_reason" type="textarea" :rows="2" placeholder="请说明未成交原因"></el-input>
                </el-form-item>

                <!-- 延迟说明 -->
                <el-form-item v-if="feedbackForm.result === 'delayed'" label="延迟说明" :rules="[{required: true, message: '请填写延迟原因'}]">
                  <el-input v-model="feedbackForm.delay_reason" type="textarea" :rows="2" placeholder="请说明项目延迟原因"></el-input>
                </el-form-item>

                <el-form-item label="备注">
                  <el-input v-model="feedbackForm.remark" type="textarea" :rows="2" placeholder="请输入备注信息（可选）"></el-input>
                </el-form-item>
              </el-form>
              <template #footer>
                <el-button @click="feedbackDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="submitFeedback">提交反馈</el-button>
              </template>
            </el-dialog>

            <!-- 澄清回复对话框 -->
            <el-dialog v-model="clarifyDialogVisible" title="回复澄清" width="700px">
              <div v-if="clarifyHistory.length > 0" style="margin-bottom: 20px; max-height: 300px; overflow-y: auto; padding: 10px; background-color: #fafafa; border-radius: 4px;">
                <div class="form-section-title">澄清记录</div>
                <el-timeline>
                  <el-timeline-item v-for="item in clarifyHistory" :key="item.id" :timestamp="item.created_at" :type="item.from === 'purchase' ? 'primary' : 'success'">
                    <div>
                      <strong>{{ item.from === 'purchase' ? '采购' : '销售' }}：</strong>
                      <p style="margin: 5px 0 0 0; color: #606266;">{{ item.content }}</p>
                    </div>
                  </el-timeline-item>
                </el-timeline>
              </div>
              <el-form :model="clarifyReplyForm" label-width="80px">
                <el-form-item label="回复内容">
                  <el-input v-model="clarifyReplyForm.content" type="textarea" :rows="4" placeholder="请输入澄清回复内容"></el-input>
                </el-form-item>
                <el-form-item label="附件">
                  <el-upload action="#" :auto-upload="false" :on-change="handleClarifyFile" :file-list="clarifyFiles" multiple>
                    <el-button type="primary" size="small">上传附件</el-button>
                  </el-upload>
                </el-form-item>
              </el-form>
              <template #footer>
                <el-button @click="clarifyDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="submitClarifyReply">提交回复</el-button>
              </template>
            </el-dialog>
          </div>

          <!-- VIEW: 待报价项目（采购） -->
          <div v-else-if="currentMenu === 'pending-quote'">
            <div class="form-section-title">待报价项目</div>

            <!-- 搜索条件 -->
            <div style="margin-bottom: 20px; padding: 15px; background-color: #fafafa; border-radius: 4px;">
              <el-form :inline="true" :model="purchaseSearchForm">
                <el-form-item label="询价编号">
                  <el-input v-model="purchaseSearchForm.inquiry_no" placeholder="询价编号" clearable style="width: 150px;"></el-input>
                </el-form-item>
                <el-form-item label="项目名称">
                  <el-input v-model="purchaseSearchForm.project_name" placeholder="项目名称" clearable style="width: 150px;"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                  <el-select v-model="purchaseSearchForm.status" placeholder="全部状态" clearable style="width: 120px;">
                    <el-option label="待处理" value="pending"></el-option>
                    <el-option label="澄清中" value="clarifying"></el-option>
                    <el-option label="可报价" value="can_quote"></el-option>
                    <el-option label="已报价" value="quoted"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="searchPurchaseList">查询</el-button>
                  <el-button @click="resetPurchaseSearch">重置</el-button>
                </el-form-item>
                <el-form-item style="float: right;">
                  <el-button type="success" @click="showImportQuoteDialog">批量导入报价</el-button>
                </el-form-item>
              </el-form>
            </div>

            <!-- 提示 -->
            <el-alert title="点击每行左侧箭头可展开查看完整物料明细和报价信息" type="info" show-icon :closable="false" style="margin-bottom: 15px;"></el-alert>

            <!-- 待报价列表 - 展开显示物料明细 -->
            <el-table :data="purchaseInquiryList" border style="width: 100%" row-key="id">
              <el-table-column type="expand">
                <template #default="props">
                  <div style="padding: 10px 20px; background-color: #fafafa;">
                    <div style="margin-bottom: 10px; font-weight: bold; color: #409eff;">询价填写内容（销售录入）</div>
                    <el-table :data="props.row.items" border size="small" style="margin-bottom: 15px;">
                      <el-table-column label="序号" width="60" align="center">
                        <template #default="scope">{{ scope.$index + 1 }}</template>
                      </el-table-column>
                      <el-table-column prop="brand" label="品牌" width="100"></el-table-column>
                      <el-table-column prop="product_name" label="商品名称" width="120"></el-table-column>
                      <el-table-column prop="specification" label="规格型号" width="120"></el-table-column>
                      <el-table-column prop="quantity" label="数量" width="80" align="center"></el-table-column>
                      <el-table-column prop="unit" label="单位" width="60" align="center"></el-table-column>
                      <el-table-column prop="medium" label="介质" width="80"></el-table-column>
                      <el-table-column prop="diameter" label="口径" width="80"></el-table-column>
                      <el-table-column prop="material" label="材质" width="100"></el-table-column>
                      <el-table-column prop="remark" label="备注" min-width="120"></el-table-column>
                      <el-table-column prop="project_result" label="项目结果" width="90">
                        <template #default="scope">
                          <el-tag v-if="scope.row.project_result === 'won'" type="success" size="small">成交</el-tag>
                          <el-tag v-else-if="scope.row.project_result === 'lost'" type="danger" size="small">未成交</el-tag>
                          <el-tag v-else-if="scope.row.project_result === 'pending'" type="warning" size="small">待定</el-tag>
                          <span v-else>-</span>
                        </template>
                      </el-table-column>
                    </el-table>
                    <div style="margin-bottom: 10px; font-weight: bold; color: #f56c6c;">报价填写内容（采购填写，仅采购可见）</div>
                    <el-table :data="props.row.items" border size="small">
                      <el-table-column label="序号" width="60" align="center">
                        <template #default="scope">{{ scope.$index + 1 }}</template>
                      </el-table-column>
                      <el-table-column prop="product_name" label="商品名称" width="120"></el-table-column>
                      <el-table-column prop="specification" label="规格型号" width="120"></el-table-column>
                      <el-table-column prop="unit" label="单位" width="60" align="center"></el-table-column>
                      <el-table-column prop="unit_price" label="单价（含税）" width="110">
                        <template #default="scope">
                          <span v-if="scope.row.unit_price" style="color: #f56c6c;">¥{{ scope.row.unit_price }}</span>
                          <span v-else style="color: #999;">待填写</span>
                        </template>
                      </el-table-column>
                      <el-table-column prop="delivery_days" label="货期" width="80">
                        <template #default="scope">
                          <span v-if="scope.row.delivery_days">{{ scope.row.delivery_days }}天</span>
                          <span v-else style="color: #999;">-</span>
                        </template>
                      </el-table-column>
                      <el-table-column prop="quote_remark" label="备注" width="120">
                        <template #default="scope">{{ scope.row.quote_remark || '-' }}</template>
                      </el-table-column>
                      <el-table-column prop="supplier_name" label="供应商" min-width="150">
                        <template #default="scope">
                          <span v-if="scope.row.supplier_name" style="color: #f56c6c;">{{ scope.row.supplier_name }}</span>
                          <span v-else style="color: #999;">待选择</span>
                        </template>
                      </el-table-column>
                    </el-table>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="inquiry_no" label="询价编号" width="150"></el-table-column>
              <el-table-column prop="project_name" label="项目名称" width="180"></el-table-column>
              <el-table-column prop="end_user" label="终端用户" width="120"></el-table-column>
              <el-table-column prop="sales_name" label="销售人员" width="100"></el-table-column>
              <el-table-column prop="material_summary" label="物料摘要" width="150"></el-table-column>
              <el-table-column prop="quote_deadline" label="报价截止" width="160">
                <template #default="scope">
                  <span :style="{color: isPurchaseDeadlineUrgent(scope.row.quote_deadline) ? '#f56c6c' : ''}">
                    {{ scope.row.quote_deadline }}
                    <el-tag v-if="isPurchaseDeadlineUrgent(scope.row.quote_deadline)" type="danger" size="small">紧急</el-tag>
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="status" label="状态" width="100">
                <template #default="scope">
                  <el-tag :type="getPurchaseStatusType(scope.row.status)">{{ getPurchaseStatusText(scope.row.status) }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="created_at" label="提交时间" width="160"></el-table-column>
              <el-table-column label="操作" width="220" fixed="right">
                <template #default="scope">
                  <el-button type="primary" size="small" link @click="viewPurchaseDetail(scope.row)">查看详情</el-button>
                  <el-button v-if="scope.row.status === 'pending' || scope.row.status === 'can_quote'" type="warning" size="small" link @click="showPurchaseClarifyDialog(scope.row)">发起澄清</el-button>
                  <el-button v-if="scope.row.status === 'can_quote' || scope.row.status === 'pending'" type="success" size="small" link @click="showPurchaseQuoteDialog(scope.row)">填写报价</el-button>
                </template>
              </el-table-column>
            </el-table>

            <el-pagination
              style="margin-top: 20px; justify-content: flex-end;"
              background
              layout="total, sizes, prev, pager, next"
              :total="purchaseInquiryTotal"
              :page-sizes="[10, 20, 50]"
              v-model:current-page="purchaseInquiryPage"
              v-model:page-size="purchaseInquiryPageSize"
            ></el-pagination>

            <!-- 采购发起澄清对话框 -->
            <el-dialog v-model="purchaseClarifyDialogVisible" title="发起澄清" width="600px">
              <el-alert title="澄清发送后，销售将收到通知。销售回复后状态自动变为【可报价】" type="info" show-icon style="margin-bottom: 20px;"></el-alert>
              <el-form :model="purchaseClarifyForm" label-width="100px">
                <el-form-item label="询价编号">
                  <el-input v-model="purchaseClarifyForm.inquiry_no" disabled></el-input>
                </el-form-item>
                <el-form-item label="澄清内容">
                  <el-input v-model="purchaseClarifyForm.content" type="textarea" :rows="4" placeholder="请输入需要销售澄清的问题"></el-input>
                </el-form-item>
                <el-form-item label="附件">
                  <el-upload action="#" :auto-upload="false" :on-change="handlePurchaseClarifyFile" :file-list="purchaseClarifyFiles" multiple>
                    <el-button type="primary" size="small">上传附件</el-button>
                  </el-upload>
                </el-form-item>
              </el-form>
              <template #footer>
                <el-button @click="purchaseClarifyDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="submitPurchaseClarify">发送澄清</el-button>
              </template>
            </el-dialog>

            <!-- 采购填写报价对话框 -->
            <el-dialog v-model="purchaseQuoteDialogVisible" title="填写报价" width="1000px">
              <div style="margin-bottom: 15px;">
                <el-descriptions :column="3" border size="small">
                  <el-descriptions-item label="询价编号">{{ purchaseQuoteForm.inquiry_no }}</el-descriptions-item>
                  <el-descriptions-item label="项目名称">{{ purchaseQuoteForm.project_name }}</el-descriptions-item>
                  <el-descriptions-item label="终端用户">{{ purchaseQuoteForm.end_user }}</el-descriptions-item>
                  <el-descriptions-item label="报价截止">{{ purchaseQuoteForm.quote_deadline }}</el-descriptions-item>
                  <el-descriptions-item label="期望交期">{{ purchaseQuoteForm.expected_delivery }}</el-descriptions-item>
                  <el-descriptions-item label="技术要求">{{ purchaseQuoteForm.technical_requirements || '无' }}</el-descriptions-item>
                </el-descriptions>
              </div>

              <div class="form-section-title">报价明细</div>
              <el-table :data="purchaseQuoteForm.items" border style="width: 100%" size="small">
                <el-table-column label="序号" width="50" align="center">
                  <template #default="scope">{{ scope.$index + 1 }}</template>
                </el-table-column>
                <el-table-column prop="brand" label="品牌" width="80"></el-table-column>
                <el-table-column prop="product_name" label="商品名称" width="120"></el-table-column>
                <el-table-column prop="specification" label="规格型号" width="100"></el-table-column>
                <el-table-column prop="quantity" label="数量" width="60"></el-table-column>
                <el-table-column prop="unit" label="单位" width="50"></el-table-column>
                <el-table-column label="供应商" width="150">
                  <template #default="scope">
                    <el-select v-model="scope.row.supplier_id" placeholder="选择供应商" size="small" style="width: 130px;">
                      <el-option v-for="sup in supplierOptions" :key="sup.id" :label="sup.name" :value="sup.id"></el-option>
                    </el-select>
                  </template>
                </el-table-column>
                <el-table-column label="单价(含税)" width="110">
                  <template #default="scope">
                    <el-input-number v-model="scope.row.unit_price" :min="0" :precision="2" size="small" controls-position="right" style="width: 90px;"></el-input-number>
                  </template>
                </el-table-column>
                <el-table-column label="总价" width="100">
                  <template #default="scope">
                    <span>{{ (scope.row.unit_price * scope.row.quantity).toFixed(2) }}</span>
                  </template>
                </el-table-column>
                <el-table-column label="交期(天)" width="90">
                  <template #default="scope">
                    <el-input-number v-model="scope.row.delivery_days" :min="1" size="small" controls-position="right" style="width: 70px;"></el-input-number>
                  </template>
                </el-table-column>
                <el-table-column label="备注" width="120">
                  <template #default="scope">
                    <el-input v-model="scope.row.quote_remark" placeholder="备注" size="small"></el-input>
                  </template>
                </el-table-column>
              </el-table>

              <el-form :model="purchaseQuoteForm" label-width="100px" style="margin-top: 20px;">
                <el-form-item label="报价备注">
                  <el-input v-model="purchaseQuoteForm.remark" type="textarea" :rows="2" placeholder="请输入报价整体备注"></el-input>
                </el-form-item>
              </el-form>

              <template #footer>
                <el-button @click="purchaseQuoteDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="submitPurchaseQuote">提交报价</el-button>
              </template>
            </el-dialog>

            <!-- 批量导入报价对话框 -->
            <el-dialog v-model="importQuoteDialogVisible" title="批量导入报价" width="500px">
              <el-upload class="upload-demo" drag action="#" :auto-upload="false" :on-change="handleQuoteExcelFile" accept=".xlsx,.xls">
                <el-icon class="el-icon--upload"><UploadFilled /></el-icon>
                <div class="el-upload__text">将Excel文件拖到此处，或<em>点击上传</em></div>
                <template #tip>
                  <div class="el-upload__tip">请按照模板格式上传Excel文件 <el-link type="primary" @click="downloadQuoteTemplate">下载报价模板</el-link></div>
                </template>
              </el-upload>
              <template #footer>
                <el-button @click="importQuoteDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="importQuoteExcel">确认导入</el-button>
              </template>
            </el-dialog>
          </div>

          <!-- VIEW: 全部询报价台账 -->
          <div v-else-if="currentMenu === 'inquiry-ledger'">
            <el-tabs v-model="ledgerTab" type="card">
              <!-- 询报价台账 -->
              <el-tab-pane label="询报价台账" name="ledger">
                <div class="form-section-title">全部询报价台账</div>

                <div style="margin-bottom: 20px; padding: 15px; background-color: #fafafa; border-radius: 4px;">
                  <el-form :inline="true" :model="ledgerSearch">
                    <el-form-item label="时间范围">
                      <el-date-picker v-model="ledgerSearch.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" style="width: 240px;"></el-date-picker>
                    </el-form-item>
                    <el-form-item label="物料">
                      <el-input placeholder="物料名称" v-model="ledgerSearch.material" style="width: 120px;"></el-input>
                    </el-form-item>
                    <el-form-item label="供应商">
                      <el-input placeholder="供应商" v-model="ledgerSearch.supplier" style="width: 120px;"></el-input>
                    </el-form-item>
                    <el-form-item label="项目">
                      <el-input placeholder="项目名称" v-model="ledgerSearch.project" style="width: 120px;"></el-input>
                    </el-form-item>
                    <el-form-item label="成交状态">
                      <el-select v-model="ledgerSearch.deal_status" placeholder="全部" style="width: 100px;" clearable>
                        <el-option label="成交" value="deal"></el-option>
                        <el-option label="未成交" value="no_deal"></el-option>
                        <el-option label="待定" value="pending"></el-option>
                        <el-option label="延迟" value="delayed"></el-option>
                      </el-select>
                    </el-form-item>
                    <el-form-item>
                      <el-button type="primary" @click="searchLedger">搜索</el-button>
                      <el-button @click="resetLedgerSearch">重置</el-button>
                      <el-button type="success" @click="exportLedger">批量导出台账</el-button>
                    </el-form-item>
                  </el-form>
                </div>

                <el-table :data="ledgerData" border style="width: 100%">
                  <el-table-column prop="inquiry_no" label="询价编号" width="150" />
                  <el-table-column prop="project_name" label="项目名称" width="150" />
                  <el-table-column prop="end_user" label="终端用户" width="120" />
                  <el-table-column prop="material" label="物料" width="120" />
                  <el-table-column prop="specification" label="规格型号" width="100" />
                  <el-table-column prop="quantity" label="数量" width="60" />
                  <el-table-column prop="supplier" label="供应商" width="120" />
                  <el-table-column prop="unit_price" label="单价" width="80" />
                  <el-table-column prop="total_price" label="总价" width="100" />
                  <el-table-column prop="history_lowest" label="历史最低" width="80" />
                  <el-table-column prop="agreement_price" label="协议价" width="80" />
                  <el-table-column prop="deal_status" label="成交状态" width="100">
                    <template #default="{ row }">
                      <el-tag :type="getLedgerStatusType(row.deal_status)">{{ row.deal_status }}</el-tag>
                    </template>
                  </el-table-column>
                  <el-table-column prop="order_no" label="订单号" width="130" />
                  <el-table-column prop="quote_date" label="报价日期" width="100" />
                </el-table>

                <el-pagination
                  style="margin-top: 20px; justify-content: flex-end;"
                  background
                  layout="total, sizes, prev, pager, next"
                  :total="ledgerTotal"
                  :page-sizes="[10, 20, 50, 100]"
                  v-model:current-page="ledgerPage"
                  v-model:page-size="ledgerPageSize"
                ></el-pagination>
              </el-tab-pane>

              <!-- 价格库 -->
              <el-tab-pane label="价格库" name="pricelib">
                <div class="form-section-title">价格库检索</div>
                <el-alert title="系统自动构建价格库，支持按物料、供应商、项目、时间范围进行检索。采购全员可查看全部数据。" type="info" show-icon style="margin-bottom: 20px;"></el-alert>

                <div style="margin-bottom: 20px; padding: 15px; background-color: #fafafa; border-radius: 4px;">
                  <el-form :inline="true" :model="priceLibSearch">
                    <el-form-item label="物料">
                      <el-input placeholder="物料名称/规格" v-model="priceLibSearch.material" style="width: 150px;"></el-input>
                    </el-form-item>
                    <el-form-item label="供应商">
                      <el-input placeholder="供应商名称" v-model="priceLibSearch.supplier" style="width: 150px;"></el-input>
                    </el-form-item>
                    <el-form-item label="项目">
                      <el-input placeholder="项目名称" v-model="priceLibSearch.project" style="width: 150px;"></el-input>
                    </el-form-item>
                    <el-form-item label="时间范围">
                      <el-date-picker v-model="priceLibSearch.dateRange" type="daterange" range-separator="至" start-placeholder="开始" end-placeholder="结束" style="width: 220px;"></el-date-picker>
                    </el-form-item>
                    <el-form-item>
                      <el-button type="primary" @click="searchPriceLib">搜索</el-button>
                      <el-button @click="resetPriceLibSearch">重置</el-button>
                    </el-form-item>
                  </el-form>
                </div>

                <el-table :data="priceLibData" border style="width: 100%">
                  <el-table-column prop="material" label="物料名称" width="150" />
                  <el-table-column prop="specification" label="规格型号" width="120" />
                  <el-table-column prop="supplier" label="供应商" width="150" />
                  <el-table-column prop="unit_price" label="单价" width="100" />
                  <el-table-column prop="project_name" label="关联项目" width="150" />
                  <el-table-column prop="quote_date" label="报价日期" width="120" />
                  <el-table-column prop="is_deal" label="是否成交" width="80">
                    <template #default="{ row }">
                      <el-tag :type="row.is_deal ? 'success' : 'info'">{{ row.is_deal ? '是' : '否' }}</el-tag>
                    </template>
                  </el-table-column>
                  <el-table-column prop="remark" label="备注" min-width="150" />
                </el-table>

                <el-pagination
                  style="margin-top: 20px; justify-content: flex-end;"
                  background
                  layout="total, sizes, prev, pager, next"
                  :total="priceLibTotal"
                  :page-sizes="[10, 20, 50]"
                  v-model:current-page="priceLibPage"
                  v-model:page-size="priceLibPageSize"
                ></el-pagination>
              </el-tab-pane>
            </el-tabs>
          </div>

          <!-- VIEW: 中标项目管理 -->
          <div v-else-if="currentMenu === 'bid-projects'">
            <div class="form-section-title">中标项目管理</div>
            <el-alert title="管理已中标（成交）的项目，可查看报价计算书和订单信息" type="success" show-icon style="margin-bottom: 20px;"></el-alert>

            <!-- 搜索 -->
            <div style="margin-bottom: 20px; padding: 15px; background-color: #fafafa; border-radius: 4px;">
              <el-form :inline="true" :model="bidProjectSearch">
                <el-form-item label="项目名称">
                  <el-input v-model="bidProjectSearch.project_name" placeholder="项目名称" clearable style="width: 150px;"></el-input>
                </el-form-item>
                <el-form-item label="客户">
                  <el-input v-model="bidProjectSearch.customer" placeholder="客户名称" clearable style="width: 150px;"></el-input>
                </el-form-item>
                <el-form-item label="中标日期">
                  <el-date-picker v-model="bidProjectSearch.dateRange" type="daterange" range-separator="至" start-placeholder="开始" end-placeholder="结束" style="width: 220px;"></el-date-picker>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="searchBidProjects">查询</el-button>
                  <el-button @click="resetBidProjectSearch">重置</el-button>
                </el-form-item>
              </el-form>
            </div>

            <el-table :data="bidProjectsData" border style="width: 100%">
              <el-table-column prop="inquiry_no" label="询价编号" width="150" />
              <el-table-column prop="project_name" label="项目名称" width="180" />
              <el-table-column prop="end_user" label="终端用户" width="120" />
              <el-table-column prop="sales_name" label="销售人员" width="100" />
              <el-table-column prop="win_date" label="成交日期" width="120" />
              <el-table-column prop="total_amount" label="成交金额" width="120" />
              <el-table-column prop="order_no" label="销售订单号/生产单号" width="150" />
              <el-table-column prop="calculation_file" label="报价计算书" width="120">
                <template #default="{ row }">
                  <el-button v-if="row.calculation_file" type="primary" size="small" link @click="downloadFile(row.calculation_file)">下载</el-button>
                  <span v-else>-</span>
                </template>
              </el-table-column>
              <el-table-column label="操作" width="100" fixed="right">
                <template #default="scope">
                  <el-button type="primary" size="small" link @click="viewBidDetail(scope.row)">详情</el-button>
                </template>
              </el-table-column>
            </el-table>

            <el-pagination
              style="margin-top: 20px; justify-content: flex-end;"
              background
              layout="total, sizes, prev, pager, next"
              :total="bidProjectTotal"
              :page-sizes="[10, 20, 50]"
              v-model:current-page="bidProjectPage"
              v-model:page-size="bidProjectPageSize"
            ></el-pagination>
          </div>

          <!-- VIEW: 吧盛支付申请 -->
          <div v-else-if="currentMenu === 'payment-approval'">
            <div class="form-section-title">吧盛支付申请</div>
            <p style="color: #909399;">支付申请功能开发中...</p>
          </div>

          <!-- VIEW: 采购订单审批 -->
          <div v-else-if="currentMenu === 'order-approval'">
            <div class="form-section-title">采购订单审批</div>

            <!-- Tab切换 -->
            <el-tabs v-model="orderApprovalTab" style="margin-bottom: 20px;">
              <el-tab-pane label="待审批订单" name="pending"></el-tab-pane>
              <el-tab-pane label="已审批订单" name="approved"></el-tab-pane>
              <el-tab-pane label="新建采购订单" name="create"></el-tab-pane>
            </el-tabs>

            <!-- 待审批订单列表 -->
            <div v-if="orderApprovalTab === 'pending'">
              <el-alert title="审批时请注意核对采购价格是否超过协议限价，超限订单需特别审批" type="warning" show-icon style="margin-bottom: 15px;"></el-alert>

              <!-- 搜索条件 -->
              <div style="margin-bottom: 15px; padding: 15px; background-color: #fafafa; border-radius: 4px;">
                <el-form :inline="true" :model="orderApprovalSearch">
                  <el-form-item label="订单编号">
                    <el-input v-model="orderApprovalSearch.order_no" placeholder="订单编号" clearable style="width: 150px;"></el-input>
                  </el-form-item>
                  <el-form-item label="供应商">
                    <el-input v-model="orderApprovalSearch.supplier" placeholder="供应商名称" clearable style="width: 150px;"></el-input>
                  </el-form-item>
                  <el-form-item label="是否超限">
                    <el-select v-model="orderApprovalSearch.over_limit" placeholder="全部" clearable style="width: 100px;">
                      <el-option label="是" value="yes"></el-option>
                      <el-option label="否" value="no"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="searchPendingOrders">查询</el-button>
                    <el-button @click="resetOrderSearch">重置</el-button>
                  </el-form-item>
                </el-form>
              </div>

              <!-- 待审批订单列表 - 直接展示明细 -->
              <div v-for="order in pendingOrderList" :key="order.id" style="margin-bottom: 20px; border: 1px solid #ebeef5; border-radius: 8px; overflow: hidden;">
                <!-- 订单头部信息 -->
                <div style="padding: 15px 20px; background: linear-gradient(135deg, #f5f7fa 0%, #e4e7ed 100%); border-bottom: 1px solid #ebeef5;">
                  <div style="display: flex; justify-content: space-between; align-items: center; flex-wrap: wrap; gap: 10px;">
                    <div style="display: flex; align-items: center; gap: 20px; flex-wrap: wrap;">
                      <span style="font-weight: bold; font-size: 16px; color: #303133;">{{ order.order_no }}</span>
                      <el-tag v-if="order.over_limit_count > 0" type="danger" size="small">{{ order.over_limit_count }}项超限</el-tag>
                      <el-tag v-else type="success" size="small">价格正常</el-tag>
                    </div>
                    <div style="display: flex; gap: 10px;">
                      <el-button type="success" size="small" @click="approveOrder(order)">通过审批</el-button>
                      <el-button type="danger" size="small" @click="rejectOrder(order)">驳回</el-button>
                    </div>
                  </div>
                  <div style="display: flex; gap: 30px; margin-top: 10px; color: #606266; font-size: 14px; flex-wrap: wrap;">
                    <span><b>供应商：</b>{{ order.supplier_name }}</span>
                    <span><b>关联项目：</b>{{ order.project_name }}</span>
                    <span><b>订单总额：</b><span style="color: #409eff; font-weight: bold;">¥{{ order.total_amount }}</span></span>
                    <span><b>申请人：</b>{{ order.applicant }}</span>
                    <span><b>申请时间：</b>{{ order.apply_time }}</span>
                  </div>
                </div>
                <!-- 订单明细表格 - 直接展示 -->
                <div style="padding: 15px;">
                  <el-table :data="order.items" border size="small" style="width: 100%;">
                    <el-table-column label="序号" width="60" align="center">
                      <template #default="scope">{{ scope.$index + 1 }}</template>
                    </el-table-column>
                    <el-table-column prop="material_name" label="物料名称" min-width="120"></el-table-column>
                    <el-table-column prop="specification" label="规格型号" min-width="100"></el-table-column>
                    <el-table-column prop="quantity" label="数量" width="70" align="center"></el-table-column>
                    <el-table-column prop="unit" label="单位" width="60" align="center"></el-table-column>
                    <el-table-column prop="unit_price" label="采购单价" width="100">
                      <template #default="scope">
                        <span :style="{color: scope.row.over_limit ? '#f56c6c' : '#67c23a', fontWeight: 'bold'}">¥{{ scope.row.unit_price }}</span>
                      </template>
                    </el-table-column>
                    <el-table-column prop="agreement_price" label="协议价" width="90">
                      <template #default="scope">
                        <span v-if="scope.row.agreement_price" style="color: #409eff;">¥{{ scope.row.agreement_price }}</span>
                        <span v-else style="color: #999;">-</span>
                      </template>
                    </el-table-column>
                    <el-table-column prop="year_lowest_price" label="本年度最低价" width="110">
                      <template #default="scope">
                        <span v-if="scope.row.year_lowest_price" style="color: #67c23a;">¥{{ scope.row.year_lowest_price }}</span>
                        <span v-else style="color: #999;">-</span>
                      </template>
                    </el-table-column>
                    <el-table-column prop="last_year_price" label="上年度最后价" width="110">
                      <template #default="scope">
                        <span v-if="scope.row.last_year_price" style="color: #909399;">¥{{ scope.row.last_year_price }}</span>
                        <span v-else style="color: #999;">-</span>
                      </template>
                    </el-table-column>
                    <el-table-column label="价格状态" width="90" align="center">
                      <template #default="scope">
                        <el-tag v-if="scope.row.over_limit" type="danger" size="small">超限</el-tag>
                        <el-tag v-else type="success" size="small">正常</el-tag>
                      </template>
                    </el-table-column>
                    <el-table-column prop="total" label="小计" width="100">
                      <template #default="scope">
                        <span style="font-weight: bold;">¥{{ (scope.row.unit_price * scope.row.quantity).toFixed(2) }}</span>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
              </div>

              <!-- 空状态提示 -->
              <el-empty v-if="pendingOrderList.length === 0" description="暂无待审批订单"></el-empty>

              <el-pagination
                style="margin-top: 20px; justify-content: flex-end;"
                background
                layout="total, sizes, prev, pager, next"
                :total="pendingOrderTotal"
                :page-sizes="[10, 20, 50]"
                v-model:current-page="pendingOrderPage"
                v-model:page-size="pendingOrderPageSize"
              ></el-pagination>
            </div>

            <!-- 已审批订单列表 -->
            <div v-else-if="orderApprovalTab === 'approved'">
              <!-- 搜索条件 -->
              <div style="margin-bottom: 15px; padding: 15px; background-color: #fafafa; border-radius: 4px;">
                <el-form :inline="true" :model="approvedOrderSearch">
                  <el-form-item label="订单编号">
                    <el-input v-model="approvedOrderSearch.order_no" placeholder="订单编号" clearable style="width: 150px;"></el-input>
                  </el-form-item>
                  <el-form-item label="审批结果">
                    <el-select v-model="approvedOrderSearch.result" placeholder="全部" clearable style="width: 100px;">
                      <el-option label="已通过" value="approved"></el-option>
                      <el-option label="已驳回" value="rejected"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="审批时间">
                    <el-date-picker v-model="approvedOrderSearch.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" style="width: 240px;"></el-date-picker>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="searchApprovedOrders">查询</el-button>
                    <el-button @click="resetApprovedOrderSearch">重置</el-button>
                  </el-form-item>
                </el-form>
              </div>

              <!-- 已审批订单表格 -->
              <el-table :data="approvedOrderList" border style="width: 100%">
                <el-table-column prop="order_no" label="订单编号" width="150"></el-table-column>
                <el-table-column prop="supplier_name" label="供应商" width="150"></el-table-column>
                <el-table-column prop="total_amount" label="订单总额" width="120">
                  <template #default="scope">
                    <span style="color: #409eff; font-weight: bold;">¥{{ scope.row.total_amount }}</span>
                  </template>
                </el-table-column>
                <el-table-column prop="applicant" label="申请人" width="100"></el-table-column>
                <el-table-column prop="approver" label="审批人" width="100"></el-table-column>
                <el-table-column prop="result" label="审批结果" width="100">
                  <template #default="scope">
                    <el-tag v-if="scope.row.result === 'approved'" type="success">已通过</el-tag>
                    <el-tag v-else type="danger">已驳回</el-tag>
                  </template>
                </el-table-column>
                <el-table-column prop="approve_time" label="审批时间" width="160"></el-table-column>
                <el-table-column prop="remark" label="审批意见" min-width="150">
                  <template #default="scope">{{ scope.row.remark || '-' }}</template>
                </el-table-column>
                <el-table-column label="操作" width="100" fixed="right">
                  <template #default="scope">
                    <el-button type="primary" size="small" link @click="viewOrderDetail(scope.row)">查看详情</el-button>
                  </template>
                </el-table-column>
              </el-table>

              <el-pagination
                style="margin-top: 20px; justify-content: flex-end;"
                background
                layout="total, sizes, prev, pager, next"
                :total="approvedOrderTotal"
                :page-sizes="[10, 20, 50]"
                v-model:current-page="approvedOrderPage"
                v-model:page-size="approvedOrderPageSize"
              ></el-pagination>
            </div>

            <!-- 新建采购订单 -->
            <div v-else-if="orderApprovalTab === 'create'">
              <el-alert title="填写采购价时，系统将自动对比协议价、历史最低价进行提示" type="warning" show-icon style="margin-bottom: 20px;"></el-alert>

              <el-form :model="orderForm" label-width="150px" style="max-width: 900px">
                <el-form-item label="关联询价单">
                  <el-select v-model="orderForm.inquiry_no" placeholder="选择询价单（可选）" style="width: 300px;" clearable filterable>
                    <el-option v-for="item in inquiryOptions" :key="item.value" :label="item.label" :value="item.value"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item label="供应商">
                  <template #label><span class="red-star">*</span>供应商</template>
                  <el-select v-model="orderForm.supplier" placeholder="选择供应商" style="width: 300px;" filterable>
                    <el-option v-for="item in supplierOptions" :key="item.id" :label="item.name" :value="item.name"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item label="项目名称">
                  <el-input v-model="orderForm.project_name" placeholder="输入项目名称" style="width: 300px;"></el-input>
                </el-form-item>

                <el-divider content-position="left">采购明细</el-divider>

                <!-- 采购明细表格 -->
                <el-table :data="orderForm.items" border style="width: 100%; margin-bottom: 15px;">
                  <el-table-column label="序号" width="60" align="center">
                    <template #default="scope">{{ scope.$index + 1 }}</template>
                  </el-table-column>
                  <el-table-column label="物料名称" width="150">
                    <template #default="scope">
                      <el-input v-model="scope.row.material_name" placeholder="物料名称" size="small"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="规格型号" width="120">
                    <template #default="scope">
                      <el-input v-model="scope.row.specification" placeholder="规格型号" size="small"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="数量" width="100">
                    <template #default="scope">
                      <el-input-number v-model="scope.row.quantity" :min="1" size="small" style="width: 80px;" @change="calcOrderItemTotal(scope.row)"></el-input-number>
                    </template>
                  </el-table-column>
                  <el-table-column label="单位" width="80">
                    <template #default="scope">
                      <el-input v-model="scope.row.unit" placeholder="单位" size="small"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="采购单价" width="120">
                    <template #default="scope">
                      <el-input v-model="scope.row.unit_price" placeholder="单价" size="small" @input="calcOrderItemTotal(scope.row)">
                        <template #prepend>¥</template>
                      </el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="协议价" width="100">
                    <template #default="scope">
                      <span v-if="scope.row.agreement_price" style="color: #409eff;">¥{{ scope.row.agreement_price }}</span>
                      <span v-else style="color: #999;">-</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="价格状态" width="100">
                    <template #default="scope">
                      <template v-if="scope.row.unit_price && scope.row.agreement_price">
                        <el-tag v-if="Number(scope.row.unit_price) > scope.row.agreement_price" type="danger" size="small">超限</el-tag>
                        <el-tag v-else type="success" size="small">正常</el-tag>
                      </template>
                      <span v-else style="color: #999;">-</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="小计" width="100">
                    <template #default="scope">
                      <span style="font-weight: bold;">¥{{ scope.row.total || 0 }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="80">
                    <template #default="scope">
                      <el-button type="danger" size="small" link @click="removeOrderItem(scope.$index)">删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>

                <el-button type="primary" size="small" @click="addOrderItem" style="margin-bottom: 20px;">+ 添加物料</el-button>

                <el-divider content-position="left">订单汇总</el-divider>

                <el-form-item label="订单总额">
                  <span style="color: #f56c6c; font-size: 18px; font-weight: bold;">¥ {{ calcOrderTotal() }}</span>
                </el-form-item>
                <el-form-item label="备注">
                  <el-input v-model="orderForm.remark" type="textarea" :rows="3" placeholder="输入备注信息" style="width: 400px;"></el-input>
                </el-form-item>

                <el-form-item>
                  <el-button type="primary" @click="submitOrder">提交审批</el-button>
                  <el-button @click="resetOrderForm">重置</el-button>
                </el-form-item>
              </el-form>
            </div>

            <!-- 审批对话框 -->
            <el-dialog v-model="orderApprovalDialogVisible" :title="orderApprovalAction === 'approve' ? '审批通过' : '审批驳回'" width="500px">
              <el-form :model="orderApprovalForm" label-width="100px">
                <el-form-item label="订单编号">
                  <el-input v-model="orderApprovalForm.order_no" disabled></el-input>
                </el-form-item>
                <el-form-item label="订单总额">
                  <span style="color: #409eff; font-weight: bold;">¥{{ orderApprovalForm.total_amount }}</span>
                </el-form-item>
                <el-form-item label="超限项数">
                  <el-tag v-if="orderApprovalForm.over_limit_count > 0" type="danger">{{ orderApprovalForm.over_limit_count }}项超限</el-tag>
                  <el-tag v-else type="success">无超限</el-tag>
                </el-form-item>
                <el-form-item :label="orderApprovalAction === 'approve' ? '审批意见' : '驳回原因'">
                  <template #label><span v-if="orderApprovalAction === 'reject'" class="red-star">*</span>{{ orderApprovalAction === 'approve' ? '审批意见' : '驳回原因' }}</template>
                  <el-input v-model="orderApprovalForm.remark" type="textarea" :rows="3" :placeholder="orderApprovalAction === 'approve' ? '可选填审批意见' : '请输入驳回原因'"></el-input>
                </el-form-item>
              </el-form>
              <template #footer>
                <el-button @click="orderApprovalDialogVisible = false">取消</el-button>
                <el-button :type="orderApprovalAction === 'approve' ? 'success' : 'danger'" @click="confirmOrderApproval">
                  {{ orderApprovalAction === 'approve' ? '确认通过' : '确认驳回' }}
                </el-button>
              </template>
            </el-dialog>
          </div>

          <!-- VIEW: 协议价管理 -->
          <div v-else-if="currentMenu === 'agreement-price'">
            <div class="form-section-title">协议价管理</div>

            <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
              <div>
                <el-input placeholder="搜索物料编码/名称" style="width: 250px; margin-right: 10px;" v-model="agreementSearch.keyword"></el-input>
                <el-button type="primary">搜索</el-button>
              </div>
              <div>
                <el-button type="success" @click="showAddAgreementDialog">+ 新增协议价</el-button>
                <el-button type="warning">Excel导入</el-button>
              </div>
            </div>

            <el-table :data="agreementPriceData" border style="width: 100%">
              <el-table-column prop="material_code" label="物料编码" width="120" />
              <el-table-column prop="material_name" label="物料名称" min-width="150" />
              <el-table-column prop="supplier" label="供应商" width="150" />
              <el-table-column prop="price" label="协议价" width="100">
                <template #default="{ row }">
                  <span style="color: #409eff; font-weight: bold;">¥ {{ row.price }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="valid_date" label="有效期至" width="120" />
              <el-table-column label="操作" width="150" fixed="right">
                <template #default>
                  <el-button type="primary" size="small">编辑</el-button>
                  <el-button type="danger" size="small">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>

          <!-- VIEW: 潜在供应商列表 -->
          <div v-else-if="currentMenu === 'temp'">
            <div class="form-section-title">潜在供应商台账</div>

            <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
              <div>
                <el-input placeholder="搜索企业名称/联系人" style="width: 250px; margin-right: 10px;" prefix-icon="Search" v-model="searchForm.name"></el-input>
                <el-button type="primary" @click="onSearch">搜索</el-button>
              </div>
              <div>
                <el-button type="success" @click="exportExcel">导出 Excel</el-button>
                <el-button type="primary" @click="currentMenu = 'apply'; supplierApplyType = 'temp';">+ 新增潜在供应商</el-button>
              </div>
            </div>

            <el-table :data="tempSupplierData" border style="width: 100%">
              <el-table-column prop="id" label="序号" width="60" fixed />
              <el-table-column label="企业名称" min-width="180">
                <template #header>
                  <span style="color: #f56c6c;">*</span> 企业名称
                </template>
                <template #default="{ row }">{{ row.enterprise_name }}</template>
              </el-table-column>
              <el-table-column prop="credit_code" label="统一社会信用代码" width="180" />
              <el-table-column label="供应商类型" width="100">
                <template #header>
                  <span style="color: #f56c6c;">*</span> 供应商类型
                </template>
                <template #default="{ row }">{{ row.entry_type }}</template>
              </el-table-column>
              <el-table-column label="供应品类" width="100">
                <template #header>
                  <span style="color: #f56c6c;">*</span> 供应品类
                </template>
                <template #default="{ row }">{{ row.category }}</template>
              </el-table-column>
              <el-table-column prop="region" label="合作区域" width="100" />
              <el-table-column label="合作行业" width="90">
                <template #header>
                  <span style="color: #f56c6c;">*</span> 合作行业
                </template>
                <template #default="{ row }">{{ row.industry }}</template>
              </el-table-column>
              <el-table-column label="合作品牌" width="100">
                <template #header>
                  <span style="color: #f56c6c;">*</span> 合作品牌
                </template>
                <template #default="{ row }">{{ row.brand }}</template>
              </el-table-column>
              <el-table-column label="联系人" width="80">
                <template #header>
                  <span style="color: #f56c6c;">*</span> 联系人
                </template>
                <template #default="{ row }">{{ row.contact_person }}</template>
              </el-table-column>
              <el-table-column label="联系电话" width="120">
                <template #header>
                  <span style="color: #f56c6c;">*</span> 联系电话
                </template>
                <template #default="{ row }">{{ row.mobile }}</template>
              </el-table-column>
              <el-table-column prop="bank_name" label="开户行" width="100" />
              <el-table-column prop="branch_name" label="开户行名称" width="130" />
              <el-table-column prop="settlement" label="结算方式" width="100" />
              <el-table-column prop="purchaser" label="采购员" width="80" />
              <el-table-column label="操作" width="80" fixed="right">
                <template #default="{ row }">
                  <el-button type="primary" size="small" link @click="viewSupplierDetail(row)">详情</el-button>
                </template>
              </el-table-column>
            </el-table>

            <div style="margin-top: 20px; text-align: right;">
              <el-pagination background layout="prev, pager, next, sizes, total" :total="tempSupplierData.length"></el-pagination>
            </div>
          </div>

          <!-- VIEW: 合格供应商列表 -->
          <div v-else-if="currentMenu === 'qualified'">
            <div class="form-section-title">合格供应商台账</div>

            <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
              <div>
                <el-input placeholder="搜索企业名称/联系人" style="width: 250px; margin-right: 10px;" prefix-icon="Search" v-model="searchForm.name"></el-input>
                <el-button type="primary" @click="onSearch">搜索</el-button>
              </div>
              <div>
                <el-button type="success" @click="exportExcel">导出 Excel</el-button>
                <el-button type="primary" @click="currentMenu = 'apply'; isTempAdd = false;">+ 新增合格供应商</el-button>
              </div>
            </div>

            <el-table :data="qualifiedSupplierData" border style="width: 100%">
              <el-table-column prop="id" label="序号" width="70" fixed />
              <el-table-column prop="enterprise_name" label="企业名称" min-width="180" fixed />
              <el-table-column prop="credit_code" label="统一社会信用代码" width="180" />
              <el-table-column prop="entry_type" label="供应商类型" width="110" />
              <el-table-column prop="category" label="供应品类" width="100" />
              <el-table-column prop="region" label="合作区域" width="100" />
              <el-table-column prop="industry" label="合作行业" width="100" />
              <el-table-column prop="brand" label="合作品牌" width="120" />
              <el-table-column prop="contact_person" label="联系人" width="100" />
              <el-table-column prop="mobile" label="联系电话" width="130" />
              <el-table-column prop="bank_name" label="开户行" width="150" />
              <el-table-column prop="branch_name" label="开户行名称" width="180" />
              <el-table-column prop="settlement" label="结算方式" width="100" />
              <el-table-column prop="purchaser" label="采购员" width="100" />
              <el-table-column label="操作" width="120" fixed="right">
                <template #default="{ row }">
                  <el-button type="primary" size="small" @click="viewSupplierDetail(row)">详情</el-button>
                </template>
              </el-table-column>
            </el-table>

            <div style="margin-top: 20px; text-align: right;">
              <el-pagination background layout="prev, pager, next, sizes, total" :total="qualifiedSupplierData.length"></el-pagination>
            </div>
          </div>

          <!-- 默认页面 -->
          <div v-else>
            <div class="form-section-title">{{ getCurrentTitle() }}</div>
            <p style="color: #909399;">该功能正在开发中...</p>
          </div>

        </div>
      </div>
    </div>

    <div class="demo-tips" v-if="showTips">
      <h4>功能更新说明</h4>
      <p>1. 侧边栏已更新为5项，包含<b>"供应商变更"</b>。</p>
      <p>2. "供应商申请"已更新字段（红色星号为必填）。</p>
      <p>3. "供应商变更"页支持筛选和信息修改/拉黑。</p>
      <el-button size="small" @click="showTips = false">知道了</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Menu, UploadFilled, Document, Delete, Edit, Search, Shop, Check, Setting, List, Trophy, Money, Warning, WarningFilled } from '@element-plus/icons-vue'
import service from '@/utils/request' // Use the project's configured axios instance

const currentMenu = ref('apply')
const activeTab = ref('basic')
const showTips = ref(true)
const changeDialogVisible = ref(false)
const isAdmin = ref(true) // TODO: 根据实际用户角色判断

// Certificate Logic
// const currentCertType = ref('') // Removed
const uploadedFiles = ref([])

// 特殊行业文件
const specialFiles = reactive({
  quality_cert: [],
  security_cert: [],
  other_cert: []
})

const form = reactive({
  enterprise_name: '杭州森森休闲用品有限公司',
  credit_code: '913301020743241989',
  entry_type: 'manufacturing',
  category: 'raw',
  region: '华东区',
  industry: 'petrochemical',
  brand: '森森',
  contact_person: '俞志勇',
  mobile: '19846768980',
  purchaser: 'Admin',
  email: '19846768980@126.com',
  settlement: 'm30',
  bank_name: '',
  branch_name: '',
  bank_account: ''
})

const searchForm = reactive({
  name: '',
  status: '',
  approvalStatus: '',
  industry: ''
})

const changeForm = reactive({
  id: '',
  enterprise_name: '',
  credit_code: '',
  entry_type: '',
  category: '',
  region: '',
  industry: '',
  brand: '',
  contact_person: '',
  mobile: '',
  bank_name: '',
  branch_name: '',
  bank_account: '',
  settlement: '',
  purchaser: '',
  reason: '',
  status: '',
  toBlacklist: false
})

const isTempAdd = ref(false)
const isApprovalMode = ref(false) // 是否为审批查看模式
const supplierApplyType = ref('qualified') // 'temp' 潜在供应商, 'qualified' 合格供应商
const tableData = ref([])
const importDialogVisible = ref(false)
const techFiles = ref([])

// 当前用户和时间
const currentUser = computed(() => 'Admin')
const currentTime = computed(() => new Date().toLocaleString())

// 询价 Tab 控制
const inquiryTab = ref('entry')
const inquiryFormRef = ref(null)

// 询价表单 - 更新字段
const inquiryForm = reactive({
  project_name: '',
  end_user: '',
  quote_deadline: '',
  expected_delivery: '',
  technical_requirements: '',
  items: [{ brand: '', product_name: '', specification: '', quantity: 1, unit: '个', medium: '', diameter: '', material: '', remark: '' }]
})

// 询价搜索表单
const inquirySearchForm = reactive({
  project_name: '',
  status: ''
})

// 我的询价列表数据
const myInquiryList = ref([
  { id: 1, inquiry_no: 'INQ202411270001', project_name: '某化工厂阀门采购项目', end_user: '某化工集团', quote_deadline: '2024-11-30 18:00', material_summary: '闸阀、球阀等', status: 'waiting_feedback', project_result: null, created_at: '2024-11-20 10:33:09', feedback_deadline: '2024-11-27 10:33:09', quote_date: '2024-11-20' },
  { id: 2, inquiry_no: 'INQ202411270002', project_name: '电厂管道配件项目', end_user: '某电力公司', quote_deadline: '2024-12-05 17:00', material_summary: '法兰、垫片', status: 'clarifying', project_result: null, created_at: '2024-11-22 14:20:00', feedback_deadline: null, quote_date: null },
  { id: 3, inquiry_no: 'INQ202411270003', project_name: '污水处理厂设备采购', end_user: '某环保科技公司', quote_deadline: '2024-12-10 12:00', material_summary: '蝶阀', status: 'pending', project_result: null, created_at: '2024-11-25 09:15:00', feedback_deadline: null, quote_date: null },
  { id: 4, inquiry_no: 'INQ202411150001', project_name: '国电阀门采购项目', end_user: '国家电网', quote_deadline: '2024-11-20 17:00', material_summary: '闸阀DN100', status: 'won', project_result: 'won', created_at: '2024-11-10 09:00:00', feedback_deadline: null, quote_date: '2024-11-13' },
  { id: 5, inquiry_no: 'INQ202411100002', project_name: '石化管件项目', end_user: '中石化', quote_deadline: '2024-11-15 12:00', material_summary: '法兰、弯头', status: 'lost', project_result: 'lost', created_at: '2024-11-05 14:30:00', feedback_deadline: null, quote_date: '2024-11-08' }
])
const myInquiryTotal = ref(5)
const myInquiryPage = ref(1)
const myInquiryPageSize = ref(10)

// 中标反馈对话框（7天规则）
const feedbackDialogVisible = ref(false)
const feedbackFormRef = ref(null)
const feedbackForm = reactive({
  inquiry_id: null,
  inquiry_no: '',
  project_name: '',
  quote_date: '',
  result: '',
  deal_amount: 0,
  order_no: '',
  no_deal_reason: '',
  delay_reason: '',
  remark: ''
})
const calculationFiles = ref([])

// 澄清回复对话框
const clarifyDialogVisible = ref(false)
const clarifyHistory = ref([])
const clarifyReplyForm = reactive({ inquiry_id: null, content: '' })
const clarifyFiles = ref([])

// ========== 采购部分 ==========
// 采购搜索表单
const purchaseSearchForm = reactive({ inquiry_no: '', project_name: '', status: '' })

// 采购待处理列表
const purchaseInquiryList = ref([
  {
    id: 1,
    inquiry_no: 'INQ202411270001',
    project_name: '某化工厂阀门采购项目',
    end_user: '某化工集团',
    sales_name: '张三',
    material_summary: '闸阀、球阀等',
    quote_deadline: '2024-11-30 18:00',
    quote_date: '2024-11-20',
    status: 'pending',
    created_at: '2024-11-20 10:33:09',
    items: [
      { brand: 'KITZ', product_name: '闸阀', specification: 'DN50 PN16', quantity: 10, unit: '个', medium: '水', diameter: 'DN50', material: '不锈钢304', remark: '需要提供质保书', project_result: null, unit_price: 120, delivery_days: 15, quote_remark: '现货供应', supplier_name: '上海阀门有限公司' },
      { brand: 'KITZ', product_name: '球阀', specification: 'DN32 PN25', quantity: 20, unit: '个', medium: '蒸汽', diameter: 'DN32', material: '碳钢', remark: '', project_result: null, unit_price: 85, delivery_days: 7, quote_remark: '', supplier_name: '上海阀门有限公司' }
    ]
  },
  {
    id: 2,
    inquiry_no: 'INQ202411270002',
    project_name: '电厂管道配件项目',
    end_user: '某电力公司',
    sales_name: '李四',
    material_summary: '法兰、管件',
    quote_deadline: '2024-12-05 17:00',
    quote_date: '2024-11-22',
    status: 'can_quote',
    created_at: '2024-11-22 14:20:00',
    items: [
      { brand: '', product_name: '法兰', specification: 'DN100 PN16', quantity: 50, unit: '片', medium: '水', diameter: 'DN100', material: '碳钢Q235', remark: '带螺栓螺母', project_result: null, unit_price: null, delivery_days: null, quote_remark: '', supplier_name: '' },
      { brand: '', product_name: '弯头', specification: '90° DN100', quantity: 30, unit: '个', medium: '水', diameter: 'DN100', material: '碳钢', remark: '', project_result: null, unit_price: null, delivery_days: null, quote_remark: '', supplier_name: '' },
      { brand: '', product_name: '三通', specification: 'DN100', quantity: 10, unit: '个', medium: '水', diameter: 'DN100', material: '碳钢', remark: '等径三通', project_result: null, unit_price: null, delivery_days: null, quote_remark: '', supplier_name: '' }
    ]
  },
  {
    id: 3,
    inquiry_no: 'INQ202411270003',
    project_name: '污水处理厂设备采购',
    end_user: '某环保科技公司',
    sales_name: '王五',
    material_summary: '蝶阀、止回阀',
    quote_deadline: '2024-12-10 12:00',
    quote_date: '2024-11-25',
    status: 'quoted',
    created_at: '2024-11-25 09:15:00',
    items: [
      { brand: '中核苏阀', product_name: '蝶阀', specification: 'DN200 PN10', quantity: 5, unit: '台', medium: '污水', diameter: 'DN200', material: '铸铁', remark: '电动执行器', project_result: 'won', unit_price: 2800, delivery_days: 25, quote_remark: '含电动头', supplier_name: '广州五金贸易公司' },
      { brand: '', product_name: '止回阀', specification: 'DN150 PN10', quantity: 8, unit: '台', medium: '污水', diameter: 'DN150', material: '铸铁', remark: '', project_result: 'won', unit_price: 650, delivery_days: 20, quote_remark: '旋启式', supplier_name: '广州五金贸易公司' }
    ]
  }
])
const purchaseInquiryTotal = ref(3)
const purchaseInquiryPage = ref(1)
const purchaseInquiryPageSize = ref(10)

// 采购澄清对话框
const purchaseClarifyDialogVisible = ref(false)
const purchaseClarifyForm = reactive({ inquiry_id: null, inquiry_no: '', content: '' })
const purchaseClarifyFiles = ref([])

// 采购报价对话框
const purchaseQuoteDialogVisible = ref(false)
const purchaseQuoteForm = reactive({ inquiry_id: null, inquiry_no: '', project_name: '', end_user: '', quote_deadline: '', expected_delivery: '', technical_requirements: '', items: [], remark: '' })

// 批量导入报价对话框
const importQuoteDialogVisible = ref(false)

// 供应商选项
const supplierOptions = ref([
  { id: 1, name: '上海阀门有限公司' },
  { id: 2, name: '北京管件制造厂' },
  { id: 3, name: '广州五金贸易公司' },
  { id: 4, name: '深圳工业设备公司' }
])

// ========== 台账与价格库部分 ==========
const ledgerTab = ref('ledger')

// 台账搜索
const ledgerSearch = reactive({ dateRange: [], material: '', supplier: '', project: '', deal_status: '' })

// 台账数据
const ledgerData = ref([
  { inquiry_no: 'INQ202411270001', project_name: '某化工厂阀门采购', end_user: '某化工集团', material: '闸阀', specification: 'DN50 PN16', quantity: 10, supplier: '上海阀门有限公司', unit_price: '¥120', total_price: '¥1,200', history_lowest: '¥115', agreement_price: '¥110', deal_status: '成交', order_no: 'SO2024112701', quote_date: '2024-11-25' },
  { inquiry_no: 'INQ202411270002', project_name: '电厂管道配件项目', end_user: '某电力公司', material: '法兰', specification: 'DN100 PN16', quantity: 50, supplier: '北京管件制造厂', unit_price: '¥45', total_price: '¥2,250', history_lowest: '¥42', agreement_price: '¥40', deal_status: '未成交', order_no: '-', quote_date: '2024-11-26' },
  { inquiry_no: 'INQ202411270003', project_name: '污水处理厂设备采购', end_user: '某环保公司', material: '蝶阀', specification: 'DN200', quantity: 5, supplier: '广州五金贸易公司', unit_price: '¥850', total_price: '¥4,250', history_lowest: '¥800', agreement_price: '-', deal_status: '待定', order_no: '-', quote_date: '2024-11-27' }
])
const ledgerTotal = ref(3)
const ledgerPage = ref(1)
const ledgerPageSize = ref(10)

// 价格库搜索
const priceLibSearch = reactive({ material: '', supplier: '', project: '', dateRange: [] })

// 价格库数据
const priceLibData = ref([
  { material: '闸阀', specification: 'DN50 PN16', supplier: '上海阀门有限公司', unit_price: '¥120', project_name: '某化工厂阀门采购', quote_date: '2024-11-25', is_deal: true, remark: '' },
  { material: '法兰', specification: 'DN100 PN16', supplier: '北京管件制造厂', unit_price: '¥45', project_name: '电厂管道配件项目', quote_date: '2024-11-26', is_deal: false, remark: '价格偏高' },
  { material: '蝶阀', specification: 'DN200', supplier: '广州五金贸易公司', unit_price: '¥850', project_name: '污水处理厂设备采购', quote_date: '2024-11-27', is_deal: false, remark: '待定中' },
  { material: '球阀', specification: 'DN25 PN25', supplier: '上海阀门有限公司', unit_price: '¥65', project_name: '某化工厂阀门采购', quote_date: '2024-11-25', is_deal: true, remark: '' }
])
const priceLibTotal = ref(4)
const priceLibPage = ref(1)
const priceLibPageSize = ref(10)

// ========== 中标项目部分 ==========
const bidProjectSearch = reactive({ project_name: '', customer: '', dateRange: [] })

// 中标项目数据
const bidProjectsData = ref([
  { inquiry_no: 'INQ202411270001', project_name: '某化工厂阀门采购项目', end_user: '某化工集团', sales_name: '张三', win_date: '2024-11-26', total_amount: '¥15,800', order_no: 'SO2024112701', calculation_file: 'calc_001.xlsx' },
  { inquiry_no: 'INQ202411150002', project_name: '国电管件供应项目', end_user: '国家电网', sales_name: '李四', win_date: '2024-11-15', total_amount: '¥32,000', order_no: 'SO2024111501', calculation_file: 'calc_002.xlsx' }
])
const bidProjectTotal = ref(2)
const bidProjectPage = ref(1)
const bidProjectPageSize = ref(10)

// ========== 采购订单审批部分 ==========
const orderApprovalTab = ref('pending')

// 待审批订单搜索
const orderApprovalSearch = reactive({ order_no: '', supplier: '', over_limit: '' })

// 待审批订单列表
// 系统自动带出：协议价、本年度最低采购价、上年度最后采购价
const pendingOrderList = ref([
  {
    id: 1,
    order_no: 'PO202411270001',
    supplier_name: '上海阀门有限公司',
    project_name: '某化工厂阀门采购项目',
    total_amount: '15,800.00',
    over_limit_count: 1,
    applicant: '张三',
    apply_time: '2024-11-25 14:30:00',
    items: [
      { material_name: '闸阀', specification: 'DN50 PN16', quantity: 10, unit: '个', unit_price: 130, agreement_price: 120, year_lowest_price: 115, last_year_price: 118, over_limit: true },
      { material_name: '球阀', specification: 'DN32 PN25', quantity: 20, unit: '个', unit_price: 85, agreement_price: 90, year_lowest_price: 82, last_year_price: 88, over_limit: false },
      { material_name: '止回阀', specification: 'DN50 PN16', quantity: 15, unit: '个', unit_price: 75, agreement_price: 80, year_lowest_price: 70, last_year_price: 78, over_limit: false }
    ]
  },
  {
    id: 2,
    order_no: 'PO202411270002',
    supplier_name: '北京管件制造厂',
    project_name: '电厂管道配件项目',
    total_amount: '8,500.00',
    over_limit_count: 0,
    applicant: '李四',
    apply_time: '2024-11-26 09:15:00',
    items: [
      { material_name: '法兰', specification: 'DN100 PN16', quantity: 50, unit: '片', unit_price: 42, agreement_price: 45, year_lowest_price: 40, last_year_price: 43, over_limit: false },
      { material_name: '弯头', specification: '90° DN100', quantity: 30, unit: '个', unit_price: 35, agreement_price: 38, year_lowest_price: 33, last_year_price: 36, over_limit: false }
    ]
  },
  {
    id: 3,
    order_no: 'PO202411270003',
    supplier_name: '广州五金贸易公司',
    project_name: '污水处理厂设备采购',
    total_amount: '28,600.00',
    over_limit_count: 2,
    applicant: '王五',
    apply_time: '2024-11-26 16:45:00',
    items: [
      { material_name: '蝶阀', specification: 'DN200 PN10', quantity: 5, unit: '台', unit_price: 3200, agreement_price: 2800, year_lowest_price: 2750, last_year_price: 2900, over_limit: true },
      { material_name: '止回阀', specification: 'DN150 PN10', quantity: 8, unit: '台', unit_price: 720, agreement_price: 650, year_lowest_price: 620, last_year_price: 680, over_limit: true }
    ]
  }
])
const pendingOrderTotal = ref(3)
const pendingOrderPage = ref(1)
const pendingOrderPageSize = ref(10)

// 已审批订单搜索
const approvedOrderSearch = reactive({ order_no: '', result: '', dateRange: [] })

// 已审批订单列表
const approvedOrderList = ref([
  { id: 1, order_no: 'PO202411200001', supplier_name: '上海阀门有限公司', total_amount: '12,000.00', applicant: '张三', approver: '审批员A', result: 'approved', approve_time: '2024-11-21 10:30:00', remark: '价格合理，同意采购' },
  { id: 2, order_no: 'PO202411200002', supplier_name: '深圳工业设备公司', total_amount: '5,600.00', applicant: '李四', approver: '审批员A', result: 'rejected', approve_time: '2024-11-22 14:15:00', remark: '价格过高，请重新询价' },
  { id: 3, order_no: 'PO202411180001', supplier_name: '北京管件制造厂', total_amount: '9,800.00', applicant: '王五', approver: '审批员B', result: 'approved', approve_time: '2024-11-19 09:00:00', remark: '' }
])
const approvedOrderTotal = ref(3)
const approvedOrderPage = ref(1)
const approvedOrderPageSize = ref(10)

// 采购订单表单（新建）
const orderForm = reactive({
  inquiry_no: '',
  supplier: '',
  project_name: '',
  items: [
    { material_name: '', specification: '', quantity: 1, unit: '', unit_price: '', agreement_price: null, total: 0 }
  ],
  remark: ''
})

// 询价单选项
const inquiryOptions = ref([
  { value: 'INQ202411270001', label: 'INQ202411270001 - 某化工厂阀门采购项目' },
  { value: 'INQ202411270002', label: 'INQ202411270002 - 电厂管道配件项目' },
  { value: 'INQ202411270003', label: 'INQ202411270003 - 污水处理厂设备采购' }
])

// 审批对话框
const orderApprovalDialogVisible = ref(false)
const orderApprovalAction = ref('approve')
const orderApprovalForm = reactive({ order_no: '', total_amount: '', over_limit_count: 0, remark: '' })

// 协议价搜索
const agreementSearch = reactive({ keyword: '' })

// 协议价数据
const agreementPriceData = ref([
  { material_code: 'M1001', material_name: '蝶阀DN100', supplier: 'A公司', price: 110, valid_date: '2025-01-01' },
  { material_code: 'M1002', material_name: '闸阀DN50', supplier: 'B公司', price: 85, valid_date: '2025-06-30' },
  { material_code: 'M1003', material_name: '球阀DN80', supplier: 'A公司', price: 95, valid_date: '2024-12-31' }
])

// 潜在供应商数据（全字段：序号/企业名称/统一社会信用代码/供应商类型/供应品类/合作区域/合作行业/合作品牌/联系人/联系电话/开户行/开户行名称/结算方式/采购员）
const tempSupplierData = ref([
  {
    id: 1,
    enterprise_name: '深圳华创电子科技有限公司',
    credit_code: '91440300MA5DLXXX01',
    entry_type: '生产商',
    category: '原材料',
    region: '华南区',
    industry: '核电',
    brand: '华创',
    contact_person: '张三',
    mobile: '13800138001',
    bank_name: '中国银行',
    branch_name: '深圳南山支行',
    settlement: '月结30天',
    purchaser: 'Admin'
  },
  {
    id: 2,
    enterprise_name: '北京中科元器件贸易有限公司',
    credit_code: '91110105MA01DXXX02',
    entry_type: '贸易商',
    category: '电子元器件',
    region: '华北区',
    industry: '军工',
    brand: '中科',
    contact_person: '李四',
    mobile: '13800138002',
    bank_name: '工商银行',
    branch_name: '北京海淀支行',
    settlement: '月结45天',
    purchaser: '王经理'
  },
  {
    id: 3,
    enterprise_name: '广州南方办公设备有限公司',
    credit_code: '91440101MA5CLXXX03',
    entry_type: '生产商',
    category: '办公用品',
    region: '华南区',
    industry: '石化',
    brand: '南方',
    contact_person: '王五',
    mobile: '13800138003',
    bank_name: '建设银行',
    branch_name: '广州天河支行',
    settlement: '月结30天',
    purchaser: 'Admin'
  }
])

// 合格供应商数据（字段：序号/企业名称/统一社会信用代码/供应商类型/供应品类/合作区域/合作行业/合作品牌/联系人/联系电话/开户行/开户行名称/结算方式/采购员）
const qualifiedSupplierData = ref([
  {
    id: 1,
    enterprise_name: '杭州森森休闲用品有限公司',
    credit_code: '913301020743241989',
    entry_type: '生产商',
    category: '原材料',
    region: '华东区',
    industry: '石化',
    brand: '森森',
    contact_person: '俞志勇',
    mobile: '19846768980',
    bank_name: '中国银行',
    branch_name: '杭州西湖支行',
    settlement: '月结30天',
    purchaser: 'Admin'
  },
  {
    id: 2,
    enterprise_name: '上海阀门制造有限公司',
    credit_code: '91310000123456789A',
    entry_type: '生产商',
    category: '原材料',
    region: '华东区',
    industry: '核电',
    brand: '沪阀',
    contact_person: '张明',
    mobile: '13900139001',
    bank_name: '工商银行',
    branch_name: '上海浦东支行',
    settlement: '预付',
    purchaser: '采购员A'
  },
  {
    id: 3,
    enterprise_name: '北京精密仪表有限公司',
    credit_code: '91110000987654321B',
    entry_type: '贸易商',
    category: '电子元器件',
    region: '华北区',
    industry: '军工',
    brand: '京仪',
    contact_person: '李华',
    mobile: '13900139002',
    bank_name: '建设银行',
    branch_name: '北京朝阳支行',
    settlement: '货到付款',
    purchaser: '采购员B'
  }
])
const approvalData = ref([
  {
    id: 1,
    code: 'GYS20231126001',
    name: '上海核电设备有限公司',
    type: '原材料供应商',
    industry: 'nuclear',
    status: '待审批',
    current_node: '质保部审核',
    created_time: '2023-11-26 10:00:00',
    unified_social_credit_code: '91310000123456789A',
    legal_representative: '张三',
    registered_capital: '1000万元',
    registered_address: '上海市浦东新区张江高科技园区',
    contact_person: '李四',
    contacts: [
      { name: '李四', department: '采购部', position: '采购经理', mobile: '13800138000', email: 'lisi@example.com' }
    ],
    banks: [
      { account_name: '上海核电设备有限公司', account_number: '1234567890123456789', bank_name: '中国银行股份有限公司上海浦东分行', branch_name: '浦东分行营业部' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-26 10:00:00', approver: '采购员-王明', comment: '资料齐全，提交审批' },
      { node: '采购部长审批', status: 'completed', time: '2023-11-26 14:30:00', approver: '采购部长-刘华', comment: '同意，转质保部审核' },
      { node: '质保部审核', status: 'current', time: '', approver: '', comment: '' },
      { node: '副总审批', status: 'pending', time: '', approver: '', comment: '' }
    ]
  },
  {
    id: 2,
    code: 'GYS20231126002',
    name: '北京军工材料供应商',
    type: '设备供应商',
    industry: 'military',
    status: '待审批',
    current_node: '采购部长审批',
    created_time: '2023-11-26 11:00:00',
    unified_social_credit_code: '91310000987654321B',
    legal_representative: '王五',
    registered_capital: '500万元',
    registered_address: '北京市朝阳区军工园区',
    contact_person: '赵六',
    contacts: [
      { name: '赵六', department: '销售部', position: '销售经理', mobile: '13900139000', email: 'zhaoliu@example.com' }
    ],
    banks: [
      { account_name: '北京军工材料供应商', account_number: '9876543210987654321', bank_name: '中国工商银行股份有限公司北京朝阳支行', branch_name: '朝阳支行营业部' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-26 11:00:00', approver: '采购员-张强', comment: '材料已核实' },
      { node: '采购部长审批', status: 'current', time: '', approver: '', comment: '' },
      { node: '质保部审核', status: 'pending', time: '', approver: '', comment: '' },
      { node: '副总审批', status: 'pending', time: '', approver: '', comment: '' }
    ]
  },
  {
    id: 3,
    code: 'GYS20231126003',
    name: '杭州普通贸易有限公司',
    type: '贸易商',
    industry: 'normal',
    status: '待审批',
    current_node: '副总审批',
    created_time: '2023-11-25 09:00:00',
    unified_social_credit_code: '91330000111222333C',
    legal_representative: '陈明',
    registered_capital: '200万元',
    registered_address: '杭州市西湖区文三路',
    contact_person: '周杰',
    contacts: [
      { name: '周杰', department: '商务部', position: '商务经理', mobile: '13700137000', email: 'zhoujie@example.com' }
    ],
    banks: [
      { account_name: '杭州普通贸易有限公司', account_number: '1111222233334444555', bank_name: '招商银行股份有限公司杭州分行', branch_name: '西湖支行' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-25 09:00:00', approver: '采购员-李伟', comment: '提交申请' },
      { node: '采购部长审批', status: 'completed', time: '2023-11-25 15:00:00', approver: '采购部长-刘华', comment: '资质符合，同意' },
      { node: '副总审批', status: 'current', time: '', approver: '', comment: '' }
    ]
  },
  {
    id: 4,
    code: 'GYS20231124001',
    name: '深圳电子元器件公司',
    type: '生产商',
    industry: 'normal',
    status: '已通过',
    current_node: '已完成',
    created_time: '2023-11-24 08:30:00',
    unified_social_credit_code: '91440000444555666D',
    legal_representative: '黄强',
    registered_capital: '800万元',
    registered_address: '深圳市南山区科技园',
    contact_person: '林芳',
    contacts: [
      { name: '林芳', department: '业务部', position: '业务主管', mobile: '13600136000', email: 'linfang@example.com' }
    ],
    banks: [
      { account_name: '深圳电子元器件公司', account_number: '4444555566667777888', bank_name: '平安银行股份有限公司深圳分行', branch_name: '南山支行' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-24 08:30:00', approver: '采购员-王明', comment: '信息完整' },
      { node: '采购部长审批', status: 'completed', time: '2023-11-24 14:00:00', approver: '采购部长-刘华', comment: '审核通过' },
      { node: '副总审批', status: 'completed', time: '2023-11-24 17:30:00', approver: '副总-赵总', comment: '同意准入' }
    ]
  }
])

// 供应商档案库数据
const supplierSearch = reactive({
  keyword: '',
  industry: '',
  status: '',
  supplierType: ''
})

const supplierList = reactive({
  total: 100,
  page: 1,
  pageSize: 10
})

const supplierListData = ref([
  {
    id: 1,
    name: '上海核电设备有限公司',
    credit_code: '91310000123456789A',
    type: '生产商',
    category: '原材料',
    region: '华东区',
    industry: 'nuclear',
    industry_text: '核电',
    brand: '核电专用',
    contact_person: '张三',
    mobile: '13800138001',
    bank_name: '中国银行',
    branch_name: '上海浦东支行',
    bank_account: '6222021234567890123',
    settlement: '月结30天',
    purchaser: '采购员A',
    supplier_status: 'qualified',
    is_blacklist: false,
    approval_status: '已通过',
    current_node: '已完成',
    apply_time: '2023-11-20 10:00:00',
    approval_time: '2023-11-22 15:30:00',
    last_change_time: '2023-11-25 09:00:00',
    attachments: [
      { name: '营业执照.pdf', url: '#' },
      { name: '质量体系认证.pdf', url: '#' },
      { name: '安全保密资质.pdf', url: '#' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-20 10:00:00', approver: '采购员A', comment: '资料齐全' },
      { node: '采购部长审批', status: 'completed', time: '2023-11-21 14:00:00', approver: '采购部长-刘华', comment: '同意' },
      { node: '质保部审核', status: 'completed', time: '2023-11-22 10:00:00', approver: '质保经理-王强', comment: '质保体系符合核电要求' },
      { node: '副总审批', status: 'completed', time: '2023-11-22 15:30:00', approver: '副总-赵总', comment: '同意准入' }
    ],
    change_history: [
      { change_time: '2023-11-25 09:00:00', change_field: '联系电话', old_value: '13800138000', new_value: '13800138001', change_by: 'Admin', change_reason: '联系人电话变更' }
    ]
  },
  {
    id: 2,
    name: '北京军工材料供应商',
    credit_code: '91110000987654321B',
    type: '贸易商',
    category: '电子元器件',
    region: '华北区',
    industry: 'military',
    industry_text: '军工',
    brand: '军工专用',
    contact_person: '李四',
    mobile: '13800138002',
    bank_name: '工商银行',
    branch_name: '北京朝阳支行',
    bank_account: '6222031234567890456',
    settlement: '预付',
    purchaser: '采购员B',
    supplier_status: 'qualified',
    is_blacklist: false,
    approval_status: '已通过',
    current_node: '已完成',
    apply_time: '2023-11-21 14:00:00',
    approval_time: '2023-11-23 16:20:00',
    last_change_time: '',
    attachments: [
      { name: '营业执照.pdf', url: '#' },
      { name: '保密资质.pdf', url: '#' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-21 14:00:00', approver: '采购员B', comment: '材料核实' },
      { node: '采购部长审批', status: 'completed', time: '2023-11-22 10:00:00', approver: '采购部长-刘华', comment: '同意' },
      { node: '质保部审核', status: 'completed', time: '2023-11-23 09:00:00', approver: '质保经理-王强', comment: '保密资质符合要求' },
      { node: '副总审批', status: 'completed', time: '2023-11-23 16:20:00', approver: '副总-赵总', comment: '同意准入' }
    ],
    change_history: []
  },
  {
    id: 3,
    name: '杭州普通贸易有限公司',
    credit_code: '91330000111222333C',
    type: '贸易商',
    category: '办公用品',
    region: '华东区',
    industry: 'normal',
    industry_text: '普通行业',
    brand: '通用品牌',
    contact_person: '王五',
    mobile: '13800138003',
    bank_name: '建设银行',
    branch_name: '杭州西湖支行',
    bank_account: '6222041234567890789',
    settlement: '货到付款',
    purchaser: '采购员C',
    supplier_status: 'qualified',
    is_blacklist: false,
    approval_status: '已驳回',
    current_node: '副总审批-驳回',
    apply_time: '2023-11-22 09:00:00',
    approval_time: '2023-11-22 17:00:00',
    last_change_time: '',
    attachments: [
      { name: '营业执照.pdf', url: '#' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-22 09:00:00', approver: '采购员C', comment: '提交申请' },
      { node: '采购部长审批', status: 'completed', time: '2023-11-22 14:00:00', approver: '采购部长-刘华', comment: '同意' },
      { node: '副总审批', status: 'rejected', time: '2023-11-22 17:00:00', approver: '副总-赵总', comment: '资质材料不完整，请补充' }
    ],
    change_history: []
  },
  {
    id: 4,
    name: '深圳电子元器件公司',
    credit_code: '91440000444555666D',
    type: '生产商',
    category: '电子元器件',
    region: '华南区',
    industry: 'normal',
    industry_text: '普通行业',
    brand: '深电',
    contact_person: '林芳',
    mobile: '13600136000',
    bank_name: '平安银行',
    branch_name: '深圳南山支行',
    bank_account: '6222051234567890012',
    settlement: '月结30天',
    purchaser: '采购员A',
    supplier_status: 'qualified',
    is_blacklist: false,
    approval_status: '已通过',
    current_node: '已完成',
    apply_time: '2023-11-24 08:30:00',
    approval_time: '2023-11-24 17:30:00',
    last_change_time: '',
    attachments: [
      { name: '营业执照.pdf', url: '#' },
      { name: 'ISO9001证书.pdf', url: '#' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-24 08:30:00', approver: '采购员A', comment: '信息完整' },
      { node: '采购部长审批', status: 'completed', time: '2023-11-24 14:00:00', approver: '采购部长-刘华', comment: '审核通过' },
      { node: '副总审批', status: 'completed', time: '2023-11-24 17:30:00', approver: '副总-赵总', comment: '同意准入' }
    ],
    change_history: []
  },
  {
    id: 5,
    name: '广州潜在供应商E',
    credit_code: '91440100777888999E',
    type: '生产商',
    category: '原材料',
    region: '华南区',
    industry: 'petrochemical',
    industry_text: '石化',
    brand: '粤石化',
    contact_person: '陈明',
    mobile: '13500135000',
    bank_name: '农业银行',
    branch_name: '广州天河支行',
    bank_account: '6222061234567890345',
    settlement: '预付',
    purchaser: '采购员B',
    supplier_status: 'temp',
    is_blacklist: false,
    approval_status: '待审批',
    current_node: '采购部长审批',
    apply_time: '2023-11-25 10:00:00',
    approval_time: '',
    last_change_time: '',
    attachments: [
      { name: '营业执照.pdf', url: '#' }
    ],
    approval_flow: [
      { node: '采购员录入', status: 'completed', time: '2023-11-25 10:00:00', approver: '采购员B', comment: '潜在供应商' },
      { node: '采购部长审批', status: 'current', time: '', approver: '', comment: '' }
    ],
    change_history: []
  }
])

const supplierDetailVisible = ref(false)
const selectedSupplier = ref(null)

// 审批详情弹窗相关
const approvalDetailVisible = ref(false)
const currentApprovalRow = ref(null)

const approvalSteps = ref([
  {
    node: '申请人提交',
    approver: '销售员-张明',
    time: '2023-11-20 10:00:00',
    status: 'completed',
    comment: '申请材料齐全'
  },
  {
    node: '采购部审批',
    approver: '采购经理-李华',
    time: '2023-11-21 14:00:00',
    status: 'completed',
    comment: '资质审核通过'
  },
  {
    node: '质保部审批',
    approver: '质保经理-王强',
    time: '2023-11-22 10:00:00',
    status: 'completed',
    comment: '质保体系符合要求'
  },
  {
    node: '总经理审批',
    approver: '总经理-赵总',
    time: '2023-11-22 15:30:00',
    status: 'completed',
    comment: '同意准入'
  }
])

const certificateOptions = computed(() => {
  if (form.entry_type === 'manufacturing') {
    return [
      '营业执照',
      '制造/设计许可证',
      'ISO 9001质量管理体系证书',
      '14001 职业健康安全管理体系证书',
      '18001 环境管理体系认证证书',
      '质量手册/质保大纲',
      '外部审查报告',
      '生产设备清单',
      '检测设备清单',
      '计量器具清单',
      '设计人员名单',
      '检验人员名单',
      '特种工艺人员名单及资质证书',
      '专职质保人员名单',
      '近三年公司业绩清单',
      '其他文件'
    ]
  } else {
    return [
      '营业执照',
      '品牌代理证书',
      'ISO 9001质量管理体系证书',
      '14001 职业健康安全管理体系证书',
      '18001 环境管理体系认证证书',
      '近三年公司业绩清单',
      '其他文件'
    ]
  }
})

// 资质证书列表（带有效期）- 使用ref存储以便编辑
const certificateDataMap = ref({}) // 存储每个证书的有效期数据

const certificateList = computed(() => {
  return certificateOptions.value.map(name => {
    // 如果没有该证书的数据，初始化
    if (!certificateDataMap.value[name]) {
      certificateDataMap.value[name] = { validStartDate: null, validEndDate: null }
    }
    return {
      name: name,
      get validStartDate() {
        return certificateDataMap.value[name]?.validStartDate
      },
      set validStartDate(val) {
        if (!certificateDataMap.value[name]) {
          certificateDataMap.value[name] = {}
        }
        certificateDataMap.value[name].validStartDate = val
      },
      get validEndDate() {
        return certificateDataMap.value[name]?.validEndDate
      },
      set validEndDate(val) {
        if (!certificateDataMap.value[name]) {
          certificateDataMap.value[name] = {}
        }
        certificateDataMap.value[name].validEndDate = val
      }
    }
  })
})

// 判断证书是否即将到期（6个月内）
const isCertExpiringSoon = (validEndDate) => {
  if (!validEndDate) return false

  const endDate = new Date(validEndDate)
  const today = new Date()
  const sixMonthsLater = new Date()
  sixMonthsLater.setMonth(sixMonthsLater.getMonth() + 6)

  // 如果结束日期在今天到6个月后之间，或者已经过期
  return endDate <= sixMonthsLater
}

// ... existing fetchData and others ...

// Helper to find file
const getCurrentTitle = () => {
  const map = {
    'apply': '供应商申请',
    'supplier-apply': '供应商准入申请',
    'supplier-list': '供应商档案库',
    'approval': '供应商审批',
    'supplier-approval': '供应商审批',
    'change': '供应商变更',
    'supplier-change': '供应商变更',
    'temp': '潜在供应商',
    'qualified': '合格供应商',
    'blacklist': '供应商黑名单',
    'my-inquiry': '我的询价',
    'pending-quote': '待报价项目',
    'inquiry-ledger': '全部询报价台账',
    'bid-projects': '中标项目管理',
    'payment-approval': '吧盛支付申请',
    'order-approval': '采购订单审批',
    'agreement-price': '协议价管理'
  }
  return map[currentMenu.value] || ''
}

const getStatusType = (status) => {
  if (status === 'qualified') return 'success'
  if (status === 'temp') return 'warning'
  return 'info'
}

const getApprovalStatusType = (status) => {
  if (status === '待审批') return 'warning'
  if (status === '已通过') return 'success'
  if (status === '已驳回') return 'danger'
  return 'info'
}

const getFileCount = (certType) => {
    return uploadedFiles.value.filter(f => f.type === certType).length
}

const getFileForCert = (certType) => {
    return uploadedFiles.value.find(f => f.type === certType)
}

const handleFileChange = (uploadFile, certType) => {
  const idx = uploadedFiles.value.findIndex(f => f.type === certType)
  if (idx !== -1) {
      uploadedFiles.value.splice(idx, 1)
  }
  
  uploadedFiles.value.push({
    type: certType,
    name: uploadFile.name
  })
  ElMessage.success(`已添加 ${certType}`)
}

const removeFile = (certType) => {
  const idx = uploadedFiles.value.findIndex(f => f.type === certType)
  if (idx !== -1) {
      uploadedFiles.value.splice(idx, 1)
      // ElMessage.success('已移除文件')
  }
}

const fetchData = async () => {
  let apiStatus = ''
  if (currentMenu.value === 'qualified') apiStatus = 'qualified'
  else if (currentMenu.value === 'temp') apiStatus = 'temp'
  
  const params = {
    page: 1,
    pageSize: 100,
  }
  if (apiStatus) params.status = apiStatus
  if (searchForm.name) params.enterprise_name = searchForm.name
  if (searchForm.status) params.status = searchForm.status

  try {
    const res = await service({
      url: '/suppliers',
      method: 'get',
      params
    })
    if (res.code === 0) {
        tableData.value = res.data.list.map(item => ({
            ...item,
            id: item.ID,
            name: item.enterprise_name,
            code: item.credit_code,
            contact: item.contact_person,
            status_base: item.status === 'blacklist' ? 'temp' : item.status,
            is_blacklist: item.status === 'blacklist' ? '是' : '否',
            blacklist_reason: item.blacklist_reason || '',
            // Map additional fields if needed by UI directly, but most are in item
            ...item
        }))
    }
  } catch (error) {
    console.error(error)
    ElMessage.error('获取数据失败')
  }
}

onMounted(() => {
  fetchData()
})

const handleSelect = (key) => {
  currentMenu.value = key
  if (key === 'apply') isTempAdd.value = false
  fetchData()
}

const nextTab = (tabName) => {
  activeTab.value = tabName
}

const saveData = async () => {
  try {
    const payload = { ...form }
    if (isTempAdd.value) payload.status = 'temp'
    else payload.status = 'qualified'
    
    const res = await service({
        url: '/suppliers',
        method: 'post',
        data: payload
    })
    
    if (res.code === 0) {
        ElMessage.success('提交申请成功！')
        fetchData()
    }
  } catch(e) {
    ElMessage.error('提交失败')
  }
}

const onSearch = () => {
  fetchData()
}

const exportExcel = () => {
  ElMessage.success('正在导出 Excel...')
}

// 询价相关方法
const addInquiryItem = () => {
  inquiryForm.items.push({ brand: '', product_name: '', specification: '', quantity: 1, unit: '个', medium: '', diameter: '', material: '', remark: '' })
}

const removeInquiryItem = (index) => {
  if (inquiryForm.items.length > 1) {
    inquiryForm.items.splice(index, 1)
  } else {
    ElMessage.warning('至少保留一条物料信息')
  }
}

const handleTechFile = (file) => {
  techFiles.value.push({ name: file.name, raw: file })
}

const showImportDialog = () => {
  importDialogVisible.value = true
}

const handleExcelFile = (file) => {
  console.log('Excel file:', file)
}

const downloadTemplate = () => {
  ElMessage.success('正在下载模板...')
}

const importExcel = () => {
  ElMessage.success('Excel导入成功')
  importDialogVisible.value = false
}

const submitInquiry = async () => {
  if (!inquiryForm.project_name || !inquiryForm.end_user) {
    ElMessage.error('请填写项目名称和终端用户名称')
    return
  }
  if (!inquiryForm.quote_deadline) {
    ElMessage.error('请选择报价截止时间')
    return
  }
  const hasValidItem = inquiryForm.items.some(item => item.product_name || item.brand)
  if (!hasValidItem) {
    ElMessage.error('请至少填写一条物料信息')
    return
  }
  ElMessage.success('询价提交成功，已发送至采购部门')
  // 重置表单
  inquiryForm.project_name = ''
  inquiryForm.end_user = ''
  inquiryForm.quote_deadline = ''
  inquiryForm.expected_delivery = ''
  inquiryForm.technical_requirements = ''
  inquiryForm.items = [{ brand: '', product_name: '', specification: '', quantity: 1, unit: '个', medium: '', diameter: '', material: '', remark: '' }]
  techFiles.value = []
}

// 我的询价列表相关方法
const searchMyInquiries = () => {
  console.log('搜索条件:', inquirySearchForm)
}

const resetInquirySearch = () => {
  inquirySearchForm.project_name = ''
  inquirySearchForm.status = ''
}

const getInquiryStatusType = (status) => {
  const map = { pending: 'info', clarifying: 'warning', quoted: 'primary', waiting_feedback: 'danger', won: 'success', lost: '' }
  return map[status] || 'info'
}

const getInquiryStatusText = (status) => {
  const map = { pending: '待采购处理', clarifying: '澄清中', quoted: '已报价', waiting_feedback: '待反馈中标', won: '已中标', lost: '未中标' }
  return map[status] || '未知'
}

const isOverdue = (deadline) => {
  if (!deadline) return false
  return new Date(deadline) < new Date()
}

const viewInquiryDetail = (row) => {
  ElMessage.info('查看询价详情: ' + row.inquiry_no)
}

// 中标反馈相关方法
const showFeedbackDialog = (row) => {
  feedbackForm.inquiry_id = row.id
  feedbackForm.inquiry_no = row.inquiry_no
  feedbackForm.project_name = row.project_name
  feedbackForm.quote_date = row.quote_date || new Date().toISOString().split('T')[0]
  feedbackForm.result = ''
  feedbackForm.deal_amount = 0
  feedbackForm.order_no = ''
  feedbackForm.no_deal_reason = ''
  feedbackForm.delay_reason = ''
  feedbackForm.remark = ''
  calculationFiles.value = []
  feedbackDialogVisible.value = true
}

const handleCalculationFile = (file) => {
  calculationFiles.value = [{ name: file.name, raw: file }]
}

const submitFeedback = () => {
  if (!feedbackForm.result) {
    ElMessage.error('请选择中标结果')
    return
  }
  // 成交时必须上传计算书和填写订单号
  if (feedbackForm.result === 'won') {
    if (!feedbackForm.order_no) {
      ElMessage.error('成交项目请填写订单号')
      return
    }
    if (calculationFiles.value.length === 0) {
      ElMessage.error('成交项目请上传计算书')
      return
    }
  }
  // 未成交时必须填写原因
  if (feedbackForm.result === 'lost' && !feedbackForm.no_deal_reason) {
    ElMessage.error('请填写未成交原因')
    return
  }
  // 项目延迟时必须填写延迟原因
  if (feedbackForm.result === 'delayed' && !feedbackForm.delay_reason) {
    ElMessage.error('请填写延迟原因')
    return
  }
  ElMessage.success('中标结果反馈成功')
  feedbackDialogVisible.value = false
}

// 澄清回复相关方法
const replyClarification = (row) => {
  clarifyReplyForm.inquiry_id = row.id
  clarifyReplyForm.content = ''
  clarifyFiles.value = []
  // 模拟加载澄清历史
  clarifyHistory.value = [
    { id: 1, from: 'purchase', content: '请确认物料规格型号是否正确？DN50是否为法兰连接？', created_at: '2024-11-23 10:30:00' },
    { id: 2, from: 'sales', content: '是的，DN50法兰连接，压力等级PN16', created_at: '2024-11-23 14:20:00' },
    { id: 3, from: 'purchase', content: '好的，另外请确认材质是否为304不锈钢？', created_at: '2024-11-24 09:15:00' }
  ]
  clarifyDialogVisible.value = true
}

const handleClarifyFile = (file) => {
  clarifyFiles.value.push({ name: file.name, raw: file })
}

const submitClarifyReply = () => {
  if (!clarifyReplyForm.content.trim()) {
    ElMessage.error('请输入回复内容')
    return
  }
  ElMessage.success('澄清回复已提交')
  clarifyDialogVisible.value = false
}

const showAddAgreementDialog = () => {
  ElMessage.info('新增协议价功能开发中...')
}

const openChangeDialog = (row) => {
  Object.assign(changeForm, {
    id: row.id,
    enterprise_name: row.enterprise_name,
    credit_code: row.credit_code,
    entry_type: row.entry_type,
    category: row.category,
    region: row.region,
    industry: row.industry,
    brand: row.brand,
    contact_person: row.contact_person,
    mobile: row.mobile,
    bank_name: row.bank_name,
    branch_name: row.branch_name,
    bank_account: row.bank_account,
    settlement: row.settlement,
    purchaser: row.purchaser,
    status: row.status,
    reason: '',
    toBlacklist: row.status === 'blacklist'
  })
  changeDialogVisible.value = true
}

const submitChange = async () => {
  const payload = {
    ID: changeForm.id,
    enterprise_name: changeForm.enterprise_name,
    credit_code: changeForm.credit_code,
    entry_type: changeForm.entry_type,
    category: changeForm.category,
    region: changeForm.region,
    industry: changeForm.industry,
    brand: changeForm.brand,
    contact_person: changeForm.contact_person,
    mobile: changeForm.mobile,
    bank_name: changeForm.bank_name,
    branch_name: changeForm.branch_name,
    bank_account: changeForm.bank_account,
    settlement: changeForm.settlement,
    purchaser: changeForm.purchaser,
    status: changeForm.toBlacklist ? 'blacklist' : (changeForm.status === 'blacklist' ? 'qualified' : changeForm.status)
  }
  
  // If specifically toggling blacklist
  if (changeForm.toBlacklist) {
      payload.status = 'blacklist'
  } else if (changeForm.status === 'blacklist' && !changeForm.toBlacklist) {
      payload.status = 'qualified'
  }

  try {
    const res = await service({
        url: `/suppliers/${changeForm.id}`,
        method: 'put',
        data: payload
    })

    if (res.code === 0) {
        ElMessage.success('变更成功')
        changeDialogVisible.value = false
        fetchData()
    }
  } catch(e) {
    ElMessage.error('变更失败')
  }
}

const submitRowChange = async (row) => {
    const status = row.is_blacklist_str === '是' ? 'blacklist' : row.status_base
    const payload = {
        ID: row.id,
        enterprise_name: row.enterprise_name,
        credit_code: row.credit_code,
        entry_type: row.entry_type,
        category: row.category,
        region: row.region,
        industry: row.industry,
        brand: row.brand,
        contact_person: row.contact_person,
        mobile: row.mobile,
        bank_name: row.bank_name,
        branch_name: row.branch_name,
        settlement: row.settlement,
        purchaser: row.purchaser,
        bank_account: row.bank_account,
        status: status
    }
    try {
        const res = await service({
            url: `/suppliers/${row.id}`,
            method: 'put',
            data: payload
        })
        if (res.code === 0) {
            ElMessage.success('保存成功')
            fetchData()
        }
    } catch(e) {
        ElMessage.error('保存失败')
    }
}

// 审批相关方法
const handleApprove = (row) => {
    ElMessageBox.confirm('确认通过该供应商的审批？', '审批通过', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'success'
    }).then(() => {
        // 更新状态
        row.status = '已通过'
        ElMessage.success('审批通过成功')
    }).catch(() => {
        // 用户取消操作
    })
}

const handleReject = (row) => {
    ElMessageBox.prompt('请输入驳回原因', '审批驳回', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        inputPattern: /^.+$/,
        inputErrorMessage: '请输入驳回原因'
    }).then(({ value }) => {
        // 更新状态
        row.status = '已驳回'
        row.reject_reason = value
        ElMessage.success('审批驳回成功')
    }).catch(() => {
        // 用户取消操作
    })
}

const viewDetails = (row) => {
    ElMessage.info('查看详情功能开发中...')
}

// 显示审批详情弹窗
const showApprovalDetail = (row) => {
    currentApprovalRow.value = row
    approvalDetailVisible.value = true
}

// 获取当前审批步骤索引
const getActiveStep = (row) => {
    if (!row || !row.approval_flow) return 0
    const currentIndex = row.approval_flow.findIndex(step => step.status === 'current')
    if (currentIndex >= 0) return currentIndex
    // 如果没有当前节点，返回最后一个已完成的节点
    const lastCompleted = row.approval_flow.filter(step => step.status === 'completed').length
    return lastCompleted
}

// 跳转到供应商申请页面查看详情并进行审批
const goToApprovalDetail = (row) => {
    // 存储当前审批行数据
    currentApprovalRow.value = row
    // 跳转到申请页面，并传递审批模式
    currentMenu.value = 'apply'
    // 设置为审批查看模式
    isApprovalMode.value = true
    // 填充表单数据
    Object.assign(form, {
        enterprise_name: row.name,
        credit_code: row.unified_social_credit_code,
        entry_type: row.type === '原材料供应商' || row.type === '生产商' ? 'manufacturing' : 'trading',
        industry: row.industry,
        contact_person: row.contact_person,
        region: row.registered_address?.split('市')[0] + '区' || '',
        brand: '',
        mobile: row.contacts?.[0]?.mobile || '',
        bank_name: row.banks?.[0]?.bank_name || '',
        branch_name: row.banks?.[0]?.branch_name || '',
        bank_account: row.banks?.[0]?.account_number || ''
    })
    ElMessage.info('已进入审批详情页面，可在此查看完整信息并进行审批操作')
}

// 快速审批 - 显示审批选项对话框
const showQuickApprovalDialog = ref(false)
const quickApprovalRow = ref(null)
const quickApprovalResult = ref('') // 'approve', 'reject_fill', 'reject_unqualified'
const quickApprovalRemark = ref('')

const quickApprove = (row) => {
    quickApprovalRow.value = row
    quickApprovalResult.value = ''
    quickApprovalRemark.value = ''
    showQuickApprovalDialog.value = true
}

// 确认快速审批
const confirmQuickApproval = () => {
    if (!quickApprovalResult.value) {
        ElMessage.warning('请选择审批结果')
        return
    }

    const row = quickApprovalRow.value
    if (!row) return

    if (quickApprovalResult.value === 'approve') {
        // 审批通过
        row.status = '已通过'
        row.current_node = '已完成'
        ElMessage.success('审批通过成功')
        showQuickApprovalDialog.value = false
    } else if (quickApprovalResult.value === 'reject_fill') {
        // 填写问题 - 跳转到供应商准入申请填写
        if (!quickApprovalRemark.value) {
            ElMessage.warning('请填写需要补充/修改的内容说明')
            return
        }
        row.status = '待补充'
        row.current_node = '待补充资料'
        row.reject_reason = quickApprovalRemark.value
        row.reject_type = '填写问题'
        ElMessage.success('已退回供应商补充资料')
        showQuickApprovalDialog.value = false
        // 跳转到供应商准入申请页面
        currentMenu.value = 'supplier-apply'
        isApprovalMode.value = true
        currentApprovalRow.value = row
        // 填充表单数据用于修改
        Object.assign(form, {
            enterprise_name: row.name || '',
            credit_code: row.credit_code || '',
            entry_type: row.supplier_type === '生产商' ? 'manufacturing' : 'trading',
            category: 'raw',
            region: '',
            industry: row.industry === '普通' ? 'normal' : 'special',
            brand: '',
            contact_person: row.contact || '',
            mobile: row.phone || '',
            purchaser: 'Admin',
            email: row.email || '',
            settlement: 'm30',
            bank_name: '',
            branch_name: '',
            bank_account: ''
        })
    } else if (quickApprovalResult.value === 'reject_unqualified') {
        // 资质不合格 - 终止
        if (!quickApprovalRemark.value) {
            ElMessage.warning('请填写不合格原因')
            return
        }
        row.status = '已驳回'
        row.current_node = '审批终止'
        row.reject_reason = quickApprovalRemark.value
        row.reject_type = '资质不合格'
        ElMessage.success('审批已终止')
        showQuickApprovalDialog.value = false
    }
}

// 重置审批搜索
const resetApprovalSearch = () => {
    searchForm.name = ''
    searchForm.approvalStatus = ''
    searchForm.industry = ''
}

// 潜在供应商转为合格供应商（进入审批流程）
const convertToQualified = (row) => {
    ElMessageBox.confirm(
        `确认将【${row.enterprise_name}】转为合格供应商？转换后将进入审批流程。`,
        '转为合格供应商',
        {
            confirmButtonText: '确认转换',
            cancelButtonText: '取消',
            type: 'info'
        }
    ).then(() => {
        // 跳转到供应商准入申请页面，并填充数据
        currentMenu.value = 'apply'
        supplierApplyType.value = 'qualified'

        // 填充表单数据
        Object.assign(form, {
            enterprise_name: row.enterprise_name || '',
            credit_code: row.credit_code || '',
            entry_type: row.entry_type === '生产商' ? 'manufacturing' : 'trading',
            category: row.category === '原材料' ? 'raw' : row.category === '电子元器件' ? 'electronic' : 'office',
            region: '',
            industry: row.industry === '核电' ? 'nuclear' : row.industry === '军工' ? 'military' : row.industry === '石化' ? 'petrochemical' : 'normal',
            brand: row.brand || '',
            contact_person: row.contact_person || '',
            mobile: row.mobile || '',
            purchaser: 'Admin',
            email: row.email || '',
            settlement: 'm30',
            bank_name: '',
            branch_name: '',
            bank_account: ''
        })

        ElMessage.success('请补充完整信息后提交审批')
    }).catch(() => {})
}

// 返回审批列表
const backToApprovalList = () => {
    isApprovalMode.value = false
    currentApprovalRow.value = null
    currentMenu.value = 'supplier-approval'
    // 重置表单
    Object.assign(form, {
        enterprise_name: '',
        credit_code: '',
        entry_type: 'manufacturing',
        category: 'raw',
        region: '',
        industry: 'normal',
        brand: '',
        contact_person: '',
        mobile: '',
        purchaser: 'Admin',
        email: '',
        settlement: 'm30',
        bank_name: '',
        branch_name: '',
        bank_account: ''
    })
}

// 审批通过当前供应商
const approveCurrentSupplier = () => {
    if (!currentApprovalRow.value) {
        ElMessage.warning('无审批数据')
        return
    }

    ElMessageBox.prompt('请输入审批意见（可选）', '审批通过', {
        confirmButtonText: '确认通过',
        cancelButtonText: '取消',
        inputPlaceholder: '审批意见...',
        type: 'success'
    }).then(({ value }) => {
        // 更新审批数据
        currentApprovalRow.value.status = '已通过'
        currentApprovalRow.value.current_node = '已完成'

        // 更新审批流程中的当前节点
        const flow = currentApprovalRow.value.approval_flow
        const currentStep = flow.find(s => s.status === 'current')
        if (currentStep) {
            currentStep.status = 'completed'
            currentStep.time = new Date().toLocaleString()
            currentStep.approver = currentUser.value
            currentStep.comment = value || '同意'
        }

        ElMessage.success('审批通过成功')
        backToApprovalList()
    }).catch(() => {})
}

// 审批驳回当前供应商
const rejectCurrentSupplier = () => {
    if (!currentApprovalRow.value) {
        ElMessage.warning('无审批数据')
        return
    }

    ElMessageBox.prompt('请输入驳回原因', '审批驳回', {
        confirmButtonText: '确认驳回',
        cancelButtonText: '取消',
        inputPattern: /^.+$/,
        inputErrorMessage: '请输入驳回原因',
        type: 'warning'
    }).then(({ value }) => {
        // 更新审批数据
        currentApprovalRow.value.status = '已驳回'
        currentApprovalRow.value.current_node = '已驳回'

        // 更新审批流程中的当前节点
        const flow = currentApprovalRow.value.approval_flow
        const currentStep = flow.find(s => s.status === 'current')
        if (currentStep) {
            currentStep.status = 'rejected'
            currentStep.time = new Date().toLocaleString()
            currentStep.approver = currentUser.value
            currentStep.comment = value
        }

        ElMessage.success('审批驳回成功')
        backToApprovalList()
    }).catch(() => {})
}

// 供应商档案库相关方法
const searchSuppliers = () => {
    ElMessage.success('搜索功能执行')
    // TODO: 实际搜索逻辑
}

// 重置供应商档案库搜索
const resetSupplierSearch = () => {
    supplierSearch.keyword = ''
    supplierSearch.industry = ''
    supplierSearch.status = ''
    supplierSearch.supplierType = ''
}

const exportSupplierList = () => {
    ElMessage.success('正在导出Excel...')
    // TODO: 实际导出逻辑
}

const viewSupplierDetail = (row) => {
    selectedSupplier.value = row
    supplierDetailVisible.value = true
}

const previewAttachment = (file) => {
    ElMessage.info(`预览文件: ${file.name}`)
    // TODO: 实际预览逻辑
}

const printApprovalForm = () => {
    ElMessage.success('正在生成PDF审批单...')
    // TODO: 实际PDF生成逻辑
}

const handleSupplierPageSizeChange = (size) => {
    supplierList.pageSize = size
    searchSuppliers()
}

const handleSupplierPageChange = (page) => {
    supplierList.page = page
    searchSuppliers()
}

// 处理行业变化
const handleIndustryChange = (value) => {
    // 如果切换到非特殊行业，清空特殊文件
    if (value !== 'nuclear' && value !== 'military') {
        specialFiles.quality_cert = []
        specialFiles.security_cert = []
        specialFiles.other_cert = []
    }
}

// 处理特殊文件上传
const handleSpecialFile = (file, type) => {
    // 检查文件大小
    const isLt5M = file.size / 1024 / 1024 < 5
    if (!isLt5M) {
        ElMessage.error('文件大小不能超过5MB')
        return false
    }

    // 添加到对应类型
    specialFiles[type].push({
        name: file.name,
        raw: file,
        url: URL.createObjectURL(file.raw || file)
    })
}

// 保存草稿
const saveDraft = () => {
    ElMessage.success('草稿保存成功')
}

// 保存潜在供应商（不进入审批流程）
const submitTempSupplier = () => {
    // 潜在供应商只需要基本信息
    if (!form.enterprise_name || !form.contact_person || !form.mobile) {
        ElMessage.error('请填写企业名称、联系人和手机号')
        return
    }

    ElMessageBox.confirm('确认保存为潜在供应商？潜在供应商不进入审批流程。', '保存潜在供应商', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'info'
    }).then(() => {
        // 添加到潜在供应商列表（字段与供应商准入申请一致）
        const categoryMap = { 'raw': '原材料', 'electronic': '电子元器件', 'office': '办公用品' }
        const industryMap = { 'normal': '普通', 'nuclear': '核电', 'military': '军工', 'petrochemical': '石化', 'other': '其他' }
        const newTemp = {
            id: tempSupplierData.value.length + 1,
            enterprise_name: form.enterprise_name,
            credit_code: form.credit_code || '',
            entry_type: form.entry_type === 'manufacturing' ? '生产商' : '贸易商',
            category: categoryMap[form.category] || '原材料',
            industry: industryMap[form.industry] || '普通',
            brand: form.brand || '',
            contact_person: form.contact_person,
            mobile: form.mobile,
            email: form.email || ''
        }
        tempSupplierData.value.unshift(newTemp)

        ElMessage.success('保存成功！已添加到潜在供应商列表')
        // 跳转到潜在供应商列表
        currentMenu.value = 'temp'
        // 重置表单
        resetForm()
    }).catch(() => {})
}

// 重置表单
const resetForm = () => {
    Object.assign(form, {
        enterprise_name: '',
        credit_code: '',
        entry_type: 'manufacturing',
        category: 'raw',
        region: '',
        industry: 'normal',
        brand: '',
        contact_person: '',
        mobile: '',
        purchaser: 'Admin',
        email: '',
        settlement: 'm30',
        bank_name: '',
        branch_name: '',
        bank_account: ''
    })
    supplierApplyType.value = 'qualified'
}

// 提交审批（合格供应商）
const submitForApproval = () => {
    // 验证必填项
    if (!form.enterprise_name || !form.credit_code || !form.industry || !form.brand || !form.contact_person || !form.mobile) {
        ElMessage.error('请填写所有必填项')
        return
    }

    // 如果是核电/军工，检查特殊文件
    if ((form.industry === 'nuclear' || form.industry === 'military')) {
        if (specialFiles.quality_cert.length === 0 || specialFiles.security_cert.length === 0) {
            ElMessage.error('核电/军工行业必须上传质保体系认证和安全保密资质')
            return
        }
    }

    ElMessageBox.confirm('提交后将进入审批流程，确认提交？', '提交审批', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        // 模拟提交
        ElMessage.success('提交成功！已进入审批流程')

        // 如果是核电/军工，提示会自动增加质保部审批节点
        if (form.industry === 'nuclear' || form.industry === 'military') {
            ElMessage.info('检测到核电/军工行业，已自动增加质保部审批节点')
        }
    })
}

// ========== 采购报价处理相关方法 ==========
const searchPurchaseList = () => {
    console.log('搜索采购待处理列表:', purchaseSearchForm)
    // TODO: 实际API调用
}

const resetPurchaseSearch = () => {
    purchaseSearchForm.inquiry_no = ''
    purchaseSearchForm.project_name = ''
    purchaseSearchForm.status = ''
}

// 判断采购时效是否紧急（距离截止时间小于24小时）
const isPurchaseDeadlineUrgent = (deadline) => {
    if (!deadline) return false
    const now = new Date()
    const deadlineDate = new Date(deadline)
    const diffHours = (deadlineDate - now) / (1000 * 60 * 60)
    return diffHours > 0 && diffHours < 24
}

const getPurchaseStatusType = (status) => {
    const map = { pending: 'warning', clarifying: 'info', quotable: 'primary', quoted: 'success' }
    return map[status] || 'info'
}

const getPurchaseStatusText = (status) => {
    const map = { pending: '待处理', clarifying: '澄清中', quotable: '可报价', quoted: '已报价' }
    return map[status] || '未知'
}

const viewPurchaseDetail = (row) => {
    ElMessage.info('查看采购询价详情: ' + row.inquiry_no)
}

// 采购发起澄清
const showPurchaseClarifyDialog = (row) => {
    purchaseClarifyForm.inquiry_id = row.id
    purchaseClarifyForm.inquiry_no = row.inquiry_no
    purchaseClarifyForm.content = ''
    purchaseClarifyFiles.value = []
    purchaseClarifyDialogVisible.value = true
}

const handlePurchaseClarifyFile = (file) => {
    purchaseClarifyFiles.value.push({ name: file.name, raw: file })
}

const submitPurchaseClarify = () => {
    if (!purchaseClarifyForm.content.trim()) {
        ElMessage.error('请输入澄清内容')
        return
    }
    ElMessage.success('澄清请求已发送给销售')
    purchaseClarifyDialogVisible.value = false
}

// 采购填写报价
const showPurchaseQuoteDialog = (row) => {
    purchaseQuoteForm.inquiry_id = row.id
    purchaseQuoteForm.inquiry_no = row.inquiry_no
    purchaseQuoteForm.project_name = row.project_name
    purchaseQuoteForm.end_user = row.end_user
    purchaseQuoteForm.quote_deadline = row.quote_deadline
    purchaseQuoteForm.expected_delivery = row.expected_delivery || ''
    purchaseQuoteForm.technical_requirements = row.technical_requirements || ''
    // 模拟物料数据
    purchaseQuoteForm.items = row.items || [
        { id: 1, brand: '西门子', product_name: '闸阀', specification: 'DN50 PN16', quantity: 10, unit: '个', medium: '水', diameter: 'DN50', material: '304不锈钢', supplier_id: null, unit_price: null, total_price: null, delivery_days: null, remark: '' }
    ]
    purchaseQuoteDialogVisible.value = true
}

// 计算报价行小计
const calcQuoteItemTotal = (item) => {
    if (item.unit_price && item.quantity) {
        item.total_price = parseFloat((item.unit_price * item.quantity).toFixed(2))
    } else {
        item.total_price = null
    }
}

const submitPurchaseQuote = () => {
    const hasQuote = purchaseQuoteForm.items.some(item => item.supplier_id && item.unit_price)
    if (!hasQuote) {
        ElMessage.error('请至少为一项物料填写报价信息')
        return
    }
    ElMessage.success('报价提交成功')
    purchaseQuoteDialogVisible.value = false
}

// 批量导入报价
const showImportQuoteDialog = () => {
    importQuoteDialogVisible.value = true
}

const handleQuoteExcelFile = (file) => {
    console.log('报价Excel文件:', file)
}

const downloadQuoteTemplate = () => {
    ElMessage.success('正在下载报价导入模板...')
}

const importQuoteExcel = () => {
    ElMessage.success('报价导入成功')
    importQuoteDialogVisible.value = false
}

// ========== 台账与价格库相关方法 ==========
const searchLedger = () => {
    console.log('搜索台账:', ledgerSearch)
    // TODO: 实际API调用
}

const resetLedgerSearch = () => {
    ledgerSearch.dateRange = []
    ledgerSearch.material = ''
    ledgerSearch.supplier = ''
    ledgerSearch.project = ''
    ledgerSearch.deal_status = ''
}

const exportLedger = () => {
    ElMessage.success('正在导出询报价台账...')
    // TODO: 实际导出逻辑
}

const getLedgerStatusType = (status) => {
    const map = { '成交': 'success', '未成交': 'danger', '待定': 'warning', '延迟': 'info' }
    return map[status] || 'info'
}

const searchPriceLib = () => {
    console.log('搜索价格库:', priceLibSearch)
    // TODO: 实际API调用
}

const resetPriceLibSearch = () => {
    priceLibSearch.material = ''
    priceLibSearch.supplier = ''
    priceLibSearch.project = ''
    priceLibSearch.dateRange = []
}

// ========== 中标项目管理相关方法 ==========
const searchBidProjects = () => {
    console.log('搜索中标项目:', bidProjectSearch)
    // TODO: 实际API调用
}

const resetBidProjectSearch = () => {
    bidProjectSearch.project_name = ''
    bidProjectSearch.customer = ''
    bidProjectSearch.dateRange = []
}

const viewBidDetail = (row) => {
    ElMessage.info('查看中标项目详情: ' + row.inquiry_no)
}

const downloadFile = (filename) => {
    ElMessage.success(`正在下载文件: ${filename}`)
    // TODO: 实际下载逻辑
}

// ========== 采购订单审批方法 ==========
const searchPendingOrders = () => {
  ElMessage.success('搜索待审批订单')
  // TODO: 调用API搜索
}

const resetOrderSearch = () => {
  orderApprovalSearch.order_no = ''
  orderApprovalSearch.supplier = ''
  orderApprovalSearch.over_limit = ''
}

const searchApprovedOrders = () => {
  ElMessage.success('搜索已审批订单')
  // TODO: 调用API搜索
}

const resetApprovedOrderSearch = () => {
  approvedOrderSearch.order_no = ''
  approvedOrderSearch.result = ''
  approvedOrderSearch.dateRange = []
}

const viewOrderDetail = (row) => {
  ElMessage.info(`查看订单详情: ${row.order_no}`)
  // TODO: 打开详情对话框或跳转详情页
}

const approveOrder = (row) => {
  orderApprovalAction.value = 'approve'
  orderApprovalForm.order_no = row.order_no
  orderApprovalForm.total_amount = row.total_amount
  orderApprovalForm.over_limit_count = row.over_limit_count
  orderApprovalForm.remark = ''
  orderApprovalDialogVisible.value = true
}

const rejectOrder = (row) => {
  orderApprovalAction.value = 'reject'
  orderApprovalForm.order_no = row.order_no
  orderApprovalForm.total_amount = row.total_amount
  orderApprovalForm.over_limit_count = row.over_limit_count
  orderApprovalForm.remark = ''
  orderApprovalDialogVisible.value = true
}

const confirmOrderApproval = () => {
  if (orderApprovalAction.value === 'reject' && !orderApprovalForm.remark) {
    ElMessage.warning('请填写驳回原因')
    return
  }
  const actionText = orderApprovalAction.value === 'approve' ? '通过' : '驳回'
  ElMessage.success(`订单 ${orderApprovalForm.order_no} 审批${actionText}成功`)
  orderApprovalDialogVisible.value = false
  // TODO: 调用API提交审批结果，刷新列表
}

const addOrderItem = () => {
  orderForm.items.push({ material_name: '', specification: '', quantity: 1, unit: '', unit_price: '', agreement_price: null, total: 0 })
}

const removeOrderItem = (index) => {
  if (orderForm.items.length > 1) {
    orderForm.items.splice(index, 1)
  } else {
    ElMessage.warning('至少保留一项物料')
  }
}

const calcOrderItemTotal = (item) => {
  const price = Number(item.unit_price) || 0
  const qty = Number(item.quantity) || 0
  item.total = (price * qty).toFixed(2)
}

const calcOrderTotal = () => {
  let total = 0
  orderForm.items.forEach(item => {
    total += Number(item.total) || 0
  })
  return total.toFixed(2)
}

const submitOrder = () => {
  if (!orderForm.supplier) {
    ElMessage.warning('请选择供应商')
    return
  }
  if (!orderForm.items[0].material_name) {
    ElMessage.warning('请填写物料信息')
    return
  }
  ElMessage.success('采购订单已提交审批')
  // TODO: 调用API提交订单
  resetOrderForm()
  orderApprovalTab.value = 'pending'
}

const resetOrderForm = () => {
  orderForm.inquiry_no = ''
  orderForm.supplier = ''
  orderForm.project_name = ''
  orderForm.items = [{ material_name: '', specification: '', quantity: 1, unit: '', unit_price: '', agreement_price: null, total: 0 }]
  orderForm.remark = ''
}

// ========== 更新中标反馈方法（支持4种状态） ==========
const updateShowFeedbackDialog = (row) => {
    feedbackForm.inquiry_id = row.id
    feedbackForm.inquiry_no = row.inquiry_no
    feedbackForm.project_name = row.project_name
    feedbackForm.quote_date = row.quote_date || new Date().toISOString().split('T')[0]
    feedbackForm.result = ''
    feedbackForm.deal_amount = 0
    feedbackForm.order_no = ''
    feedbackForm.no_deal_reason = ''
    feedbackForm.delay_reason = ''
    feedbackForm.remark = ''
    calculationFiles.value = []
    feedbackDialogVisible.value = true
}

const updateSubmitFeedback = () => {
    if (!feedbackForm.result) {
        ElMessage.error('请选择中标结果')
        return
    }
    // 成交时必须上传计算书和填写订单号
    if (feedbackForm.result === 'won') {
        if (!feedbackForm.order_no) {
            ElMessage.error('成交项目请填写订单号')
            return
        }
        if (calculationFiles.value.length === 0) {
            ElMessage.error('成交项目请上传计算书')
            return
        }
    }
    // 未成交时必须填写原因
    if (feedbackForm.result === 'lost' && !feedbackForm.no_deal_reason) {
        ElMessage.error('请填写未成交原因')
        return
    }
    // 项目延迟时必须填写延迟原因
    if (feedbackForm.result === 'delayed' && !feedbackForm.delay_reason) {
        ElMessage.error('请填写延迟原因')
        return
    }

    ElMessage.success('中标结果反馈成功')
    feedbackDialogVisible.value = false
}
</script>

<style scoped>
.supplier-container {
  height: 100vh;
  display: flex;
  background-color: #f5f7fa; /* Softer background */
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
}

.aside {
  background-color: #001529; /* Darker, more professional sidebar */
  color: #fff;
  transition: width 0.3s;
  width: 220px;
  display: flex;
  flex-direction: column;
  box-shadow: 2px 0 6px rgba(0,21,41,.35);
}
.logo {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
  background-color: #002140;
  color: #fff;
}

.el-menu {
  border-right: none;
}

.header {
  background-color: #fff;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  height: 64px;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
}
.breadcrumb {
  font-size: 14px;
  color: #595959;
}

.main {
  background-color: #f0f2f5;
  padding: 24px;
  flex: 1;
  overflow-y: auto;
}
.content-card {
  background-color: #fff;
  padding: 32px;
  border-radius: 4px; /* Slightly reduced radius for 'stable' look */
  min-height: calc(100vh - 150px);
  box-shadow: 0 2px 4px 0 rgba(0,0,0,0.05); /* Very subtle shadow */
}

.form-section-title {
  font-size: 18px;
  font-weight: 500;
  color: #262626; /* Dark gray instead of black */
  margin-bottom: 24px;
  padding-bottom: 12px;
  border-bottom: 1px solid #f0f0f0; /* Subtle bottom line */
  /* padding-left: 0; Remove left border */
  /* border-left: none; */
  display: flex;
  align-items: center;
}
.form-section-title::before {
    content: '';
    display: block;
    width: 4px;
    height: 16px;
    background-color: #1890ff; /* Enterprise blue */
    margin-right: 8px;
    border-radius: 2px;
}

.demo-tips {
  position: fixed;
  bottom: 24px;
  right: 24px;
  background: rgba(0,0,0,0.85);
  color: white;
  padding: 16px 24px;
  border-radius: 4px;
  max-width: 320px;
  z-index: 9999;
  font-size: 14px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
}

.red-star {
  color: #ff4d4f;
  margin-right: 4px;
  font-family: SimSun, sans-serif;
}

/* Adjust Element Plus overrides for this page */
:deep(.el-form-item__label) {
    color: #595959;
    font-weight: 400;
}
:deep(.el-input__wrapper), :deep(.el-select__wrapper) {
    box-shadow: 0 0 0 1px #d9d9d9 inset;
}
:deep(.el-input__wrapper:hover), :deep(.el-select__wrapper:hover) {
    box-shadow: 0 0 0 1px #409eff inset;
}
:deep(.el-tabs__item) {
    font-size: 15px;
    color: #595959;
}
:deep(.el-tabs__item.is-active) {
    color: #1890ff;
    font-weight: 500;
}

/* 审批时间线样式 */
.approval-timeline {
    background: #fafafa;
    padding: 20px;
    border-radius: 8px;
}

.timeline-content {
    padding: 5px 0;
}

.timeline-title {
    font-weight: 500;
    color: #303133;
    margin-bottom: 5px;
}

.timeline-approver {
    font-size: 12px;
    color: #606266;
    margin-bottom: 3px;
}

.timeline-comment {
    font-size: 12px;
    color: #909399;
    font-style: italic;
}
</style>
