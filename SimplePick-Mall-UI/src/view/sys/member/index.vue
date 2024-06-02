<template>
  <!-- 搜索 -->
  <el-card>
    <el-form :inline="true">
      <el-form-item label="用户名:">
        <el-input
          placeholder="请输入搜索的用户名"
          v-model="username"
        ></el-input>
      </el-form-item>
      <el-form-item label="昵称:">
        <el-input placeholder="请输入搜索的昵称" v-model="nickname"></el-input>
      </el-form-item>
      <el-form-item label="手机号:">
        <el-input placeholder="请输入搜索的手机号" v-model="phone"></el-input>
      </el-form-item>
      <el-form-item label="邮箱:">
        <el-input placeholder="请输入搜索的邮箱" v-model="email"></el-input>
      </el-form-item>
      <el-form-item label="性别:">
        <el-select v-model="gender" class="m-2" placeholder="请选择性别">
          <el-option label="保密" value="0" />
          <el-option label="男" value="1" />
          <el-option label="女" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态:">
        <el-select v-model="status" class="m-2" placeholder="请选择状态">
          <el-option label="正常" value="1" />
          <el-option label="禁用" value="0" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          @click="search"
          :disabled="
            email.length ||
            username.length ||
            nickname.length ||
            phone.length ||
            gender.length ||
            status.length
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
  <!-- 数据列表 -->
  <el-card>
    <el-button type="success" size="default" @click="addUser"> 添加 </el-button>
    <el-button
      type="danger"
      size="default"
      @click="deleteSelectUser"
      :disabled="selectIdArr.length ? false : true"
    >
      批量删除
    </el-button>
    <el-table
      border
      :data="userArr"
      @selection-change="selectChange"
      style="margin: 15px 0"
    >
      <!--复选框 -->
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
        label="用户昵称"
        align="center"
        prop="nickname"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="性别"
        align="center"
        prop="gender"
        show-overflow-tooltip
        width="70px"
      >
        <template #="{ row }">
          <template v-if="row.gender === '0'">
            <el-tag type="warning">保密</el-tag></template
          >
          <template v-else-if="row.gender === '1'">
            <el-tag type="primary">男</el-tag>
          </template>
          <template v-else> <el-tag type="success">女 </el-tag></template>
        </template>
      </el-table-column>
      <el-table-column
        label="邮箱"
        align="center"
        prop="email"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="手机号"
        align="center"
        prop="phone"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="用户角色"
        align="center"
        prop="roleName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="创建时间"
        align="center"
        prop="creat_at"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="更新时间"
        align="center"
        prop="update_at"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="状态"
        align="center"
        prop="status"
        show-overflow-tooltip
        width="60px"
      >
        <template #="{ row }">
          <el-switch
            v-model="row.status"
            class="ml-2"
            style="
              --el-switch-on-color: #13ce66;
              --el-switch-off-color: #f06455;
            "
            active-value="1"
            inactive-value="0"
            @change="handleChange(row)"
          />
        </template>
      </el-table-column>
      <el-table-column label="操作" width="400px" align="center">
        <template #="{ row }">
          <el-button
            type="success"
            size="small"
            icon="User"
            @click="setRole(row)"
          >
            角色
          </el-button>
          <el-popconfirm
            :title="`你确定要重置${row.username}的密码吗？`"
            width="260px"
            @confirm="ResetPassword(row.id)"
          >
            <template #reference>
              <el-button type="warning" size="small" icon="Star">
                重置
              </el-button>
            </template>
          </el-popconfirm>
          <el-button
            type="primary"
            size="small"
            icon="Edit"
            @click="updateUser(row)"
          >
            编辑
          </el-button>
          <el-popconfirm
            :title="`你确定删除${row.username}`"
            width="260px"
            @confirm="deleteUser(row.id)"
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
  <!-- 添加｜修改  -->
  <el-dialog v-model="drawer" :title="userParams.id ? '更新' : '添加'">
    <el-form :model="userParams" :rules="rules" ref="formRef">
      <el-form-item label="用户名" prop="username">
        <el-input
          placeholder="请您输入用户名"
          v-model="userParams.username"
        ></el-input>
      </el-form-item>
      <el-form-item label="昵称" prop="nickname">
        <el-input
          placeholder="请您输入昵称"
          v-model="userParams.nickname"
        ></el-input>
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input
          placeholder="请您输入手机号"
          v-model="userParams.phone"
        ></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input
          placeholder="请您输入邮箱"
          v-model="userParams.email"
        ></el-input>
      </el-form-item>
      <el-form-item label="性别" prop="gender">
        <el-select
          v-model="userParams.gender"
          class="m-2"
          placeholder="请选择性别"
        >
          <el-option label="保密" value="0" />
          <el-option label="男" value="1" />
          <el-option label="女" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select
          v-model="userParams.status"
          class="m-2"
          placeholder="请选择状态"
        >
          <el-option label="正常" value="1" />
          <el-option label="禁用" value="0" />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="save">确定</el-button>
      </div>
    </template>
  </el-dialog>
  <!--  职位分配 -->
  <el-drawer v-model="drawer1">
    <template #header>
      <h4>分配角色</h4>
    </template>
    <template #default>
      <el-form>
        <el-form-item label="用户姓名">
          <el-input v-model="userParams.username" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="职位列表">
          <el-checkbox
            v-model="checkAll"
            :indeterminate="isIndeterminate"
            @change="handleCheckAllChange"
          >
            全选
          </el-checkbox>
          <el-checkbox-group
            v-model="userRole"
            @change="handleCheckedUsersChange"
          >
            <el-checkbox
              v-for="(role, index) in allRole"
              :key="index"
              :label="role"
            >
              {{ role.name }}
            </el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
    </template>
    <!-- 按钮 -->
    <template #footer>
      <div style="flex: auto">
        <el-button @click="drawer1 = false">取消</el-button>
        <el-button type="primary" @click="confirmClick">确定</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup lang="ts">
