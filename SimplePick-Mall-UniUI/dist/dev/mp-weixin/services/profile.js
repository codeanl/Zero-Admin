"use strict";
const utils_http = require("../utils/http.js");
const getMemberProfileAPI = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/auth/info",
    data
  });
};
const putMemberProfileAPI = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/auth/update",
    data
  });
};
exports.getMemberProfileAPI = getMemberProfileAPI;
exports.putMemberProfileAPI = putMemberProfileAPI;
//# sourceMappingURL=profile.js.map
