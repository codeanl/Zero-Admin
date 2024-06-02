<template>
  <!-- 商品信息 -->
  <view class="goods">
    <view class="item">
      <navigator class="navigator" v-for="item in order?.skuList" :key="item.id"
        :url="`/pages/goods/goods?id=${item.spuId}`" hover-class="none">
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
  </view>
  <!-- 用户信息 -->
  <view class="related">
    <view class="item">
      <text class="text">退货原因</text>
      <view @tap="popup?.open?.()">{{ place ? place.name : "请填写退货原因" }} </view>
    </view>
    <view class="item">
      <text class="text">描述</text>
      <input class="input" :cursor-spacing="30" placeholder="请填写描述" v-model="description" />
    </view>
    <view class="item">
      <text class="text">退货金额</text>
      <text class="account placeholder">¥1887</text>
    </view>
    <view class="item">
      <text class="text">凭证</text>
    </view>
    <!-- <view>
      <uni-file-picker limit="1" title="请上传凭证" @select="upload"></uni-file-picker>
    </view> -->
    <view @tap="onAvatarChange">
      <image class="pz" :src="pz" mode="aspectFill" />
    </view>
  </view>
  <!--  -->
  <view class="toolbar" :style="{ paddingBottom: safeAreaInsets?.bottom + 'px' }">
    <view class="button delete" @click="limit11"> 提交 </view>
  </view>
  <!--  -->
  <uni-popup ref="popup" type="bottom" background-color="#fff">
    <view class="popup-root">
      <view class="title">退货原因</view>
      <view class="description">
        <view class="tips">请选退货原因：</view>
        <!-- @tap="reason = item" -->
        <view class="cell" v-for=" item  in  reasonList " :key="item" @click="ddd(item)">
          <text class="text">{{ item.name }}</text>
          <text class="icon" :class="{ checked: item === reason }"></text>
        </view>
      </view>
      <view class="footer">
        <view class="button primary" @tap="popup?.close?.(), dddd">确认</view>
      </view>
    </view>
  </uni-popup>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { onLoad } from '@dcloudio/uni-app'
import { orderInfo } from '@/services/order'
const popup = ref<UniHelper.UniPopupInstance>()
let reason = ref('')
const description = ref('')
const proofPics = ref('')
onLoad(() => {
  getData()
})
// 获取页面参数
const query = defineProps<{
  id: number
}>()
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()
import { listreturnReason, addreturnApply, returnApplyInfo } from '@/services/return'
const reasonList = ref<any>()
const order = ref<any>()
const getData = async () => {
  const res = await listreturnReason({})
  if (res.code === 200) {
    reasonList.value = res.data
  }
  const res1 = await orderInfo({ id: parseInt(query.id) })
  if (res1.code === 200) {
    order.value = res1.data
  }

}

let place = ref<any>()
let ddd = (item: any) => {
  reason.value = item
  place.value = item
}

let dddd = () => {
  console.log(place);
}

import { useMemberStore } from '@/stores'
const memberStore = useMemberStore()
let pz = ref('')
let limit11 = async () => {
  let res = await addreturnApply({
    userID: memberStore?.profile.info.id,
    orderID: order.value.orderInfo.id,
    returnReasonID: place.value.id,
    status: '0',
    description: description.value,
    proofPics: pz.value,
    returnAmount: order.value.orderInfo.totalAmount,
  })
  if (res.code == 200) {
    uni.redirectTo({ url: `/pages/detail/index?id=${parseInt(query.id)}` })
  }
}

// 
const onAvatarChange = () => {
  uni.chooseMedia({
    // 文件个数
    count: 1,
    // 文件类型
    mediaType: ['image'],
    success: (res) => {
      // 本地路径
      const { tempFilePath } = res.tempFiles[0]
      // 上传
      uploadFile(tempFilePath)
    },
  })
}
const uploadFile = (file: string) => {
  uni.uploadFile({
    url: '/api/upload',
    name: 'file',
    filePath: file,
    success: (res: any) => {
      pz.value = JSON.parse(res.data).data
      console.log(pz.value);
    },
  })
}

</script>

<style lang="scss" scoped>
.goods {
  margin: 20rpx 20rpx 0;
  padding: 0 20rpx;
  border-radius: 10rpx;
  background-color: #fff;

  .pz {
    width: 200rpx;
    height: 200rpx;
    background-color: #e4dddd;
  }

  .item {
    padding: 30rpx 0;
    border-bottom: 1rpx solid #eee;

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
}

//
.related {
  margin: 20rpx;
  padding: 0 20rpx;
  border-radius: 10rpx;
  background-color: #fff;

  .pz {
    width: 300rpx;
    height: 300rpx;
    background-color: #f5f5f5;
  }

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
</style>