import {
  reqUserInfo,
  reqAddOrUpdateUser,
  reqAllRole,
  reqRemoveUser,
  reqMenuListByUser,
  resetPassword,
  Rbac,
} from "@/api/sys/user";
import type { Records, User } from "@/api/sys/user/type";
import { ref, onMounted, reactive, nextTick } from "vue";
import { ElMessage } from "element-plus";
//默认个数 默认页码
let pageNo = ref<number>(1);
let pageSize = ref<number>(10);
let total = ref<number>(0);

//收集用户查找的关键字
let username = ref<string>("");
let nickname = ref<string>("");
let phone = ref<string>("");
let email = ref<string>("");
let status = ref<string>("");
let gender = ref<string>("");

//收集用户信息
let userParams = reactive<any>({
  username: "",
  nickname: "",
  phone: "",
  email: "",
  status: "",
  gender: "",
});

//组件实例
let formRef = ref<any>();
//全选框
const checkAll = ref<boolean>(false);
//不确定样式
const isIndeterminate = ref<boolean>(true);
//所有role
let allRole = ref<any>([]);
// 用户role
let userRole = ref<any>([]);
//准备批量删除用户的id
let selectIdArr = ref<User[]>([]);
//收集删除的id
let ids = ref<any[]>([]);
//setting仓库
import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();

//组件挂载完毕
onMounted(() => {
  getHasUser();
});

