//路由鉴权
import router from '@/router'
import * as nprogress from 'nprogress';
import 'nprogress/nprogress.css'
import pinia from './store'
import useUserStore from './store/sys/user'
let userStore = useUserStore(pinia)
import setting from './setting'
// 全局前置守卫  保护访问路由的访问权限,也就是对路由权限进行控制
// 只有在允许特定的情况下才能访问某个固定的路由,具备查看某些页面的权限,
router.beforeEach(async (to, from, next) => {
    console.log(from);
    document.title = to.meta.title + ` | ${setting.title}`
    nprogress.start()
    let token = userStore.token
    let username = userStore.username
    if (to.path === '/login' || to.path === '/merchantApply') {
        next()
    } else if (token) {
        if (username) {
            next()
        } else {
            try {
                await userStore.userInfo()
                next({ ...to, replace: true })
            } catch (error) {
                console.error(error)
                next({ path: '/login' })
            }
        }
    } else {
        next({ path: '/login' })
        try {
            await userStore.userLogout()
        } catch (error) {
            console.error(error)
        }
    }
})
//全局后置守卫
router.afterEach((to: any, from: any) => {
    console.log(to + from);
    nprogress.done()
})