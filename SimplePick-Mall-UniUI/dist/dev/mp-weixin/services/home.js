"use strict";
const utils_http = require("../utils/http.js");
const getHomeBannerAPI = () => {
  return utils_http.http({
    method: "GET",
    url: "/api/index/homeAdvertiseList"
  });
};
const getHomeCategoryAPI = () => {
  return utils_http.http({
    method: "GET",
    url: "/api/index/categoryList"
  });
};
const getHomeHotAPI = () => {
  return utils_http.http({
    method: "GET",
    url: "/api/index/subjectList"
  });
};
const getHomeGoodsGuessLikeAPI = (data) => {
  return utils_http.http({
    method: "GET",
    url: "/api/index/productList",
    data
  });
};
const getProductInfoAPI = (data) => {
  return utils_http.http({
    method: "POST",
    url: "/api/index/productInfo",
    data
  });
};
exports.getHomeBannerAPI = getHomeBannerAPI;
exports.getHomeCategoryAPI = getHomeCategoryAPI;
exports.getHomeGoodsGuessLikeAPI = getHomeGoodsGuessLikeAPI;
exports.getHomeHotAPI = getHomeHotAPI;
exports.getProductInfoAPI = getProductInfoAPI;
//# sourceMappingURL=home.js.map
