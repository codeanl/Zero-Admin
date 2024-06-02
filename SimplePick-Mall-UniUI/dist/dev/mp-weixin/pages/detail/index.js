"use strict";
const common_vendor = require("../../common/vendor.js");
const composables_index = require("../../composables/index.js");
const services_order = require("../../services/order.js");
const services_return = require("../../services/return.js");
require("../../stores/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
require("../../utils/http.js");
if (!Array) {
  const _component_PageSkeleton = common_vendor.resolveComponent("PageSkeleton");
  const _easycom_uni_popup2 = common_vendor.resolveComponent("uni-popup");
  (_component_PageSkeleton + _easycom_uni_popup2)();
}
const _easycom_uni_popup = () => "../../node-modules/@dcloudio/uni-ui/lib/uni-popup/uni-popup.js";
if (!Math) {
  _easycom_uni_popup();
}
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  props: {
    id: null
  },
  setup(__props) {
    const query = __props;
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const { guessRef, onScrolltolower } = composables_index.useGuessList();
    const popup = common_vendor.ref();
    const popup1 = common_vendor.ref();
    let reason = common_vendor.ref("");
    const onCopy = (id) => {
      common_vendor.index.setClipboardData({ data: id });
    };
    common_vendor.onLoad(() => {
      getData();
    });
    const order = common_vendor.ref();
    const reasonList = common_vendor.ref();
    const returnApply = common_vendor.ref();
    let idd = parseInt(query.id);
    const getData = async () => {
      const res = await services_order.orderInfo({ id: idd });
      if (res.code === 200) {
        order.value = res.data;
      }
      const res1 = await services_return.listreturnReason({});
      if (res1.code === 200) {
        reasonList.value = res1.data;
      }
      const res2 = await services_return.returnApplyInfo({ orderID: parseInt(query.id) });
      if (res2.code === 200) {
        returnApply.value = res2.data;
      }
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
    const onOrderPay = async () => {
      const currentDateTime = getCurrentDateTime();
      let res = await services_order.updateOrder({
        id: idd,
        status: "1",
        paymentTime: currentDateTime
      });
      if (res.code == 200) {
        common_vendor.index.redirectTo({ url: `/pages/payment/index?id=${query.id}` });
      }
    };
    const onCancel = async () => {
      let res = await services_order.updateOrder({
        id: idd,
        status: "5"
      });
      if (res.code == 200) {
        common_vendor.index.redirectTo({ url: `/pages/detail/index?id=${query.id}` });
      }
    };
    const onConfirm = async () => {
      const currentDateTime = getCurrentDateTime();
      common_vendor.index.showModal({
        content: "为保障您的权益，请收到货并确认无误后，再确认收货",
        confirmColor: "#27BA9B",
        success: async (success) => {
          if (success.confirm) {
            let res = await services_order.updateOrder({
              id: idd,
              status: "4",
              receiveTime: currentDateTime
            });
            if (res.code == 200) {
              common_vendor.index.redirectTo({ url: `/pages/detail/index?id=${query.id}` });
            }
          }
        }
      });
    };
    stores_modules_member.useMemberStore();
    let onOrderCancel = async () => {
      console.log(returnReason.value.id);
    };
    let returnReason = common_vendor.ref();
    let ddd = (item) => {
      reason.value = item;
      returnReason.value = item;
    };
    return (_ctx, _cache) => {
      var _a, _b, _c, _d, _e, _f, _g, _h, _i, _j, _k, _l;
      return common_vendor.e({}, {
        a: ((_a = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _a.top) + "px"
      }, common_vendor.e({
        b: order.value.orderInfo.status == "0"
      }, order.value.orderInfo.status == "0" ? {
        c: common_vendor.t(order.value.orderInfo.payAmount)
      } : {}, {
        d: order.value.orderInfo.status == "1"
      }, order.value.orderInfo.status == "1" ? {} : {}, {
        e: order.value.orderInfo.status == "2"
      }, order.value.orderInfo.status == "2" ? {} : {}, {
        f: order.value.orderInfo.status == "3"
      }, order.value.orderInfo.status == "3" ? {} : {}, {
        g: order.value.orderInfo.status == "4"
      }, order.value.orderInfo.status == "4" ? {} : {}, {
        h: order.value.orderInfo.status == "5"
      }, order.value.orderInfo.status == "5" ? {} : {}, {
        i: common_vendor.unref(safeAreaInsets).top + 20 + "px",
        j: common_vendor.t((_b = order.value) == null ? void 0 : _b.merchantInfo.name),
        k: common_vendor.f((_c = order.value) == null ? void 0 : _c.skuList, (item, k0, i0) => {
          return {
            a: item.pic,
            b: common_vendor.t(item.name),
            c: common_vendor.t(item.tag),
            d: common_vendor.t(item.price),
            e: common_vendor.t(item.count),
            f: item.id,
            g: `/pages/goods/index?id=${item.productId}`
          };
        }),
        l: common_vendor.t((_d = order.value) == null ? void 0 : _d.orderInfo.payAmount),
        m: common_vendor.t((_e = order.value) == null ? void 0 : _e.orderInfo.payAmount),
        n: common_vendor.t((_f = order.value) == null ? void 0 : _f.orderInfo.orderSn),
        o: common_vendor.o(($event) => {
          var _a2;
          return onCopy((_a2 = order.value) == null ? void 0 : _a2.orderInfo.orderSn);
        }),
        p: common_vendor.t((_g = order.value) == null ? void 0 : _g.orderInfo.createTime),
        q: common_vendor.t((_h = order.value) == null ? void 0 : _h.placeInfo.name),
        r: common_vendor.t((_i = order.value) == null ? void 0 : _i.placeInfo.place),
        s: common_vendor.t((_j = order.value) == null ? void 0 : _j.placeInfo.phone),
        t: ((_k = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _k.bottom) + "px",
        v: order.value.orderInfo.status == "0"
      }, order.value.orderInfo.status == "0" ? {
        w: common_vendor.o(onOrderPay),
        x: common_vendor.o(onCancel)
      } : {}, {
        y: order.value.orderInfo.status == "1"
      }, order.value.orderInfo.status == "1" ? {
        z: `/pagesOrder/create/create?orderId=${query.id}`
      } : {}, {
        A: order.value.orderInfo.status == "2"
      }, order.value.orderInfo.status == "2" ? {
        B: `/pagesOrder/create/create?orderId=${query.id}`
      } : {}, {
        C: order.value.orderInfo.status == "3"
      }, order.value.orderInfo.status == "3" ? {
        D: common_vendor.o(onConfirm)
      } : {}, {
        E: order.value.orderInfo.status == "4"
      }, order.value.orderInfo.status == "4" ? common_vendor.e({
        F: !returnApply.value
      }, !returnApply.value ? common_vendor.e({
        G: !returnApply.value
      }, !returnApply.value ? {
        H: `/pages/returnApply/index?id=${query.id}`
      } : {}) : common_vendor.e({
        I: returnApply.value.status == "0"
      }, returnApply.value.status == "0" ? {} : {}, {
        J: returnApply.value.status == "1"
      }, returnApply.value.status == "1" ? {} : {}, {
        K: returnApply.value.status == "2"
      }, returnApply.value.status == "2" ? {} : {}, {
        L: returnApply.value.status == "3"
      }, returnApply.value.status == "3" ? {} : {})) : {}, {
        M: ((_l = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _l.bottom) + "px"
      }), {
        N: common_vendor.o(
          //@ts-ignore
          (...args) => common_vendor.unref(onScrolltolower) && common_vendor.unref(onScrolltolower)(...args)
        ),
        O: common_vendor.f(reasonList.value, (item, k0, i0) => {
          return {
            a: common_vendor.t(item),
            b: item === common_vendor.unref(reason) ? 1 : "",
            c: item,
            d: common_vendor.o(($event) => common_vendor.isRef(reason) ? reason.value = item : reason = item, item)
          };
        }),
        P: common_vendor.o(($event) => {
          var _a2, _b2;
          return (_b2 = (_a2 = popup.value) == null ? void 0 : _a2.close) == null ? void 0 : _b2.call(_a2);
        }),
        Q: common_vendor.sr(popup, "ffb446d6-1", {
          "k": "popup"
        }),
        R: common_vendor.p({
          type: "bottom",
          ["background-color"]: "#fff"
        }),
        S: common_vendor.f(reasonList.value, (item, k0, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: item === common_vendor.unref(reason) ? 1 : "",
            c: item,
            d: common_vendor.o(($event) => common_vendor.unref(ddd)(item), item)
          };
        }),
        T: common_vendor.o(($event) => {
          var _a2, _b2;
          return (_b2 = (_a2 = popup1.value) == null ? void 0 : _a2.close) == null ? void 0 : _b2.call(_a2);
        }),
        U: common_vendor.o(($event) => {
          var _a2, _b2;
          return (_b2 = (_a2 = popup1.value) == null ? void 0 : _a2.close) == null ? void 0 : _b2.call(_a2), common_vendor.unref(onOrderCancel)();
        }),
        V: common_vendor.sr(popup1, "ffb446d6-2", {
          "k": "popup1"
        }),
        W: common_vendor.p({
          type: "bottom",
          ["background-color"]: "#fff"
        })
      });
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/detail/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
