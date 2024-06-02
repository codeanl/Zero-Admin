<script setup lang="ts">
import { useGuessList } from '@/composables'
import { ref } from 'vue'
import { onLoad } from '@dcloudio/uni-app'
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()
// 猜你喜欢
const { guessRef, onScrolltolower } = useGuessList()
// 弹出层组件
const popup = ref<UniHelper.UniPopupInstance>()
const popup1 = ref<UniHelper.UniPopupInstance>()

// 订单取消原因
let reason = ref('')
// 复制内容
const onCopy = (id: string) => {
  // 设置系统剪贴板的内容
  uni.setClipboardData({ data: id })
}

import { orderInfo, updateOrder } from '@/services/order'
import { listreturnReason, addreturnApply, returnApplyInfo } from '@/services/return'

onLoad(() => {
  getData()
})
// 获取页面参数
const query = defineProps<{
  id: number
}>()
const order = ref<any>()
const reasonList = ref<any>()
const returnApply = ref<any>()
let idd = parseInt(query.id)
const getData = async () => {
  const res = await orderInfo({ id: idd })
  if (res.code === 200) {
    order.value = res.data
  }
  const res1 = await listreturnReason({})
  if (res1.code === 200) {
    reasonList.value = res1.data
  }
  const res2 = await returnApplyInfo({ orderID: parseInt(query.id) })
  if (res2.code === 200) {
    returnApply.value = res2.data
  }
}

let getCurrentDateTime = () => {
  const now = new Date()
  const year = now.getFullYear()
  const month = String(now.getMonth() + 1).padStart(2, '0') // 月份从 0 开始，需要加 1，并补零
  const day = String(now.getDate()).padStart(2, '0') // 补零
  const hours = String(now.getHours()).padStart(2, '0') // 补零
  const minutes = String(now.getMinutes()).padStart(2, '0') // 补零
  const seconds = String(now.getSeconds()).padStart(2, '0') // 补零
  const dateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
  return dateTime
}
// 订单支付
const onOrderPay = async () => {
  const currentDateTime = getCurrentDateTime()
  let res = await updateOrder({
    id: idd,
    status: '1',
    paymentTime: currentDateTime,
  })
  if (res.code == 200) {
    // 关闭当前页，再跳转支付结果页
    uni.redirectTo({ url: `/pages/payment/index?id=${query.id}` })
  }
}

const onCancel = async () => {
  let res = await updateOrder({
    id: idd,
    status: '5',
  })
  if (res.code == 200) {
    uni.redirectTo({ url: `/pages/detail/index?id=${query.id}` })
  }
}

const onConfirm = async () => {
  const currentDateTime = getCurrentDateTime()
  // 二次确认弹窗
  uni.showModal({
    content: '为保障您的权益，请收到货并确认无误后，再确认收货',
    confirmColor: '#27BA9B',
    success: async (success) => {
      if (success.confirm) {
        let res = await updateOrder({
          id: idd,
          status: '4',
          receiveTime: currentDateTime,
        })
        if (res.code == 200) {
          uni.redirectTo({ url: `/pages/detail/index?id=${query.id}` })
        }
      }
    },
  })
}

const onComment = async () => {
  const currentDateTime = getCurrentDateTime()
  let res = await updateOrder({
    id: idd,
    status: '4',
    commentTime: currentDateTime,
  })
  if (res.code == 200) {
    uni.redirectTo({ url: `/pages/detail/index?id=${query.id}` })
    uni.showToast({ icon: 'success', title: '评论完成' })
  }
}

import { useMemberStore } from '@/stores'
const memberStore = useMemberStore()

let onOrderCancel = async () => {
  console.log(returnReason.value.id)
  // await addreturnApply({
  //   userID: memberStore.profile.info.id,
  //   orderID: 1,
  //   returnReasonID: 1,
  //   status: 1,
  //   description: 1,
  //   proofPics: 1,
  //   returnAmount: 1,
  // })
}

let returnReason = ref<any>()
let ddd = (item: any) => {
  reason.value = item
  returnReason.value = item
}
</script>

