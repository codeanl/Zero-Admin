import { http } from '@/utils/http'

/**
 * 首页-广告区域-小程序
 * @param distributionSite 广告区域展示位置（投放位置 投放位置，1为首页，2为分类商品页） 默认是1
 */
export const getHomeBannerAPI = () => {
  return http<any[]>({
    method: 'GET',
    url: '/api/index/homeAdvertiseList',
  })
}

/**
 * 首页-前台分类-小程序
 */
export const getHomeCategoryAPI = () => {
  return http<any[]>({
    method: 'GET',
    url: '/api/index/categoryList',
  })
}

/**
 * 首页-热门推荐-小程序
 */
export const getHomeHotAPI = () => {
  return http<any[]>({
    method: 'GET',
    url: '/api/index/subjectList',
  })
}

/**
 * 猜你喜欢-小程序
 */
export const getHomeGoodsGuessLikeAPI = (data?: any) => {
  return http<any>({
    method: 'GET',
    url: '/api/index/productList',
    data,
  })
}


/**
 * 商品详情-小程序
 */
export const getProductInfoAPI = (data: any) => {
  return http<any>({
    method: 'POST',
    url: '/api/index/productInfo',
    data,
  })
}


export const placeList = (data: any) => {
  return http<any>({
    method: 'POST',
    url: '/api/index/placeList',
    data,
  })
}


