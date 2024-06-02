import { http } from '@/utils/http'

/**
 * 分类列表-小程序
 */
export const getCategoryTopAPI = () => {
  return http<any[]>({
    method: 'GET',
    url: '/api/index/categoryList',
  })
}
