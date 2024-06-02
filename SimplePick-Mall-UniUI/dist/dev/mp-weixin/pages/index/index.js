"use strict";
const common_vendor = require("../../common/vendor.js");
const services_home = require("../../services/home.js");
const composables_index = require("../../composables/index.js");
require("../../utils/http.js");
require("../../stores/index.js");
require("../../stores/modules/member.js");
if (!Math) {
  (CustomNavbar + Swiper + CategoryPanel + HotPanel + Guess)();
}
const CustomNavbar = () => "./components/CustomNavbar.js";
const CategoryPanel = () => "./components/CategoryPanel.js";
const HotPanel = () => "./components/HotPanel.js";
const Swiper = () => "../../components/Swiper.js";
const Guess = () => "../../components/Guess.js";
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  setup(__props) {
    const bannerList = common_vendor.ref([]);
    const getHomeBannerData = async () => {
      const res = await services_home.getHomeBannerAPI();
      bannerList.value = res.data;
    };
    const categoryList = common_vendor.ref([]);
    const getHomeCategoryData = async () => {
      const res = await services_home.getHomeCategoryAPI();
      categoryList.value = res.data;
    };
    const hotList = common_vendor.ref([]);
    const getHomeHotData = async () => {
      const res = await services_home.getHomeHotAPI();
      hotList.value = res.data;
    };
    common_vendor.onLoad(async () => {
      await Promise.all([getHomeBannerData(), getHomeCategoryData(), getHomeHotData()]);
    });
    const { guessRef, onScrolltolower } = composables_index.useGuessList();
    const isTriggered = common_vendor.ref(false);
    const onRefresherrefresh = async () => {
      var _a, _b;
      isTriggered.value = true;
      (_a = guessRef.value) == null ? void 0 : _a.resetData();
      await Promise.all([
        getHomeBannerData(),
        getHomeCategoryData(),
        getHomeHotData(),
        (_b = guessRef.value) == null ? void 0 : _b.getMore()
      ]);
      isTriggered.value = false;
    };
    return (_ctx, _cache) => {
      return {
        a: common_vendor.p({
          list: bannerList.value
        }),
        b: common_vendor.p({
          list: categoryList.value
        }),
        c: common_vendor.p({
          list: hotList.value
        }),
        d: common_vendor.sr(guessRef, "609f1c38-4", {
          "k": "guessRef"
        }),
        e: common_vendor.o(onRefresherrefresh),
        f: isTriggered.value,
        g: common_vendor.o(
          //@ts-ignore
          (...args) => common_vendor.unref(onScrolltolower) && common_vendor.unref(onScrolltolower)(...args)
        )
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/index/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
