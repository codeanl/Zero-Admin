"use strict";
const common_vendor = require("../../common/vendor.js");
const services_login = require("../../services/login.js");
require("../../stores/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
require("../../utils/http.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  setup(__props) {
    common_vendor.onLoad(async () => {
      const res = await common_vendor.wx$1.login();
      res.code;
    });
    const onGetphonenumberSimple = async () => {
      const myInfo = {
        username: "codeanl",
        password: "123456"
      };
      const res = await services_login.login(myInfo);
      if (res.code == 200) {
        loginSuccess(res.data);
      }
    };
    const loginSuccess = (profile) => {
      const memberStore = stores_modules_member.useMemberStore();
      memberStore.setProfile(profile);
      common_vendor.index.showToast({ icon: "success", title: "登录成功" });
      setTimeout(() => {
        common_vendor.index.navigateBack();
      }, 500);
    };
    const form = common_vendor.ref({
      username: "codeanl",
      password: "123456"
    });
    const onSubmit = async () => {
      console.log(form);
      const res = await services_login.login(form.value);
      if (res.code == 200) {
        loginSuccess(res.data);
      }
    };
    return (_ctx, _cache) => {
      return {
        a: form.value.username,
        b: common_vendor.o(($event) => form.value.username = $event.detail.value),
        c: form.value.password,
        d: common_vendor.o(($event) => form.value.password = $event.detail.value),
        e: common_vendor.o(onSubmit),
        f: common_vendor.o(onGetphonenumberSimple)
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/login/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
