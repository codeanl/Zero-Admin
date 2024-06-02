<template>
  <view class="caption">
    <navigator class="search">
      <input type="text" v-model="inputValue" placeholder="请输入内容" />
      <navigator class="icon-scan" :url="`/pages/productList/index?name=${inputValue}`">
      </navigator>
    </navigator>
  </view>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
let inputValue = ref<any>('')

</script>

<style lang="scss" scoped>
.caption {
  display: flex;
  justify-content: center;
  line-height: 1;
  padding: 16rpx 0 40rpx;
  font-size: 32rpx;
  color: #262626;

  .search {
    width: 90%;
    background-color: #f2f2f2;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 10rpx 0 26rpx;
    height: 64rpx;
    margin: 0rpx 10rpx;
    color: #242424;
    font-size: 28rpx;
    border-radius: 32rpx;

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
}
</style>