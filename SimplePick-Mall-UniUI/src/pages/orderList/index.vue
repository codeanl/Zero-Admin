<script setup lang="ts">
import { ref } from 'vue'
import { onLoad } from '@dcloudio/uni-app'
import { useMemberStore } from '@/stores'
import { orderStateList } from '@/services/constants'
// 获取会员Store
const memberStore = useMemberStore()
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()
// tabs 数据
const orderTabs = ref([
  { orderState: 100, title: '全部', isRender: false },
  { orderState: 0, title: '待付款', isRender: false },
  { orderState: 1, title: '待发货', isRender: false },
  { orderState: 2, title: '运输中', isRender: false },
  { orderState: 3, title: '待提货', isRender: false },
])

// 高亮下标
const activeIndex = ref(orderTabs.value.findIndex((v) => v.orderState === parseInt(query.status)))
// 默认渲染容器
orderTabs.value[activeIndex.value].isRender = true
const query = defineProps<{
  status: string
}>()

import { listOrder } from '@/services/order'
onLoad(() => {
  getData(query.status)
})
const order = ref<any>()
const getData = async (status: number) => {
  order.value = []
  if (status < 100) {
    const res = await listOrder({ status: status, userID: memberStore.profile.info.id })
    if (res.code === 200) {
      order.value = res.data
    }
  } else {
    const res = await listOrder({ userID: memberStore.profile.info.id })
    if (res.code === 200) {
      order.value = res.data
    }
  }
}

