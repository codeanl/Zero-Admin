<template>
  <!-- 上边搜索 -->
  <el-card>
    <el-form :inline="true">
      <el-form-item label="用户名:">
        <el-input
          placeholder="请输入搜索的用户名"
          v-model="username"
        ></el-input>
      </el-form-item>
      <el-form-item label="手机号:">
        <el-input placeholder="请输入搜索的手机号" v-model="phone"></el-input>
      </el-form-item>
      <el-form-item label="昵称:">
        <el-input placeholder="请输入搜索的昵称" v-model="nickname"></el-input>
      </el-form-item>
      <el-form-item label="状态:">
        <el-select v-model="status" class="m-2" placeholder="请选择状态">
          <el-option label="正常" value="1" />
          <el-option label="禁用" value="0" />
        </el-select>
      </el-form-item>
      <el-form-item label="性别:">
        <el-select v-model="gender" class="m-2" placeholder="请选择性别">
          <el-option label="保密" value="0" />
          <el-option label="男" value="1" />
          <el-option label="女" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          @click="search"
          :disabled="
            username.length ||
            phone.length ||
            status.length ||
            nickname.length ||
            gender.length
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
      :data="ListArr"
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
        label="手机号"
        align="center"
        prop="phone"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="用户昵称"
        align="center"
        prop="nickname"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="工作"
        align="center"
        prop="job"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="邮箱"
        align="center"
        prop="email"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="城市"
        align="center"
        prop="city"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="个性签名"
        align="center"
        prop="signature"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="性别"
        align="center"
        prop="gender"
        show-overflow-tooltip
        width="60px"
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
              --el-switch-off-color: #ff4949;
            "
            active-value="1"
            inactive-value="0"
            @change="handleChange(row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="创建时间"
        align="center"
        prop="creatTIme"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column label="操作" width="280px" align="center">
        <template #="{ row }">
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
          <el-button
            type="primary"
            size="small"
            icon="Edit"
            @click="loginLog(row.id)"
          >
            登录日志
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
  <!-- 抽屉  完成 添加｜修改 的窗口 -->
  <el-dialog v-model="drawer" :title="userParams.id ? '更新会员' : '添加会员'">
    <el-form :model="userParams" ref="formRef">
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
      <el-form-item label="城市" prop="city">
        <el-input
          placeholder="请您输入城市"
          v-model="userParams.city"
        ></el-input>
      </el-form-item>
      <el-form-item label="个性签名" prop="signature">
        <el-input
          placeholder="请您输入个性签名"
          v-model="userParams.signature"
        ></el-input>
      </el-form-item>
      <el-form-item label="工作" prop="job">
        <el-input
          placeholder="请您输入工作"
          v-model="userParams.job"
        ></el-input>
      </el-form-item>
      <el-form-item label="性别" prop="gender">
        <el-select
          v-model="userParams.gender"
          class="m-2"
          placeholder="请选择性别"
        >
          <el-option label="未知" value="0" />
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
      <el-form-item label="头像" prop="avatar">
        <el-upload
          class="avatar-uploader"
          action="/api/api/sys/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img
            v-if="userParams.avatar"
            :src="userParams.avatar"
            class="avatar"
          />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="save">确定</el-button>
      </div>
    </template>
  </el-dialog>
  <!-- 登录日志 -->
  <el-dialog v-model="drawer1" title="登录日志">
    <el-table border :data="loginLogListArr">
      <el-table-column
        label="id"
        align="center"
        prop="id"
        width="50px"
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
      <el-table-column label="操作" width="150px" align="center">
        <template #="{ row }">
          <el-popconfirm
            title="你确定删除吗"
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
      v-model:current-page="pageNo1"
      v-model:page-size="pageSize1"
      :page-sizes="[5, 10, 15, 20]"
      :background="true"
      layout="prev, pager, next, jumper, -> , sizes, total"
      :total="total1"
      @current-change="getLoginLog"
      @size-change="getLoginLog"
    />
  </el-dialog>
</template>

<script setup lang="ts">
import {
  reqMemberAll,
  reqAddOrUpdateMember,
  reqDeleteMember,
} from "@/api/ums/member";
import { ref, onMounted, reactive, nextTick } from "vue";
import {
  reqMemberLoginLogAll,
  reqDeleteMemberLoginLog,
} from "@/api/ums/memberLoginLog";
import { ElMessage } from "element-plus";
//默认页码
let pageNo = ref<number>(1);
//默认个数
let pageSize = ref<number>(10);
let total = ref<number>(0);
//数据列表
let ListArr = ref<any>([]);
//准备批量删除用户的id
let selectIdArr = ref<any>([]);
//收集删除的id
let ids = ref<number[]>([]);
//复选框选择
const selectChange = (value: any) => {
  selectIdArr.value = value;
};
//组件实例
let formRef = ref<any>();
//定义响应式数据 抽屉的显示隐藏
let drawer1 = ref<boolean>(false);
let drawer = ref<boolean>(false);
//收集用户信息
let userParams = reactive<any>({
  username: "",
  nickname: "",
  phone: "",
  email: "",
  status: "",
  gender: "",
  job: "",
  city: "",
  signature: "",
  avatar: "",
});
//收集用户查找的关键字
let username = ref<string>("");
let phone = ref<string>("");
let status = ref<string>("");
let nickname = ref<string>("");
let gender = ref<string>("");

