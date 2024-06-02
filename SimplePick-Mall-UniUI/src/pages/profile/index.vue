<script setup lang="ts">
import { getMemberProfileAPI, putMemberProfileAPI } from '../../services/profile'
import { useMemberStore } from '@/stores'
// import { formatDate } from '@/utils'
import { onLoad } from '@dcloudio/uni-app'
import { ref } from 'vue'

// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()

// 获取个人信息，修改个人信息需提供初始值
const profile = ref()
const getMemberProfileData = async () => {
  const requestData = { id: memberStore.profile.info.id }
  const res = await getMemberProfileAPI(requestData)
  profile.value = res.data
  // 同步 Store 的头像和昵称，用于我的页面展示
  memberStore.profile!.info.avatar = res.data.avatar
  memberStore.profile!.info.nickname = res.data.nickname
}

onLoad(() => {
  getMemberProfileData()
})

const memberStore = useMemberStore()
// 修改头像
const onAvatarChange = () => {
  // 调用拍照/选择图片
  // 选择图片条件编译
  // #ifdef H5 || APP-PLUS
  // 微信小程序从基础库 2.21.0 开始， wx.chooseImage 停止维护，请使用 uni.chooseMedia 代替
  uni.chooseImage({
    count: 1,
    success: (res) => {
      // 文件路径
      const tempFilePaths = res.tempFilePaths
      // 上传
      uploadFile(tempFilePaths[0])
    },
  })
  // #endif
  // #ifdef MP-WEIXIN
  // uni.chooseMedia 仅支持微信小程序端
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
  // #endif
}
// 文件上传-兼容小程序端、H5端、App端
const uploadFile = (file: string) => {
  uni.uploadFile({
    url: '/api/upload',
    name: 'file',
    filePath: file,
    success: (res: any) => {
      updateAvatar(JSON.parse(res.data))
    },
  })
}
let updateAvatar = async (data: any) => {
  console.log(data.data)
  memberStore.profile!.info.avatar = data.data
  let res = await putMemberProfileAPI({
    id: memberStore.profile.info.id,
    avatar: data.data,
  })
  if (res.code == 200) {
    uni.showToast({ icon: 'success', title: '更新成功' })
    profile.value.avatar = data.data
  } else {
    uni.showToast({ icon: 'error', title: '出现错误' })
  }
}

// 修改性别
const onGenderChange: UniHelper.RadioGroupOnChange = (ev) => {
  profile.value.gender = ev.detail.value
}

// 点击保存提交表单
const onSubmit = async () => {
  const { nickname, gender, email, signature } = profile.value
  const res = await putMemberProfileAPI({
    id: memberStore.profile.info.id,
    nickname,
    gender,
    email,
    signature,
  })
  if (res.code == 200) {
    uni.showToast({ icon: 'success', title: '保存成功' })
    setTimeout(() => {
      uni.navigateBack()
    }, 400)
    // 更新Store昵称
    memberStore.profile!.nickname = res.data.nickname
    memberStore.profile!.email = res.data.email
    memberStore.profile!.gender = res.data.gender
    memberStore.profile!.signature = res.data.signature
  }
}
</script>

<template>
  <view class="viewport">
    <!-- 导航栏 -->
    <view class="navbar" :style="{ paddingTop: safeAreaInsets?.top + 'px' }">
      <navigator open-type="navigateBack" class="back icon-left" hover-class="none"></navigator>
      <view class="title">个人信息</view>
    </view>
    <!-- 头像 -->
    <view class="avatar">
      <view @tap="onAvatarChange" class="avatar-content">
        <image class="image" :src="profile?.avatar" mode="aspectFill" />
        <text class="text">点击修改头像</text>
      </view>
    </view>
    <!-- 表单 -->
    <view class="form">
      <!-- 表单内容 -->
      <view class="form-content">
        <view class="form-item">
          <text class="label">账号</text>
          <text class="account placeholder">{{ profile?.username }}</text>
        </view>
        <view class="form-item">
          <text class="label">昵称</text>
          <input class="input" type="text" placeholder="请填写昵称" v-model="profile!.nickname" />
        </view>
        <view class="form-item">
          <text class="label">邮箱</text>
          <input class="input" type="text" placeholder="请填写邮箱" v-model="profile!.email" />
        </view>
        <view class="form-item">
          <text class="label">性别</text>
          <radio-group @change="onGenderChange">
            <label class="radio">
              <radio value="0" color="#27ba9b" :checked="profile?.gender === '0'" />
              未知
            </label>
            <label class="radio">
              <radio value="1" color="#27ba9b" :checked="profile?.gender === '1'" />
              男
            </label>
            <label class="radio">
              <radio value="2" color="#27ba9b" :checked="profile?.gender === '2'" />
              女
            </label>
          </radio-group>
        </view>

        <!-- #ifdef MP-WEIXIN -->
        <!-- #endif -->
        <view class="form-item">
          <text class="label">个性签名</text>
          <input
            class="input"
            type="text"
            placeholder="请填写个性签名"
            v-model="profile!.signature"
          />
        </view>
      </view>
      <!-- 提交按钮 -->
      <button @tap="onSubmit" class="form-button">保 存</button>
    </view>
  </view>
</template>

<style lang="scss">
page {
  background-color: #f4f4f4;
}

.viewport {
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: #fe9333;
  //background-image: url(https://pcapi-xiaotuxian-front-devtest.itheima.net/miniapp/images/order_bg.png);
  background-size: auto 420rpx;
  background-repeat: no-repeat;
}

// 导航栏
.navbar {
  position: relative;

  .title {
    height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 16px;
    font-weight: 500;
    color: #fff;
  }

  .back {
    position: absolute;
    height: 40px;
    width: 40px;
    left: 0;
    font-size: 20px;
    color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
  }
}

// 头像
.avatar {
  text-align: center;
  width: 100%;
  height: 260rpx;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  .image {
    width: 160rpx;
    height: 160rpx;
    border-radius: 50%;
    background-color: #eee;
  }

  .text {
    display: block;
    padding-top: 20rpx;
    line-height: 1;
    font-size: 26rpx;
    color: #fff;
  }
}

// 表单
.form {
  background-color: #f4f4f4;

  &-content {
    margin: 20rpx 20rpx 0;
    padding: 0 20rpx;
    border-radius: 10rpx;
    background-color: #fff;
  }

  &-item {
    display: flex;
    height: 96rpx;
    line-height: 46rpx;
    padding: 25rpx 10rpx;
    background-color: #fff;
    font-size: 28rpx;
    border-bottom: 1rpx solid #ddd;

    &:last-child {
      border: none;
    }

    .label {
      width: 180rpx;
      color: #333;
    }

    .account {
      color: #666;
    }

    .input {
      flex: 1;
      display: block;
      height: 46rpx;
    }

    .radio {
      margin-right: 20rpx;
    }

    .picker {
      flex: 1;
    }

    .placeholder {
      color: #808080;
    }
  }

  &-button {
    height: 80rpx;
    text-align: center;
    line-height: 80rpx;
    margin: 30rpx 20rpx;
    color: #fff;
    border-radius: 80rpx;
    font-size: 30rpx;
    background-color: #27ba9b;
  }
}
</style>
