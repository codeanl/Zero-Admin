<template>
  <!-- 上边搜索 -->
  <el-card>
    <el-form :inline="true">
      <el-form-item label="名称:">
        <el-input placeholder="请输入搜索的用户名" v-model="name"></el-input>
      </el-form-item>
      <el-form-item label="地址:">
        <el-input placeholder="请输入搜索的用户名" v-model="place"></el-input>
      </el-form-item>
      <el-form-item label="手机号:">
        <el-input placeholder="请输入搜索的用户名" v-model="phone"></el-input>
      </el-form-item>
      <el-form-item label="负责人:">
        <el-input
          placeholder="请输入搜索的负责人"
          v-model="principal"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          @click="search"
          :disabled="
            name.length || place.length || phone.length || principal.length
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
    <!-- <el-button type="success" size="default" @click="add"> 添加 </el-button> -->
    <el-button
      type="danger"
      size="default"
      @click="deleteSelect"
      :disabled="selectIdArr.length ? false : true"
    >
      批量删除
    </el-button>
    <el-badge
      v-if="total1 != 0"
      :value="total1"
      :max="10"
      class="item"
      style="margin-left: 10px"
    >
      <el-button type="primary" size="default" icon="Edit" @click="look"
        >查看审核</el-button
      >
    </el-badge>
    <el-button v-else type="primary" size="default" icon="Edit" @click="look"
      >查看审核</el-button
    >
    <el-table
      border
      :data="listArr"
      @selection-change="selectChange"
      style="margin-top: 15px"
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
        label="门店图片"
        align="center"
        prop="pic"
        show-overflow-tooltip
        width="120px"
      >
        <template #="{ row }">
          <img :src="row.pic" alt="头像" style="width: 100px; height: auto" />
        </template>
      </el-table-column>
      <el-table-column
        label="名称"
        align="center"
        prop="name"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="联系方式"
        align="center"
        prop="phone"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="详细地址"
        align="center"
        prop="place"
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
              --el-switch-off-color: #ff4949;
            "
            active-value="1"
            inactive-value="0"
            @change="handleChange(row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="负责人"
        align="center"
        prop="principal"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column label="操作" width="300px" align="center">
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
      @current-change="getHasPlace"
      @size-change="handler"
    />
  </el-card>
  <!-- 抽屉  完成 添加｜修改 的窗口 -->
  <el-dialog v-model="drawer" :title="Params.id ? '更新' : '添加'">
    <el-form :model="Params" ref="formRef">
      <el-form-item label="名称" prop="name">
        <el-input placeholder="请您输入名称" v-model="Params.name"></el-input>
      </el-form-item>
      <el-form-item label="联系方式" prop="place">
        <el-input
          placeholder="请您输入联系方式"
          v-model="Params.place"
        ></el-input>
      </el-form-item>
      <el-form-item label="详细地址" prop="phone">
        <el-input
          placeholder="请您输入详细地址"
          v-model="Params.phone"
        ></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="Params.status" class="m-2" placeholder="请选择状态">
          <el-option label="暂停营业" value="0" />
          <el-option label="正常营业" value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="负责人" prop="principal">
        <el-input
          placeholder="请您输入负责人"
          v-model="Params.principal"
        ></el-input>
      </el-form-item>
      <el-form-item label="门店图片" prop="pic">
        <el-upload
          class="avatar-uploader"
          action="/api/api/sys/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="Params.pic" :src="Params.pic" class="avatar" />
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
  <!--  -->
  <el-dialog v-model="drawer1" width="70%" :before-close="beforeCloseHandler">
    <el-select
      v-model="applyStatus"
      class="m-2"
      placeholder="请选择状态"
      @change="applyStatusChange"
    >
      <el-option label="待审核" value="0" />
      <el-option label="已同意" value="1" />
      <el-option label="已拒绝" value="2" />
    </el-select>
    <!-- 数据 -->
    <el-table
      border
      :data="listArr1"
      @selection-change="selectChange"
      style="margin: 15px 0"
    >
      <el-table-column
        label="id"
        align="center"
        prop="id"
        width="50px"
      ></el-table-column>
      <el-table-column
        label="店铺图片"
        align="center"
        prop="pic"
        show-overflow-tooltip
        width="120px"
      >
        <template #="{ row }">
          <img
            :src="row.pic"
            alt="店铺图片"
            style="width: 100px; height: auto"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="名称"
        align="center"
        prop="name"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="负责人"
        align="center"
        prop="principalName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="联系电话"
        align="center"
        prop="principalPhone"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="地址"
        align="center"
        prop="address"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column label="操作" width="300px" align="center">
        <template #="{ row }">
          <el-button
            type="primary"
            size="small"
            icon="Edit"
            @click="lookInfo(row)"
          >
            查看详情
          </el-button>
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
      @current-change="getHas1"
      @size-change="handler1"
    />
  </el-dialog>
  <!--  -->
  <el-dialog v-model="drawer2" width="70%" title="店铺申请详情">
    <el-descriptions direction="horizontal" :column="1" :size="size" border>
      <el-descriptions-item label="图片">
        <img :src="merchantsApply.pic" alt="" style="width: 200px" />
      </el-descriptions-item>
      <el-descriptions-item label="收货人">{{
        merchantsApply.principalName
      }}</el-descriptions-item>
      <el-descriptions-item label="联系电话">{{
        merchantsApply.principalPhone
      }}</el-descriptions-item>
      <el-descriptions-item label="地址">{{
        merchantsApply.address
      }}</el-descriptions-item>
      <el-descriptions-item label="身份证正反面">
        <img :src="merchantsApply?.IDCardFront" alt="" style="width: 200px" />
        <img :src="merchantsApply?.IDCardReverse" alt="" style="width: 200px" />
      </el-descriptions-item>
      <el-descriptions-item label="备注">
        <el-input v-model="input" placeholder="请填写备注" clearable />
      </el-descriptions-item>
    </el-descriptions>
    <template #footer>
      <div style="flex: auto" v-if="isShow">
        <el-button @click="notOk">拒绝</el-button>
        <el-button type="primary" @click="ok">同意</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, nextTick } from "vue";
