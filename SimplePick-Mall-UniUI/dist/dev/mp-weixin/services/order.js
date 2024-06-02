"use strict";
const utils_http = require("../utils/http.js");
const addOrder = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/order/add",
    data
  });
};
const updateOrder = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/order/update",
    data
  });
};
const listOrder = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/order/list",
    data
  });
};
const orderInfo = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/order/info",
    data
  });
};
exports.addOrder = addOrder;
exports.listOrder = listOrder;
exports.orderInfo = orderInfo;
exports.updateOrder = updateOrder;
//# sourceMappingURL=order.js.map
