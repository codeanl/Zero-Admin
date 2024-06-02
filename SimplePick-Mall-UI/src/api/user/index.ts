// 统一管理用户相关接口
import request from '@/util/request'
import type {
    LoginFormData,
    LoginResponseData,
    userInfoResponseData,
} from './type'
enum API {
    LOGIN_URL = '/api/sys/user/login',
    USERINFO_URL = '/api/sys/user/info',
    LOGOUT_URL = '/admin/acl/index/logout',
}
//登录
export const reqLogin = (data: LoginFormData) => request.post<any, LoginResponseData>(API.LOGIN_URL, data)
//获取用户信息 
export const reqUserInfo = () => request.get<any, userInfoResponseData>(API.USERINFO_URL)

export const reqLogout = () => request.post<any, any>(API.LOGOUT_URL)