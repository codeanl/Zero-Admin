"use strict";
const common_vendor = require("../../common/vendor.js");
const services_profile = require("../../services/profile.js");
require("../../stores/index.js");
const stores_modules_member = require("../../stores/modules/member.js");
require("../../utils/http.js");
const _sfc_main = /* @__PURE__ */ common_vendor.defineComponent({
  __name: "index",
  setup(__props) {
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    const profile = common_vendor.ref();
    const getMemberProfileData = async () => {
      const requestData = { id: memberStore.profile.info.id };
      const res = await services_profile.getMemberProfileAPI(requestData);
      profile.value = res.data;
      memberStore.profile.info.avatar = res.data.avatar;
      memberStore.profile.info.nickname = res.data.nickname;
    };
    common_vendor.onLoad(() => {
      getMemberProfileData();
    });
    const memberStore = stores_modules_member.useMemberStore();
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
          updateAvatar(JSON.parse(res.data));
        }
      });
    };
    let updateAvatar = async (data) => {
      console.log(data.data);
      memberStore.profile.info.avatar = data.data;
      let res = await services_profile.putMemberProfileAPI({
        id: memberStore.profile.info.id,
        avatar: data.data
      });
      if (res.code == 200) {
        common_vendor.index.showToast({ icon: "success", title: "更新成功" });
        profile.value.avatar = data.data;
      } else {
        common_vendor.index.showToast({ icon: "error", title: "出现错误" });
      }
    };
    const onGenderChange = (ev) => {
      profile.value.gender = ev.detail.value;
    };
    const onSubmit = async () => {
      const { nickname, gender, email, signature } = profile.value;
      const res = await services_profile.putMemberProfileAPI({
        id: memberStore.profile.info.id,
        nickname,
        gender,
        email,
        signature
      });
      if (res.code == 200) {
        common_vendor.index.showToast({ icon: "success", title: "保存成功" });
        setTimeout(() => {
          common_vendor.index.navigateBack();
        }, 400);
        memberStore.profile.nickname = res.data.nickname;
        memberStore.profile.email = res.data.email;
        memberStore.profile.gender = res.data.gender;
        memberStore.profile.signature = res.data.signature;
      }
    };
    return (_ctx, _cache) => {
      var _a, _b, _c, _d, _e, _f;
      return {
        a: ((_a = common_vendor.unref(safeAreaInsets)) == null ? void 0 : _a.top) + "px",
        b: (_b = profile.value) == null ? void 0 : _b.avatar,
        c: common_vendor.o(onAvatarChange),
        d: common_vendor.t((_c = profile.value) == null ? void 0 : _c.username),
        e: profile.value.nickname,
        f: common_vendor.o(($event) => profile.value.nickname = $event.detail.value),
        g: profile.value.email,
        h: common_vendor.o(($event) => profile.value.email = $event.detail.value),
        i: ((_d = profile.value) == null ? void 0 : _d.gender) === "0",
        j: ((_e = profile.value) == null ? void 0 : _e.gender) === "1",
        k: ((_f = profile.value) == null ? void 0 : _f.gender) === "2",
        l: common_vendor.o(onGenderChange),
        m: profile.value.signature,
        n: common_vendor.o(($event) => profile.value.signature = $event.detail.value),
        o: common_vendor.o(onSubmit)
      };
    };
  }
});
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/pages/profile/index.vue"]]);
wx.createPage(MiniProgramPage);
//# sourceMappingURL=index.js.map
