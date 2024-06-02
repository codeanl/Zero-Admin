<template>
  <!-- 搜索 -->
  <el-card style="height: 80px">
    <el-form :inline="true" class="form">
      <el-form-item label="角色:">
        <el-input
          placeholder="请你输入搜索角色名称"
          v-model="keyword"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          :disabled="keyword ? false : true"
          @click="search"
        >
          搜索
        </el-button>
        <el-button size="default" @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <!--  -->
  <el-card>
    <el-button type="success" size="default" icon="Plus" @click="addRole">
      添加
    </el-button>
    <el-button
      type="danger"
      size="default"
      @click="deleteSelectRole"
      :disabled="selectIdArr.length ? false : true"
    >
      批量删除
    </el-button>
    <el-table
      border
      style="margin: 15px 0"
      :data="allRole"
      @selection-change="selectChange"
    >
      <el-table-column
        type="selection"
        align="center"
        width="40px"
      ></el-table-column>
      <el-table-column
        label="ID"
        align="center"
        prop="id"
        width="50px"
      ></el-table-column>
      <el-table-column
        label="角色名称"
        align="center"
        show-overflow-tooltip
        prop="name"
      ></el-table-column>
      <el-table-column
        label="创建者"
        align="center"
        show-overflow-tooltip
        prop="create_by"
      ></el-table-column>
      <el-table-column
        label="创建时间"
        align="center"
        show-overflow-tooltip
        prop="create_at"
      ></el-table-column>
      <el-table-column
        label="更新者"
        align="center"
        show-overflow-tooltip
        prop="update_by"
      ></el-table-column>
      <el-table-column
        label="更新时间"
        align="center"
        show-overflow-tooltip
        prop="update_at"
      ></el-table-column>
      <el-table-column
        label="备注"
        align="center"
        show-overflow-tooltip
        prop="remark"
      ></el-table-column>
      <el-table-column label="操作" width="280px" align="center">
        <template #="{ row }">
          <el-button
            size="small"
            type="warning"
            icon="User"
            @click="setPermission(row)"
          >
            分配权限
          </el-button>
          <el-button
            type="primary"
            size="small"
            icon="Edit"
            @click="updateRole(row)"
          >
            编辑
          </el-button>
          <el-popconfirm
            :title="`你确定要删除${row.name}?`"
            width="260px"
            @confirm="removeRole(row.id)"
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
    <!--  -->
    <el-pagination
      v-model:current-page="pageNo"
      v-model:page-size="pageSize"
      :page-sizes="[10, 20, 30, 40]"
      :background="true"
      layout="prev, pager, next, jumper , ->, sizes, total, "
      :total="total"
      @current-change="getHasRole"
      @size-change="sizeHandler"
    />
  </el-card>
  <!-- 对话框 添加｜修改-->
  <el-dialog v-model="dialogVisible" :title="RoleParams.id ? '更新' : '添加'">
    <el-form :model="RoleParams" :rules="rules" ref="form">
      <el-form-item label="角色名称" prop="name">
        <el-input
          placeholder="请你输入角色名称"
          v-model="RoleParams.name"
        ></el-input>
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input
          placeholder="请你输入角色名称"
          v-model="RoleParams.remark"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button size="default" @click="dialogVisible = false">取消</el-button>
      <el-button type="primary" size="default" @click="save">确定</el-button>
    </template>
  </el-dialog>
  <!-- 分配权限抽屉 -->
  <el-drawer v-model="drawer">
    <template #header>
      <h4>分配菜单与按钮的权限</h4>
    </template>
    <template #default>
      <!-- default-expand-all默认展开 -->
      <!-- default-checked-keys	默认勾选的节点的 key 的数组 -->
      <el-tree
        ref="tree"
        :data="menuArr"
        show-checkbox
        node-key="id"
        default-expand-all
        :default-checked-keys="selectArr"
        :props="defaultProps"
      />
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="drawer = false">取消</el-button>
        <el-button type="primary" @click="handler">确定</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup lang="ts">
import {
  reqRemoveRole,
  reqAllRoleList,
  reqAddOrUpdateRole,
  reqAllMenuList,
  reqMenuByRoleID,
  reqUpdateRoleMenu,
} from "@/api/sys/role";
import type { Records, RoleData } from "@/api/sys/role/type";

import { ElMessage } from "element-plus";

import { ref, onMounted, reactive, nextTick } from "vue";
let pageNo = ref<number>(1);
let pageSize = ref<number>(10);
let total = ref<number>(0);
//收集关键字
let keyword = ref<string>("");