let now = ref()
let handleTabClick = (index: any, id: any) => {
  console.log(index, id)
  now.value = id
  activeIndex.value = index
  orderTabs.value.forEach((item, i) => {
    item.isRender = i === index
  })
  getData(id.toString())
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

import { orderInfo, updateOrder } from '@/services/order'
// 订单支付
const onOrderPay = async (id: number) => {
  const currentDateTime = getCurrentDateTime()
  let res = await updateOrder({
    id: id,
    status: '1',
    paymentTime: currentDateTime,
  })
  if (res.code == 200) {
    // 关闭当前页，再跳转支付结果页
    uni.redirectTo({ url: `/pages/payment/index?id=${id}` })
  }
}
// 模拟发货
const onDeliver = async (id: number) => {
  let res = await updateOrder({
    id: id,
    status: '3',
  })
  if (res.code == 200) {
    uni.showToast({ icon: 'success', title: '模拟发货完成' })
  }
}

const onCancel = async (id: number) => {
  console.log(now.value)
  let res = await updateOrder({
    id: id,
    status: '5',
  })
  if (res.code == 200) {
    getData(now.value.toString())
    uni.showToast({ icon: 'success', title: '取消订单完成' })
  }
}
const onConfirm = async (id: number) => {
  const currentDateTime = getCurrentDateTime()
  // 二次确认弹窗
  uni.showModal({
    content: '为保障您的权益，请收到货并确认无误后，再确认收货',
    confirmColor: '#27BA9B',
    success: async (success) => {
      if (success.confirm) {
        let res = await updateOrder({
          id: id,
          status: '4',
          receiveTime: currentDateTime,
        })
        if (res.code == 200) {
          getData(now.value.toString())
          uni.showToast({ icon: 'success', title: '确认收货完成' })
        }
      }
    },
  })
}
const onComment = async (id: number) => {
  const currentDateTime = getCurrentDateTime()
  let res = await updateOrder({
    id: id,
    status: '4',
    commentTime: currentDateTime,
  })
  if (res.code == 200) {
    uni.showToast({ icon: 'success', title: '评论完成' })
  }
}
const isTriggered = ref(false)
const onRefresherrefresh = async () => {
  isTriggered.value = true

  // isTriggered.value = false
}

const createOrder = () => {
  console.log('111')
}
</script>

<template>
  <view class="viewport">
    <!-- tabs -->
    <view class="tabs">
      <text
        class="item"
        v-for="(item, index) in orderTabs"
        :key="item.title"
        @tap="handleTabClick(index, item.orderState)"
      >
        {{ item.title }}
      </text>
      <!-- 游标 -->
      <view class="cursor" :style="{ left: activeIndex * 20 + '%' }"></view>
    </view>
    <!-- 滑动容器 -->
    <swiper class="swiper" :current="activeIndex" @change="activeIndex = $event.detail.current">
      <!-- 滑动项 -->
      <swiper-item v-for="item in orderTabs" :key="item.title">
        <!-- 订单列表 -->
        <scroll-view scroll-y class="orders" @refresherrefresh="onRefresherrefresh">
          <view class="card" v-for="item in order" :key="item">
            <!-- 订单信息 -->
            <view class="status">
              <navigator class="dianpu" url="/pages/login/index">
                <text class="label">店铺</text>
                <text class="name">{{ item?.merchantInfo.name }} ></text>
              </navigator>
              <!-- <text class="date">{{ item.createTime }}</text> -->
              <!-- 订单状态文字 -->
              <!-- 待评价/已完成/已取消 状态: 展示删除订单 -->
              <text class="statusText">{{ orderStateList[item.status].text }}</text>
              <text class="icon-delete"></text>
            </view>
            <!-- 商品信息，点击商品跳转到订单详情，不是商品详情 -->
            <navigator
              v-for="sku in item.skuList"
              :key="sku.id"
              class="goods"
              :url="`/pages/detail/index?id=${item.id}`"
              hover-class="none"
            >
              <view class="cover">
                <image class="image" mode="aspectFit" :src="sku.pic"></image>
              </view>
              <view class="meta">
                <view class="name ellipsis">{{ sku.productName }}</view>
                <view class="type">{{ sku.tag }}</view>
              </view>
            </navigator>
            <!-- 支付信息 -->
            <view class="payment">
              <text class="quantity">共1件商品</text>
              <text>实付</text>
              <text class="amount"> <text class="symbol">¥</text>{{ item.payAmount }}</text>
            </view>
            <!-- 订单操作按钮 -->
            <view class="action">
              <!-- 待付款状态：显示去支付按钮 -->
              <template v-if="item.status == '0'">
                <view class="button" @tap="onCancel(item.id)">取消订单</view>
                <view class="button primary" @tap="onOrderPay(item.id)">去支付</view>
              </template>
              <template v-if="item.status == '1'">
                <navigator
                  class="button secondary"
                  :url="`/pages/orderCreate/index?sku=${encodeURIComponent(
                    JSON.stringify(item.skuList),
                  )}`"
                  hover-class="none"
                >
                  再次购买
                </navigator>
              </template>
              <template v-if="item.status == '2'">
                <navigator
                  class="button secondary"
                  :url="`/pages/orderCreate/index?sku=${encodeURIComponent(
                    JSON.stringify(item.skuList),
                  )}`"
                  hover-class="none"
                >
                  再次购买
                </navigator>
              </template>
              <template v-if="item.status == '3'">
                <view class="button primary" @tap="onConfirm(item.id)">确认收货</view>
              </template>
            </view>
          </view>
          <!-- 底部提示文字 -->
          <view class="loading-text" :style="{ paddingBottom: safeAreaInsets?.bottom + 'px' }">
            {{ true ? '没有更多数据~' : '正在加载...' }}
          </view>
        </scroll-view>
      </swiper-item>
    </swiper>
  </view>
</template>

<style lang="scss">
page {
  height: 100%;
  overflow: hidden;
}

.viewport {
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: #fff;
}

// tabs
.tabs {
  display: flex;
  justify-content: space-around;
  line-height: 60rpx;
  margin: 0 10rpx;
  background-color: #fff;
  box-shadow: 0 4rpx 6rpx rgba(240, 240, 240, 0.6);
  position: relative;
  z-index: 9;

  .item {
    flex: 1;
    text-align: center;
    padding: 20rpx;
    font-size: 28rpx;
    color: #262626;
  }

  .cursor {
    position: absolute;
    left: 0;
    bottom: 0;
    width: 20%;
    height: 6rpx;
    padding: 0 50rpx;
    background-color: #27ba9b;
    /* 过渡效果 */
    transition: all 0.4s;
  }
}

// swiper
.swiper {
  background-color: #f7f7f8;
}

// 订单列表
.orders {
  .card {
    min-height: 100rpx;
    padding: 20rpx;
    margin: 20rpx 20rpx 0;
    border-radius: 10rpx;
    background-color: #fff;

    &:last-child {
      padding-bottom: 40rpx;
    }
  }

  .status {
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 28rpx;
    color: #999;
    margin-bottom: 15rpx;

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

    .date {
      color: #666;
      flex: 1;
    }

    .primary {
      color: #ff9240;
    }

    .statusText {
      margin-left: 60rpx;
    }

    .icon-delete {
      line-height: 1;
      margin-left: 10rpx;
      padding-left: 10rpx;
      border-left: 1rpx solid #e3e3e3;
    }
  }

  .goods {
    display: flex;
    margin-bottom: 20rpx;

    .cover {
      width: 170rpx;
      height: 170rpx;
      margin-right: 20rpx;
      border-radius: 10rpx;
      overflow: hidden;
      position: relative;
    }

    .quantity {
      position: absolute;
      bottom: 0;
      right: 0;
      line-height: 1;
      padding: 6rpx 4rpx 6rpx 8rpx;
      font-size: 24rpx;
      color: #fff;
      border-radius: 10rpx 0 0 0;
      background-color: rgba(0, 0, 0, 0.6);
    }

    .meta {
      flex: 1;
      display: flex;
      flex-direction: column;
      justify-content: center;
    }

    .name {
      height: 80rpx;
      font-size: 26rpx;
      color: #444;
    }

    .type {
      line-height: 1.8;
      padding: 0 15rpx;
      margin-top: 10rpx;
      font-size: 24rpx;
      align-self: flex-start;
      border-radius: 4rpx;
      color: #888;
      background-color: #f7f7f8;
    }

    .more {
      flex: 1;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 22rpx;
      color: #333;
    }
  }

  .payment {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    line-height: 1;
    padding: 20rpx 0;
    text-align: right;
    color: #999;
    font-size: 28rpx;
    border-bottom: 1rpx solid #eee;

    .quantity {
      font-size: 24rpx;
      margin-right: 16rpx;
    }

    .amount {
      color: #444;
      margin-left: 6rpx;
    }

    .symbol {
      font-size: 20rpx;
    }
  }

  .action {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    padding-top: 20rpx;

    .button {
      width: 180rpx;
      height: 60rpx;
      display: flex;
      justify-content: center;
      align-items: center;
      margin-left: 20rpx;
      border-radius: 60rpx;
      border: 1rpx solid #ccc;
      font-size: 26rpx;
      color: #444;
    }

    .secondary {
      color: #27ba9b;
      border-color: #27ba9b;
    }

    .primary {
      color: #fff;
      background-color: #27ba9b;
    }
  }

  .loading-text {
    text-align: center;
    font-size: 28rpx;
    color: #666;
    padding: 20rpx 0;
  }
}
</style>
