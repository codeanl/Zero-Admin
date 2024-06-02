<template>
  <!-- 上边搜索 -->
  <el-card>
    <el-form :inline="true" class="form">
      <el-form-item label="名称:">
        <el-input placeholder="请输入搜索的用户名" v-model="name"></el-input>
      </el-form-item>
      <el-form-item label="分类:">
        <el-tree-select
          v-model="CateID"
          :data="CateListArr"
          check-strictly
          :props="{ key: 'categoryId', label: 'name' }"
          node-key="id"
          :render-after-expand="false"
        />
      </el-form-item>
      <el-form-item label="价格区间:">
        <el-input
          style="width: 150px"
          v-model="minPrice"
          placeholder="请输入最小值"
        />--
        <el-input
          style="width: 150px"
          v-model="maxPrice"
          placeholder="请输入最大值"
        />
      </el-form-item>
      <el-form-item label="排序方式:">
        <el-select
          v-model="searchType"
          class="m-2"
          placeholder="请选择排序方式"
        >
          <el-option label="发布时间" value="1" />
          <el-option label="销量" value="2" />
          <el-option label="点击数" value="3" />
          <el-option label="价格低->高" value="4" />
          <el-option label="价格高->低" value="5" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="default"
          @click="search"
          :disabled="
            name || CateID || searchType || minPrice || maxPrice ? false : true
          "
        >
          搜索
        </el-button>
        <el-button size="default" @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
  <!-- 数据展示 -->
  <el-card>
    <!-- 上边按钮 -->
    <el-button type="success" size="default" @click="add"> 添加 </el-button>
    <el-button
      type="danger"
      size="default"
      @click="deleteSelect"
      :disabled="selectIdArr.length ? false : true"
    >
      批量删除
    </el-button>
    <!-- 数据 -->
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
        label="商品封面"
        align="center"
        prop="pic"
        show-overflow-tooltip
        width="120px"
      >
        <template #="{ row }">
          <img
            :src="row.pic"
            alt="商品封面"
            style="width: 100px; height: auto"
          />
        </template>
      </el-table-column>
      <el-table-column
        label="商品名称"
        align="center"
        prop="name"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="货号"
        align="center"
        width="150px"
        prop="productSn"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="商品描述"
        align="center"
        prop="desc"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="价格"
        align="center"
        width="100px"
        prop="price"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="市场价"
        align="center"
        width="100px"
        prop="originalPrice"
        show-overflow-tooltip
      ></el-table-column>
      <!-- <el-table-column
        label="销量"
        align="center"
        prop="sale"
        show-overflow-tooltip
      ></el-table-column> -->
      <!-- <el-table-column
        label="单位"
        align="center"
        prop="unit"
        show-overflow-tooltip
      ></el-table-column> -->
      <el-table-column label="操作" width="250px" align="center">
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
            @confirm="deleteproduct(row.id)"
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
            @click="look(row.id)"
          >
            查看
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
  <!-- 更新｜添加 -->
  <el-dialog v-model="drawer" :title="Params.id ? '更新' : '添加'">
    <el-form :model="Params" ref="formRef">
      <el-form-item label="名称" prop="name">
        <el-input placeholder="请您输入名称" v-model="Params.name"></el-input>
      </el-form-item>
      <el-form-item label="所属分类">
        <el-tree-select
          v-model="Params.categoryId"
          :data="CateListArr"
          check-strictly
          :props="{ key: 'categoryId', label: 'name' }"
          node-key="id"
          :render-after-expand="false"
        />
      </el-form-item>
      <el-form-item label="所属属性分类">
        <el-tree-select
          v-model="Params.attributeCategoryID"
          :data="attrCateListArr"
          check-strictly
          :props="{ key: 'attributeCategoryID', label: 'name' }"
          node-key="id"
          :render-after-expand="false"
          @change="changeAttrCate"
        />
      </el-form-item>
      <el-form-item label="货号" prop="productSn">
        <el-input
          placeholder="请您输入货号"
          v-model="Params.productSn"
        ></el-input>
      </el-form-item>
      <el-form-item label="商品描述" prop="desc">
        <el-input placeholder="商品描述" v-model="Params.desc"></el-input>
      </el-form-item>
      <el-form-item label="市场价" prop="originalPrice">
        <el-input
          placeholder="市场价"
          v-model="Params.originalPrice"
        ></el-input>
      </el-form-item>
      <el-form-item label="价格" prop="price">
        <el-input placeholder="价格" v-model="Params.price"></el-input>
      </el-form-item>
      <!-- <el-form-item label="单位" prop="unit">
        <el-input placeholder="单位" v-model="Params.unit"></el-input>
      </el-form-item> -->
      <el-form-item label="商品封面" prop="pic">
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
    <!-- 照片墙 -->
    <el-form :model="Params" :inline="true">
      <h2>照片墙</h2>
      <el-upload
        v-model:file-list="fileList"
        action="/api/api/sys/upload"
        list-type="picture-card"
        :on-preview="handlePictureCardPreview"
        :on-remove="handleRemove"
        :on-success="handleAvatarSuccess1"
      >
        <el-icon>
          <Plus />
        </el-icon>
      </el-upload>
      <el-dialog v-model="dialogVisible">
        <img w-full :src="dialogImageUrl" alt="Preview Image" />
      </el-dialog>
    </el-form>

    <!-- 介绍照片墙 -->
    <el-form :model="Params" :inline="true">
      <h2>介绍照片墙</h2>
      <el-upload
        v-model:file-list="introduceImgUrlList"
        action="/api/api/sys/upload"
        list-type="picture-card"
        :on-preview="handlePictureCardPreview1"
        :on-remove="handleRemove1"
        :on-success="handleAvatarSuccess2"
      >
        <el-icon>
          <Plus />
        </el-icon>
      </el-upload>
      <el-dialog v-model="dialogVisible1">
        <img w-full :src="dialogImageUrl1" alt="Preview Image" />
      </el-dialog>
    </el-form>
    <!--  -->
    <!-- ---------------------------------------- -->
    <el-form :model="Params" :inline="true">
      <h2>属性</h2>
      <el-form-item
        prop="attrList"
        v-for="(attribute, index) in Params.attributeType1"
        :key="index"
        :label="attribute.name"
      >
        <el-select v-model="selectedValue[index]" placeholder="请选择">
          <el-option
            v-for="(valueObj, valueIndex) in convertToArray(attribute.value)"
            :key="valueIndex"
            :label="valueObj"
            :value="valueObj"
          ></el-option>
        </el-select>
      </el-form-item>
    </el-form>
    <!--  -->
    <el-form :model="Params" :inline="true">
      <h2>规格（修改时如果不修改sku无需填写）</h2>
      <el-checkbox-group
        v-model="selectedValue1[index]"
        v-for="(attribute, index) in Params.attributeType2"
        :label="attribute.name"
      >
        <el-checkbox
          v-for="(valueObj, valueIndex) in convertToArray(attribute.value)"
          :label="valueObj"
          :key="valueIndex"
          :value="valueObj"
        />
      </el-checkbox-group>
    </el-form>
    <!-- ---------------------------------------- -->
    <!-- 按钮 -->
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="save">确定</el-button>
      </div>
    </template>
  </el-dialog>
  <!-- sku -->
  <el-dialog v-model="drawer1" title="SKU列表" width="70%">
    <el-table :data="skuArr" v-model="SkuParams">
      <el-table-column label="SKU图片">
        <template #="{ row }">
          <img :src="row.pic" alt="" style="width: 100px; height: 100px" />
        </template>
      </el-table-column>
      <el-table-column
        label="sku编号"
        width="120px"
        prop="skuSn"
      ></el-table-column>
      <el-table-column label="SKU名字" prop="name"></el-table-column>
      <el-table-column
        label="SKU价格"
        width="80px"
        prop="price"
      ></el-table-column>
      <el-table-column label="库存" width="80px" prop="stock"></el-table-column>
      <el-table-column label="销量" width="80px" prop="sale"></el-table-column>
      <el-table-column label="TAG" prop="tag"></el-table-column>
      <el-table-column label="描述" prop="description"></el-table-column>
      <el-table-column label="操作" width="100px" align="center">
        <template #="{ row }">
          <el-button
            type="primary"
            size="small"
            icon="Edit"
            @click="updatesku(row)"
          >
            编辑
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>
  <!--  -->
  <el-dialog v-model="drawer2" title="SKU" width="70%">
    <el-form :model="SkuParams" ref="formRef">
      <el-form-item label="名称" prop="name">
        <el-input
          placeholder="请您输入名称"
          v-model="SkuParams.name"
        ></el-input>
      </el-form-item>
      <el-form-item label="sku编号" prop="skuSn">
        <el-input placeholder="请您输入" v-model="SkuParams.skuSn"></el-input>
      </el-form-item>
      <el-form-item label="商品描述" prop="description">
        <el-input
          placeholder="商品描述"
          v-model="SkuParams.description"
        ></el-input>
      </el-form-item>
      <el-form-item label="库存" prop="stock">
        <el-input
          placeholder="库存"
          v-model.number="SkuParams.stock"
        ></el-input>
      </el-form-item>
      <el-form-item label="价格" prop="price">
        <el-input placeholder="价格" v-model="SkuParams.price"></el-input>
      </el-form-item>
      <el-form-item
        v-for="(attr, index) in SkuParams.sizeList"
        :label="attr.name"
        :prop="`sizeList${index}`"
      >
        <el-select
          v-model="SkuParams.AttributeShopValueID[index]"
          placeholder="请选择属性"
        >
          <el-option
            v-for="(option, optionIndex) in attr.sizeValueList"
            :key="optionIndex"
            :label="option.name"
            :value="option.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="商品封面" prop="pic">
        <el-upload
          class="avatar-uploader"
          action="/api/api/sys/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="SkuParams.pic" :src="SkuParams.pic" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancelsku">取消</el-button>
        <el-button type="primary" @click="savesku">确定</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from "vue";
