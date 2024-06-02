"use strict";
const utils_http = require("../utils/http.js");
const addreturnApply = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/returnApply/add",
    data
  });
};
const listreturnReason = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/returnReason/list",
    data
  });
};
const returnApplyInfo = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/returnApply/info",
    data
  });
};
exports.addreturnApply = addreturnApply;
exports.listreturnReason = listreturnReason;
exports.returnApplyInfo = returnApplyInfo;
//# sourceMappingURL=return.js.map
