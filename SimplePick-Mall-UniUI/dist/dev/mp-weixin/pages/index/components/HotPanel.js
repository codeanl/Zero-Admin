"use strict";
const common_vendor = require("../../../common/vendor.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "HotPanel",
  props: {
    list: null
  },
  setup(__props) {
    return (_ctx, _cache) => {
      return {
        a: common_vendor.f(__props.list, (item, k0, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: item.pic,
            c: `/pages/hot/index?id=${item.id}&&pic=${item.pic}&&name=${item.name} `,
            d: item.id
          };
        })
      };
    };
  }
});
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/index/components/HotPanel.vue"]]);
wx.createComponent(Component);
//# sourceMappingURL=HotPanel.js.map
