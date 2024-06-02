"use strict";
const common_vendor = require("../../common/vendor.js");
const services_category = require("../../services/category.js");
const services_home = require("../../services/home.js");
require("../../utils/http.js");
require("../../stores/index.js");
require("../../stores/modules/member.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "category",
  setup(__props) {
    const bannerList = common_vendor.ref([]);
    const getBannerData = async () => {
      const res = await services_home.getHomeBannerAPI();
      bannerList.value = res.data;
    };
    const categoryList = common_vendor.ref([]);
    const activeIndex = common_vendor.ref(0);
    const getCategoryTopData = async () => {
      const res = await services_category.getCategoryTopAPI();
      categoryList.value = res.data;
    };
    const isFinish = common_vendor.ref(false);
    common_vendor.onLoad(async () => {
      await Promise.all([getBannerData(), getCategoryTopData()]);
      isFinish.value = true;
    });
    const subCategoryList = common_vendor.computed(() => {
      var _a;
      return ((_a = categoryList.value[activeIndex.value]) == null ? void 0 : _a.children) || [];
    });
    return (_ctx, _cache) => {
      return {
        a: common_vendor.f(categoryList.value, (item, index, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: item.id,
            c: index === activeIndex.value ? 1 : "",
            d: common_vendor.o(($event) => activeIndex.value = index, item.id)
          };
        }),
        b: common_vendor.f(common_vendor.unref(subCategoryList), (item, k0, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: `/pages/productList/index?id=${item.id}`,
            c: common_vendor.f(item.productList, (goods, k1, i1) => {
              return {
                a: goods.pic,
                b: common_vendor.t(goods.name),
                c: common_vendor.t(goods.price),
                d: goods.id,
                e: `/pages/goods/index?id=${goods.id}`
              };
            }),
            d: item.id
          };
        })
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/category/category.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=category.js.map
