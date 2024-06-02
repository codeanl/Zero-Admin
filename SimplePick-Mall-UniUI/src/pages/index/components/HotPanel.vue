<script setup lang="ts">

// 定义 props 接收数据
defineProps<{
  list: any[]
}>()
</script>

<template>
  <!-- 推荐专区 -->
  <view class="panel hot">
    <view class="item" v-for="item in list" :key="item.id">
      <view class="title">
        <text class="title-text">{{ item.name }}</text>
      </view>
      <navigator hover-class="none" :url="`/pages/hot/index?id=${item.id}&&pic=${item.pic}&&name=${item.name} `"
        class="cards">
        <image class="image" mode="aspectFit" :src="item.pic">
        </image>
      </navigator>
    </view>
  </view>
</template>

<style lang="scss">
@import '../styles/hot.scss';
</style>
