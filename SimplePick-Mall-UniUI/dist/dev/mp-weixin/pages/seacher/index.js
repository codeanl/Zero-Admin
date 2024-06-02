"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  setup(__props) {
    let inputValue = common_vendor.ref("");
    return (_ctx, _cache) => {
      return {
        a: common_vendor.unref(inputValue),
        b: common_vendor.o(($event) => common_vendor.isRef(inputValue) ? inputValue.value = $event.detail.value : inputValue = $event.detail.value),
        c: `/pages/productList/index?name=${common_vendor.unref(inputValue)}`
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-923f0508"], ["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/seacher/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
