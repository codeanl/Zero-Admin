"use strict";
const common_vendor = require("../../common/vendor.js");
const services_home = require("../../services/home.js");
require("../../utils/http.js");
require("../../stores/index.js");
require("../../stores/modules/member.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  props: {
    id: null,
    name: null
  },
  setup(__props) {
    const query = __props;
    const guessList = common_vendor.ref([]);
    const getHomeGoodsGuessLikeData = async () => {
      if (query.id) {
        const res = await services_home.getHomeGoodsGuessLikeAPI({ categoryId: query.id, searchType: searchType.value });
        guessList.value = res.data;
      }
      if (query.name) {
        const res = await services_home.getHomeGoodsGuessLikeAPI({ name: query.name, searchType: searchType.value });
        guessList.value = res.data;
      }
    };
    common_vendor.onMounted(() => {
      getHomeGoodsGuessLikeData();
    });
    let selectedOption = common_vendor.ref(0);
    let searchType = common_vendor.ref(0);
    let tabShow = (showid, searchTypeID) => {
      console.log("展示的是" + showid + ",type是" + searchTypeID);
      selectedOption.value = showid;
      searchType.value = searchTypeID;
      guessList.value = [];
      getHomeGoodsGuessLikeData();
    };
    return (_ctx, _cache) => {
      return {
        a: common_vendor.o(($event) => common_vendor.unref(tabShow)(0, 0)),
        b: common_vendor.unref(selectedOption) === 0 ? 1 : "",
        c: common_vendor.o(($event) => common_vendor.unref(tabShow)(1, 1)),
        d: common_vendor.unref(selectedOption) === 1 ? 1 : "",
        e: common_vendor.o(($event) => common_vendor.unref(tabShow)(2, 2)),
        f: common_vendor.unref(selectedOption) === 2 ? 1 : "",
        g: common_vendor.o(($event) => common_vendor.unref(tabShow)(4, 4)),
        h: common_vendor.unref(selectedOption) === 4 ? 1 : "",
        i: common_vendor.o(($event) => common_vendor.unref(tabShow)(5, 5)),
        j: common_vendor.unref(selectedOption) === 5 ? 1 : "",
        k: common_vendor.f(guessList.value, (item, k0, i0) => {
          return {
            a: item.pic,
            b: common_vendor.t(item.name),
            c: common_vendor.t(item.price),
            d: item.id,
            e: `/pages/goods/index?id=${item.id}`
          };
        })
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/productList/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
