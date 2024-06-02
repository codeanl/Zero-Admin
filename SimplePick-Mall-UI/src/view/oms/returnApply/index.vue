<template>
  <!--  上边搜索 -->
  <el-card>
    <el-form :inline="true">
      <el-form-item label="状态:">
        <el-select v-model="status" class="m-2" placeholder="请选择状态">
          <!-- ：0->待处理；1->退货中；2->已完成；3->已拒绝' -->
          <el-option label="待处理" value="0" />
          <el-option label="退货中" value="1" />
          <el-option label="已完成" value="2" />
          <el-option label="已拒绝" value="3" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          @click="search"
          :disabled="status.length ? false : true"
        >
          搜索
        </el-button>
        <el-button size="default" @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <!--  -->
  <el-card>
    <el-table
      border
      :data="listArr"
      @selection-change="selectChange"
      style="margin: 15px 0"
    >
      <el-table-column
        type="selection"
        align="center"
        width="40px"
      ></el-table-column>
      <el-table-column
        label="id"
        align="center"
        prop="id"
        width="50px"
      ></el-table-column>
      <el-table-column
        label="订单号"
        align="center"
        prop="order.orderInfo.orderSn"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="退款金额"
        align="center"
        prop="returnAmount"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <p style="color: red">¥{{ row.returnAmount }}</p>
        </template>
      </el-table-column>
      <el-table-column
        label="退货原因"
        align="center"
        prop="returnReasonName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="用户名"
        align="center"
        prop="user.username"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="状态"
        align="center"
        prop="status"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.status === '0'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              待处理
            </el-tag>
          </template>
          <template v-if="row.status === '1'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              退货中
            </el-tag>
          </template>
          <template v-if="row.status === '2'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              已完成
            </el-tag>
          </template>
          <template v-if="row.status === '3'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              已拒绝
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="300px" align="center">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="look(row)">
            查看详情
          </el-button>
          <el-button
            v-if="!isZTD"
            type="danger"
            size="small"
            icon="Edit"
            @click="update(row)"
          >
            退货处理
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页 -->
    <el-pagination
      v-model:current-page="pageNo"
      v-model:page-size="pageSize"
      :page-sizes="[5, 10, 15, 20]"
      :background="true"
      layout="prev, pager, next, jumper, -> , sizes, total"
      :total="total"
      @current-change="getHas"
      @size-change="handler"
    />
  </el-card>
  <el-dialog v-model="drawer" title="订单详情" width="70%">
    <el-table :data="skuList">
      <el-table-column
        label="商品图片"
        align="center"
        prop="pic"
        show-overflow-tooltip
        width="120px"
      >
        <template #="{ row }">
          <img
            :src="row.pic"
            alt="商品图片"
            style="width: 100px; height: auto"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="货号"
        align="center"
        prop="skuSn"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="商品名称"
        align="center"
        prop="productName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="规格"
        align="center"
        prop="tag"
        show-overflow-tooltip
        width="300px"
      ></el-table-column>
      <el-table-column
        label="价格"
        align="center"
        prop="price"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <p style="color: red">¥{{ row.price }}</p>
        </template>
      </el-table-column>
    </el-table>
    <!--  -->
    <el-descriptions
      title="服务单信息"
      direction="horizontal"
      :column="1"
      :size="size"
      border
      style="margin: 40px 0 0 20px"
    >
      <el-descriptions-item label="订单号">{{
        orderInfo.order.orderInfo.orderSn
      }}</el-descriptions-item>
      <el-descriptions-item label="用户账号">{{
        orderInfo.user.username
      }}</el-descriptions-item>
      <el-descriptions-item label="联系人">{{
        orderInfo.order.orderInfo.receiverName
      }}</el-descriptions-item>
      <el-descriptions-item label="联系电话">{{
        orderInfo.order.orderInfo.receiverPhone
      }}</el-descriptions-item>
      <el-descriptions-item label="退货原因">{{
        orderInfo.returnReasonName
      }}</el-descriptions-item>
      <el-descriptions-item label="问题描述">{{
        orderInfo.description
      }}</el-descriptions-item>
      <el-descriptions-item label="订单金额">{{
        orderInfo.order.orderInfo.payAmount
      }}</el-descriptions-item>
      <el-descriptions-item label="退货金额">
        <p style="color: red">¥{{ orderInfo.returnAmount }}</p>
      </el-descriptions-item>
      <el-descriptions-item label="凭证图片">
        <img :src="orderInfo.proofPics" alt="" style="width: 400px" />
      </el-descriptions-item>
    </el-descriptions>
  </el-dialog>
  <el-dialog v-model="drawer1" title="退货处理">
    <el-radio-group v-model="radio1" size="large" @change="change">
      <el-radio-button label="待处理" disabled="true" />
      <el-radio-button label="允许退货" />
      <el-radio-button label="完成退货" />
      <el-radio-button label="拒绝退货" />
    </el-radio-group>
  </el-dialog>