<template>
  <!-- 自定义导航栏: 默认透明不可见, scroll-view 滚动到 50 时展示 -->
  <view class="navbar" :style="{ paddingTop: safeAreaInsets?.top + 'px' }">
    <view class="wrap">
      <navigator v-if="true" open-type="navigateBack" class="back icon-left"></navigator>
      <navigator v-else url="/pages/index/index" open-type="switchTab" class="back icon-home">
      </navigator>
      <view class="title">订单详情</view>
    </view>
  </view>
  <scroll-view scroll-y class="viewport" id="scroller" @scrolltolower="onScrolltolower">
    <template v-if="true">
      <!-- 订单状态 -->
      <view class="overview" :style="{ paddingTop: safeAreaInsets!.top + 20 + 'px' }">
        <!-- 待付款状态:展示去支付按钮和倒计时 -->
        <template v-if="order.orderInfo.status == '0'">
          <view class="status icon-clock">等待付款</view>
          <view class="tips">
            <text class="money">应付金额: {{ order.orderInfo.payAmount }} 元</text>
          </view>
        </template>
        <view class="status" v-if="order.orderInfo.status == '1'"> 待发货 </view>
        <view class="status" v-if="order.orderInfo.status == '2'"> 运输中 </view>
        <view class="status" v-if="order.orderInfo.status == '3'"> 待提货 </view>
        <view class="status" v-if="order.orderInfo.status == '4'"> 订单完成 </view>
        <view class="status" v-if="order.orderInfo.status == '5'"> 订单无效 </view>
      </view>

      <!-- 商品信息 -->
      <view class="goods">
        <view class="item">
          <navigator class="dianpu" url="/pages/login/index">
            <text class="label">店铺</text>
            <text class="name">{{ order?.merchantInfo.name }} ></text>
          </navigator>
          <navigator
            class="navigator"
            v-for="item in order?.skuList"
            :key="item.id"
            :url="`/pages/goods/index?id=${item.productId}`"
            hover-class="none"
          >
            <image class="cover" :src="item.pic"></image>
            <view class="meta">
              <view class="name ellipsis">{{ item.name }}</view>
              <view class="type">{{ item.tag }}</view>
              <view class="price">
                <view class="actual">
                  <text class="symbol">¥</text>
                  <text>{{ item.price }}</text>
                </view>
              </view>
              <view class="quantity">x{{ item.count }}</view>
            </view>
          </navigator>
        </view>
        <!-- 合计 -->
        <view class="total">
          <view class="row">
            <view class="text">商品总价: </view>
            <view class="symbol">{{ order?.orderInfo.payAmount }}</view>
          </view>
          <view class="row">
            <view class="text">运费: </view>
            <view class="symbol">0</view>
          </view>
          <view class="row">
            <view class="text">应付金额: </view>
            <view class="symbol">{{ order?.orderInfo.payAmount }}</view>
          </view>
        </view>
      </view>

      <!-- 订单信息 -->
      <view class="detail">
        <view class="title">订单信息</view>
        <view class="row">
          <view class="item">
            订单编号: {{ order?.orderInfo.orderSn }}
            <text class="copy" @tap="onCopy(order?.orderInfo.orderSn)">复制</text>
          </view>
          <view class="item">下单时间:{{ order?.orderInfo.createTime }}</view>
        </view>
      </view>

      <!-- 订单信息 -->
      <view class="detail">
        <view class="title">自提点信息</view>
        <view class="row">
          <view class="item">名称:{{ order?.placeInfo.name }}</view>
          <view class="item">详细地址:{{ order?.placeInfo.place }}</view>
          <view class="item">联系电话:{{ order?.placeInfo.phone }}</view>
        </view>
      </view>

      <!-- 底部操作栏 -->
      <view class="toolbar-height" :style="{ paddingBottom: safeAreaInsets?.bottom + 'px' }"></view>
      <view class="toolbar" :style="{ paddingBottom: safeAreaInsets?.bottom + 'px' }">
        <!-- 待付款状态:展示支付按钮 -->
        <template v-if="order.orderInfo.status == '0'">
          <view class="button primary" @tap="onOrderPay"> 去支付 </view>
          <view class="button" @tap="onCancel"> 取消订单 </view>
        </template>
        <template v-if="order.orderInfo.status == '1'">
          <navigator
            class="button secondary"
            :url="`/pagesOrder/create/create?orderId=${query.id}`"
            hover-class="none"
          >
            再次购买
          </navigator>
        </template>
        <template v-if="order.orderInfo.status == '2'">
          <navigator
            class="button secondary"
            :url="`/pagesOrder/create/create?orderId=${query.id}`"
            hover-class="none"
          >
            再次购买
          </navigator>
        </template>
        <template v-if="order.orderInfo.status == '3'">
          <view class="button primary" @tap="onConfirm"> 确认收货 </view>
        </template>
        <template v-if="order.orderInfo.status == '4'">
          <template v-if="!returnApply">
            <navigator
              v-if="!returnApply"
              class="button primary"
              :url="`/pages/returnApply/index?id=${query.id}`"
              hover-class="none"
            >
              申请售后
            </navigator>
          </template>
          <template v-else>
            <view v-if="returnApply.status == '0'" class="button primary">退货待处理</view>
            <view v-if="returnApply.status == '1'" class="button primary">退货中</view>
            <view v-if="returnApply.status == '2'" class="button primary">退货已完成</view>
            <view v-if="returnApply.status == '3'" class="button primary">退货已拒绝</view>
          </template>
        </template>
      </view>
    </template>
    <template v-else>
      <!-- 骨架屏组件 -->
      <PageSkeleton />
    </template>
  </scroll-view>
  <!-- 取消订单弹窗 -->
  <uni-popup ref="popup" type="bottom" background-color="#fff">
    <view class="popup-root">
      <view class="title">订单取消</view>
      <view class="description">
        <view class="tips">请选择取消订单的原因：</view>
        <view class="cell" v-for="item in reasonList" :key="item" @tap="reason = item">
          <text class="text">{{ item }}</text>
          <text class="icon" :class="{ checked: item === reason }"></text>
        </view>
      </view>
      <view class="footer">
        <view class="button" @tap="popup?.close?.()">取消</view>
        <view class="button primary">确认</view>
      </view>
    </view>
  </uni-popup>
  <!-- 取消订单弹窗 -->
  <uni-popup ref="popup1" type="bottom" background-color="#fff">
    <view class="popup-root">
      <view class="title">订单取消</view>
      <view class="description">
        <view class="tips">请选择取消订单的原因：</view>
        <view class="cell" v-for="item in reasonList" :key="item" @click="ddd(item)">
          <text class="text">{{ item.name }}</text>
          <text class="icon" :class="{ checked: item === reason }"></text>
        </view>
      </view>
      <view class="footer">
        <view class="button" @tap="popup1?.close?.()">取消</view>
        <view class="button primary" @tap="popup1?.close?.(), onOrderCancel()">确认</view>
      </view>
    </view>
  </uni-popup>
