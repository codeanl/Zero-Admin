<template>
  <!-- 上边搜索 -->
  <el-card>
    <el-form :inline="true">
      <el-form-item label="名称:">
        <el-input placeholder="请输入搜索的用户名" v-model="name"></el-input>
      </el-form-item>
      <el-form-item label="优惠券类型:">
        <el-select v-model="type" class="m-2" placeholder="请选择状态">
          <el-option label="全场赠券" value="0" />
          <el-option label="购物赠券" value="1" />
          <el-option label="注册赠券" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="使用类型:">
        <el-select v-model="useType" class="m-2" placeholder="请选择状态">
          <el-option label="全场通用" value="0" />
          <el-option label="指定分类" value="1" />
          <el-option label="指定商品" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          @click="search"
          :disabled="
            name.length || type.length || useType.length ? false : true
          "
        >
          搜索
        </el-button>
        <el-button size="default" @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <!--  -->
  <!--  -->
  <el-card>
    <el-button type="success" size="default" @click="add"> 添加 </el-button>
    <el-button
      type="danger"
      size="default"
      @click="deleteSelect"
      :disabled="selectIdArr.length ? false : true"
    >
      批量删除
    </el-button>
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
        label="类型"
        align="center"
        prop="type"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.type === '0'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              全场赠券
            </el-tag>
          </template>
          <template v-if="row.type === '1'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              购物赠券
            </el-tag>
          </template>
          <template v-if="row.type === '2'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              注册赠券
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <el-table-column
        label="使用类型"
        align="center"
        prop="useType"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.useType === '0'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              全场通用
            </el-tag>
          </template>
          <template v-if="row.useType === '1'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              指定分类
            </el-tag>
          </template>
          <template v-if="row.useType === '2'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              指定商品
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <el-table-column
        label="名称"
        align="center"
        prop="name"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="数量"
        align="center"
        prop="count"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="金额"
        align="center"
        prop="amount"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="每人限领张数"
        align="center"
        prop="perLimit"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="使用门槛"
        align="center"
        prop="minPoint"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="开始时间"
        align="center"
        prop="startTime"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="结束时间"
        align="center"
        prop="endTime"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="开始领取时间"
        align="center"
        prop="enableTime"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="备注"
        align="center"
        prop="note"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="优惠码"
        align="center"
        prop="code"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column label="操作" width="200px" align="center">
        <template #="{ row }">
          <el-button
            type="primary"
            size="small"
            icon="Edit"
            @click="update(row)"
          >
            编辑
          </el-button>
          <el-popconfirm
            :title="`你确定删除${row.name}`"
            width="260px"
            @confirm="deletePlace(row.id)"
          >
            <template #reference>
              <el-button type="danger" size="small" icon="Delete">
                删除
              </el-button>
            </template>
          </el-popconfirm>
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

  <!-- 抽屉  完成 添加｜修改 的窗口 -->
  <el-dialog v-model="drawer" :title="Params.id ? '更新' : '添加'">
    <el-form :model="Params" ref="formRef">
      <el-form-item label="名称" prop="name">
        <el-input placeholder="请您输入名称" v-model="Params.name"></el-input>
      </el-form-item>
      <el-form-item label="类型" prop="type">
        <el-select v-model="Params.type" class="m-2" placeholder="请选择性别">
          <el-option label="全场赠券" value="0" />
          <el-option label="购物赠券" value="1" />
          <el-option label="注册赠券" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="使用类型" prop="useType">
        <el-select
          v-model="Params.useType"
          class="m-2"
          placeholder="请选择使用类型"
        >
          <el-option label="全场通用" value="0" />
          <el-option label="指定分类" value="1" />
          <el-option label="指定商品" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="数量" prop="count">
        <el-input placeholder="请您输入数量" v-model="Params.count"></el-input>
      </el-form-item>
      <el-form-item label="金额" prop="amount">
        <el-input placeholder="请您输入金额" v-model="Params.amount"></el-input>
      </el-form-item>
      <el-form-item label="每人限领张数" prop="perLimit">
        <el-input
          placeholder="请您输入每人限领张数"
          v-model="Params.perLimit"
        ></el-input>
      </el-form-item>
      <el-form-item label="使用门槛" prop="minPoint">
        <el-input
          placeholder="请您输入使用门槛"
          v-model="Params.minPoint"
        ></el-input>
      </el-form-item>
      <el-form-item label="起止时间" prop="endTime">
        <div class="block">
          <el-date-picker
            v-model="Params.starEndTime"
            type="datetimerange"
            range-separator="到"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </div>
      </el-form-item>

      <el-form-item label="领取时间" prop="enableTime">
        <div class="block">
          <el-date-picker
            v-model="Params.enableTime"
            type="datetime"
            placeholder="请您输入开始领取时间"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </div>
      </el-form-item>
      <el-form-item label="备注" prop="note">
        <el-input placeholder="请您输入备注" v-model="Params.note"></el-input>
      </el-form-item>
      <el-form-item label="优惠码" prop="code">
        <el-input placeholder="请您输入优惠码" v-model="Params.code"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="save">确定</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, nextTick } from "vue";
