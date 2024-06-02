<template>
  <!--  上边搜索 -->
  <el-card>
    <el-form :inline="true">
      <el-form-item label="订单号:">
        <el-input placeholder="请输入搜索的用户名" v-model="orderSn"></el-input>
      </el-form-item>
      <el-form-item label="用户名:">
        <el-input
          placeholder="请输入搜索的用户名"
          v-model="memberUsername"
        ></el-input>
      </el-form-item>
      <el-form-item label="状态:">
        <el-select v-model="status" class="m-2" placeholder="请选择状态">
          <el-option label="待付款" value="0" />
          <el-option label="待发货" value="1" />
          <el-option label="运输中" value="2" />
          <el-option label="待提货" value="3" />
          <el-option label="订单完成" value="4" />
          <el-option label="无效订单" value="5" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          @click="search"
          :disabled="
            orderSn.length || memberUsername.length || status.length
              ? false
              : true
          "
        >
          搜索
        </el-button>
        <el-button size="default" @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <!--  -->
  <el-card>
    <el-table border :data="listArr" style="margin: 15px 0">
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
        prop="orderSn"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="用户名"
        align="center"
        prop="memberUserName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="支付方式"
        align="center"
        prop="payType"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.payType === '1'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              微信
            </el-tag>
          </template>
          <template v-if="row.payType === '2'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              支付宝
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <el-table-column
        label="订单类型"
        align="center"
        prop="orderType"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.orderType === '1'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              正常订单
            </el-tag>
          </template>
          <template v-if="row.orderType === '2'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              秒杀订单
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <el-table-column label="用户信息" align="center" show-overflow-tooltip>
        <template #="{ row }">
          <span>{{ row.receiverName + row.receiverPhone }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="总金额"
        align="center"
        prop="totalAmount"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <p style="color: red">¥{{ row.totalAmount }}</p>
        </template>
      </el-table-column>
      <el-table-column
        label="因付金额"
        align="center"
        prop="payAmount"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <p style="color: red">¥{{ row.payAmount }}</p>
        </template>
      </el-table-column>
      <el-table-column
        label="状态"
        align="center"
        prop="status"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.status === '0'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              待付款
            </el-tag>
          </template>
          <template v-if="row.status === '1'">
            <el-tag key="item.label" class="mx-1" type="primary" effect="light">
              待发货
            </el-tag>
          </template>
          <template v-if="row.status === '2'">
            <el-tag key="item.label" class="mx-1" type="Primary" effect="light">
              运输中
            </el-tag>
          </template>
          <template v-if="row.status === '3'">
            <el-tag key="item.label" class="mx-1" type="Warning" effect="light">
              待提货
            </el-tag>
          </template>
          <template v-if="row.status === '4'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              订单完成
            </el-tag>
          </template>
          <template v-if="row.status === '5'">
            <el-tag key="item.label" class="mx-1" type="Info" effect="light">
              无效订单
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="300px" align="center">
        <template #="{ row }">
          <el-button type="primary" size="small" icon="Edit" @click="look(row)">
            查看订单
          </el-button>
          <el-button
            v-if="row.status == '1' && !isZTD"
            type="danger"
            size="small"
            icon="Edit"
            @click="toDeliver(row)"
          >
            订单发货
          </el-button>
          <el-button
            v-if="row.status == '2' && !isSJ"
            type="warning"
            size="small"
            icon="Edit"
            @click="addone(row)"
          >
            已到店
          </el-button>
          <el-button
            v-if="row.status == '3' && !isSJ"
            type="danger"
            size="small"
            icon="Edit"
            @click="addone(row)"
          >
            完成提货
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

  <!-- 订单详情 -->
  <el-dialog v-model="drawer" title="订单详情" width="70%">
    <el-steps :active="orderInfo?.status" align-center finish-status="success">
      <el-step title="支付订单" description="Some description" />
      <el-step title="运输中" description="Some description" />
      <el-step title="已到店" description="Some description" />
      <el-step title="订单完成" description="Some description" />
    </el-steps>
    <!--  -->
    <!-- 商品 -->
    <h2>商品信息</h2>
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
      title="订单详情"
      direction="horizontal"
      :column="1"
      :size="size"
      border
      style="margin: 40px 0 0 20px"
    >
      <el-descriptions-item label="订单号">{{
        orderInfo?.orderSn
      }}</el-descriptions-item>
      <el-descriptions-item label="用户名">{{
        orderInfo?.memberUserName
      }}</el-descriptions-item>
      <el-descriptions-item label="总金额">
        <p style="color: red">¥{{ orderInfo?.totalAmount }}</p>
      </el-descriptions-item>
      <el-descriptions-item label="因付金额">
        <p style="color: red">¥{{ orderInfo?.payAmount }}</p>
      </el-descriptions-item>
      <el-descriptions-item label="支付方式">
        <el-tag class="mx-1" effect="success" v-if="orderInfo?.payType == '1'">
          微信
        </el-tag>
        <el-tag class="mx-1" effect="dark" v-if="orderInfo?.payType == '2'">
          支付宝
        </el-tag>
      </el-descriptions-item>
      <el-descriptions-item label="订单类型">
        <el-tag
          class="mx-1"
          effect="success"
          v-if="orderInfo?.orderType == '1'"
        >
          正常订单
        </el-tag>
        <el-tag
          class="mx-1"
          type="dark"
          effect="dark"
          v-if="orderInfo?.orderType == '2'"
        >
          秒杀订单
        </el-tag>
      </el-descriptions-item>
      <el-descriptions-item label="备注">{{
        orderInfo?.note
      }}</el-descriptions-item>
      <el-descriptions-item label="支付时间">{{
        orderInfo?.paymentTime
      }}</el-descriptions-item>
      <el-descriptions-item label="发货时间">{{
        orderInfo?.deliveryTime
      }}</el-descriptions-item>
      <el-descriptions-item label="确认收货时间">{{
        orderInfo?.receiveTime
      }}</el-descriptions-item>
      <el-descriptions-item label="评论时间">{{
        orderInfo?.commentTime
      }}</el-descriptions-item>
    </el-descriptions>
    <!--  -->
    <el-descriptions
      title="收货人信息"
      direction="horizontal"
      :column="1"
      :size="size"
      border
      style="margin: 40px 0 0 20px"
    >
      <el-descriptions-item label="收货人">{{
        orderInfo?.receiverName
      }}</el-descriptions-item>
      <el-descriptions-item label="联系电话">{{
        orderInfo?.receiverPhone
      }}</el-descriptions-item>
    </el-descriptions>
    <el-descriptions
      title="商家信息"
      direction="horizontal"
      :column="1"
      :size="size"
      border
      style="margin: 40px 0 0 20px"
    >
      <el-descriptions-item label="自提点图片">
        <img :src="merchantInfo?.pic" alt="" style="width: 200px" />
      </el-descriptions-item>
      <el-descriptions-item label="名称">{{
        merchantInfo?.name
      }}</el-descriptions-item>
      <el-descriptions-item label="详细地址">{{
        merchantInfo?.address
      }}</el-descriptions-item>
      <el-descriptions-item label="负责人">{{
        merchantInfo?.principal
      }}</el-descriptions-item>
      <el-descriptions-item label="联系电话">{{
        merchantInfo?.phone
      }}</el-descriptions-item>
    </el-descriptions>
    <el-descriptions
      title="自提点信息"
      direction="horizontal"
      :column="1"
      :size="size"
      border
      style="margin: 40px 0 0 20px"
    >
      <el-descriptions-item label="自提点图片">
        <img :src="placeInfo?.pic" alt="" style="width: 200px" />
      </el-descriptions-item>
      <el-descriptions-item label="名称">{{
        placeInfo?.name
      }}</el-descriptions-item>
      <el-descriptions-item label="详细地址">{{
        placeInfo?.place
      }}</el-descriptions-item>
      <el-descriptions-item label="联系电话">{{
        placeInfo?.phone
      }}</el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import {
  reqOrderAll,
  reqOrderInfo,
  reqAddOrUpdateOrder,
} from "@/api/oms/order";
import { ElMessage } from "element-plus";
//setting仓库
import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();
const size = ref("");
//默认页码
let pageNo = ref<number>(1);
//默认个数
let pageSize = ref<number>(10);
let total = ref<number>(0);
//数据列表
let listArr = ref<any>([]);
//订单详情
// 定义响应式数据
const orderInfo = ref<any>(null);
const skuList = ref<any[]>([]);
const placeInfo = ref<any>(null);
const merchantInfo = ref<any>(null);
//收集用户查找的关键字
let orderSn = ref<string>("");
let memberUsername = ref<string>("");
let status = ref<string>("");
//定义响应式数据 抽屉的显示隐藏
//定义响应式数据 抽屉的显示隐藏
let drawer = ref<boolean>(false);
//组件挂载完毕
onMounted(() => {
  getHas();
});
//获取信息
const getHas = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqOrderAll(
    pageNo.value,
    pageSize.value,
    orderSn.value,
    memberUsername.value,
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
//搜索按钮
const search = () => {
  getHas();
  orderSn.value = "";
  memberUsername.value = "";
};
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};

//查看订单详情
let look = async (row: any) => {
  drawer.value = true;
  let res: any = await reqOrderInfo({ id: row.id });
  if (res.code === 200) {
    orderInfo.value = res.data.orderInfo;
    skuList.value = res.data.skuList;
    placeInfo.value = res.data.placeInfo;
    merchantInfo.value = res.data.merchantInfo;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
let addone = async (row: any) => {
  let data: any = {
    id: row.id as number,
    status: String(parseInt(row.status) + 1),
  };
  let res = await reqAddOrUpdateOrder(data);
  if (res.code === 200) {
    ElMessage({ type: "success", message: res.message });
    window.location.reload();
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

let getCurrentDateTime = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, "0"); // 月份从 0 开始，需要加 1，并补零
  const day = String(now.getDate()).padStart(2, "0"); // 补零
  const hours = String(now.getHours()).padStart(2, "0"); // 补零
  const minutes = String(now.getMinutes()).padStart(2, "0"); // 补零
  const seconds = String(now.getSeconds()).padStart(2, "0"); // 补零
  const dateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
  return dateTime;
};
const currentDateTime = getCurrentDateTime();
let toDeliver = async (row: any) => {
  let data: any = {
    id: row.id as number,
    status: "2",
    deliveryTime: currentDateTime,
  };
  let res = await reqAddOrUpdateOrder(data);
  if (res.code === 200) {
    ElMessage({ type: "success", message: "修改成功" });
    window.location.reload();
  }
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

<style scoped>
h2 {
  font-size: 20px;
  margin: 20px 0 40px 0;
}
</style>
