"use strict";
const common_vendor = require("../../../common/vendor.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "CategoryPanel",
  props: {
    list: null
  },
  setup(__props) {
    return (_ctx, _cache) => {
      return {
        a: common_vendor.f(__props.list, (item, k0, i0) => {
          return {
            a: item.icon,
            b: common_vendor.t(item.name),
            c: `/pages/productList/index?id=${item.id}`,
            d: item.id
          };
        })
      };
    };
  }
});
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/index/components/CategoryPanel.vue"]]);
wx.createComponent(Component);
//# sourceMappingURL=CategoryPanel.js.map