//获取用户信息
let userArr = ref<Records>([]);
const getHasUser = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqUserInfo(
    pageNo.value,
    pageSize.value,
    username.value,
    phone.value,
    nickname.value,
    status.value,
    gender.value,
    email.value
  );
  if (res.code == 200) {
    total.value = res.total;
    userArr.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//下拉改变
const handler = () => {
  getHasUser();
};

//定义响应式数据 抽屉的显示隐藏
let drawer = ref<boolean>(false);
let drawer1 = ref<boolean>(false);

//添加用户按钮
const addUser = () => {
  drawer.value = true;
  //存储收集已有的账号信息
  Object.assign(userParams, {
    id: 0,
    username: "",
    nickname: "",
    phone: "",
    email: "",
    status: "1",
    gender: "",
  });
  //清除上一次的提示信息
  nextTick(() => {
    formRef.value.clearValidate("username");
    formRef.value.clearValidate("nicknem");
    formRef.value.clearValidate("phone");
    formRef.value.clearValidate("status");
    formRef.value.clearValidate("gender");
  });
};

// 编辑按钮
const updateUser = (row: User) => {
  drawer.value = true;
  Object.assign(userParams, row);
  nextTick(() => {
    formRef.value.clearValidate("username");
    formRef.value.clearValidate("nicknem");
    formRef.value.clearValidate("phone");
    formRef.value.clearValidate("status");
    formRef.value.clearValidate("gender");
  });
};

////表单校验 自定义
//const validatorUserName = (rule: any, value: any, callBack: any) => {
//  if (value.trim().length >= 5 || value.trim().length <= 20) {
//    callBack();
//  } else {
//    callBack(new Error("用户名需在5-20位之间"));
//  }
//};
//const validatorNickname = (rule: any, value: any, callBack: any) => {
//  console.log(rule);
//  if (value.trim().length >= 5 || value.trim().length <= 20) {
//   callBack();
// } else {
//   callBack(new Error("用户昵称需在5-20位之间"));
//  }
//};
//const validatorPhone = (rule: any, value: any, callBack: any) => {
//console.log(rule);
//const phoneRegex = /^1[0-9]{10}$/; // 手机号码正则表达式
//if (phoneRegex.test(value.trim())) {
//  callBack();
// } else {
//   callBack(new Error("请输入有效的手机号码"));
// }
//};
const rules = {
  username: [{ required: true, trigger: "blur" }],
  nickname: [{ required: true, trigger: "blur" }],
  phone: [{ required: true, trigger: "blur" }],
  gender: [{ required: true, trigger: "blur" }],
  status: [{ required: true, trigger: "blur" }],
};

//窗口保存按钮
const save = async () => {
  await formRef.value.validate();
  let res: any = await reqAddOrUpdateUser(userParams);
  if (res.code == 200) {
    drawer.value = false;
    ElMessage({ type: "success", message: res.message });
    //更新留在当前页 添加回到第一页
    // getHasUser(userParams.id ? pageNo.value : 1)
    window.location.reload();
  } else {
    drawer.value = false;
    ElMessage({ type: "error", message: res.message });
  }
};
//取消按钮
const cancel = () => {
  drawer.value = false;
};

//分配角色按钮
const setRole = async (row: User) => {
  drawer1.value = true;
  Object.assign(userParams, row);
  let res: any = await reqAllRole();
  if (res.code === 200) {
    allRole.value = res.data;
  }
  let res1: any = await reqMenuListByUser({ userID: userParams.id });
  if (res1.code === 200) {
    userRole.value = res1.data;
  }
};

//点击全选框
const handleCheckAllChange = (val: boolean) => {
  userRole.value = val ? allRole.value : []; //全选就把所有的role给它 否者为空
  isIndeterminate.value = false;
};
//底部复选框change事件
const handleCheckedUsersChange = (value: string[]) => {
  // 已经勾选的长度
  const checkedCount = value.length;
  checkAll.value = checkedCount === allRole.value.length; //boolean 不确定样式
  isIndeterminate.value =
    checkedCount > 0 && checkedCount < allRole.value.length;
};

//分配角色 确定按钮
const confirmClick = async () => {
  let data: any = {
    id: userParams.id as number,
    roleId: userRole.value.map((item: any) => {
      return item.id as number;
    }),
  };
  let res: any = await Rbac(data);
  if (res.code === 200) {
    drawer1.value = false;
    getHasUser(pageNo.value);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
// 删除用户按钮
const deleteUser = async (userId: number) => {
  ids.value.push(userId);
  const requestData: any = { ids: ids.value }; // 提取 ids 引用的值并构造请求数据对象
  let res: any = await reqRemoveUser(requestData);
  if (res.code == 200) {
    getHasUser(userArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
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
const deleteSelectUser = async () => {
  ids.value = selectIdArr.value.map((item) => {
    return item.id;
  });
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemoveUser(requestData);
  if (res.code === 200) {
    getHasUser(userArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//搜索按钮
const search = () => {
  getHasUser();
  username.value = "";
  nickname.value = "";
  email.value = "";
  status.value = "";
  gender.value = "";
};
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};
//状态修改用户
let handleChange = async (row: any) => {
  let data: any = {
    id: row.id as number,
    status: row.status,
  };
  let res = await reqAddOrUpdateUser(data);
  if (res.code === 200) {
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

let ResetPassword = async (id: number) => {
  let res = await resetPassword({ id: id });
  if (res.code === 200) {
    if (res.code === 200) {
      ElMessage({ type: "success", message: res.message });
    } else {
      ElMessage({ type: "error", message: res.message });
    }
  }
};
</script>

<style scoped lang="scss"></style>
