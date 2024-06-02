"use strict";
const common_vendor = require("../../common/vendor.js");
require("../../stores/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "my",
  setup(__props) {
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const orderTypes = [
      { type: "100", name: "全部", icon: "icon-currency" },
      { type: "0", name: "待付款", icon: "icon-currency" },
      { type: "1", name: "待发货", icon: "icon-gift" },
      { type: "2", name: "运输中", icon: "icon-check" },
      { type: "3", name: "待提货", icon: "icon-comment" }
    ];
    const memberStore = stores_modules_member.useMemberStore();
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.unref(memberStore).profile
      }, common_vendor.unref(memberStore).profile ? {
        b: common_vendor.unref(memberStore).profile.info.avatar,
        c: common_vendor.t(common_vendor.unref(memberStore).profile.info.nickname)
      } : {}, {
        d: common_vendor.unref(safeAreaInsets).top + "px",
        e: common_vendor.f(orderTypes, (item, k0, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: item.type,
            c: common_vendor.n(item.icon),
            d: `/pages/orderList/index?status=${item.type}`
          };
        })
      });
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/my/my.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=my.js.map
