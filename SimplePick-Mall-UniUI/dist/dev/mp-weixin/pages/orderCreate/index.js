"use strict";
const common_vendor = require("../../common/vendor.js");
require("../../utils/http.js");
const services_cart = require("../../services/cart.js");
const services_order = require("../../services/order.js");
const services_place = require("../../services/place.js");
require("../../stores/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
if (!Array) {
  const _easycom_uni_popup2 = common_vendor.resolveComponent("uni-popup");
  const _easycom_uni_popup_dialog2 = common_vendor.resolveComponent("uni-popup-dialog");
  (_easycom_uni_popup2 + _easycom_uni_popup_dialog2)();
}
const _easycom_uni_popup = () => "../../node-modules/@dcloudio/uni-ui/lib/uni-popup/uni-popup.js";
const _easycom_uni_popup_dialog = () => "../../node-modules/@dcloudio/uni-ui/lib/uni-popup-dialog/uni-popup-dialog.js";
if (!Math) {
  (_easycom_uni_popup + _easycom_uni_popup_dialog)();
}
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  props: {
    sku: null
  },
  setup(__props) {
    const query = __props;
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const buyerMessage = common_vendor.ref("");
    const userName = common_vendor.ref("");
    const userPhone = common_vendor.ref("");
    const deliveryList = common_vendor.ref([
      { type: 1, text: "时间不限 (周一至周日)" },
      { type: 2, text: "工作日送 (周一至周五)" },
      { type: 3, text: "周末配送 (周六至周日)" }
    ]);
    const activeIndex = common_vendor.ref(0);
    common_vendor.computed(() => deliveryList.value[activeIndex.value]);
    let skus = common_vendor.ref();
    let addPri = 0;
    common_vendor.onLoad(() => {
      skus.value = JSON.parse(query.sku);
      for (let index = 0; index < skus.value.length; index++) {
        addPri += skus.value[index].price * skus.value[index].count;
      }
      getMemberCartData();
    });
    const memberStore = stores_modules_member.useMemberStore();
    const PlaceList = common_vendor.ref();
    const getMemberCartData = async () => {
      const res = await services_place.listplace({});
      if (res.code === 200) {
        PlaceList.value = res.data;
      }
    };
    let generateOrderNumber = () => {
      const prefix = "ORD";
      const randomLength = 4;
      const timestamp = Date.now();
      const random = Math.floor(Math.random() * Math.pow(10, randomLength)).toString();
      const orderNumber = prefix + timestamp.toString() + random;
      return orderNumber;
    };
    let getCurrentDateTime = () => {
      const now = new Date();
      const year = now.getFullYear();
      const month = String(now.getMonth() + 1).padStart(2, "0");
      const day = String(now.getDate()).padStart(2, "0");
      const hours = String(now.getHours()).padStart(2, "0");
      const minutes = String(now.getMinutes()).padStart(2, "0");
      const seconds = String(now.getSeconds()).padStart(2, "0");
      const dateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
      return dateTime;
    };
    const popup = common_vendor.ref();
    const alertDialog = common_vendor.ref();
    common_vendor.ref([
      "商品无货",
      "不想要了",
      "商品信息填错了",
      "地址信息填写错误",
      "商品降价",
      "其它"
    ]);
    let reason = common_vendor.ref("");
    let isture = common_vendor.ref(false);
    let place = common_vendor.ref();
    let ddd = (item) => {
      reason.value = item;
      place.value = item;
      isture.value = true;
    };
    let check = () => {
      if (userName.value === "" || userPhone.value === "") {
        common_vendor.index.showToast({
          title: "请填写完整信息",
          icon: "none"
        });
        return;
      }
      const phoneRegex = /^1[0-9]{10}$/;
      if (!phoneRegex.test(userPhone.value)) {
        common_vendor.index.showToast({
          title: "请填写有效手机号",
          icon: "none"
        });
        return;
      }
      if (!place.value) {
        common_vendor.index.showToast({
          title: "请选择自提点",
          icon: "none"
        });
        return;
      }
    };
    const currentDateTime = getCurrentDateTime();
    let toDoPay = async () => {
      check();
      const orderNumber = generateOrderNumber();
      let skuIDs = [];
      let ids = common_vendor.ref([]);
      for (let index = 0; index < (skus == null ? void 0 : skus.value.length); index++) {
        ids.value.push(parseInt(skus.value[index].id));
        skuIDs.push(
          {
            skuID: skus.value[index].skuID,
            count: skus.value[index].count
          }
        );
      }
      const res = await services_order.addOrder({
        memberId: memberStore == null ? void 0 : memberStore.profile.info.id,
        placeId: place.value.id,
        //
        couponId: 0,
        //
        orderSn: orderNumber,
        memberUserName: memberStore.profile.info.nickname,
        totalAmount: addPri,
        payAmount: addPri,
        freightAmount: 0,
        couponAmount: 0,
        payType: "1",
        status: "0",
        orderType: "1",
        receiverName: userName.value,
        receiverPhone: userPhone.value,
        note: buyerMessage.value,
        skus: skuIDs
      });
      if (res.code == 200) {
        await services_cart.deleteCart({ ids: ids.value });
        common_vendor.index.redirectTo({ url: `/pages/orderList/index?status=0` });
      }
    };
    let toPay = async () => {
      check();
      const orderNumber = generateOrderNumber();
      let skuIDs = [];
      let ids = common_vendor.ref([]);
      for (let index = 0; index < (skus == null ? void 0 : skus.value.length); index++) {
        ids.value.push(parseInt(skus.value[index].id));
        skuIDs.push(
          {
            skuID: skus.value[index].skuID,
            count: skus.value[index].count
          }
        );
      }
      const res = await services_order.addOrder({
        memberId: memberStore == null ? void 0 : memberStore.profile.info.id,
        placeId: place.value.id,
        //
        couponId: 0,
        //
        orderSn: orderNumber,
        memberUserName: memberStore.profile.info.nickname,
        totalAmount: addPri,
        payAmount: addPri,
        freightAmount: 0,
        couponAmount: 0,
        payType: "1",
        status: "0",
        orderType: "1",
        receiverName: userName.value,
        receiverPhone: userPhone.value,
        note: buyerMessage.value,
        skus: skuIDs
      });
      if (res.code == 200) {
        await services_cart.deleteCart({ ids: ids.value });
        for (const i in res.data.orderID) {
          await services_order.updateOrder({
            id: res.data.orderID[i],
            status: "1",
            paymentTime: currentDateTime
          });
        }
        common_vendor.index.redirectTo({ url: `/pages/payment/index?status=1` });
      }
    };
    return (_ctx, _cache) => {
      var _a;
      return {
        a: common_vendor.f(common_vendor.unref(skus), (item, index, i0) => {
          return {
            a: item.pic,
            b: common_vendor.t(item.name),
            c: common_vendor.t(item.tag),
            d: common_vendor.t(item.price),
            e: common_vendor.t(item.price),
            f: common_vendor.t(item.count),
            g: index,
            h: `/pages/goods/goods?id=${item.id}`
          };
        }),
        b: buyerMessage.value,
        c: common_vendor.o(($event) => buyerMessage.value = $event.detail.value),
        d: userName.value,
        e: common_vendor.o(($event) => userName.value = $event.detail.value),
        f: userPhone.value,
        g: common_vendor.o(($event) => userPhone.value = $event.detail.value),
        h: common_vendor.t(common_vendor.unref(place) ? common_vendor.unref(place).name : "请填写自提点"),
        i: common_vendor.o(($event) => {
          var _a2, _b;
          return (_b = (_a2 = popup.value) == null ? void 0 : _a2.open) == null ? void 0 : _b.call(_a2);
        }),
        j: common_vendor.t(common_vendor.unref(addPri)),
        k: common_vendor.t(common_vendor.unref(addPri)),
        l: common_vendor.o(($event) => {
          var _a2, _b;
          return (_b = (_a2 = alertDialog.value) == null ? void 0 : _a2.open) == null ? void 0 : _b.call(_a2);
        }),
        m: ((_a = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _a.bottom) + "px",
        n: common_vendor.f(PlaceList.value, (item, k0, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: item === common_vendor.unref(reason) ? 1 : "",
            c: item,
            d: common_vendor.o(($event) => common_vendor.unref(ddd)(item), item)
          };
        }),
        o: common_vendor.o(($event) => {
          var _a2, _b;
          return (_b = (_a2 = popup.value) == null ? void 0 : _a2.close) == null ? void 0 : _b.call(_a2);
        }),
        p: common_vendor.sr(popup, "d5ed9088-0", {
          "k": "popup"
        }),
        q: common_vendor.p({
          type: "bottom",
          ["background-color"]: "#fff"
        }),
        r: common_vendor.o(common_vendor.unref(toPay)),
        s: common_vendor.o(common_vendor.unref(toDoPay)),
        t: common_vendor.p({
          type: "success",
          cancelText: "稍后支付",
          confirmText: "现在支付",
          content: "请支付订单!"
        }),
        v: common_vendor.sr(alertDialog, "d5ed9088-1", {
          "k": "alertDialog"
        }),
        w: common_vendor.p({
          type: "dialog"
        })
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/orderCreate/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
