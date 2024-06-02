"use strict";
const common_vendor = require("../../common/vendor.js");
const services_hot = require("../../services/hot.js");
require("../../utils/http.js");
require("../../stores/index.js");
require("../../stores/modules/member.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  props: {
    id: null,
    pic: null,
    name: null
  },
  setup(__props) {
    const query = __props;
    common_vendor.index.setNavigationBarTitle({ title: query.name });
    const productList = common_vendor.ref([]);
    const getHomeHotData = async () => {
      const res = await services_hot.getHotProductApi({ subjectId: parseInt(query.id), status: "1" });
      productList.value = res.data;
    };
    common_vendor.onLoad(() => {
      getHomeHotData();
    });
    return (_ctx, _cache) => {
      return {
        a: query.pic,
        b: common_vendor.f(productList.value, (goods, k0, i0) => {
          return {
            a: goods.productInfo.pic,
            b: common_vendor.t(goods.productInfo.name),
            c: common_vendor.t(goods.productInfo.price),
            d: goods.id,
            e: `/pages/goods/index?id=${goods.productInfo.id}`
          };
        })
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/hot/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
