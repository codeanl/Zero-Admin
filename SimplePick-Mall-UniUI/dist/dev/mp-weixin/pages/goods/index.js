"use strict";
const common_vendor = require("../../common/vendor.js");
const services_home = require("../../services/home.js");
const services_cart = require("../../services/cart.js");
require("../../utils/http.js");
require("../../stores/index.js");
require("../../stores/modules/member.js");
if (!Array) {
  const _easycom_vk_data_goods_sku_popup2 = common_vendor.resolveComponent("vk-data-goods-sku-popup");
  const _easycom_uni_popup2 = common_vendor.resolveComponent("uni-popup");
  (_easycom_vk_data_goods_sku_popup2 + _easycom_uni_popup2)();
}
const _easycom_vk_data_goods_sku_popup = () => "../../components/vk-data-goods-sku-popup/vk-data-goods-sku-popup.js";
const _easycom_uni_popup = () => "../../node-modules/@dcloudio/uni-ui/lib/uni-popup/uni-popup.js";
if (!Math) {
  (_easycom_vk_data_goods_sku_popup + AddressPanel + ServicePanel + _easycom_uni_popup)();
}
const AddressPanel = () => "./components/AddressPanel.js";
const ServicePanel = () => "./components/ServicePanel.js";
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  props: {
    id: null
  },
  setup(__props) {
    const query = __props;
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const goods = common_vendor.ref();
    const getGoodsByIdData = async () => {
      let id = Number(query.id);
      const res = await services_home.getProductInfoAPI({ id });
      goods.value = res.data;
      localdata.value = {
        _id: res.data.productInfo.id,
        name: res.data.productInfo.name,
        goods_thumb: res.data.productInfo.pic,
        spec_list: res.data.sizeList.map((v) => {
          return {
            name: v.name,
            list: v.values
          };
        }),
        sku_list: res.data.skuList.map((v) => {
          return {
            _id: v.id,
            goods_id: v.productId,
            goods_name: v.name,
            image: v.pic,
            price: v.price * 100,
            // 注意：需要乘以 100
            stock: v.stock,
            sku_name_arr: v.specs.map((vv) => vv.valueName)
          };
        })
      };
    };
    common_vendor.onLoad(() => {
      getGoodsByIdData();
    });
    const currentIndex = common_vendor.ref(0);
    const onChange = (ev) => {
      currentIndex.value = ev.detail.current;
    };
    const onTapImage = (url) => {
      common_vendor.index.previewImage({
        current: url,
        urls: goods.value.imgUrl
      });
    };
    const popup = common_vendor.ref();
    const popupName = common_vendor.ref();
    const isShowSku = common_vendor.ref(false);
    const localdata = common_vendor.ref({});
    const mode = common_vendor.ref(
      2
      /* Cart */
    );
    const openSkuPopup = (val) => {
      isShowSku.value = true;
      mode.value = val;
    };
    const skuPopupRef = common_vendor.ref();
    common_vendor.computed(() => {
      var _a, _b;
      return ((_b = (_a = skuPopupRef.value) == null ? void 0 : _a.selectArr) == null ? void 0 : _b.join(" ").trim()) || "请选择商品规格";
    });
    const onAddCart = async (ev) => {
      let res = await services_cart.addCart({ userID: 103, skuID: ev._id, count: ev.buy_num });
      if (res.code === 200) {
        common_vendor.index.showToast({ title: "添加成功" });
        isShowSku.value = false;
      } else {
        common_vendor.index.showToast({ title: "添加失败" });
      }
    };
    const onBuyNow = (ev) => {
      let sku = [];
      sku.push({
        id: ev._id,
        count: ev.buy_num,
        skuID: ev._id,
        name: goods.value.productInfo.name,
        pic: goods.value.productInfo.pic,
        price: ev.price / 100,
        tag: ev.sku_name_arr,
        productId: goods.value.productInfo.id
      });
      common_vendor.index.navigateTo({
        url: "/pages/orderCreate/index?sku=" + encodeURIComponent(JSON.stringify(sku))
      });
    };
    return (_ctx, _cache) => {
      var _a, _b, _c, _d, _e, _f, _g, _h, _i, _j, _k;
      return common_vendor.e({
        a: common_vendor.sr(skuPopupRef, "4d4f8b68-0", {
          "k": "skuPopupRef"
        }),
        b: common_vendor.o(onAddCart),
        c: common_vendor.o(onBuyNow),
        d: common_vendor.o(($event) => isShowSku.value = $event),
        e: common_vendor.p({
          localdata: localdata.value,
          mode: mode.value,
          ["add-cart-background-color"]: "#FFA868",
          ["buy-now-background-color"]: "#27BA9B",
          ["actived-style"]: {
            color: "#27BA9B",
            borderColor: "#27BA9B",
            backgroundColor: "#E9F8F5"
          },
          modelValue: isShowSku.value
        }),
        f: common_vendor.f((_a = goods.value) == null ? void 0 : _a.imgUrl, (item, k0, i0) => {
          return {
            a: common_vendor.o(($event) => onTapImage(item), item),
            b: item,
            c: item
          };
        }),
        g: common_vendor.o(onChange),
        h: common_vendor.t(currentIndex.value + 1),
        i: common_vendor.t((_b = goods.value) == null ? void 0 : _b.imgUrl.length),
        j: common_vendor.t((_c = goods.value) == null ? void 0 : _c.productInfo.price),
        k: common_vendor.t((_d = goods.value) == null ? void 0 : _d.productInfo.sale),
        l: common_vendor.t((_e = goods.value) == null ? void 0 : _e.productInfo.name),
        m: common_vendor.t((_f = goods.value) == null ? void 0 : _f.productInfo.desc),
        n: (_g = goods.value) == null ? void 0 : _g.merchantInfo.pic,
        o: common_vendor.t((_h = goods.value) == null ? void 0 : _h.merchantInfo.name),
        p: common_vendor.f((_i = goods.value) == null ? void 0 : _i.attributeList, (item, k0, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: common_vendor.t(item.value),
            c: item.name
          };
        }),
        q: common_vendor.f((_j = goods.value) == null ? void 0 : _j.introduceImgUrl, (item, k0, i0) => {
          return {
            a: item,
            b: item
          };
        }),
        r: goods.value
      }, goods.value ? {
        s: common_vendor.o(($event) => openSkuPopup(
          2
          /* Cart */
        )),
        t: common_vendor.o(($event) => openSkuPopup(
          3
          /* Buy */
        )),
        v: ((_k = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _k.bottom) + "px"
      } : {}, {
        w: popupName.value === "address"
      }, popupName.value === "address" ? {
        x: common_vendor.o(($event) => {
          var _a2;
          return (_a2 = popup.value) == null ? void 0 : _a2.close();
        })
      } : {}, {
        y: popupName.value === "service"
      }, popupName.value === "service" ? {
        z: common_vendor.o(($event) => {
          var _a2;
          return (_a2 = popup.value) == null ? void 0 : _a2.close();
        })
      } : {}, {
        A: common_vendor.sr(popup, "4d4f8b68-1", {
          "k": "popup"
        }),
        B: common_vendor.p({
          type: "bottom",
          ["background-color"]: "#fff"
        })
      });
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/goods/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
