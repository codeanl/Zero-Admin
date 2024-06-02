<template>
  <div class="logo" v-if="setting.logoHidden">
    <img :src="setting.logo" alt="" />
    <p :class="{ menu: LayoutSettingStore.fold }">{{ setting.title }}</p>
  </div>
</template>

<script setup lang="ts">
import setting from "@/setting";
import useLayoutSettingStore from "@/store/setting";
let LayoutSettingStore = useLayoutSettingStore();
</script>

<style lang="scss">
.logo {
  height: $base-menu-logo-height;
  width: 100%;
  display: flex;
  align-items: center;
  padding: 10px;

  img {
    width: 40px;
  }
  p {
    font-size: $base-logo-title-fontSize;
    margin-left: 5px;
  }
  .menu {
    display: none;
  }
}
</style>
