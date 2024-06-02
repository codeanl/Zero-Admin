"use strict";
const common_vendor = require("../common/vendor.js");
const services_home = require("../services/home.js");
require("../utils/http.js");
require("../stores/index.js");
require("../stores/modules/member.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "Guess",
  setup(__props, { expose }) {
    const pageParams = {
      current: 1,
      pageSize: 10
    };
    const guessList = common_vendor.ref([]);
    const finish = common_vendor.ref(false);
    const getHomeGoodsGuessLikeData = async () => {
      if (finish.value === true) {
        return common_vendor.index.showToast({ icon: "none", title: "没有更多数据~" });
      }
      const res = await services_home.getHomeGoodsGuessLikeAPI(pageParams);
      guessList.value.push(...res.data);
      if (pageParams.current < Math.ceil(res.total / pageParams.pageSize)) {
        pageParams.current++;
      } else {
        finish.value = true;
      }
    };
    const resetData = () => {
      pageParams.current = 1;
      guessList.value = [];
      finish.value = false;
    };
    common_vendor.onMounted(() => {
      getHomeGoodsGuessLikeData();
    });
    expose({
      resetData,
      getMore: getHomeGoodsGuessLikeData
    });
    return (_ctx, _cache) => {
      return {
        a: common_vendor.f(guessList.value, (item, k0, i0) => {
          return {
            a: item.pic,
            b: common_vendor.t(item.name),
            c: common_vendor.t(item.price),
            d: item.id,
            e: `/pages/goods/index?id=${item.id}`
          };
        }),
        b: common_vendor.t(finish.value ? "没有更多数据~" : "正在加载...")
      };
    };
  }
});
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/components/Guess.vue"]]);
wx.createComponent(Component);
//# sourceMappingURL=Guess.js.map
