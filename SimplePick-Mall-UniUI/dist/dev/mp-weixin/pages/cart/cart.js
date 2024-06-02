"use strict";
const common_vendor = require("../../common/vendor.js");
const services_cart = require("../../services/cart.js");
require("../../stores/index.js");
const composables_index = require("../../composables/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
require("../../utils/http.js");
if (!Array) {
  const _easycom_vk_data_input_number_box2 = common_vendor.resolveComponent("vk-data-input-number-box");
  const _easycom_uni_swipe_action_item2 = common_vendor.resolveComponent("uni-swipe-action-item");
  const _easycom_uni_swipe_action2 = common_vendor.resolveComponent("uni-swipe-action");
  (_easycom_vk_data_input_number_box2 + _easycom_uni_swipe_action_item2 + _easycom_uni_swipe_action2)();
}
const _easycom_vk_data_input_number_box = () => "../../components/vk-data-input-number-box/vk-data-input-number-box.js";
const _easycom_uni_swipe_action_item = () => "../../node-modules/@dcloudio/uni-ui/lib/uni-swipe-action-item/uni-swipe-action-item.js";
const _easycom_uni_swipe_action = () => "../../node-modules/@dcloudio/uni-ui/lib/uni-swipe-action/uni-swipe-action.js";
if (!Math) {
  (_easycom_vk_data_input_number_box + _easycom_uni_swipe_action_item + _easycom_uni_swipe_action + Guess)();
}
const Guess = () => "../../components/Guess.js";
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "cart",
  props: {
    safeAreaInsetBottom: { type: Boolean }
  },
  setup(__props) {
    const memberStore = stores_modules_member.useMemberStore();
    const cartList = common_vendor.ref();
    const showCartList = common_vendor.ref(true);
    const getMemberCartData = async () => {
      const res = await services_cart.listCart({ userID: 103 });
      if (res.code === 200) {
        cartList.value = res.data;
        showCartList.value = res.data.length > 0;
      }
    };
    common_vendor.onShow(() => {
      if (memberStore.profile) {
        getMemberCartData();
        console.log(cartList.value);
      }
    });
    const { guessRef, onScrolltolower } = composables_index.useGuessList();
    const onDeleteCart = (id) => {
      console.log(id);
      common_vendor.index.showModal({
        content: "是否删除",
        confirmColor: "#27BA9B",
        success: async (res) => {
          if (res.confirm) {
            await services_cart.deleteCart({ ids: [id] });
            getMemberCartData();
          }
        }
      });
    };
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const onChangeCount = (ev) => {
      putMemberCartBySkuIdAPI(ev.index, { count: ev.value });
    };
    const onChangeSelected = (item) => {
      item.selected = !item.selected;
      putMemberCartBySkuIdAPI(item.skuId, { selected: item.selected });
    };
    const isSelectedAll = common_vendor.computed(() => {
      return cartList.value.length && cartList.value.every((v) => v.selected);
    });
    const onChangeSelectedAll = () => {
      const _isSelectedAll = !isSelectedAll.value;
      cartList.value.forEach((item) => {
        item.selected = _isSelectedAll;
      });
    };
    const selectedCartList = common_vendor.computed(() => {
      return cartList.value.filter((v) => v.selected);
    });
    const selectedCartListCount = common_vendor.computed(() => {
      return selectedCartList.value.reduce((sum, item) => sum + item.count, 0);
    });
    const selectedCartListMoney = common_vendor.computed(() => {
      return selectedCartList.value.reduce((sum, item) => sum + item.count * item.listSkuData.price, 0).toFixed(2);
    });
    const gotoPayment = () => {
      let sku = [];
      if (selectedCartListCount.value === 0) {
        return common_vendor.index.showToast({
          icon: "none",
          title: "请选择商品"
        });
      }
      let dd = cartList.value.filter((v) => v.selected);
      for (let index = 0; index < dd.length; index++) {
        sku.push(
          {
            id: dd[index].id,
            count: dd[index].count,
            skuID: dd[index].skuID,
            name: dd[index].listSkuData.name,
            pic: dd[index].listSkuData.pic.replace(/"/g, ""),
            price: dd[index].listSkuData.price,
            tag: dd[index].listSkuData.tag,
            productId: dd[index].listSkuData.productId
          }
        );
      }
      common_vendor.index.navigateTo({
        url: "/pages/orderCreate/index?sku=" + encodeURIComponent(JSON.stringify(sku))
      });
    };
    return (_ctx, _cache) => {
      var _a;
      return common_vendor.e({
        a: common_vendor.unref(memberStore).profile
      }, common_vendor.unref(memberStore).profile ? common_vendor.e({
        b: cartList.value
      }, cartList.value ? {
        c: common_vendor.f(cartList.value, (item, k0, i0) => {
          return {
            a: common_vendor.o(($event) => onChangeSelected(item), item.skuId),
            b: item.selected ? 1 : "",
            c: common_vendor.t(item == null ? void 0 : item.merchantData.name),
            d: item == null ? void 0 : item.listSkuData.pic,
            e: common_vendor.t(item == null ? void 0 : item.listSkuData.name),
            f: common_vendor.t(item == null ? void 0 : item.listSkuData.tag),
            g: common_vendor.t(item == null ? void 0 : item.listSkuData.price),
            h: `/pages/goods/index?id=${item.listSkuData.productId}`,
            i: common_vendor.o(onChangeCount, item.skuId),
            j: "6db479fe-2-" + i0 + "," + ("6db479fe-1-" + i0),
            k: common_vendor.o(($event) => item.count = $event, item.skuId),
            l: common_vendor.p({
              min: 1,
              max: item.stock,
              index: item.skuId,
              modelValue: item.count
            }),
            m: common_vendor.o(($event) => onDeleteCart(item.id), item.skuId),
            n: item.skuId,
            o: "6db479fe-1-" + i0 + ",6db479fe-0"
          };
        })
      } : {}, {
        d: cartList.value
      }, cartList.value ? {
        e: common_vendor.o(onChangeSelectedAll),
        f: common_vendor.unref(isSelectedAll) ? 1 : "",
        g: common_vendor.t(common_vendor.unref(selectedCartListMoney)),
        h: common_vendor.t(common_vendor.unref(selectedCartListCount)),
        i: common_vendor.o(gotoPayment),
        j: common_vendor.unref(selectedCartListCount) === 0 ? 1 : "",
        k: __props.safeAreaInsetBottom ? ((_a = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _a.bottom) + "px" : 0
      } : {}) : {}, {
        l: common_vendor.sr(guessRef, "6db479fe-3", {
          "k": "guessRef"
        }),
        m: common_vendor.o(
          //@ts-ignore
          (...args) => common_vendor.unref(onScrolltolower) && common_vendor.unref(onScrolltolower)(...args)
        )
      });
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/cart/cart.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=cart.js.map
