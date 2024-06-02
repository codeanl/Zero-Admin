<template>
  <el-card>
    <el-button
      type="danger"
      size="default"
      @click="deleteSelectLoginLog"
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
        width="80px"
      ></el-table-column>
      <el-table-column
        label="头像"
        align="center"
        prop="avatar"
        show-overflow-tooltip
        width="60px"
      >
        <template #="{ row }">
          <img
            :src="row.avatar"
            alt="头像"
            style="width: 40px; height: 40px; border-radius: 50%"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="用户名字"
        align="center"
        prop="username"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="IP"
        align="center"
        prop="ip"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="创建时间"
        align="center"
        prop="createTime"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column label="操作" align="center">
        <template #="{ row }">
          <el-popconfirm
            :title="`你确定删除${row.username}`"
            width="260px"
            @confirm="deleteLog(row.id)"
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
      @current-change="getHasUser"
      @size-change="handler"
    />
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { reqLoginLog, reqRemoveLoginLog } from "@/api/sys/log";
import { ElMessage } from "element-plus";
//setting仓库
// import useLayoutSettingStore from "@/store/setting";
// let settingStore = useLayoutSettingStore();
//默认页码
let pageNo = ref<number>(1);
//默认个数
let pageSize = ref<number>(10);
let total = ref<number>(0);
//数据列表
let listArr = ref<any>([]);
//收集用户查找的关键字
let status = ref<string>("");
//收集删除的id
let ids = ref<number[]>([]);
//多选框选择的id
let selectIdArr = ref<any[]>([]);
//复选框选择
const selectChange = (value: any) => {
  selectIdArr.value = value;
};
//组件挂载完毕
onMounted(() => {
  getHasUser();
});
//获取用户信息
const getHasUser = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqLoginLog(pageNo.value, pageSize.value, status.value);
  if (res.code == 200) {
    ElMessage({ type: "success", message: res.message });
    total.value = res.total;
    listArr.value = res.data;
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
// //搜索按钮
// const search = () => {
//   getHasUser();
//   status.value = "";
// };

//下拉改变
const handler = () => {
  getHasUser();
};

// 删除用户按钮
const deleteLog = async (id: number) => {
  ids.value.push(id);
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemoveLoginLog(requestData);
  if (res.code == 200) {
    ElMessage({ type: "success", message: res.message });
    getHasUser(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//批量删除用户按钮
const deleteSelectLoginLog = async () => {
  ids.value = selectIdArr.value.map((item) => {
    return item.id;
  });
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemoveLoginLog(requestData);
  if (res.code === 200) {
    ElMessage({ type: "success", message: res.message });
    getHasUser(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
// //重置按钮
// const reset = () => {
//   settingStore.refresh = !settingStore.refresh;
// };
</script>

<style scoped lang="scss">
.form {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
