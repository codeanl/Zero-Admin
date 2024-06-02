<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import { onShow, onLoad } from '@dcloudio/uni-app'
import { placeList } from '@/services/home'
import { deleteCart } from '@/services/cart'
import { addOrder } from '@/services/order'
import { listplace } from '@/services/place'
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()
// 订单备注
const buyerMessage = ref('')
const userName = ref('')
const userPhone = ref('')
// 配送时间
const deliveryList = ref([
  { type: 1, text: '时间不限 (周一至周日)' },
  { type: 2, text: '工作日送 (周一至周五)' },
  { type: 3, text: '周末配送 (周六至周日)' },
])
// 当前配送时间下标
const activeIndex = ref(0)
// 当前配送时间
const activeDelivery = computed(() => deliveryList.value[activeIndex.value])
// 修改配送时间
const onChangeDelivery: UniHelper.SelectorPickerOnChange = (ev) => {
  activeIndex.value = ev.detail.value
}
const query = defineProps<{
  sku: any
}>()
let skus = ref<any>()
let addPri = 0
onLoad(() => {
  skus.value = JSON.parse(query.sku);
  for (let index = 0; index < skus.value.length; index++) {
    addPri += skus.value[index].price * skus.value[index].count
  }
  getMemberCartData()
})

import { useMemberStore } from '@/stores'
// 获取会员Store
const memberStore = useMemberStore()
const PlaceList = ref<any>()
const getMemberCartData = async () => {
  const res = await listplace({})
  if (res.code === 200) {
    PlaceList.value = res.data
  }
}




let generateOrderNumber = () => {
  const prefix = 'ORD'; // 订单号前缀
  const randomLength = 4; // 随机数长度
  const timestamp = Date.now(); // 当前时间戳
  // 生成随机数
  const random = Math.floor(Math.random() * Math.pow(10, randomLength)).toString();
  // 将时间戳和随机数拼接为订单号
  const orderNumber = prefix + timestamp.toString() + random;
  return orderNumber;
}
let getCurrentDateTime = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0'); // 月份从 0 开始，需要加 1，并补零
  const day = String(now.getDate()).padStart(2, '0'); // 补零
  const hours = String(now.getHours()).padStart(2, '0'); // 补零
  const minutes = String(now.getMinutes()).padStart(2, '0'); // 补零
  const seconds = String(now.getSeconds()).padStart(2, '0'); // 补零
  const dateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
  return dateTime;
}


// 提交订单
const onOrderSubmit = async () => {

  //

}

const popup = ref<UniHelper.UniPopupInstance>()
const alertDialog = ref<UniHelper.UniPopupInstance>()
// 取消原因列表
const reasonList = ref([
  '商品无货',
  '不想要了',
  '商品信息填错了',
  '地址信息填写错误',
  '商品降价',
  '其它',
])
// 订单取消原因
let reason = ref('')

let isture = ref(false)
let place = ref<any>()
let ddd = (item: any) => {
  reason.value = item
  place.value = item
  isture.value = true
}


let check = () => {
  //校验
  if (userName.value === '' || userPhone.value === '') {
    uni.showToast({
      title: '请填写完整信息',
      icon: 'none'
    });
    return;
  }
  // 校验手机号格式
  const phoneRegex = /^1[0-9]{10}$/;
  if (!phoneRegex.test(userPhone.value)) {
    uni.showToast({
      title: '请填写有效手机号',
      icon: 'none'
    });
    return;
  }
  if (!place.value) {
    uni.showToast({
      title: '请选择自提点',
      icon: 'none'
    });
    return;
  }
}


const currentDateTime = getCurrentDateTime();
import { updateOrder } from '@/services/order'
//稍后支付
let toDoPay = async () => {
  check()
  const orderNumber = generateOrderNumber();
  let skuIDs: any = []
  let ids = ref<any>([])
  for (let index = 0; index < skus?.value.length; index++) {
    ids.value.push(parseInt(skus.value[index].id))
    skuIDs.push(
      {
        skuID: skus.value[index].skuID as number,
        count: skus.value[index].count as number,
      }
    )
  }
  const res = await addOrder({
    memberId: memberStore?.profile.info.id,
    placeId: place!.value.id,        //
    couponId: 0,        //
    orderSn: orderNumber,
    memberUserName: memberStore.profile.info.nickname,
    totalAmount: addPri,
    payAmount: addPri,
    freightAmount: 0,
    couponAmount: 0,
    payType: "1",
    status: "0",
    orderType: "1",
    receiverName: userName.value,
    receiverPhone: userPhone.value,
    note: buyerMessage.value,
    skus: skuIDs,
  })
  if (res.code == 200) {
    await deleteCart({ ids: ids.value })
    uni.redirectTo({ url: `/pages/orderList/index?status=0` })
  }
}

