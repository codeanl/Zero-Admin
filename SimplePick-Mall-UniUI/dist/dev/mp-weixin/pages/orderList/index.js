"use strict";
const common_vendor = require("../../common/vendor.js");
require("../../stores/index.js");
const services_constants = require("../../services/constants.js");
const services_order = require("../../services/order.js");
const stores_modules_member = require("../../stores/modules/member.js");
require("../../utils/http.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  props: {
    status: null
  },
  setup(__props) {
    const query = __props;
    const memberStore = stores_modules_member.useMemberStore();
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const orderTabs = common_vendor.ref([
      { orderState: 100, title: "全部", isRender: false },
      { orderState: 0, title: "待付款", isRender: false },
      { orderState: 1, title: "待发货", isRender: false },
      { orderState: 2, title: "运输中", isRender: false },
      { orderState: 3, title: "待提货", isRender: false }
    ]);
    const activeIndex = common_vendor.ref(orderTabs.value.findIndex((v) => v.orderState === parseInt(query.status)));
    orderTabs.value[activeIndex.value].isRender = true;
    common_vendor.onLoad(() => {
      getData(query.status);
    });
    const order = common_vendor.ref();
    const getData = async (status) => {
      order.value = [];
      if (status < 100) {
        const res = await services_order.listOrder({ status, userID: memberStore.profile.info.id });
        if (res.code === 200) {
          order.value = res.data;
        }
      } else {
        const res = await services_order.listOrder({ userID: memberStore.profile.info.id });
        if (res.code === 200) {
          order.value = res.data;
        }
      }
    };
    let now = common_vendor.ref();
    let handleTabClick = (index, id) => {
      console.log(index, id);
      now.value = id;
      activeIndex.value = index;
      orderTabs.value.forEach((item, i) => {
        item.isRender = i === index;
      });
      getData(id.toString());
    };
    let getCurrentDateTime = () => {
      const now2 = new Date();
      const year = now2.getFullYear();
      const month = String(now2.getMonth() + 1).padStart(2, "0");
      const day = String(now2.getDate()).padStart(2, "0");
      const hours = String(now2.getHours()).padStart(2, "0");
      const minutes = String(now2.getMinutes()).padStart(2, "0");
      const seconds = String(now2.getSeconds()).padStart(2, "0");
      const dateTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
      return dateTime;
    };
    const onOrderPay = async (id) => {
      const currentDateTime = getCurrentDateTime();
      let res = await services_order.updateOrder({
        id,
        status: "1",
        paymentTime: currentDateTime
      });
      if (res.code == 200) {
        common_vendor.index.redirectTo({ url: `/pages/payment/index?id=${id}` });
      }
    };
    const onCancel = async (id) => {
      console.log(now.value);
      let res = await services_order.updateOrder({
        id,
        status: "5"
      });
      if (res.code == 200) {
        getData(now.value.toString());
        common_vendor.index.showToast({ icon: "success", title: "取消订单完成" });
      }
    };
    const onConfirm = async (id) => {
      const currentDateTime = getCurrentDateTime();
      common_vendor.index.showModal({
        content: "为保障您的权益，请收到货并确认无误后，再确认收货",
        confirmColor: "#27BA9B",
        success: async (success) => {
          if (success.confirm) {
            let res = await services_order.updateOrder({
              id,
              status: "4",
              receiveTime: currentDateTime
            });
            if (res.code == 200) {
              getData(now.value.toString());
              common_vendor.index.showToast({ icon: "success", title: "确认收货完成" });
            }
          }
        }
      });
    };
    const isTriggered = common_vendor.ref(false);
    const onRefresherrefresh = async () => {
      isTriggered.value = true;
    };
    return (_ctx, _cache) => {
      var _a;
      return {
        a: common_vendor.f(orderTabs.value, (item, index, i0) => {
          return {
            a: common_vendor.t(item.title),
            b: item.title,
            c: common_vendor.o(($event) => common_vendor.unref(handleTabClick)(index, item.orderState), item.title)
          };
        }),
        b: activeIndex.value * 20 + "%",
        c: common_vendor.f(orderTabs.value, (item, k0, i0) => {
          return {
            a: common_vendor.f(order.value, (item2, k1, i1) => {
              return common_vendor.e({
                a: common_vendor.t(item2 == null ? void 0 : item2.merchantInfo.name),
                b: common_vendor.t(common_vendor.unref(services_constants.orderStateList)[item2.status].text),
                c: common_vendor.f(item2.skuList, (sku, k2, i2) => {
                  return {
                    a: sku.pic,
                    b: common_vendor.t(sku.productName),
                    c: common_vendor.t(sku.tag),
                    d: sku.id
                  };
                }),
                d: `/pages/detail/index?id=${item2.id}`,
                e: common_vendor.t(item2.payAmount),
                f: item2.status == "0"
              }, item2.status == "0" ? {
                g: common_vendor.o(($event) => onCancel(item2.id), item2),
                h: common_vendor.o(($event) => onOrderPay(item2.id), item2)
              } : {}, {
                i: item2.status == "1"
              }, item2.status == "1" ? {
                j: `/pages/orderCreate/index?sku=${encodeURIComponent(JSON.stringify(item2.skuList))}`
              } : {}, {
                k: item2.status == "2"
              }, item2.status == "2" ? {
                l: `/pages/orderCreate/index?sku=${encodeURIComponent(JSON.stringify(item2.skuList))}`
              } : {}, {
                m: item2.status == "3"
              }, item2.status == "3" ? {
                n: common_vendor.o(($event) => onConfirm(item2.id), item2)
              } : {}, {
                o: item2
              });
            }),
            b: common_vendor.o(onRefresherrefresh, item.title),
            c: item.title
          };
        }),
        d: common_vendor.t("没有更多数据~"),
        e: ((_a = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _a.bottom) + "px",
        f: activeIndex.value,
        g: common_vendor.o(($event) => activeIndex.value = $event.detail.current)
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/orderList/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
