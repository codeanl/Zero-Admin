<template>
  <!-- 数据展示 -->
  <el-card>
    <el-button type="success" size="default" icon="Plus" @click="addPermission">
      添加菜单
    </el-button>
    <!--  -->
    <el-table style="margin: 20px 0" row-key="id" border :data="PermissionArr">
      <el-table-column prop="name" label="名称" />
      <!-- <el-table-column prop="type" label="类型" /> -->
      <el-table-column
        label="类型"
        align="center"
        prop="type"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.type === '1'">
            <el-tag key="item.label" class="mx-1" type="success" effect="light">
              目录
            </el-tag>
          </template>
          <template v-if="row.type === '2'">
            <el-tag key="item.label" class="mx-1" type="danger" effect="light">
              菜单
            </el-tag>
          </template>
          <template v-if="row.type === '3'">
            <el-tag key="item.label" class="mx-1" type="warning" effect="light">
              按钮
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="路径" />
      <el-table-column prop="tag" label="标识" />
      <el-table-column prop="orderNum" label="排序" />
      <el-table-column prop="remark" label="备注" />
      <el-table-column label="操作" width="260px">
        <template #="{ row }">
          <el-button type="primary" size="small" @click="updatePermission(row)">
            编辑
          </el-button>
          <el-popconfirm
            :title="`你确定要删除${row.name}?`"
            width="260px"
            @confirm="removeMenu(row.id)"
          >
            <template #reference>
              <el-button type="danger" size="small"> 删除 </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  <!-- 添加｜编辑 -->
  <el-dialog v-model="dialogVisible" :title="menuData.id ? '更新' : '添加'">
    <el-form :rules="rules" ref="form">
      <el-form-item label="名称" prop="name">
        <el-input
          placeholder="请你输入菜单的名称"
          v-model="menuData.name"
        ></el-input>
      </el-form-item>
      <el-form-item label="类型">
        <el-select v-model="menuData.type" class="m-2" placeholder="请选择类型">
          <el-option label="目录" value="1" />
          <el-option label="菜单" value="2" />
          <el-option label="按钮" value="3" />
        </el-select>
      </el-form-item>
      <el-form-item label="父级id">
        <el-tree-select
          v-model="menuData.parentId"
          :data="PermissionArrWithRoot"
          check-strictly
          :props="{ key: 'id', label: 'name' }"
          node-key="id"
          :render-after-expand="false"
        />
      </el-form-item>
      <el-form-item label="路径">
        <el-input placeholder="请你输入路径" v-model="menuData.url"></el-input>
      </el-form-item>
      <el-form-item label="排序">
        该菜单排在第<el-input-number
          v-model="menuData.orderNum"
          :min="1"
          size="small"
          controls-position="right"
        />位
      </el-form-item>
      <el-form-item label="标识">
        <el-input placeholder="请你输入标识" v-model="menuData.tag"></el-input>
      </el-form-item>
      <el-form-item label="备注">
        <el-input
          placeholder="请你输入备注"
          v-model="menuData.remark"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="save">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, computed } from "vue";
import { ElTreeSelect, ElMessage } from "element-plus";
import {
  reqAllPermission,
  reqAddOrUpdateMenu,
  reqRemoveMenu,
} from "@/api/sys/menu";

let PermissionArr = ref<any>([]);
let dialogVisible = ref<boolean>(false);
let ids = ref<number[]>([]);

let menuData = reactive<any>({
  name: "",
  url: "",
  tag: "",
  orderNum: 0,
  remark: "",
  parentId: 0,
  type: "",
});

onMounted(() => {
  getHasPermission();
});

const getHasPermission = async () => {
  let res: any = await reqAllPermission();
  if (res.code === 200) {
    PermissionArr.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
const PermissionArrWithRoot = computed(() => {
  return [{ id: 0, name: "无上级" }, ...PermissionArr.value];
});

const addPermission = () => {
  Object.assign(menuData, {
    id: 0,
    name: "",
    url: "",
    tag: "",
    parentId: 0,
    orderNum: 0,
    remark: "",
    type: "",
  });
  dialogVisible.value = true;
};

const updatePermission = (row: any) => {
  dialogVisible.value = true;
  menuData.parentId = parseInt(row.parentId);
  Object.assign(menuData, row);
};

const save = async () => {
  dialogVisible.value = false;
  let res: any = await reqAddOrUpdateMenu(menuData);
  if (res.code === 200) {
    dialogVisible.value = false;
    getHasPermission();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

const removeMenu = async (id: number) => {
  ids.value.push(id);
  const requestData: any = { ids: ids.value };
  let res = await reqRemoveMenu(requestData);
  if (res.code == 200) {
    ids.value = [];
    getHasPermission();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//form组件实例
let form = ref<any>();
const rules = reactive<any>({});
</script>

<style lang="scss" scoped></style>
