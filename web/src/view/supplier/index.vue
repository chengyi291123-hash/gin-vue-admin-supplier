<template>
  <div class="supplier-container">
    <!-- Sidebar -->
    <div class="aside">
      <div class="logo">
        <span>Gin-Vue-Admin 演示</span>
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
        <div class="breadcrumb">
          首页 / 供应商管理 / {{ getCurrentTitle() }}
          <el-tag v-if="isApprovalMode" type="warning" size="small" style="margin-left: 10px;">审批模式</el-tag>
        </div>
        <div style="display: flex; align-items: center;">
          <template v-if="isApprovalMode && (currentMenu === 'apply' || currentMenu === 'supplier-apply')">
            <el-button size="small" @click="backToApprovalList">返回列表</el-button>
            <el-button type="success" size="small" @click="approveCurrentSupplier">审批通过</el-button>
            <el-button type="danger" size="small" @click="rejectCurrentSupplier" style="margin-right: 15px;">审批驳回</el-button>
          </template>
          <template v-else-if="currentMenu === 'apply' || currentMenu === 'supplier-apply'">
            <el-button size="small" @click="saveDraft">保存草稿</el-button>
            <el-button type="primary" size="small" @click="submitForApproval" style="margin-right: 15px;">提交审批</el-button>
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
                    <el-table :data="certificateOptions.map(c => ({ name: c }))" border style="width: 100%">
                        <el-table-column prop="name" label="证书名称"></el-table-column>
                        <el-table-column label="已上传份数" width="120" align="center">
                            <template #default="scope">
                                {{ getFileCount(scope.row.name) }} 份
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" width="300">
                            <template #default="scope">
                                <div v-if="getFileForCert(scope.row.name)" style="display: flex; align-items: center;">
                                    <el-icon style="margin-right: 5px;"><document /></el-icon>
                                    <span style="margin-right: 10px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; max-width: 150px;">{{ getFileForCert(scope.row.name).name }}</span>
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

              <el-table-column label="是否黑名单" min-width="120">
                <template #default="scope">
                    <el-input v-model="scope.row.is_blacklist_str" size="small" placeholder="输入是/否"></el-input>
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
            <div class="form-section-title">询价需求录入</div>
            <el-form :model="inquiryForm" label-width="120px" style="max-width: 900px">
              <el-form-item label="项目名称">
                <el-input v-model="inquiryForm.project_name" placeholder="请输入项目名称"></el-input>
              </el-form-item>
              <el-form-item label="客户名称">
                <el-input v-model="inquiryForm.customer_name" placeholder="请输入客户名称"></el-input>
              </el-form-item>

              <el-form-item label="物料信息">
                <el-table :data="inquiryForm.items" border style="width: 100%">
                  <el-table-column label="物料编码" width="150">
                    <template #default="scope">
                      <el-input v-model="scope.row.material_code" placeholder="输入编码" size="small"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="物料名称" width="200">
                    <template #default="scope">
                      <el-input v-model="scope.row.material_name" placeholder="输入名称" size="small"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="规格型号" width="150">
                    <template #default="scope">
                      <el-input v-model="scope.row.specification" placeholder="输入规格" size="small"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="数量" width="120">
                    <template #default="scope">
                      <el-input-number v-model="scope.row.quantity" :min="1" size="small"></el-input-number>
                    </template>
                  </el-table-column>
                  <el-table-column label="单位" width="100">
                    <template #default="scope">
                      <el-input v-model="scope.row.unit" placeholder="单位" size="small"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="80">
                    <template #default="scope">
                      <el-button type="danger" size="small" @click="removeInquiryItem(scope.$index)">删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-button type="primary" size="small" @click="addInquiryItem" style="margin-top: 10px;">+ 添加物料</el-button>
              </el-form-item>

              <el-form-item label="期望交期">
                <el-date-picker v-model="inquiryForm.expected_delivery" type="date" placeholder="选择日期" style="width: 100%"></el-date-picker>
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

              <el-form-item label="录入人">
                <el-input :value="currentUser" disabled></el-input>
              </el-form-item>

              <el-form-item label="录入时间">
                <el-input :value="currentTime" disabled></el-input>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="submitInquiry">单条提交</el-button>
                <el-button type="success" @click="showImportDialog">批量导入</el-button>
              </el-form-item>
            </el-form>

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
          </div>

          <!-- VIEW: 待报价项目（采购） -->
          <div v-else-if="currentMenu === 'pending-quote'">
            <div class="form-section-title">待报价项目</div>
            <el-alert title="这是采购人员的主要工作台，处理销售提交的询价需求" type="info" show-icon style="margin-bottom: 20px;"></el-alert>

            <div style="margin-bottom: 20px;">
              <el-input placeholder="搜索项目名称/客户" style="width: 250px; margin-right: 10px;" v-model="quoteSearch.keyword"></el-input>
              <el-select v-model="quoteSearch.status" placeholder="全部状态" style="width: 150px; margin-right: 10px;">
                <el-option label="全部" value=""></el-option>
                <el-option label="待报价" value="pending"></el-option>
                <el-option label="已澄清" value="clarified"></el-option>
                <el-option label="已报价" value="quoted"></el-option>
              </el-select>
              <el-button type="primary">搜索</el-button>
              <el-button type="success" style="float: right;">批量导入报价</el-button>
            </div>

            <el-table :data="pendingQuoteData" border style="width: 100%">
              <el-table-column prop="project_name" label="项目名称" min-width="180" />
              <el-table-column prop="customer" label="客户" width="150" />
              <el-table-column prop="material" label="物料" width="150" />
              <el-table-column prop="quantity" label="数量" width="100" />
              <el-table-column prop="submit_time" label="销售提交时间" width="180" />
              <el-table-column prop="status" label="状态" width="100">
                <template #default="{ row }">
                  <el-tag :type="row.status === '待报价' ? 'warning' : row.status === '已澄清' ? 'info' : 'success'">{{ row.status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="操作" width="150" fixed="right">
                <template #default>
                  <el-button type="primary" size="small">报价详情</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>

          <!-- VIEW: 全部询报价台账 -->
          <div v-else-if="currentMenu === 'inquiry-ledger'">
            <div class="form-section-title">全部询报价台账</div>

            <div style="margin-bottom: 20px;">
              <el-row :gutter="10">
                <el-col :span="6">
                  <el-date-picker v-model="ledgerSearch.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" style="width: 100%;"></el-date-picker>
                </el-col>
                <el-col :span="4">
                  <el-input placeholder="物料" v-model="ledgerSearch.material"></el-input>
                </el-col>
                <el-col :span="4">
                  <el-input placeholder="供应商" v-model="ledgerSearch.supplier"></el-input>
                </el-col>
                <el-col :span="3">
                  <el-select v-model="ledgerSearch.status" placeholder="项目状态" style="width: 100%;">
                    <el-option label="全部" value=""></el-option>
                    <el-option label="进行中" value="ongoing"></el-option>
                    <el-option label="已完成" value="completed"></el-option>
                  </el-select>
                </el-col>
                <el-col :span="3">
                  <el-select v-model="ledgerSearch.isWin" placeholder="是否中标" style="width: 100%;">
                    <el-option label="全部" value=""></el-option>
                    <el-option label="已中标" value="yes"></el-option>
                    <el-option label="未中标" value="no"></el-option>
                  </el-select>
                </el-col>
                <el-col :span="4">
                  <el-button type="primary">搜索</el-button>
                  <el-button type="success">批量导出台账</el-button>
                </el-col>
              </el-row>
            </div>

            <el-table :data="ledgerData" border style="width: 100%">
              <el-table-column prop="project" label="项目" min-width="150" />
              <el-table-column prop="material" label="物料" width="120" />
              <el-table-column prop="inquiry_qty" label="询价数量" width="100" />
              <el-table-column prop="purchase_price" label="采购报价" width="100" />
              <el-table-column prop="lowest_price" label="历史最低价" width="100" />
              <el-table-column prop="agreement_price" label="协议价" width="100" />
              <el-table-column prop="deal_status" label="成交状态" width="100">
                <template #default="{ row }">
                  <el-tag :type="row.deal_status === '成交' ? 'success' : 'info'">{{ row.deal_status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="order_no" label="订单号" width="150" />
            </el-table>
          </div>

          <!-- VIEW: 中标项目管理 -->
          <div v-else-if="currentMenu === 'bid-projects'">
            <div class="form-section-title">中标项目管理</div>
            <el-alert title="管理已中标的项目，跟踪合同执行情况" type="success" show-icon style="margin-bottom: 20px;"></el-alert>

            <el-table :data="bidProjectsData" border style="width: 100%">
              <el-table-column prop="project_name" label="项目名称" min-width="180" />
              <el-table-column prop="customer" label="客户" width="150" />
              <el-table-column prop="win_date" label="中标日期" width="120" />
              <el-table-column prop="amount" label="合同金额" width="120" />
              <el-table-column prop="status" label="执行状态" width="100">
                <template #default="{ row }">
                  <el-tag :type="row.status === '执行中' ? 'primary' : 'success'">{{ row.status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="操作" width="150" fixed="right">
                <template #default>
                  <el-button type="primary" size="small">查看详情</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>

          <!-- VIEW: 吧盛支付申请 -->
          <div v-else-if="currentMenu === 'payment-approval'">
            <div class="form-section-title">吧盛支付申请</div>
            <p style="color: #909399;">支付申请功能开发中...</p>
          </div>

          <!-- VIEW: 采购订单审批 -->
          <div v-else-if="currentMenu === 'order-approval'">
            <div class="form-section-title">采购订单审批</div>
            <el-alert title="填写采购价时，系统将自动对比协议价、历史最低价进行提示" type="warning" show-icon style="margin-bottom: 20px;"></el-alert>

            <el-form :model="orderForm" label-width="150px" style="max-width: 800px">
              <el-form-item label="物料">
                <el-input v-model="orderForm.material" placeholder="选择或输入物料"></el-input>
              </el-form-item>
              <el-form-item label="供应商">
                <el-input v-model="orderForm.supplier" placeholder="选择供应商"></el-input>
              </el-form-item>
              <el-form-item label="采购数量">
                <el-input-number v-model="orderForm.quantity" :min="1"></el-input-number>
              </el-form-item>
              <el-form-item label="采购单价">
                <el-input v-model="orderForm.price" placeholder="输入采购单价">
                  <template #prepend>¥</template>
                </el-input>
              </el-form-item>

              <el-divider content-position="left">价格参考</el-divider>
              <el-form-item label="本年度最低采购价">
                <span style="color: #67c23a; font-weight: bold;">¥ 120</span>
              </el-form-item>
              <el-form-item label="上年度最后采购价">
                <span style="color: #909399;">¥ 130</span>
              </el-form-item>
              <el-form-item label="协议价">
                <span :style="{ color: orderForm.price > 110 ? '#f56c6c' : '#67c23a', fontWeight: 'bold' }">
                  ¥ 110
                  <el-icon v-if="orderForm.price > 110" style="color: #f56c6c; margin-left: 5px;"><Warning /></el-icon>
                  <span v-if="orderForm.price > 110" style="color: #f56c6c; font-size: 12px; margin-left: 5px;">高于协议限价</span>
                </span>
              </el-form-item>

              <el-form-item>
                <el-button type="primary">提交审批</el-button>
              </el-form-item>
            </el-form>
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
                <el-button type="primary" @click="currentMenu = 'apply'; isTempAdd = true;">+ 新增潜在供应商</el-button>
              </div>
            </div>

            <el-table :data="tempSupplierData" border style="width: 100%">
              <el-table-column prop="id" label="序号" width="70" />
              <el-table-column prop="entry_type" label="供应商类型" width="110" />
              <el-table-column prop="category" label="供应品类" width="100" />
              <el-table-column prop="industry" label="合作行业" width="100" />
              <el-table-column prop="brand" label="合作品牌" width="120" />
              <el-table-column prop="contact_person" label="联系人" width="100" />
              <el-table-column prop="mobile" label="联系电话" width="130" />
              <el-table-column label="操作" width="120" fixed="right">
                <template #default="{ row }">
                  <el-button type="primary" size="small" @click="viewSupplierDetail(row)">详情</el-button>
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
import { Menu, UploadFilled, Document, Delete, Edit, Search, Shop, Check, Setting, List, Trophy, Money, Warning } from '@element-plus/icons-vue'
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
const tableData = ref([])
const importDialogVisible = ref(false)
const techFiles = ref([])

// 当前用户和时间
const currentUser = computed(() => 'Admin')
const currentTime = computed(() => new Date().toLocaleString())

// 询价表单
const inquiryForm = reactive({
  project_name: '',
  customer_name: '',
  expected_delivery: '',
  technical_requirements: '',
  items: [{ material_code: '', material_name: '', specification: '', quantity: 1, unit: '个' }]
})

// 待报价搜索
const quoteSearch = reactive({ keyword: '', status: '' })

// 待报价数据
const pendingQuoteData = ref([
  { project_name: '中核项目-阀门采购', customer: '中核集团', material: '蝶阀DN100', quantity: 50, submit_time: '2023-11-25 10:00', status: '待报价' },
  { project_name: '石化项目-管件', customer: '中石化', material: '弯头90°', quantity: 100, submit_time: '2023-11-24 14:30', status: '已澄清' },
  { project_name: '电力项目-仪表', customer: '国家电网', material: '压力表', quantity: 30, submit_time: '2023-11-23 09:00', status: '已报价' }
])

// 台账搜索
const ledgerSearch = reactive({ dateRange: [], material: '', supplier: '', status: '', isWin: '' })

// 台账数据
const ledgerData = ref([
  { project: '中核项目', material: '蝶阀', inquiry_qty: 50, purchase_price: '¥120', lowest_price: '¥115', agreement_price: '¥110', deal_status: '成交', order_no: 'PO2023112501' },
  { project: '石化项目', material: '弯头', inquiry_qty: 100, purchase_price: '¥45', lowest_price: '¥42', agreement_price: '¥40', deal_status: '未成交', order_no: '-' }
])

// 中标项目数据
const bidProjectsData = ref([
  { project_name: '中核集团阀门采购项目', customer: '中核集团', win_date: '2023-11-20', amount: '¥580,000', status: '执行中' },
  { project_name: '国电管件供应项目', customer: '国家电网', win_date: '2023-11-15', amount: '¥320,000', status: '已完成' }
])

// 采购订单表单
const orderForm = reactive({ material: '', supplier: '', quantity: 1, price: 0 })

// 协议价搜索
const agreementSearch = reactive({ keyword: '' })

// 协议价数据
const agreementPriceData = ref([
  { material_code: 'M1001', material_name: '蝶阀DN100', supplier: 'A公司', price: 110, valid_date: '2025-01-01' },
  { material_code: 'M1002', material_name: '闸阀DN50', supplier: 'B公司', price: 85, valid_date: '2025-06-30' },
  { material_code: 'M1003', material_name: '球阀DN80', supplier: 'A公司', price: 95, valid_date: '2024-12-31' }
])

// 潜在供应商数据（字段：序号/供应商类型/供应品类/合作行业/合作品牌/联系人/联系电话）
const tempSupplierData = ref([
  { id: 1, entry_type: '生产商', category: '原材料', industry: '核电', brand: '品牌A', contact_person: '张三', mobile: '13800138001' },
  { id: 2, entry_type: '贸易商', category: '电子元器件', industry: '军工', brand: '品牌B', contact_person: '李四', mobile: '13800138002' },
  { id: 3, entry_type: '生产商', category: '办公用品', industry: '石化', brand: '品牌C', contact_person: '王五', mobile: '13800138003' }
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
  status: ''
})

const supplierList = reactive({
  total: 100,
  page: 1,
  pageSize: 10
})

const supplierListData = ref([
  {
    id: 1,
    name: '中核集团供应商A',
    industry: 'nuclear',
    industry_text: '核电',
    status: 'approved',
    contact_person: '张三',
    mobile: '13800138001',
    apply_time: '2023-11-20 10:00:00',
    approval_time: '2023-11-22 15:30:00',
    type: '生产商',
    credit_code: '91310000123456789A',
    attachments: [
      { name: '营业执照.pdf', url: '#' },
      { name: '质量体系认证.pdf', url: '#' }
    ]
  },
  {
    id: 2,
    name: '军工供应商B',
    industry: 'military',
    industry_text: '军工',
    status: 'approved',
    contact_person: '李四',
    mobile: '13800138002',
    apply_time: '2023-11-21 14:00:00',
    approval_time: '2023-11-23 16:20:00',
    type: '贸易商',
    credit_code: '91310000987654321B',
    attachments: [
      { name: '营业执照.pdf', url: '#' },
      { name: '保密资质.pdf', url: '#' }
    ]
  },
  {
    id: 3,
    name: '普通供应商C',
    industry: 'normal',
    industry_text: '普通行业',
    status: 'rejected',
    contact_person: '王五',
    mobile: '13800138003',
    apply_time: '2023-11-22 09:00:00',
    approval_time: '2023-11-22 17:00:00',
    type: '生产商',
    credit_code: '91310000111111111C',
    attachments: [
      { name: '营业执照.pdf', url: '#' }
    ]
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
            is_blacklist_str: item.status === 'blacklist' ? '是' : '否',
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
  inquiryForm.items.push({ material_code: '', material_name: '', specification: '', quantity: 1, unit: '个' })
}

const removeInquiryItem = (index) => {
  inquiryForm.items.splice(index, 1)
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

const submitInquiry = () => {
  if (!inquiryForm.project_name || !inquiryForm.customer_name) {
    ElMessage.error('请填写项目名称和客户名称')
    return
  }
  ElMessage.success('询价提交成功')
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

// 快速审批
const quickApprove = (row) => {
    ElMessageBox.confirm(`确认通过【${row.name}】的供应商申请？`, '快速审批', {
        confirmButtonText: '通过',
        cancelButtonText: '取消',
        type: 'success'
    }).then(() => {
        row.status = '已通过'
        row.current_node = '已完成'
        ElMessage.success('审批通过成功')
    }).catch(() => {})
}

// 重置审批搜索
const resetApprovalSearch = () => {
    searchForm.name = ''
    searchForm.approvalStatus = ''
    searchForm.industry = ''
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

// 提交审批
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
