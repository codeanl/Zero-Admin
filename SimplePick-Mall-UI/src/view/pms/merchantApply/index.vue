<template>
  <div v-if="isSuccess == false">
    <h1 style="text-align: center; margin-top: 40px; font-size: 24px">
      商家｜自提点入驻
    </h1>
    <el-form :model="form" label-width="120px" class="box">
      <el-form-item label="类型">
        <el-radio-group v-model="form.type">
          <el-radio label="商户" value="0" />
          <el-radio label="自提点" value="1" />
        </el-radio-group>
      </el-form-item>
      <el-form-item label="商户/自提点名称">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="详细地址">
        <el-input v-model="form.address" />
      </el-form-item>
      <el-form-item label="负责人姓名">
        <el-input v-model="form.principalName" />
      </el-form-item>
      <el-form-item label="负责人联系方式">
        <el-input v-model="form.principalPhone" />
      </el-form-item>
      <el-form-item label="门店图片" prop="pic">
        <el-upload
          class="avatar-uploader"
          action="/api/api/sys/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="form.pic" :src="form.pic" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="身份证正面照" prop="IDCardFront">
        <el-upload
          class="avatar-uploader"
          action="/api/api/sys/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess1"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="form.IDCardFront" :src="form.IDCardFront" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="身份证反面照" prop="IDCardReverse">
        <el-upload
          class="avatar-uploader"
          action="/api/api/sys/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess2"
          :before-upload="beforeAvatarUpload"
        >
          <img
            v-if="form.IDCardReverse"
            :src="form.IDCardReverse"
            class="avatar"
          />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
      </el-form-item>
      <div style="float: right">
        <el-button type="primary" @click="save">提交申请</el-button>
      </div>
    </el-form>
  </div>
  <div v-else class="box">
    <el-col>
      <el-result
        icon="success"
        title="申请成功"
        sub-title="您已申请成功，等待管理员审核！审核结果将通过手机短信发送至您的手机！"
      >
        <template #extra>
          <el-button type="primary">查看进度</el-button>
        </template>
      </el-result>
    </el-col>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { ElMessage } from "element-plus";
import { reqAddOrUpdateMerchantApply } from "@/api/pms/merchantApply";
const form = reactive({
  principalName: "",
  principalPhone: "",
  IDCardFront: "",
  IDCardReverse: "",
  name: "",
  address: "",
  type: "商户",
  pic: "",
});
let isSuccess = ref<any>(false);
import type { UploadProps } from "element-plus";
//图片上传成功的钩子
const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  form.pic = response.data;
};
const handleAvatarSuccess1: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  form.IDCardFront = response.data;
};
const handleAvatarSuccess2: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  form.IDCardReverse = response.data;
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

let save = async () => {
  if (form.type == "商户") {
    form.type = "0";
  }
  if (form.type == "自提点") {
    form.type = "1";
  }
  let res = await reqAddOrUpdateMerchantApply(form);
  if (res.code == 200) {
    isSuccess.value = true;
  }
};
</script>
<style lang="scss">
.box {
  width: 400px;
  margin: 30px auto;
}
</style>
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
