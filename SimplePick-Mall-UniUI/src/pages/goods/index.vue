// src/pages/goods/goods.vue
<script setup lang="ts">
import type {
  SkuPopupEvent,
  SkuPopupInstance,
  SkuPopupLocaldata,
} from '@/components/vk-data-goods-sku-popup/vk-data-goods-sku-popup'
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()
import { onLoad } from '@dcloudio/uni-app'
import { computed, ref } from 'vue'
import { getProductInfoAPI } from '@/services/home'
import { addCart } from '@/services/cart'
import AddressPanel from './components/AddressPanel.vue'
import ServicePanel from './components/ServicePanel.vue'
// 接收页面参数
const query = defineProps<{
  id: number
}>()

// 获取商品详情信息
const goods = ref<any>()
const getGoodsByIdData = async () => {
  let id = Number(query.id)
  const res = await getProductInfoAPI({ id: id })
  goods.value = res.data
  // SKU组件所需格式
  localdata.value = {
    _id: res.data.productInfo.id,
    name: res.data.productInfo.name,
    goods_thumb: res.data.productInfo.pic,
    spec_list: res.data.sizeList.map((v) => {
      return {
        name: v.name,
        list: v.values,
      }
    }),
    sku_list: res.data.skuList.map((v) => {
      return {
        _id: v.id,
        goods_id: v.productId,
        goods_name: v.name,
        image: v.pic,
        price: v.price * 100, // 注意：需要乘以 100
        stock: v.stock,
        sku_name_arr: v.specs.map((vv) => vv.valueName),
      }
    }),
  }
}

// 页面加载
onLoad(() => {
  getGoodsByIdData()
})

// 轮播图变化时
const currentIndex = ref(0)
const onChange: UniHelper.SwiperOnChange = (ev) => {
  currentIndex.value = ev.detail.current
}

// 点击图片时
const onTapImage = (url: string) => {
  // 大图预览
  uni.previewImage({
    current: url,
    urls: goods.value!.imgUrl,
  })
}

// uni-ui 弹出层组件 ref
const popup = ref<{
  open: (type?: UniHelper.UniPopupType) => void
  close: () => void
}>()

// 弹出层条件渲染
const popupName = ref<'address' | 'service'>()
const openPopup = (name: typeof popupName.value) => {
  // 修改弹出层名称
  popupName.value = name
  popup.value?.open()
}
// 是否显示SKU组件
const isShowSku = ref(false)
// 商品信息
const localdata = ref({} as SkuPopupLocaldata)
// 按钮模式
enum SkuMode {
  Both = 1,
  Cart = 2,
  Buy = 3,
}
const mode = ref<SkuMode>(SkuMode.Cart)
// 打开SKU弹窗修改按钮模式
const openSkuPopup = (val: SkuMode) => {
  // 显示SKU弹窗
  isShowSku.value = true
  // 修改按钮模式
  mode.value = val
}
// SKU组件实例
const skuPopupRef = ref<SkuPopupInstance>()
// 计算被选中的值
const selectArrText = computed(() => {
  return skuPopupRef.value?.selectArr?.join(' ').trim() || '请选择商品规格'
})
// 加入购物车事件
const onAddCart = async (ev: SkuPopupEvent) => {
  let res = await addCart({ userID: 103, skuID: ev._id, count: ev.buy_num })
  if (res.code === 200) {
    uni.showToast({ title: '添加成功' })
    isShowSku.value = false
  } else {
    uni.showToast({ title: '添加失败' })
  }
}
// 立即购买
const onBuyNow = (ev: SkuPopupEvent) => {
  // uni.navigateTo({ url: `/pagesOrder/create/create?skuId=${ev._id}&count=${ev.buy_num}` })
  let sku: any = []
  sku.push({
    id: ev._id,
    count: ev.buy_num,
    skuID: ev._id,
    name: goods.value.productInfo.name,
    pic: goods.value.productInfo.pic,
    price: ev.price / 100,
    tag: ev.sku_name_arr,
    productId: goods.value.productInfo.id,
  })
  // 跳转到结算页
  uni.navigateTo({
    url: '/pages/orderCreate/index?sku=' + encodeURIComponent(JSON.stringify(sku)),
  })
}
</script>