import { ElMessage, UploadUserFile } from "element-plus";
import {
  reqAllProduct,
  reqAddOrUpdate,
  reqRemove,
  reqSku,
  reqAddOrUpdateSKU,
  reqProductInfo,
} from "@/api/pms/product";
import { reqAll } from "@/api/pms/category";
import { reqAllAttribute } from "@/api/pms/attribute";
import { reqAllattributeCategory } from "@/api/pms/attributeCategory";
//数据
let name = ref<string>("");
let CateID = ref<number>();
let searchType = ref<number>();
let minPrice = ref<number>();
let maxPrice = ref<number>();

//分页数据
let pageNo = ref<number>(1);
let pageSize = ref<number>(10);
let total = ref<number>(0);
//下拉改变
const handler = () => {
  getHas();
};

const Params = reactive<any>({
  id: 0,
  categoryId: 0,
  name: "",
  pic: "",
  desc: "",
  price: "",
  originalPrice: 0,
  unit: "",
  productSn: "",
  subTitle: "",
  attributeValue: [],
  size: [{ name: "", sizeValue: [] }],
  imgUrl: [],
  introduceImgUrl: [],
  attributeType1: [],
  attributeType2: [],
  attributeValueList: [],
  attributeCategoryID: 0,
});
let SkuParams = reactive<any>({
  productId: 0,
  name: "",
  pic: "",
  skuSn: "",
  subTitle: "",
  description: "",
  stock: 0,
  price: "",
  AttributeShopValueID: [],
  sizeList: [],
});
//组件挂载完毕
onMounted(() => {
  getHas();
});
//数据列表
let listArr = ref<any>([]);
//分类
let CateListArr = ref<any>([]);
let attrCateListArr = ref<any>([]);
//获取信息
const getHas = async (pager = 1) => {
  pageNo.value = pager;
  let res: any = await reqAllProduct(
    pageNo.value,
    pageSize.value,
    name.value,
    CateID.value as number,
    searchType.value as number,
    minPrice.value as number,
    maxPrice.value as number
  );
  if (res.code == 200) {
    total.value = res.total;
    listArr.value = res.data;
  }
  let res1: any = await reqAll();
  if (res1.code === 200) {
    CateListArr.value = res1.data;
  }
};
let fetchAttributeList = async (id: number) => {
  let res = await reqAllAttribute(1, 100, "", "1", id);
  let res1 = await reqAllAttribute(1, 100, "", "2", id);
  if (res.code == 200) {
    Params.attributeType1 = res.data;
  }
  if (res1.code == 200) {
    Params.attributeType2 = res1.data;
  }
  getAllattributeCategory();
};
let getAllattributeCategory = async () => {
  let res2: any = await reqAllattributeCategory();
  if (res2.code === 200) {
    attrCateListArr.value = res2.data;
  }
};