import { ElMessage } from "element-plus";
import {
  reqcouponList,
  reqRemovecoupon,
  reqAddOrUpdate,
} from "@/api/sms/coupon";
//setting仓库
import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();

//默认页码
let pageNo = ref<number>(1);
//默认个数
let pageSize = ref<number>(10);
let total = ref<number>(0);
//数据列表
let listArr = ref<any>([]);
//收集用户查找的关键字
let name = ref<string>("");
let type = ref<string>("");
let useType = ref<string>("");
//收集删除的id
let ids = ref<number[]>([]);
//多选框选择的id
let selectIdArr = ref<any[]>([]);
//复选框选择
const selectChange = (value: any) => {
  selectIdArr.value = value;
};
//组件实例
let formRef = ref<any>();
//定义响应式数据 抽屉的显示隐藏
let drawer = ref<boolean>(false);
let Params = reactive<any>({
  type: "",
  name: "",
  count: "",
  amount: "",
  perLimit: "",
  minPoint: "",
  startTime: "",
  endTime: "",
  enableTime: "",
  useType: "",
  note: "",
  code: "",
  starEndTime: [],
});
//组件挂载完毕
onMounted(() => {
  getHas();
});
//获取信息
const getHas = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqcouponList(
    pageNo.value,
    pageSize.value,
    type.value,
    name.value,
    useType.value
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
//取消按钮
const cancel = () => {
  drawer.value = false;
};
//窗口保存按钮
const save = async () => {
  Params.amount = parseFloat(Params.amount);
  Params.minPoint = parseFloat(Params.minPoint);
  Params.count = parseInt(Params.count);
  Params.perLimit = parseInt(Params.perLimit);
  Params.startTime = Params.starEndTime[0];
  Params.endTime = Params.starEndTime[1];
  let res: any = await reqAddOrUpdate(Params);
  if (res.code == 200) {
    drawer.value = false;
    ElMessage({ type: "success", message: res.message });
    getHas();
  } else {
    drawer.value = false;
    ElMessage({ type: "error", message: res.message });
  }
};
//添加按钮
const add = () => {
  drawer.value = true;
  //存储收集已有的账号信息
  Object.assign(Params, {
    id: 0,
    type: "",
    name: "",
    count: "",
    amount: "",
    perLimit: "",
    minPoint: "",
    startTime: "",
    endTime: "",
    enableTime: "",
    useType: "",
    note: "",
    code: "",
    starEndTime: [],
  });
  nextTick(() => {
    formRef.value.clearValidate("name");
    formRef.value.clearValidate("place");
    formRef.value.clearValidate("status");
    formRef.value.clearValidate("pic");
    formRef.value.clearValidate("phone");
    formRef.value.clearValidate("principal");
  });
};
// 编辑按钮
const update = (row: any) => {
  drawer.value = true;
  Object.assign(Params, row);
  Params.starEndTime.push(Params.startTime);
  Params.starEndTime.push(Params.endTime);
  nextTick(() => {
    formRef.value.clearValidate("name");
    formRef.value.clearValidate("place");
    formRef.value.clearValidate("status");
    formRef.value.clearValidate("pic");
    formRef.value.clearValidate("phone");
    formRef.value.clearValidate("principal");
  });
};
// 删除按钮
const deletePlace = async (id: number) => {
  ids.value.push(id);
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemovecoupon(requestData);
  if (res.code == 200) {
    getHas(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//批量删除用户按钮
const deleteSelect = async () => {
  ids.value = selectIdArr.value.map((item) => {
    return item.id;
  });
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemovecoupon(requestData);
  if (res.code === 200) {
    getHas(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};
//搜索按钮
const search = () => {
  getHas();
  name.value = "";
  type.value = "";
  useType.value = "";
};
</script>

<style scoped>
.block {
  border-right: solid 1px var(--el-border-color);
  flex: 1;
}

.block:last-child {
  border-right: none;
}

.block .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
}
</style>