<template>
  <!-- SKU弹窗组件 -->
  <vk-data-goods-sku-popup
    v-model="isShowSku"
    :localdata="localdata"
    :mode="mode"
    add-cart-background-color="#FFA868"
    buy-now-background-color="#27BA9B"
    ref="skuPopupRef"
    :actived-style="{
      color: '#27BA9B',
      borderColor: '#27BA9B',
      backgroundColor: '#E9F8F5',
    }"
    @add-cart="onAddCart"
    @buy-now="onBuyNow"
  />
  <scroll-view scroll-y class="viewport">
    <!-- 基本信息 -->
    <view class="goods">
      <!-- 商品主图 -->
      <view class="preview">
        <swiper @change="onChange" circular>
          <swiper-item v-for="item in goods?.imgUrl" :key="item">
            <image class="image" @tap="onTapImage(item)" mode="aspectFill" :src="item" />
          </swiper-item>
        </swiper>
        <view class="indicator">
          <text class="current">{{ currentIndex + 1 }}</text>
          <text class="split">/</text>
          <text class="total">{{ goods?.imgUrl.length }}</text>
        </view>
      </view>

      <!-- 商品简介 -->
      <view class="meta">
        <view class="price">
          <text class="symbol">¥</text>
          <text class="number">{{ goods?.productInfo.price }}</text>
          <text class="xiaoliang">已售 {{ goods?.productInfo.sale }}</text>
        </view>
        <view class="name ellipsis">{{ goods?.productInfo.name }}</view>
        <view class="desc"> {{ goods?.productInfo.desc }} </view>
      </view>

      <!-- 操作面板 -->
      <view class="action">
        <!-- <view class="item arrow">
          <text class="label">选择</text>
          <text class="text ellipsis"> 请选择商品规格 </text>
        </view>
        <view class="item arrow">
          <text class="label">送至</text>
          <text class="text ellipsis"> 请选择收获地址 </text>
        </view> -->
        <view class="item arrow">
          <text class="label">服务</text>
          <text class="text ellipsis"> 无忧退 快速退款 免费包邮 </text>
        </view>
      </view>

      <!-- 商品详情 -->
      <view class="merchant">
        <image class="image" :src="goods?.merchantInfo.pic"></image>
        <text class="info">
          <text class="name">{{ goods?.merchantInfo.name }}</text>
        </text>
        <!-- <navigator class="button" url="/pages/login/index" open-type="navigate"> 进店 </navigator> -->
      </view>

      <!-- 商品详情 -->
      <view class="detail panel">
        <view class="title">
          <text>详情</text>
        </view>
        <view class="content">
          <view class="properties">
            <view class="properties">
              <!-- 属性详情 -->
              <view class="item" v-for="item in goods?.attributeList" :key="item.name">
                <text class="label">{{ item.name }}</text>
                <text class="value">{{ item.value }}</text>
              </view>
            </view>
          </view>
          <!-- 图片详情 -->
          <image
            class="image"
            v-for="item in goods?.introduceImgUrl"
            :key="item"
            mode="widthFix"
            :src="item"
          ></image>
        </view>
      </view>
    </view>
  </scroll-view>

  <!-- 用户操作 -->
  <view v-if="goods" class="toolbar" :style="{ paddingBottom: safeAreaInsets?.bottom + 'px' }">
    <view class="icons">
      <button class="icons-button"><text class="icon-heart"></text>收藏</button>
      <!-- #ifdef MP-WEIXIN -->
      <button class="icons-button" open-type="contact">
        <text class="icon-handset"></text>客服
      </button>
      <!-- #endif -->
      <navigator class="icons-button" url="/pages/cart/cart2" open-type="navigate">
        <text class="icon-cart"></text>购物车
      </navigator>
    </view>
    <view class="buttons">
      <view @tap="openSkuPopup(SkuMode.Cart)" class="addcart"> 加入购物车 </view>
      <view @tap="openSkuPopup(SkuMode.Buy)" class="payment"> 立即购买 </view>
    </view>
  </view>

  <!-- uni-ui 弹出层 -->
  <uni-popup ref="popup" type="bottom" background-color="#fff">
    <AddressPanel v-if="popupName === 'address'" @close="popup?.close()" />
    <ServicePanel v-if="popupName === 'service'" @close="popup?.close()" />
  </uni-popup>
</template>