//setting仓库
import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();
//组件挂载完毕
onMounted(() => {
  getHas();
});
//获取用户信息
const getHas = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqMemberAll(
    pageNo.value,
    pageSize.value,
    username.value,
    phone.value,
    status.value,
    nickname.value,
    gender.value
  );
  if (res.code == 200) {
    total.value = res.total;
    ListArr.value = res.data;
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
  username.value = "";
  phone.value = "";
  status.value = "";
};
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};
//取消按钮
const cancel = () => {
  drawer.value = false;
};
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
    status: "",
    gender: "",
    job: "",
    city: "",
    signature: "",
    avatar: "",
  });
  userParams.status = "1";
  //清除上一次的提示信息
  nextTick(() => {
    formRef.value.clearValidate("username");
    formRef.value.clearValidate("name");
    formRef.value.clearValidate("password");
  });
};
import type { UploadProps } from "element-plus";
//图片上传成功的钩子
const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  userParams.avatar = response.data;
  formRef.value.clearValidate("avatar");
};
//上传图片之前出发的钩子函数
const beforeAvatarUpload: UploadProps["beforeUpload"] = (rawFile) => {
  if (
    rawFile.type === "image/png" ||
    rawFile.type === "image/jpeg" ||
    rawFile.type === "image/gif"
  ) {
    if (rawFile.size / 1024 / 1024 < 4) {
      return true;
    } else {
      ElMessage({
        type: "error",
        message: "上传的文件大小应小于4M",
      });
    }
  } else {
    ElMessage({
      type: "error",
      message: "上传的文件类型必须是PNG|JPG|GIF",
    });
    return false;
  }
};
// 编辑按钮
const updateUser = (row: any) => {
  drawer.value = true;
  Object.assign(userParams, row);
  nextTick(() => {
    formRef.value.clearValidate("username");
    formRef.value.clearValidate("name");
  });
};
//窗口保存按钮
const save = async () => {
  await formRef.value.validate();
  let res: any = await reqAddOrUpdateMember(userParams);
  if (res.code == 200) {
    drawer.value = false;
    ElMessage({ type: "success", message: res.message });

    getHas(userParams.id ? pageNo.value : 1);
    // window.location.reload()
  } else {
    drawer.value = false;
    ElMessage({ type: "error", message: res.message });
  }
};
// 删除用户按钮
const deleteUser = async (userId: number) => {
  ids.value.push(userId);
  const requestData: any = { ids: ids.value }; // 提取 ids 引用的值并构造请求数据对象
  let res: any = await reqDeleteMember(requestData);
  if (res.code == 200) {
    getHas(ListArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//批量删除用户按钮
const deleteSelectUser = async () => {
  ids.value = selectIdArr.value.map((item: any) => {
    return item.id;
  });
  const requestData: any = { ids: ids.value };
  let res: any = await reqDeleteMember(requestData);
  if (res.code === 200) {
    getHas(ListArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//默认页码
let pageNo1 = ref<number>(1);
//默认个数
let pageSize1 = ref<number>(10);
let total1 = ref<number>(0);
let loginLogListArr = ref<any>([]);
let nowMember = ref<any>();
let loginLog = (id: number) => {
  nowMember.value = id;
  getLoginLog();
};
//登录日志
let getLoginLog = async () => {
  drawer1.value = true;
  let res = await reqMemberLoginLogAll(
    pageNo1.value,
    pageSize1.value,
    nowMember.value
  );
  if (res.code == 200) {
    total1.value = res.total;
    loginLogListArr.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//状态修改用户
let handleChange = async (row: any) => {
  let data: any = {
    id: row.id as number,
    status: row.status,
  };
  let res = await reqAddOrUpdateMember(data);
  if (res.code == 200) {
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
// 删除按钮
const deleteLog = async (id: number) => {
  ids.value = [];
  ids.value.push(id);
  const requestData: any = { ids: ids.value }; // 提取 ids 引用的值并构造请求数据对象
  let res: any = await reqDeleteMemberLoginLog(requestData);
  if (res.code == 200) {
    getLoginLog();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
</script>

<style scoped></style>
<style>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>
