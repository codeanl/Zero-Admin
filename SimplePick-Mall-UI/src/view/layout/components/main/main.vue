<template>
  <div class="main">
    <!-- 路由组件出口的位置 -->
    <!-- 过渡动画 -->
    <router-view v-slot="{ Component }">
      <transition name="fade">
        <!-- 渲染layout一级路由组件的子路由 -->
        <component :is="Component" v-if="flag" />
      </transition>
    </router-view>
  </div>
</template>

<script setup lang="ts">
import useLayoutSettingStore from "@/store/setting";
import { watch, ref, nextTick } from "vue";
let layoutSettingStore = useLayoutSettingStore();
let flag = ref(true);
watch(
  () => layoutSettingStore.refresh,
  () => {
    flag.value = false;
    nextTick(() => {
      flag.value = true;
    });
  }
);
</script>

<style scoped>
.main {
  overflow-y: auto;

  .fade-enter-from {
    opacity: 0;
    transform: scale(0);
  }

  .fade-enter-active {
    transition: all 0.3s;
  }

  .fade-enter-to {
    opacity: 1;
    transform: scale(1);
  }
}
</style>
