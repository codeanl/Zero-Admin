import { http } from '@/utils/http'

/**
 * 获取个人信息
 */
export const getMemberProfileAPI = (data: any) => {
  return http<any>({
    method: 'POST',
    url: '/api/auth/info',
    data,
  })
}

/**
 * 修改个人信息
 * @param data 请求体参数
 */
export const putMemberProfileAPI = (data: any) => {
  return http<any>({
    method: 'POST',
    url: '/api/auth/update',
    data,
  })
}
