import { http } from '@/utils/http'

type LoginWxMinParams = {
    code: string
    encryptedData?: string
    iv?: string
}
/**
 * 小程序登录
 * @param data 请求参数
 */
export const postLoginWxMinAPI = (data: LoginWxMinParams) => {
    return http<any>({
        method: 'POST',
        url: '/login/wxMin',
        data,
    })
}

/**
 * 小程序登录_内测版
 * @param phoneNumber 模拟手机号码
 */
export const login = (data:LoginParams) => {
    return http<any>({
        method: 'POST',
        url: '/api/auth/login',
        data,
    })
}

type LoginParams = {
    username: string
    password: string
}
/**
 * 传统登录-用户名+密码
 * @param data 请求参数
 */
export const postLoginAPI = (data: LoginParams) => {
    return http<any>({
        method: 'POST',
        url: '/login',
        data,
    })
}
