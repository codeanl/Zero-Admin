<template>
  <div class="tabbar">
    <!-- 左侧 -->
    <div class="left">
      <el-icon style="margin-right: 20px" @click="changMenu">
        <component
          :is="layoutSettingStore.fold ? 'Fold' : 'Expand'"
        ></component>
      </el-icon>
      <!-- 面包屑 -->
      <el-breadcrumb separator-icon="ArrowRight">
        <el-breadcrumb-item
          v-for="(item, index) in $route.matched"
          :key="index"
          v-show="item.meta.title"
          :to="item.path"
        >
          <el-icon style="margin: 0px 4px">
            <component :is="item.meta.icon"></component>
          </el-icon>
          <span>{{ item.meta.title }}</span>
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <!--  -->
    <div class="right">
      <el-button
        type="primary"
        size="small"
        icon="Refresh"
        circle
        @click="updateRefresh"
      />
      <el-button
        type="primary"
        size="small"
        icon="FullScreen"
        circle
        @click="FullScreen"
      />
      <!-- 系统设置 -->
      <el-popover
        placement="bottom"
        title="主题设置"
        :width="200"
        trigger="hover"
      >
        <el-form>
          <el-form-item label="主题颜色">
            <el-color-picker
              v-model="color"
              show-alpha
              :predefine="predefineColors"
              size="small"
              @change="setColor"
            />
          </el-form-item>
          <el-form-item label="暗黑模式">
            <el-switch
              v-model="dark"
              size="small"
              inline-prompt
              active-icon="Moon"
              inactive-icon="Sunny"
              @change="changeDark"
            />
          </el-form-item>
        </el-form>
        <template #reference>
          <el-button type="primary" size="small" icon="Setting" circle />
        </template>
      </el-popover>
      <!--  -->
      <img
        :src="userStore.avatar"
        style="width: 24px; height: 24px; margin: 0px 12px; border-radius: 50%"
      />
      <!-- 下拉菜单 -->
      <el-dropdown>
        <span class="el-dropdown-link">
          {{ userStore.username }}
          <el-icon class="el-icon--right">
            <arrow-down />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="profile">个人信息</el-dropdown-item>
          </el-dropdown-menu>
          <el-dropdown-menu>
            <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup lang="ts">
import useLayoutSettingStore from "@/store/setting";
import { ref } from "vue";
let layoutSettingStore = useLayoutSettingStore();
import { useRoute } from "vue-router";
let $route = useRoute();
import { useRouter } from "vue-router";
let $router = useRouter();
//用户相关信息
import useUserStore from "@/store/sys/user";
let userStore = useUserStore();

//收回menu
const changMenu = () => {
  //图标切换
  layoutSettingStore.fold = !layoutSettingStore.fold;
};
//刷新按钮
const updateRefresh = () => {
  layoutSettingStore.refresh = !layoutSettingStore.refresh;
};
//全屏按钮
const FullScreen = () => {
  let full = document.fullscreenElement;
  if (!full) {
    //文档跟节点requestFullscreen实现全屏模式
    document.documentElement.requestFullscreen();
  } else {
    //退出全屏
    document.exitFullscreen();
  }
};
//退出登录
const logout = async () => {
  await userStore.userLogout();
  $router.push({ path: "/login" });
};
//个人信息
const profile = async () => {
  $router.push({ path: "/setting/profile" });
};
const color = ref("rgba(255, 69, 0, 0.68)");
const predefineColors = ref([
  "#ff4500",
  "#ff8c00",
  "#ffd700",
  "#90ee90",
  "#00ced1",
  "#1e90ff",
  "#c71585",
  "rgba(255, 69, 0, 0.68)",
  "rgb(255, 120, 0)",
  "hsv(51, 100, 98)",
  "hsva(120, 40, 94, 0.5)",
  "hsl(181, 100%, 37%)",
  "hsla(209, 100%, 56%, 0.73)",
  "#c7158577",
]);

//收集开关的数据
let dark = ref<boolean>(false);

//暗黑模式
// 创建一个开关来控制 暗黑模式 的 class 类名。
const changeDark = () => {
  let html = document.documentElement;
  dark.value ? (html.className = "dark") : (html.className = "");
};
//设置主题颜色
const setColor = () => {
  let html = document.documentElement;
  html.style.setProperty("--el-color-primary", color.value);
};
</script>

<style scoped lang="scss">
.tabbar {
  height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;

  .left {
    display: flex;
    align-items: center;
  }

  .right {
    display: flex;
    align-items: center;
  }
}
</style>
