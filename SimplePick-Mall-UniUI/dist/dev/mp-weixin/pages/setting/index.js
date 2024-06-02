"use strict";
const common_vendor = require("../../common/vendor.js");
require("../../stores/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  setup(__props) {
    const memberStore = stores_modules_member.useMemberStore();
    const onLogout = () => {
      common_vendor.index.showModal({
        content: "是否退出登录？",
        confirmColor: "#27BA9B",
        success: (res) => {
          if (res.confirm) {
            memberStore.clearProfile();
            common_vendor.index.navigateBack();
          }
        }
      });
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.unref(memberStore).profile.token
      }, common_vendor.unref(memberStore).profile.token ? {
        b: common_vendor.o(onLogout)
      } : {});
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/setting/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