//立即支付
let toPay = async () => {
  check()
  const orderNumber = generateOrderNumber();
  let skuIDs: any = []
  let ids = ref<any>([])
  for (let index = 0; index < skus?.value.length; index++) {
    ids.value.push(parseInt(skus.value[index].id))
    skuIDs.push(
      {
        skuID: skus.value[index].skuID as number,
        count: skus.value[index].count as number,
      }
    )
  }
  const res = await addOrder({
    memberId: memberStore?.profile.info.id,
    placeId: place!.value.id,        //
    couponId: 0,        //
    orderSn: orderNumber,
    memberUserName: memberStore.profile.info.nickname,
    totalAmount: addPri,
    payAmount: addPri,
    freightAmount: 0,
    couponAmount: 0,
    payType: "1",
    status: "0",
    orderType: "1",
    receiverName: userName.value,
    receiverPhone: userPhone.value,
    note: buyerMessage.value,
    skus: skuIDs,
  })
  if (res.code == 200) {
    await deleteCart({ ids: ids.value })
    for (const i in res.data.orderID) {
      let res1 = await updateOrder({
        id: res.data.orderID[i],
        status: "1",
        paymentTime: currentDateTime
      })
    }
    uni.redirectTo({ url: `/pages/payment/index?status=1` })
  }
}

const dddd=()=>{
  console.log(111);
  
}
</script>
<template>
  <scroll-view scroll-y class="viewport">
    <!-- 商品信息 -->
    <view class="goods">
      <navigator v-for="(item, index) in skus" :key="index" :url="`/pages/goods/goods?id=${item.id}`" class="item"
        hover-class="none">
        <image class="picture" :src="item.pic" />
        <view class="meta">
          <view class="name ellipsis"> {{ item.name }} </view>
          <view class="attrs">{{ item.tag }}</view>
          <view class="prices">
            <view class="pay-price symbol">{{ item.price }}</view>
            <view class="price symbol">{{ item.price }}</view>
          </view>
          <view class="count">x{{ item.count }}</view>
        </view>
      </navigator>
    </view>

    <!-- 配送及支付方式 -->
    <view class="related">
      <view class="item">
        <text class="text">订单备注</text>
        <input class="input" :cursor-spacing="30" placeholder="留言前先与商家沟通确认" v-model="buyerMessage" />
      </view>
    </view>

    <!-- 用户信息 -->
    <view class="related">
      <view class="item">
        <text class="text">姓名</text>
        <input class="input" :cursor-spacing="30" placeholder="请填写姓名" v-model="userName" />
      </view>
      <view class="item">
        <text class="text">联系电话</text>
        <input class="input" :cursor-spacing="30" placeholder="请填写联系电话" v-model="userPhone" />
      </view>
      <view class="item">
        <text class="text">自提点</text>
        <view @tap="popup?.open?.()">{{ place ? place.name : "请填写自提点" }} </view>
      </view>
    </view>

    <!-- 支付金额 -->
    <view class="settlement">
      <view class="item">
        <text class="text">商品总价: </text>
        <text class="number symbol">{{ addPri }}</text>
      </view>
    </view>
  </scroll-view>

  <!-- 吸底工具栏 -->
  <view class="toolbar" :style="{ paddingBottom: safeAreaInsets?.bottom + 'px' }">
    <view class="total-pay symbol">
      <text class="number"> {{ addPri }} </text>
    </view>
    <view class="button" @tap=" alertDialog?.open?.()">
      提交订单
    </view>
  </view>

  <uni-popup ref="popup" type="bottom" background-color="#fff">
    <view class="popup-root">
      <view class="title">自提点</view>
      <view class="description">
        <view class="tips">请选自提点：</view>
        <!-- @tap="reason = item" -->
        <view class="cell" v-for="item in PlaceList" :key="item" @click="ddd(item)">
          <text class="text">{{ item.name }}</text>
          <text class="icon" :class="{ checked: item === reason }"></text>
        </view>
      </view>
      <view class="footer">
        <!-- <view class="button" @tap="popup?.close?.()">取消</view> -->
        <view class="button primary" @tap="popup?.close?.() ">确认</view>
      </view>
    </view>
  </uni-popup>

  <uni-popup ref="alertDialog" type="dialog">
    <uni-popup-dialog type="success" cancelText="稍后支付" confirmText="现在支付" content="请支付订单!" @confirm="toPay"
      @close="toDoPay"></uni-popup-dialog>
  </uni-popup>
