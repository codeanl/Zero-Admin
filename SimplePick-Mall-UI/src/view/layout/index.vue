<template>
  <div class="layout">
    <el-container>
      <!-- 左边 logo和菜单 -->
      <!-- fold: LayoutSettingStore.fold为菜单是否折叠 fold的布尔存放在仓库 当true时有一个class为fold的样式 -->
      <el-aside
        class="layout_slider"
        :class="{ fold: LayoutSettingStore.fold ? true : false }"
      >
        <Logo class="logo"></Logo>
        <!-- 滚动条 -->
        <el-scrollbar class="scrollbar">
          <!-- 菜单组件 -->
          <el-menu
            background-color="#4f6a84"
            text-color="#FFFFFF"
            active-text-color="#3e98f3"
            :default-active="$route.path"
            :collapse="LayoutSettingStore.fold ? true : false"
            :collapse-transition="false"
          >
            <Menu :menuList="userStore.menuRoutes"></Menu>
          </el-menu>
        </el-scrollbar>
      </el-aside>
      <el-container>
        <!-- 顶部 -->
        <el-header class="layout_tabble">
          <Tabbar></Tabbar>
        </el-header>
        <!--  -->
        <el-main class="layout_main">
          <Main></Main>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script setup lang="ts">
import Logo from "./components/logo/logo.vue";
import Menu from "./components/menu/menu.vue";
import Main from "./components/main/main.vue";
import Tabbar from "./components/tabbar/tabbar.vue";

//获取相关仓库
import useLayoutSettingStore from "@/store/setting";
let LayoutSettingStore = useLayoutSettingStore();
import useUserStore from "@/store/sys/user";
let userStore = useUserStore();
</script>
<style scoped lang="scss">
.layout {
  width: 100%;
  height: 100vh;

  .layout_slider {
    background: $base-menu-nackground;
    height: 100vh;
    width: 200px;
    border: 0;
    color: white;

    .scrollbar {
      width: 100%;
      height: calc(100vh - $base-menu-logo-height);

      .el-menu {
        border-right: none;
      }
    }

    &.fold {
      width: 65px;
      height: 100vh;
    }
  }

  .layout_tabble {
    height: 60px;
  }

  .layout_main {
    height: 100vh;
    height: calc(100vh - 60px);
  }
}
</style>