//多选框选择的id
let selectIdArr = ref<any[]>([]);
//复选框选择
const selectChange = (value: any) => {
  selectIdArr.value = value;
};
//收集删除的id
let ids = ref<number[]>([]);
//批量删除用户按钮
const deleteSelect = async () => {
  ids.value = selectIdArr.value.map((item) => {
    return item.id;
  });
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemove(requestData);
  if (res.code === 200) {
    getHas(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};
// 删除按钮
const deleteproduct = async (id: number) => {
  ids.value.push(id);
  const requestData: any = { ids: ids.value };
  let res: any = await reqRemove(requestData);
  if (res.code == 200) {
    getHas(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1);
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

//定义响应式数据 抽屉的显示隐藏
let drawer = ref<boolean>(false);
let drawer1 = ref<boolean>(false);
let drawer2 = ref<boolean>(false);
//取消按钮
const cancel = () => {
  drawer.value = false;
};

const fileList = ref<UploadUserFile[]>([]);
const introduceImgUrlList = ref<UploadUserFile[]>([]);
//添加用户按钮
const add = () => {
  drawer.value = true;
  //存储收集已有的账号信息
  fileList.value = [];
  introduceImgUrlList.value = [];
  // 清空Params对象
  Object.keys(Params).forEach((key) => {
    Params[key] = Array.isArray(Params[key]) ? [] : "";
  });
  selectedValue.value = [];
  selectedValue1.value = [];
  getAllattributeCategory();
};

// 编辑按钮
const update = async (row: any) => {
  console.log(row.id);
  drawer.value = true;
  // 清空Params对象
  Object.keys(Params).forEach((key) => {
    Params[key] = Array.isArray(Params[key]) ? [] : "";
  });
  selectedValue.value = [];
  selectedValue1.value = [];
  fileList.value = [];
  introduceImgUrlList.value = [];
  let res = await reqProductInfo({ id: row.id });
  if (res.code == 200) {
    // 初始化 Params.size
    Params.size = [];
    Object.assign(Params, {
      id: res.data.product_info?.id,
      categoryId: res.data.product_info?.categoryId,
      name: res.data.product_info?.name,
      pic: res.data.product_info?.pic,
      desc: res.data.product_info?.desc,
      price: res.data.product_info?.price,
      originalPrice: res.data.product_info?.originalPrice,
      unit: res.data.product_info?.unit,
      productSn: res.data.product_info?.productSn,
      attributeValue: res.data.attributeValue,
      imgUrl: res.data.imgUrl,
      introduceImgUrl: res.data.introduceImgUrl,
      attributeCategoryID: res.data.product_info.attributeCategoryID,
    });
    fetchAttributeList(res.data.product_info?.attributeCategoryID);
    fileList.value = Params.imgUrl.map((url: string) => ({
      name: url.substring(url.lastIndexOf("/") + 1),
      url: url,
    }));
    introduceImgUrlList.value = Params.introduceImgUrl.map((url: string) => ({
      name: url.substring(url.lastIndexOf("/") + 1),
      url: url,
    }));
    if (res.data.attributeValue.attributeType1.length > 0) {
      selectedValue.value = res.data.attributeValue.attributeType1.map(
        (item: { values: { value: any }[] }) => item.values[0].value
      );
    }
    // if (res.data.attributeValue.attributeType2.length > 0) {
    //     for (let i = 0; i < res.data.attributeValue.attributeType2.length; i++) {
    //         let valuess = []
    //         let values = res.data.attributeValue.attributeType2[i].values;
    //         for (let j = 0; j < values.length; j++) {
    //             let value = values[j].value;
    //             valuess.push(value);
    //         }
    //         console.log(valuess);
    //         selectedValue1.value[i] = valuess
    //     }
    // }
    // selectedValue1.value[0] = ['大芒果9斤装单', '大芒果5斤装单'];
  }
};
//type1
let selectedValue = ref<any>();
//type2
let selectedValue1 = ref<any>();
let convertToArray = (value: any) => {
  return value.split(","); // 将字符串转换为数组
};
const save = async () => {
  selectedValue.value = selectedValue.value.map(
    (Value: any, index: string | number) => {
      const attribute = Params.attributeType1[index];
      return {
        attributeID: attribute.id,
        value: [Value],
      };
    }
  );
  selectedValue1.value = selectedValue1.value.map(
    (Value: any, index: string | number) => {
      const attribute1 = Params.attributeType2[index];
      return {
        attributeID: attribute1.id,
        value: Value,
      };
    }
  );
  if (selectedValue1.value[0]) {
    for (let i = 0; i < selectedValue1.value.length; i++) {
      Params.attributeValueList.push(selectedValue1.value[i]);
    }
  }
  for (let i = 0; i < selectedValue.value.length; i++) {
    Params.attributeValueList.push(selectedValue.value[i]);
  }
  Params.originalPrice = parseFloat(Params.originalPrice);
  Params.price = parseFloat(Params.price);
  let res: any = await reqAddOrUpdate(Params);
  if (res.code == 200) {
    drawer.value = false;
    // window.location.reload()
    getHas();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

const dialogImageUrl = ref("");
const dialogImageUrl1 = ref("");
const dialogVisible = ref(false);
const dialogVisible1 = ref(false);
const handlePictureCardPreview: UploadProps["onPreview"] = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url!;
  dialogVisible.value = true;
};
const handlePictureCardPreview1: UploadProps["onPreview"] = (uploadFile) => {
  dialogImageUrl1.value = uploadFile.url!;
  dialogVisible1.value = true;
};
const handleRemove: UploadProps["onRemove"] = (uploadFile, uploadFiles) => {
  console.log(uploadFiles);
  Params.imgUrl = Params.imgUrl.filter(
    (item: string | undefined) => item !== uploadFile.url
  );
};
const handleRemove1: UploadProps["onRemove"] = (uploadFile, uploadFiles) => {
  console.log(uploadFiles);
  Params.introduceImgUrl = Params.introduceImgUrl.filter(
    (item: string | undefined) => item !== uploadFile.url
  );
};
//图片上传
import type { UploadProps } from "element-plus";
//组件实例
let formRef = ref<any>();
//图片上传成功的钩子
const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  Params.pic = response.data;
  SkuParams.pic = response.data;
  formRef.value.clearValidate("pic");
  SkuParams.value.clearValidate("pic");
};
const handleAvatarSuccess1: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  Params.imgUrl.push(response.data);
  Params.imgUrl.value.clearValidate("imgUrl");
};
const handleAvatarSuccess2: UploadProps["onSuccess"] = (response) => {
  //上传返回的数据 图片url  uploadFile
  // Params.introduceImgUrl.push(response.data)
  if (Params.introduceImgUrl === null) {
    Params.introduceImgUrl = []; // 创建一个新的空数组
  }
  Params.introduceImgUrl.push(response.data);
  Params.introduceImgUrl.value.clearValidate("introduceImgUrl");
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

let skuArr = ref<any>([]);
let look = async (id: number) => {
  drawer1.value = true;
  let res: any = await reqSku({ productID: id });
  if (res.code == 200) {
    skuArr.value = res.data;
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

// 更新sku按钮
const updatesku = (row: any) => {
  drawer2.value = true;
  drawer1.value = false;
  Object.assign(SkuParams, row);
};
//sku取消按钮
const cancelsku = () => {
  drawer2.value = false;
  drawer1.value = true;
};
//sku确定按钮
const savesku = async () => {
  SkuParams.stock = parseInt(SkuParams.stock);
  SkuParams.price = parseFloat(SkuParams.price);
  let res: any = await reqAddOrUpdateSKU(SkuParams);
  if (res.code == 200) {
    drawer2.value = false;
    look(SkuParams.productId);
    // window.location.reload()
    getHas();
    ElMessage({ type: "success", message: res.message });
  } else {
    ElMessage({ type: "error", message: res.message });
  }
};

let changeAttrCate = () => {
  Params.attributeType1 = [];
  Params.attributeType2 = [];
  selectedValue.value = [];
  fetchAttributeList(Params.attributeCategoryID);
};

import useLayoutSettingStore from "@/store/setting";
let settingStore = useLayoutSettingStore();
//重置按钮
const reset = () => {
  settingStore.refresh = !settingStore.refresh;
};
//搜索按钮
const search = () => {
  getHas();
  name.value = "";
  CateID.value = 0;
};
</script>

<style scoped>
h2 {
  font-size: 16px;
  padding: 20px;
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
