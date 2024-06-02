"use strict";
const utils_http = require("../utils/http.js");
const addCart = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/cart/add",
    data
  });
};
const deleteCart = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/cart/delete",
    data
  });
};
const listCart = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/cart/list",
    data
  });
};
exports.addCart = addCart;
exports.deleteCart = deleteCart;
exports.listCart = listCart;
//# sourceMappingURL=cart.js.map
