"use strict";
const utils_http = require("../utils/http.js");
const getHotProductApi = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/index/subjectProductList",
    data
  });
};
exports.getHotProductApi = getHotProductApi;
//# sourceMappingURL=hot.js.map