</template>

<script setup lang="ts">
// ------------------------------------------------------------
import { ref, onMounted } from "vue";
import {
  reqReturnApplyAll,
  reqAddOrUpdateReturnApply,
} from "@/api/oms/returnApply";
import { ElMessage } from "element-plus";
//setting仓库
import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();
//
const size = ref("");
//默认页码
let pageNo = ref<number>(1);
//默认个数
let pageSize = ref<number>(10);
let total = ref<number>(0);
//数据列表
let listArr = ref<any>([]);
let skuList = ref<any>([]);
let orderInfo = ref<any>([]);

let memberUsername = ref<any>("");
let status = ref<any>("");
//收集删除的id
//多选框选择的id
let selectIdArr = ref<any[]>([]);
//复选框选择
const selectChange = (value: any) => {
  selectIdArr.value = value;
};
//组件实例
//定义响应式数据 抽屉的显示隐藏
let drawer = ref<boolean>(false);
let drawer1 = ref<boolean>(false);

//组件挂载完毕
onMounted(() => {
  getHas();
});
//获取信息
const getHas = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqReturnApplyAll(
    pageNo.value,
    pageSize.value,
    status.value
  );
  if (res.code == 200) {
    total.value = res.total;
    listArr.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//下拉改变
const handler = () => {
  getHas();
};

//查看订单详情
let look = async (row: any) => {
  drawer.value = true;
  skuList.value = row.order.skuList;
  orderInfo.value = row;
};

let nowID = ref<any>(null);
let radio1 = ref<any>(null);
let update = async (row: any) => {
  nowID.value = row.id;
  drawer1.value = true;
  if (row.status == "0") {
    radio1.value = "待处理";
  } else if (row.status == "1") {
    radio1.value = "允许退货";
  } else if (row.status == "2") {
    radio1.value = "完成退货";
  } else if (row.status == "3") {
    radio1.value = "拒绝退货";
  }
};
let change = async () => {
  let data: any = {
    id: nowID.value as number,
    status: getStatusValue(radio1.value),
  };
  let res = await reqAddOrUpdateReturnApply(data);
  if (res.code === 200) {
    getHas();
    drawer1.value = false;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
function getStatusValue(label: string): string | number {
  const options = [
    { label: "待处理", value: "0" },
    { label: "允许退货", value: "1" },
    { label: "完成退货", value: "2" },
    { label: "拒绝退货", value: "3" },
  ];
  const option = options.find((opt) => opt.label === label);
  return option ? option.value : label;
}

//搜索按钮
const search = () => {
  getHas();
  status.value = "";
  memberUsername.value = "";
};
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};

import useUserStore from "@/store/sys/user";
let userStore = useUserStore();
let isSJ = ref<any>(false);
let isZTD = ref<any>(false);
let isAD = ref<any>(false);
for (let index = 0; index < userStore.roles.length; index++) {
  if (userStore.roles[index].includes("商家")) {
    isSJ.value = true;
  }
  if (userStore.roles[index].includes("自提点")) {
    isZTD.value = true;
  }
  if (userStore.roles[index].includes("超级管理员")) {
    isAD.value = true;
  }
}
</script>

<style scoped></style>
