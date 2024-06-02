"use strict";
const common_vendor = require("../../common/vendor.js");
const services_order = require("../../services/order.js");
const services_return = require("../../services/return.js");
require("../../stores/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
require("../../utils/http.js");
if (!Array) {
  const _easycom_uni_popup2 = common_vendor.resolveComponent("uni-popup");
  _easycom_uni_popup2();
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
    const popup = common_vendor.ref();
    let reason = common_vendor.ref("");
    const description = common_vendor.ref("");
    common_vendor.ref("");
    common_vendor.onLoad(() => {
      getData();
    });
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const reasonList = common_vendor.ref();
    const order = common_vendor.ref();
    const getData = async () => {
      const res = await services_return.listreturnReason({});
      if (res.code === 200) {
        reasonList.value = res.data;
      }
      const res1 = await services_order.orderInfo({ id: parseInt(query.id) });
      if (res1.code === 200) {
        order.value = res1.data;
      }
    };
    let place = common_vendor.ref();
    let ddd = (item) => {
      reason.value = item;
      place.value = item;
    };
    let dddd = () => {
      console.log(place);
    };
    const memberStore = stores_modules_member.useMemberStore();
    let pz = common_vendor.ref("");
    let limit11 = async () => {
      let res = await services_return.addreturnApply({
        userID: memberStore == null ? void 0 : memberStore.profile.info.id,
        orderID: order.value.orderInfo.id,
        returnReasonID: place.value.id,
        status: "0",
        description: description.value,
        proofPics: pz.value,
        returnAmount: order.value.orderInfo.totalAmount
      });
      if (res.code == 200) {
        common_vendor.index.redirectTo({ url: `/pages/detail/index?id=${parseInt(query.id)}` });
      }
    };
    const onAvatarChange = () => {
      common_vendor.index.chooseMedia({
        // 文件个数
        count: 1,
        // 文件类型
        mediaType: ["image"],
        success: (res) => {
          const { tempFilePath } = res.tempFiles[0];
          uploadFile(tempFilePath);
        }
      });
    };
    const uploadFile = (file) => {
      common_vendor.index.uploadFile({
        url: "/api/upload",
        name: "file",
        filePath: file,
        success: (res) => {
          pz.value = JSON.parse(res.data).data;
          console.log(pz.value);
        }
      });
    };
    return (_ctx, _cache) => {
      var _a, _b;
      return {
        a: common_vendor.f((_a = order.value) == null ? void 0 : _a.skuList, (item, k0, i0) => {
          return {
            a: item.pic,
            b: common_vendor.t(item.name),
            c: common_vendor.t(item.tag),
            d: common_vendor.t(item.price),
            e: common_vendor.t(item.count),
            f: item.id,
            g: `/pages/goods/goods?id=${item.spuId}`
          };
        }),
        b: common_vendor.t(common_vendor.unref(place) ? common_vendor.unref(place).name : "请填写退货原因"),
        c: common_vendor.o(($event) => {
          var _a2, _b2;
          return (_b2 = (_a2 = popup.value) == null ? void 0 : _a2.open) == null ? void 0 : _b2.call(_a2);
        }),
        d: description.value,
        e: common_vendor.o(($event) => description.value = $event.detail.value),
        f: common_vendor.unref(pz),
        g: common_vendor.o(onAvatarChange),
        h: common_vendor.o(
          //@ts-ignore
          (...args) => common_vendor.unref(limit11) && common_vendor.unref(limit11)(...args)
        ),
        i: ((_b = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _b.bottom) + "px",
        j: common_vendor.f(reasonList.value, (item, k0, i0) => {
          return {
            a: common_vendor.t(item.name),
            b: item === common_vendor.unref(reason) ? 1 : "",
            c: item,
            d: common_vendor.o(($event) => common_vendor.unref(ddd)(item), item)
          };
        }),
        k: common_vendor.o(($event) => {
          var _a2, _b2;
          return (_b2 = (_a2 = popup.value) == null ? void 0 : _a2.close) == null ? void 0 : _b2.call(_a2), common_vendor.unref(dddd);
        }),
        l: common_vendor.sr(popup, "e4e2b222-0", {
          "k": "popup"
        }),
        m: common_vendor.p({
          type: "bottom",
          ["background-color"]: "#fff"
        })
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-e4e2b222"], ["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/returnApply/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