</template>

<style lang="scss">
page {
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow: hidden;
}

.navbar {
  width: 750rpx;
  color: #000;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9;
  /* background-color: #f8f8f8; */
  background-color: transparent;

  .wrap {
    position: relative;

    .title {
      height: 44px;
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 32rpx;
      /* color: #000; */
      color: transparent;
    }

    .back {
      position: absolute;
      left: 0;
      height: 44px;
      width: 44px;
      font-size: 44rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      /* color: #000; */
      color: #fff;
    }
  }
}

.viewport {
  background-color: #f7f7f8;
}

.overview {
  display: flex;
  flex-direction: column;
  align-items: center;

  line-height: 1;
  padding-bottom: 30rpx;
  color: #fff;
  background-image: url(https://pcapi-xiaotuxian-front-devtest.itheima.net/miniapp/images/order_bg.png);
  background-size: cover;

  .status {
    font-size: 36rpx;
  }

  .status::before {
    margin-right: 6rpx;
    font-weight: 500;
  }

  .tips {
    margin: 30rpx 0;
    display: flex;
    font-size: 14px;
    align-items: center;

    .money {
      margin-right: 30rpx;
    }
  }

  .button-group {
    margin-top: 30rpx;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .button {
    width: 260rpx;
    height: 64rpx;
    margin: 0 10rpx;
    text-align: center;
    line-height: 64rpx;
    font-size: 28rpx;
    color: #27ba9b;
    border-radius: 68rpx;
    background-color: #fff;
  }
}

.shipment {
  line-height: 1.4;
  padding: 0 20rpx;
  margin: 20rpx 20rpx 0;
  border-radius: 10rpx;
  background-color: #fff;

  .locate,
  .item {
    min-height: 120rpx;
    padding: 30rpx 30rpx 25rpx 75rpx;
    background-size: 50rpx;
    background-repeat: no-repeat;
    background-position: 6rpx center;
  }

  .locate {
    background-image: url(https://pcapi-xiaotuxian-front-devtest.itheima.net/miniapp/images/locate.png);

    .user {
      font-size: 26rpx;
      color: #444;
    }

    .address {
      font-size: 24rpx;
      color: #666;
    }
  }

  .item {
    background-image: url(https://pcapi-xiaotuxian-front-devtest.itheima.net/miniapp/images/car.png);
    border-bottom: 1rpx solid #eee;
    position: relative;

    .message {
      font-size: 26rpx;
      color: #444;
    }

    .date {
      font-size: 24rpx;
      color: #666;
    }
  }
}

.goods {
  margin: 20rpx 20rpx 0;
  padding: 0 20rpx;
  border-radius: 10rpx;
  background-color: #fff;

  .item {
    padding: 30rpx 0;
    border-bottom: 1rpx solid #eee;

    .dianpu {
      margin-bottom: 10rpx;

      .name {
        height: 72rpx;
        font-size: 28rpx;
        color: #444;
      }

      .label {
        color: #fff;
        padding: 7rpx 15rpx 5rpx;
        border-radius: 4rpx;
        font-size: 24rpx;
        background-color: #dc9c4f;
        margin-right: 10rpx;
      }
    }

    .navigator {
      display: flex;
      margin: 20rpx 0;
    }

    .cover {
      width: 170rpx;
      height: 170rpx;
      border-radius: 10rpx;
      margin-right: 20rpx;
    }

    .meta {
      flex: 1;
      display: flex;
      flex-direction: column;
      justify-content: center;
      position: relative;
    }

    .name {
      height: 80rpx;
      font-size: 26rpx;
      color: #444;
    }

    .type {
      line-height: 1.8;
      padding: 0 15rpx;
      margin-top: 6rpx;
      font-size: 24rpx;
      align-self: flex-start;
      border-radius: 4rpx;
      color: #888;
      background-color: #f7f7f8;
    }

    .price {
      display: flex;
      margin-top: 6rpx;
      font-size: 24rpx;
    }

    .symbol {
      font-size: 20rpx;
    }

    .original {
      color: #999;
      text-decoration: line-through;
    }

    .actual {
      margin-left: 10rpx;
      color: #444;
    }

    .text {
      font-size: 22rpx;
    }

    .quantity {
      position: absolute;
      bottom: 0;
      right: 0;
      font-size: 24rpx;
      color: #444;
    }

    .action {
      display: flex;
      flex-direction: row-reverse;
      justify-content: flex-start;
      padding: 30rpx 0 0;

      .button {
        width: 200rpx;
        height: 60rpx;
        text-align: center;
        justify-content: center;
        line-height: 60rpx;
        margin-left: 20rpx;
        border-radius: 60rpx;
        border: 1rpx solid #ccc;
        font-size: 26rpx;
        color: #444;
      }

      .primary {
        color: #27ba9b;
        border-color: #27ba9b;
      }
    }
  }

  .total {
    line-height: 1;
    font-size: 26rpx;
    padding: 20rpx 0;
    color: #666;

    .row {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 10rpx 0;
    }

    .symbol::before {
      content: '¥';
      font-size: 80%;
      margin-right: 3rpx;
    }

    .primary {
      color: #cf4444;
      font-size: 36rpx;
    }
  }
}

.detail {
  line-height: 1;
  padding: 30rpx 20rpx 0;
  margin: 20rpx 20rpx 0;
  font-size: 26rpx;
  color: #666;
  border-radius: 10rpx;
  background-color: #fff;

  .title {
    font-size: 30rpx;
    color: #444;
  }

  .row {
    padding: 20rpx 0;

    .item {
      padding: 10rpx 0;
      display: flex;
      align-items: center;
    }

    .copy {
      border-radius: 20rpx;
      font-size: 20rpx;
      border: 1px solid #ccc;
      padding: 5rpx 10rpx;
      margin-left: 10rpx;
    }
  }
}

.toolbar-height {
  height: 100rpx;
  box-sizing: content-box;
}

.toolbar {
  position: fixed;
  left: 0;
  right: 0;
  bottom: calc(var(--window-bottom));
  z-index: 1;

  height: 100rpx;
  padding: 0 20rpx;
  display: flex;
  align-items: center;
  flex-direction: row-reverse;
  border-top: 1rpx solid #ededed;
  border-bottom: 1rpx solid #ededed;
  background-color: #fff;
  box-sizing: content-box;

  .button {
    display: flex;
    justify-content: center;
    align-items: center;

    width: 200rpx;
    height: 72rpx;
    margin-left: 15rpx;
    font-size: 26rpx;
    border-radius: 72rpx;
    border: 1rpx solid #ccc;
    color: #444;
  }

  .delete {
    order: 4;
  }

  .button {
    order: 3;
  }

  .secondary {
    order: 2;
    color: #27ba9b;
    border-color: #27ba9b;
  }

  .primary {
    order: 1;
    color: #fff;
    background-color: #27ba9b;
  }
}

.popup-root {
  padding: 30rpx 30rpx 0;
  border-radius: 10rpx 10rpx 0 0;
  overflow: hidden;

  .title {
    font-size: 30rpx;
    text-align: center;
    margin-bottom: 30rpx;
  }

  .description {
    font-size: 28rpx;
    padding: 0 20rpx;

    .tips {
      color: #444;
      margin-bottom: 12rpx;
    }

    .cell {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 15rpx 0;
      color: #666;
    }

    .icon::before {
      content: '\e6cd';
      font-family: 'erabbit' !important;
      font-size: 38rpx;
      color: #999;
    }

    .icon.checked::before {
      content: '\e6cc';
      font-size: 38rpx;
      color: #27ba9b;
    }
  }

  .footer {
    display: flex;
    justify-content: space-between;
    padding: 30rpx 0 40rpx;
    font-size: 28rpx;
    color: #444;

    .button {
      flex: 1;
      height: 72rpx;
      text-align: center;
      line-height: 72rpx;
      margin: 0 20rpx;
      color: #444;
      border-radius: 72rpx;
      border: 1rpx solid #ccc;
    }

    .primary {
      color: #fff;
      background-color: #27ba9b;
      border: none;
    }
  }
}
</style>
