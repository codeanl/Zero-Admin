"use strict";
const utils_http = require("../utils/http.js");
const getCategoryTopAPI = () => {
  return utils_http.http({
    method: "GET",
    url: "/api/index/categoryList"
  });
};
exports.getCategoryTopAPI = getCategoryTopAPI;
//# sourceMappingURL=category.js.map
