<script setup lang="ts">
import { getHomeGoodsGuessLikeAPI } from '@/services/home'
import { onMounted, ref } from 'vue'

// 获取页面参数
const query = defineProps<{
  id: number
  name: string
}>()


const guessList = ref<any[]>([])
const getHomeGoodsGuessLikeData = async () => {
  if (query.id) {
    const res = await getHomeGoodsGuessLikeAPI({ categoryId: query.id, searchType: searchType.value })
    guessList.value = res.data
  }
  if (query.name) {
    const res = await getHomeGoodsGuessLikeAPI({ name: query.name, searchType: searchType.value })
    guessList.value = res.data
  }
}
onMounted(() => {
  getHomeGoodsGuessLikeData()
})

let selectedOption = ref(0)
let searchType = ref(0)
let tabShow = (showid: number, searchTypeID: number) => {
  console.log("展示的是" + showid + ",type是" + searchTypeID);
  selectedOption.value = showid
  searchType.value = searchTypeID
  guessList.value = []
  getHomeGoodsGuessLikeData()
}

</script>

<template>
  <view class="caption">
    <navigator class="search" url="/pages/seacher/index">
      <text class="icon-search">搜索商品</text>
      <text class="icon-scan"></text>
    </navigator>
  </view>
  <view class="shaixuan">
    <view :class="{ 'show': selectedOption === 0 }">
      <text @click="tabShow(0, 0)">推荐</text>
    </view>
    <view :class="{ 'show': selectedOption === 1 }">
      <text @click="tabShow(1, 1)">最新</text>
    </view>
    <view :class="{ 'show': selectedOption === 2 }">
      <text @click="tabShow(2, 2)">销量</text>
    </view>
    <view :class="{ 'show': selectedOption === 4 }">
      <text @click="tabShow(4, 4)">价格⬆</text>
    </view>
    <view :class="{ 'show': selectedOption === 5 }">
      <text @click="tabShow(5, 5)">价格⬇</text>
    </view>
  </view>
  <view class=" guess">
    <navigator class="guess-item" v-for="item in guessList" :key="item.id" :url="`/pages/goods/index?id=${item.id}`">
      <image class="image" mode="aspectFill" :src="item.pic"></image>
      <view class="name"> {{ item.name }} </view>
      <view class="price">
        <text class="small">¥</text>
        <text>{{ item.price }}</text>
      </view>
    </navigator>
  </view>
</template>

<style lang="scss">
:host {
  display: block;
}

page {
  background-color: #f0efef;
}

.dropdown-menu {
  position: absolute;
  top: 30px;
  /* 调整下拉菜单的位置 */
  background-color: #fff;
  border: 1px solid #ccc;
  padding: 10px;
}

.menu-item {
  padding: 8px;
  cursor: pointer;
}

/* 分类标题 */
.caption {
  display: flex;
  justify-content: center;
  line-height: 1;
  padding: 16rpx 0 20rpx 0rpx;
  font-size: 32rpx;
  color: #262626;

  .search {
    width: 90%;
    background-color: #fff;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 10rpx 0 26rpx;
    height: 64rpx;
    margin: 0rpx 5rpx;
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

  .text {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0 28rpx 0 30rpx;

    &::before,
    &::after {
      content: '';
      width: 20rpx;
      height: 20rpx;
      background-image: url(@/static/images/bubble.png);
      background-size: contain;
      margin: 0 10rpx;
    }
  }
}

//筛选
.shaixuan {
  display: flex;

  .show {
    color: #c95b5b;

    .icon {}
  }

  view {
    margin: 25rpx;
    font-size: 32rpx;
    color: #5c5c5c;
  }
}

/* 猜你喜欢 */
.guess {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  padding: 0 20rpx;

  .guess-item {
    width: 345rpx;
    padding: 24rpx 20rpx 20rpx;
    margin-bottom: 20rpx;
    border-radius: 10rpx;
    overflow: hidden;
    background-color: #fff;
  }

  .image {
    width: 304rpx;
    height: 304rpx;
  }

  .name {
    height: 75rpx;
    margin: 10rpx 0;
    font-size: 26rpx;
    color: #262626;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }

  .price {
    line-height: 1;
    padding-top: 4rpx;
    color: #cf4444;
    font-size: 26rpx;
  }

  .small {
    font-size: 80%;
  }
}

// 加载提示文字
.loading-text {
  text-align: center;
  font-size: 28rpx;
  color: #666;
  padding: 20rpx 0;
}
</style>
