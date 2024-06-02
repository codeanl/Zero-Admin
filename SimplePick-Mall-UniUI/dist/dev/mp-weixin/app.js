"use strict";
Object.defineProperty(exports, Symbol.toStringTag, { value: "Module" });
const common_vendor = require("./common/vendor.js");
const stores_index = require("./stores/index.js");
require("./stores/modules/member.js");
if (!Math) {
  "./pages/index/index.js";
  "./pages/my/my.js";
  "./pages/cart/cart.js";
  "./pages/category/category.js";
  "./pages/login/index.js";
  "./pages/setting/index.js";
  "./pages/profile/index.js";
  "./pages/goods/index.js";
  "./pages/hot/index.js";
  "./pages/orderCreate/index.js";
  "./pages/detail/index.js";
  "./pages/payment/index.js";
  "./pages/orderList/index.js";
  "./pages/seacher/index.js";
  "./pages/productList/index.js";
  "./pages/returnApply/index.js";
}
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "App",
  setup(__props) {
    common_vendor.onLaunch(() => {
      console.log("App Launch");
    });
    common_vendor.onShow(() => {
      console.log("App Show");
    });
    common_vendor.onHide(() => {
      console.log("App Hide");
    });
    return () => {
    };
  }
});
const App = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/App.vue"]]);
function createApp() {
  const app = common_vendor.createSSRApp(App);
  app.use(stores_index.pinia);
  return {
    app
  };
}
createApp().app.mount("#app");
exports.createApp = createApp;
//# sourceMappingURL=app.js.map
