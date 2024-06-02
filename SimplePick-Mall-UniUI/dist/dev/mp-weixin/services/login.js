"use strict";
const utils_http = require("../utils/http.js");
const login = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/auth/login",
    data
  });
};
exports.login = login;
//# sourceMappingURL=login.js.map
