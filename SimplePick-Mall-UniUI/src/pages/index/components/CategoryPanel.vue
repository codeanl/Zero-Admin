<script setup lang="ts">
// 定义 props 接收数据
defineProps<{
  list: any[]
}>()
</script>

<template>
  <view class="category">
    <navigator class="category-item" hover-class="none" :url="`/pages/productList/index?id=${item.id}`"
      v-for="item in list" :key="item.id">
      <image class="icon" :src="item.icon"></image>
      <text class="text">{{ item.name }}</text>
    </navigator>
  </view>
</template>

<style lang="scss">
@import '../styles/category.scss';
</style>