</template>

<style lang="scss">
//
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

//
page {
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow: hidden;
  background-color: #f4f4f4;
}

.symbol::before {
  content: '¥';
  font-size: 80%;
  margin-right: 5rpx;
}

.shipment {
  margin: 20rpx;
  padding: 30rpx 30rpx 30rpx 84rpx;
  font-size: 26rpx;
  border-radius: 10rpx;
  background: url(https://pcapi-xiaotuxian-front-devtest.itheima.net/miniapp/images/locate.png) 20rpx center / 50rpx no-repeat #fff;
  position: relative;

  .icon {
    font-size: 36rpx;
    color: #333;
    transform: translateY(-50%);
    position: absolute;
    top: 50%;
    right: 20rpx;
  }

  .user {
    color: #333;
    margin-bottom: 5rpx;
  }

  .address {
    color: #666;
  }
}

.goods {
  margin: 20rpx;
  padding: 0 20rpx;
  border-radius: 10rpx;
  background-color: #fff;

  .item {
    display: flex;
    padding: 30rpx 0;
    border-top: 1rpx solid #eee;

    &:first-child {
      border-top: none;
    }

    .picture {
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

    .attrs {
      line-height: 1.8;
      padding: 0 15rpx;
      margin-top: 6rpx;
      font-size: 24rpx;
      align-self: flex-start;
      border-radius: 4rpx;
      color: #888;
      background-color: #f7f7f8;
    }

    .prices {
      display: flex;
      align-items: baseline;
      margin-top: 6rpx;
      font-size: 28rpx;

      .pay-price {
        margin-right: 10rpx;
        color: #cf4444;
      }

      .price {
        font-size: 24rpx;
        color: #999;
        text-decoration: line-through;
      }
    }

    .count {
      position: absolute;
      bottom: 0;
      right: 0;
      font-size: 26rpx;
      color: #444;
    }
  }
}

.related {
  margin: 20rpx;
  padding: 0 20rpx;
  border-radius: 10rpx;
  background-color: #fff;

  .item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    min-height: 80rpx;
    font-size: 26rpx;
    color: #333;
  }

  .input {
    flex: 1;
    text-align: right;
    margin: 20rpx 0;
    padding-right: 20rpx;
    font-size: 26rpx;
    color: #999;
  }

  .item .text {
    width: 125rpx;
  }

  .picker {
    color: #666;
  }

  .picker::after {
    content: '\e6c2';
  }
}

/* 结算清单 */
.settlement {
  margin: 20rpx;
  padding: 0 20rpx;
  border-radius: 10rpx;
  background-color: #fff;

  .item {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 80rpx;
    font-size: 26rpx;
    color: #333;
  }

  .danger {
    color: #cf4444;
  }
}

/* 吸底工具栏 */
.toolbar {
  position: fixed;
  left: 0;
  right: 0;
  bottom: calc(var(--window-bottom));
  z-index: 1;

  background-color: #fff;
  height: 100rpx;
  padding: 0 20rpx;
  border-top: 1rpx solid #eaeaea;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-sizing: content-box;

  .total-pay {
    font-size: 40rpx;
    color: #cf4444;

    .decimal {
      font-size: 75%;
    }
  }

  .button {
    width: 220rpx;
    text-align: center;
    line-height: 72rpx;
    font-size: 26rpx;
    color: #fff;
    border-radius: 72rpx;
    background-color: #27ba9b;
  }

  .disabled {
    opacity: 0.6;
  }
}
</style>