//控制对话框显示
let dialogVisible = ref<boolean>(false);
//收集删除的id
let ids = ref<number[]>([]);
//收集数据
let RoleParams = reactive<any>({
  name: "",
  remark: "",
});
//form组件实例
let form = ref<any>();
//权限抽屉
let drawer = ref<boolean>(false);
//菜单数据收集
let menuArr = ref<any>([]);
//准备数组 存储勾选的四级  选择的权限
let selectArr = ref<any[]>([]);
//复选框的数据
let selectIdArr = ref<any[]>([]);
//获取tree组件实例
let tree = ref<any>();
//仓库
import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();
//组件挂载完毕 获取列表
onMounted(() => {
  getHasRole();
});
//获取列表
let allRole = ref<Records>([]);
const getHasRole = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqAllRoleList(
    pageNo.value,
    pageSize.value,
    keyword.value
  );
  if (res.code === 200) {
    total.value = res.total;
    allRole.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//下拉改变
const sizeHandler = () => {
  getHasRole();
};
//搜索按钮
const search = () => {
  getHasRole();
  keyword.value = "";
};
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};

//添加按钮
const addRole = () => {
  dialogVisible.value = true;
  //清空
  Object.assign(RoleParams, {
    name: "",
    remark: "",
    id: 0,
  });
  nextTick(() => {
    form.value.clearValidate("name");
  });
};
//更新按钮
const updateRole = (row: RoleData) => {
  dialogVisible.value = true;
  //存储已有的角色
  Object.assign(RoleParams, row);
  nextTick(() => {
    form.value.clearValidate("name");
  });
};

//表单校验
const validateName = (rule: any, value: any, callBack: any) => {
  console.log(rule);
  if (value.trim().length >= 2) {
    callBack();
  } else {
    callBack(new Error("角色名称至少两位"));
  }
};
const rules = {
  name: [{ required: true, trigger: "blur", validator: validateName }],
};

//确定按钮
const save = async () => {
  await form.value.validate();
  let res: any = await reqAddOrUpdateRole(RoleParams);
  if (res.code === 200) {
    dialogVisible.value = false;
    getHasRole(RoleParams.id ? pageNo.value : 1); //更新留在当前 添加回到1
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//分配权限按钮
const setPermission = async (row: RoleData) => {
  menuArr.value = "";
  selectArr.value = [];
  drawer.value = true;
  Object.assign(RoleParams, row);
  let res: any = await reqAllMenuList();
  if (res.code == 200) {
    menuArr.value = res.data;
  }
  let res1: any = await reqMenuByRoleID({ id: row.id });
  if (res1.code === 200) {
    selectArr.value = filterMenuIds(res.data, res1.data);
  }
};

function filterMenuIds(menuArr: any, selectArr: any) {
  const filteredIds: any = [];
  const traverseMenu = (menu: any) => {
    for (const item of menu) {
      if (selectArr.includes(item.id) && item.children.length === 0) {
        filteredIds.push(item.id);
      }
      if (item.children && item.children.length > 0) {
        traverseMenu(item.children);
      }
    }
  };
  traverseMenu(menuArr);
  return filteredIds;
}

const defaultProps = {
  children: "children",
  label: "name",
};

//分配按钮
const handler = async () => {
  const roleId = RoleParams.id as number;
  //当前节点id
  let arr = tree.value.getCheckedKeys();
  //半选id 把父id给传进去
  let arr1 = tree.value.getHalfCheckedKeys();
  // 合并id
  let permissionId = arr.concat(arr1);
  let res: any = await reqUpdateRoleMenu({
    roleId: roleId,
    menuIds: permissionId,
  });
  if (res.code === 200) {
    drawer.value = false;
    window.location.reload();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//删除按钮
const removeRole = async (id: number) => {
  ids.value.push(id);
  const requestData: any = { ids: ids.value }; // 提取 ids 引用的值并构造请求数据对象
  let res: any = await reqRemoveRole(requestData);
  if (res.code === 200) {
    getHasRole(allRole.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//复选框选择
const selectChange = (value: any) => {
  selectIdArr.value = value;
};
//批量删除用户按钮
const deleteSelectRole = async () => {
  ids.value = selectIdArr.value.map((item) => {
    return item.id;
  });
  const requestData = { ids: ids.value };
  let res: any = await reqRemoveRole(requestData);
  if (res.code === 200) {
    ids.value = [];
    getHasRole(allRole.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
</script>
<style lang="scss" scoped>
.form {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
