<script setup lang="ts">
import { ref } from 'vue'

const activeIndex = ref(0)

// 当 swiper 下标发生变化时触发
const onChange: UniHelper.SwiperOnChange = (ev) => {
  activeIndex.value = ev.detail.current
}
// 定义 props 接收
defineProps<{
  list: any[]
}>()
</script>

<template>
  <view class="carousel">
    <swiper :circular="true" :autoplay="false" :interval="3000" @change="onChange">
      <swiper-item v-for="item in list" :key="item.id">
        <navigator url="/pages/index/index" hover-class="none" class="navigator">
          <image mode="aspectFill" class="image" :src="item.pic"></image>
        </navigator>
      </swiper-item>
    </swiper>
    <!-- 指示点 -->
    <view class="indicator">
      <text v-for="(item, index) in list" :key="item.id" class="dot" :class="{ active: index === activeIndex }"></text>
    </view>
  </view>
</template>

<!-- <style lang="scss">
@import './styles/Swiper.scss';
</style> -->
<style lang="scss">
:host {
  display: block;
  height: 280rpx;
  background-image: linear-gradient(#a6e531 0, #a6e531 100rpx, transparent 100rpx);

}

/* 轮播图 */
.carousel {
  height: 100%;
  position: relative;
  overflow: hidden;
  transform: translateY(0);
  background-color: #f5f4f4;
  width: 97%;
  margin-left: auto;
  margin-right: auto;
  border-radius: 20rpx;

  .indicator {
    position: absolute;
    left: 0;
    right: 0;
    bottom: 16rpx;
    display: flex;
    justify-content: center;

    .dot {
      width: 30rpx;
      height: 6rpx;
      margin: 0 8rpx;
      border-radius: 6rpx;
      background-color: rgba(255, 255, 255, 0.4);
    }

    .active {
      background-color: #fff;
    }
  }

  .navigator,
  .image {
    width: 100%;
    height: 100%;
  }
}
</style>