import { reqPlaceList, reqRemovePlace, reqAddOrUpdate } from "@/api/pms/place";
import {
  reqAllMerchantApply,
  reqAddOrUpdateMerchantApply,
} from "@/api/pms/merchantApply";
import { ElMessage } from "element-plus";

//setting仓库
import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();

//默认页码
let pageNo = ref<number>(1);
//默认个数
let pageSize = ref<number>(10);
let total = ref<number>(0);
let pageNo1 = ref<number>(1);
let pageSize1 = ref<number>(5);
let total1 = ref<number>(0);
//数据列表
let listArr = ref<any>([]);
//收集用户查找的关键字
let name = ref<string>("");
let place = ref<string>("");
let phone = ref<string>("");
let principal = ref<string>("");
//收集删除的id
let ids = ref<number[]>([]);
//多选框选择的id
let selectIdArr = ref<any[]>([]);
//复选框选择
const selectChange = (value: any) => {
  selectIdArr.value = value;
};
const size = ref("");
//组件实例
let formRef = ref<any>();
//定义响应式数据 抽屉的显示隐藏
let drawer = ref<boolean>(false);
let drawer1 = ref<boolean>(false);
let drawer2 = ref<boolean>(false);
let Params = reactive<any>({
  name: "",
  place: "",
  status: "",
  pic: "",
  phone: "",
  principal: "",
});
//组件挂载完毕
onMounted(() => {
  getHasPlace();
  getHas1();
});
//获取信息
const getHasPlace = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqPlaceList(
    pageNo.value,
    pageSize.value,
    name.value,
    place.value,
    phone.value,
    principal.value
  );
  if (res.code == 200) {
    total.value = res.total;
    listArr.value = res.data;
  }
};
//下拉改变
const handler = () => {
  getHasPlace();
};
const handler1 = () => {
  getHas1();
};
//搜索按钮
const search = () => {
  getHasPlace();
  name.value = "";
  place.value = "";
  phone.value = "";
};
// 删除按钮
const deletePlace = async (id: number) => {
  ids.value.push(id);
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemovePlace(requestData);
  if (res.code == 200) {
    getHasPlace(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
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
  let res: any = await reqRemovePlace(requestData);
  if (res.code === 200) {
    getHasPlace(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};
//取消按钮
const cancel = () => {
  drawer.value = false;
};
// 编辑按钮
const update = (row: any) => {
  drawer.value = true;
  Object.assign(Params, row);
  nextTick(() => {
    formRef.value.clearValidate("name");
    formRef.value.clearValidate("place");
    formRef.value.clearValidate("status");
    formRef.value.clearValidate("pic");
    formRef.value.clearValidate("phone");
    formRef.value.clearValidate("principal");
  });
};
//添加用户按钮
// const add = () => {
//   drawer.value = true;
//   //存储收集已有的账号信息
//   Object.assign(Params, {
//     id: 0,
//     name: "",
//     place: "",
//     status: "1",
//     pic: "",
//     phone: "",
//     principal: "",
//   });
//   nextTick(() => {
//     formRef.value.clearValidate("name");
//     formRef.value.clearValidate("place");
//     formRef.value.clearValidate("status");
//     formRef.value.clearValidate("pic");
//     formRef.value.clearValidate("phone");
//     formRef.value.clearValidate("principal");
//   });
// };

//窗口保存按钮
const save = async () => {
  let res: any = await reqAddOrUpdate(Params);
  if (res.code == 200) {
    drawer.value = false;
    ElMessage({ type: "success", message: res.message });

    // window.location.reload()
    getHasPlace();
  } else {
    drawer.value = false;
    ElMessage({ type: "error", message: res.message });
  }
};

import type { UploadProps } from "element-plus";
//图片上传成功的钩子
const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  Params.pic = response.data;
  formRef.value.clearValidate("pic");
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

//状态修改用户
let handleChange = async (row: any) => {
  let data: any = {
    id: row.id as number,
    status: row.status,
  };
  let res = await reqAddOrUpdate(data);
  if (res.code === 200) {
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

let applyStatusChange = () => {
  console.log(applyStatus.value);
  getHas1();
};
let listArr1 = ref<any>([]);
let applyStatus = ref<string>("0");
let name1 = ref<string>("");

const getHas1 = async (pager = 1) => {
  pageNo1.value = pager;
  let res: any = await reqAllMerchantApply(
    pageNo1.value,
    pageSize1.value,
    name1.value,
    applyStatus.value,
    "1"
  );
  if (res.code == 200) {
    total1.value = res.total;
    listArr1.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//查看审核
let look = () => {
  drawer1.value = true;
};
const input = ref("");
let merchantsApply = ref<any>();
let isShow = ref<any>(false);
let lookInfo = (row: any) => {
  if (row.status == "0") {
    isShow.value = true;
  }
  input.value = "";
  drawer2.value = true;
  merchantsApply.value = row;
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

import useUserStore from "@/store/sys/user";
let userStore = useUserStore();
let notOk = async () => {
  let res = await reqAddOrUpdateMerchantApply({
    id: merchantsApply.value.id,
    status: "2",
    remarks: input.value,
    approver: userStore.username,
    approvalTime: currentDateTime,
  });
  if (res.code === 200) {
    drawer2.value = false;
    getHas1();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
import { reqAddOrUpdateUser, Rbac } from "@/api/sys/user";
let ok = async () => {
  drawer2.value = false;
  let res = await reqAddOrUpdateMerchantApply({
    id: merchantsApply.value.id,
    status: "1",
    remarks: input.value,
    approver: userStore.username,
    approvalTime: currentDateTime,
  });
  if (res.code == 200) {
    let res2 = await reqAddOrUpdateUser({
      username: merchantsApply.value.principalPhone,
      phone: merchantsApply.value.principalPhone,
      nickname: merchantsApply.value.principalName,
      status: "1",
      gender: "0",
    });
    if (res2.code == 200) {
      drawer2.value = false;
      getHas1();
      getHasPlace();
      await Rbac({
        id: res2.data,
        roleId: [3],
      });
      await reqAddOrUpdate({
        name: merchantsApply.value.name,
        principal: merchantsApply.value.principalName,
        phone: merchantsApply.value.principalPhone,
        place: merchantsApply.value.address,
        pic: merchantsApply.value.pic,
        userID: res2.data,
        status: "1",
      });
      getHas1();
      getHasPlace();
    }
  }
};

//恢复到原来
let beforeCloseHandler = (done: any) => {
  applyStatus.value = "0";
  getHas1();
  done(); // 允许对话框关闭
};
</script>

<style scoped lang="scss"></style>
<style>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
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
