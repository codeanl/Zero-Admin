<template>
  <!-- 数据 -->
  <el-card>
    <el-button type="success" size="default" icon="Plus" @click="add">
      添加
    </el-button>
    <!--  -->
    <el-table style="margin: 15px 0" row-key="id" border :data="ListArr">
      <el-table-column prop="name" label="名称" />
      <!-- <el-table-column prop="type" label="类型" /> -->
      <el-table-column
        label="图标"
        align="center"
        prop="icon"
        show-overflow-tooltip
        width="120px"
      >
        <template #="{ row }">
          <img :src="row.icon" alt="图标" style="width: 60px; height: auto" />
        </template>
      </el-table-column>
      <el-table-column
        label="类型"
        align="center"
        prop="level"
        show-overflow-tooltip
      >
        <template #="{ row }">
          <template v-if="row.level === '1'">
            <el-tag class="mx-1" type="success" effect="light">
              一级目录
            </el-tag>
          </template>
          <template v-if="row.level === '2'">
            <el-tag class="mx-1" type="warning" effect="light">
              二级目录
            </el-tag>
          </template>
        </template>
      </el-table-column>
      <!-- <el-table-column prop="productCount" label="产品数量" /> -->
      <el-table-column prop="description" label="描述" />
      <!-- <el-table-column prop="count" label="商品数量" /> -->
      <el-table-column prop="sort" label="排序">
        <template #="{ row }"> 当前排{{ row.sort }}位 </template>
      </el-table-column>
      <el-table-column label="操作" width="260px">
        <template #="{ row }">
          <el-button type="primary" size="small" @click="update(row)">
            编辑
          </el-button>
          <el-popconfirm
            :title="`你确定要删除${row.name}?`"
            width="260px"
            @confirm="remove(row.id)"
          >
            <template #reference>
              <el-button type="danger" size="small"> 删除 </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  <!-- 添加｜更新 -->
  <el-dialog v-model="dialogVisible" :title="Data.id ? '更新' : '添加'">
    <el-form ref="formRef">
      <el-form-item label="名称">
        <el-input
          placeholder="请你输入菜单的名称"
          v-model="Data.name"
        ></el-input>
      </el-form-item>
      <el-form-item label="类型">
        <el-select v-model="Data.level" class="m-2" placeholder="请选择类型">
          <el-option label="一级目录" value="1" />
          <el-option label="二级目录" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="父级id">
        <el-tree-select
          v-model="Data.parentId"
          :data="ListArrWithRoot"
          check-strictly
          :props="{ key: 'id', label: 'name' }"
          node-key="id"
          :render-after-expand="false"
        />
      </el-form-item>
      <el-form-item label="排序">
        <!-- <el-input placeholder="请你输入排序" v-model="Data.sort"></el-input> -->
        排在第<el-input-number
          v-model="Data.sort"
          :min="1"
          size="small"
          controls-position="right"
        />位
      </el-form-item>
      <el-form-item label="描述">
        <el-input
          placeholder="请你输入描述"
          v-model="Data.description"
        ></el-input>
      </el-form-item>
      <el-form-item label="图标" prop="pic">
        <el-upload
          class="avatar-uploader"
          action="/api/api/sys/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="Data.icon" :src="Data.icon" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
        <!-- <el-input placeholder="请您输入门店图片" v-model="Params.pic"></el-input> -->
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
import { reqAll, reqAddOrUpdate, reqRemove } from "@/api/pms/category";
//组件实例
let formRef = ref<any>();
let ListArr = ref<any>([]);
let dialogVisible = ref<boolean>(false);
let ids = ref<number[]>([]);

onMounted(() => {
  getHas();
});
const getHas = async () => {
  let res: any = await reqAll();
  if (res.code === 200) {
    ListArr.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
const ListArrWithRoot = computed(() => {
  return [{ id: 0, name: "无上级" }, ...ListArr.value];
});

let Data = reactive<any>({
  description: "",
  icon: "",
  keywords: "",
  level: "",
  name: "",
  navStatus: "",
  parentId: 0,
  productCount: 0,
  productUnit: "",
  showStatus: "",
  sort: 0,
});
//添加
const add = () => {
  Object.assign(Data, {
    id: 0,
    description: "",
    icon: "",
    keywords: "",
    level: "",
    name: "",
    navStatus: "",
    parentId: 0,
    productCount: 0,
    productUnit: "",
    showStatus: "",
    sort: 0,
  });
  dialogVisible.value = true;
};

const update = (row: any) => {
  dialogVisible.value = true;
  Data.parentId = parseInt(row.parentId);
  Object.assign(Data, row);
};

const save = async () => {
  dialogVisible.value = false;
  let res: any = await reqAddOrUpdate(Data);
  if (res.code === 200) {
    dialogVisible.value = false;
    getHas();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
const remove = async (id: number) => {
  ids.value.push(id);
  const requestData: any = { ids: ids.value };
  let res = await reqRemove(requestData);
  if (res.code == 200) {
    ids.value = [];
    getHas();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

import type { UploadProps } from "element-plus";
//图片上传成功的钩子
const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  Data.icon = response.data;
  formRef.value.clearValidate("icon");
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
</script>

<style lang="scss" scoped></style>
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
