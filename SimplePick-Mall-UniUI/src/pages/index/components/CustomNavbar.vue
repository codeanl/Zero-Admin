<script setup lang="ts">
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()
</script>

<template>
  <view class="navbar" :style="{ paddingTop: safeAreaInsets!.top + 10 + 'px' }">
    <!-- logo文字 -->
    <view class="logo">
      <!-- <image class="logo-image" src="@/static/images/logo.png"></image> -->
      <text class="logo-text">假的商城V1.0</text>
    </view>
    <!-- 搜索条 -->
    <navigator class="search" url="/pages/seacher/index">
      <text class="icon-search">搜索商品</text>
      <text class="icon-scan"></text>
    </navigator>
  </view>
</template>

<style lang="scss">
/* 自定义导航条 */
.navbar {
  //background-image: url(@/static/images/navigator_bg.png);
  background-color: #a6e531;
  //background-color: #5ad912;
  background-size: cover;
  position: relative;
  display: flex;
  flex-direction: column;
  padding-top: 20px;

  .logo {
    display: flex;
    align-items: center;
    height: 64rpx;
    padding-left: 30rpx;

    .logo-image {
      width: 166rpx;
      height: 39rpx;
    }

    .logo-text {
      flex: 1;
      line-height: 28rpx;
      color: #fff;
      margin: 2rpx 0 0 20rpx;
      padding-left: 20rpx;
      font-size: 26rpx;
    }
  }

  .search {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 10rpx 0 26rpx;
    height: 64rpx;
    margin: 16rpx 20rpx;
    color: #fff;
    font-size: 28rpx;
    border-radius: 32rpx;
    background-color: rgba(255, 255, 255, 0.5);
  }

  .icon-search {
    &::before {
      margin-right: 10rpx;
    }
  }

  .icon-scan {
    font-size: 30rpx;
    padding: 15rpx;
  }
}

.icon-scan {
  font-size: 30rpx;
  padding: 15rpx;
}
</style>
