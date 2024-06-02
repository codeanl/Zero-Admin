// 登录接口需要携带参数ts类型
export interface LoginFormData {
    username?: string
    password?: string
}

export interface ResponseData {
    code?: number
    message?: string
    ok?: boolean
}

export interface LoginResponseData extends ResponseData {
    code: number
    data?: string
}

export interface userInfoResponseData extends ResponseData {
    data: {
        routes: string[]
        buttons: string[]
        roles: string[]
        name: string
        avatar: string
        token: string

    }
}