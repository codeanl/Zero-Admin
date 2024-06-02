import { defineStore } from 'pinia'
import router from '@/router'
import { reqLogin, reqUserInfo } from '@/api/user'
import type {
    LoginFormData,
    LoginResponseData,
} from '@/api/user/type'
import { SET_TOKEN, GET_TOKEN, REMOVE_TOKEN } from '@/util/token'
import { constantRoute, asyncRoute, anyRoute } from '@/router/routes'

// @ts-ignore 深拷贝
import cloneDeep from 'lodash/cloneDeep'
//过滤展示的异步路由
function filterAsyncRoute(asyncRoute: any, routes: any) {
    return asyncRoute.filter((item: any) => {
        if (routes.includes(item.name)) {
            if (item.children && item.children.length > 0) {
                item.children = filterAsyncRoute(item.children, routes)
            }
            return true
        }
    })
}

let useUserStore = defineStore('User', {
    // 小仓库存储数据的地方
    state: (): any => {
        return {
            token: GET_TOKEN()!,
            menuRoutes: constantRoute,
            username: '',
            avatar: '',
            buttons: [],
            routes: [],
            roles:[],
        }
    },
    // 异步|逻辑的地方
    actions: {
        //用户登录方法
        async userLogin(data: LoginFormData) {
            let res: LoginResponseData = await reqLogin(data)
            // success=>token
            // error=>error.message
            if (res.code === 200) {
                this.token = res.data as string
                // 持久化
                SET_TOKEN(res.data as string)
                return 'ok'
            } else {
                // return Promise.reject(new Error(res.data as string))
                return res.message
            }
        },
        async userInfo() {
            let res: any = await reqUserInfo()
            if (res.code === 200) {
                this.username = res.data.userInfo.username as string
                this.avatar = res.data.userInfo.avatar as string
                this.buttons = res.data.buttons
                this.routes = res.data.routes
                this.roles = res.data.roles
                //计算用户展示的异步路由
                let userAsyncRoute = filterAsyncRoute(
                    cloneDeep(asyncRoute),
                    res.data.routes,
                )
                //数据整理
                this.menuRoutes = [...constantRoute, ...userAsyncRoute, anyRoute];
                //动态追加
                [...userAsyncRoute, anyRoute].forEach((route: any) => {
                    router.addRoute(route)
                })
                return 'ok'
            }
            else {
                REMOVE_TOKEN()
                // return Promise.reject(new Error(res.message))
                return res.message
            }
        },
        async userLogout() {
            this.token = ''
            this.username = ''
            this.avatar = ''
            REMOVE_TOKEN()
        },
    },
    getters: {},
})

export default useUserStore
