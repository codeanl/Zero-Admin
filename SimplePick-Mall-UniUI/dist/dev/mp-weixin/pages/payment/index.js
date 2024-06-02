"use strict";
const common_vendor = require("../../common/vendor.js");
const composables_index = require("../../composables/index.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  props: {
    id: null
  },
  setup(__props) {
    const { guessRef, onScrolltolower } = composables_index.useGuessList();
    return (_ctx, _cache) => {
      return {
        a: common_vendor.o(
          //@ts-ignore
          (...args) => common_vendor.unref(onScrolltolower) && common_vendor.unref(onScrolltolower)(...args)
        )
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/payment/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