<style lang="scss">
page {
  height: 100%;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.viewport {
  background-color: #f4f4f4;
}

.panel {
  margin-top: 20rpx;
  background-color: #fff;

  .title {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 90rpx;
    line-height: 1;
    padding: 30rpx 60rpx 30rpx 6rpx;
    position: relative;

    text {
      padding-left: 10rpx;
      font-size: 28rpx;
      color: #333;
      font-weight: 600;
      border-left: 4rpx solid #27ba9b;
    }

    navigator {
      font-size: 24rpx;
      color: #666;
    }
  }
}

.arrow {
  &::after {
    position: absolute;
    top: 50%;
    right: 30rpx;
    content: '\e6c2';
    color: #ccc;
    font-family: 'erabbit' !important;
    font-size: 32rpx;
    transform: translateY(-50%);
  }
}

/* 商品信息 */
.goods {
  background-color: #fff;

  .preview {
    height: 750rpx;
    position: relative;

    .image {
      width: 750rpx;
      height: 750rpx;
    }

    .indicator {
      height: 40rpx;
      padding: 0 24rpx;
      line-height: 40rpx;
      border-radius: 30rpx;
      color: #fff;
      font-family: Arial, Helvetica, sans-serif;
      background-color: rgba(0, 0, 0, 0.3);
      position: absolute;
      bottom: 30rpx;
      right: 30rpx;

      .current {
        font-size: 26rpx;
      }

      .split {
        font-size: 24rpx;
        margin: 0 1rpx 0 2rpx;
      }

      .total {
        font-size: 24rpx;
      }
    }
  }

  .meta {
    position: relative;
    border-bottom: 1rpx solid #eaeaea;

    .price {
      height: 130rpx;
      padding: 25rpx 30rpx 0;
      color: #fff;
      font-size: 34rpx;
      box-sizing: border-box;
      //background-color: #35c8a9;
      background-color: #e49364;
    }

    .number {
      font-size: 56rpx;
    }

    .xiaoliang {
      float: right;
      font-size: 26rpx;
      margin-right: 20rpx;
      margin-top: 20rpx;
      color: #444;
    }

    .brand {
      width: 160rpx;
      height: 80rpx;
      overflow: hidden;
      position: absolute;
      top: 26rpx;
      right: 30rpx;
    }

    .name {
      max-height: 88rpx;
      line-height: 1.4;
      margin: 20rpx;
      font-size: 32rpx;
      color: #333;
    }

    .desc {
      line-height: 1;
      padding: 0 20rpx 30rpx;
      font-size: 24rpx;
      color: #cf4444;
    }
  }

  .action {
    padding-left: 20rpx;

    .item {
      height: 90rpx;
      padding-right: 60rpx;
      border-bottom: 1rpx solid #eaeaea;
      font-size: 26rpx;
      color: #333;
      position: relative;
      display: flex;
      align-items: center;

      &:last-child {
        border-bottom: 0 none;
      }
    }

    .label {
      width: 60rpx;
      color: #898b94;
      margin: 0 16rpx 0 10rpx;
    }

    .text {
      flex: 1;
      -webkit-line-clamp: 1;
    }
  }
}

/* 商家详情 */
.merchant {
  margin: 20rpx 0;
  background-color: #f4f1f1;
  padding: 20rpx 20rpx 20rpx 20px;
  display: flex;
  align-items: center;

  .image {
    width: 120rpx;
    height: 120rpx;
    border-radius: 50%;
  }

  .info {
    margin-left: 30rpx;

    .name {
      height: 72rpx;
      font-size: 32rpx;
      color: #444;
    }
  }

  .button {
    margin-left: 28rpx;
    font-size: 30rpx;
    color: #444;
    border-radius: 72rpx;
    width: 140rpx;
    border: 1rpx solid #ccc;
    text-align: center;
  }
}

/* 商品详情 */
.detail {
  padding-left: 20rpx;

  .content {
    margin-left: -20rpx;

    .image {
      width: 100%;
    }
  }

  .properties {
    padding: 0 20rpx;
    margin-bottom: 30rpx;

    .item {
      display: flex;
      line-height: 2;
      padding: 10rpx;
      font-size: 26rpx;
      color: #333;
      border-bottom: 1rpx dashed #ccc;
    }

    .label {
      width: 200rpx;
    }

    .value {
      flex: 1;
    }
  }
}

/* 同类推荐 */
.similar {
  .content {
    padding: 0 20rpx 20rpx;
    background-color: #f4f4f4;
    display: flex;
    flex-wrap: wrap;

    .goods {
      width: 340rpx;
      padding: 24rpx 20rpx 20rpx;
      margin: 20rpx 7rpx;
      border-radius: 10rpx;
      background-color: #fff;
    }

    .image {
      width: 300rpx;
      height: 260rpx;
    }

    .name {
      height: 80rpx;
      margin: 10rpx 0;
      font-size: 26rpx;
      color: #262626;
    }

    .price {
      line-height: 1;
      font-size: 20rpx;
      color: #cf4444;
    }

    .number {
      font-size: 26rpx;
      margin-left: 2rpx;
    }
  }

  navigator {
    &:nth-child(even) {
      margin-right: 0;
    }
  }
}

/* 底部工具栏 */
.toolbar {
  position: fixed;
  left: 0;
  right: 0;
  bottom: calc((var(--window-bottom)));
  z-index: 1;
  background-color: #fff;
  height: 100rpx;
  padding: 0 20rpx;
  border-top: 1rpx solid #eaeaea;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-sizing: content-box;

  .buttons {
    display: flex;

    & > view {
      width: 220rpx;
      text-align: center;
      line-height: 72rpx;
      font-size: 26rpx;
      color: #fff;
      border-radius: 72rpx;
    }

    .addcart {
      background-color: #ffa868;
    }

    .payment {
      background-color: #27ba9b;
      margin-left: 20rpx;
    }
  }

  .icons {
    padding-right: 20rpx;
    display: flex;
    align-items: center;
    flex: 1;

    // 兼容 H5 端和 App 端的导航链接样式
    .navigator-wrap,
    .icons-button {
      flex: 1;
      text-align: center;
      line-height: 1.4;
      padding: 0;
      margin: 0;
      border-radius: 0;
      font-size: 20rpx;
      color: #333;
      background-color: #fff;

      &::after {
        border: none;
      }
    }

    text {
      display: block;
      font-size: 34rpx;
    }
  }
}
</style>
