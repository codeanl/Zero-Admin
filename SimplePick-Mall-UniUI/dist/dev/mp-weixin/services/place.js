"use strict";
const utils_http = require("../utils/http.js");
const listplace = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/place/list",
    data
  });
};
exports.listplace = listplace;
//# sourceMappingURL=place.js.map
