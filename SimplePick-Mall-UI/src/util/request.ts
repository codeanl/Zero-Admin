
import { ElMessage } from 'element-plus'
import useUserStore from '@/store/sys/user'
// 二次封装axios
import axios from 'axios'
//创建axios实例
let request = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_API,
    // baseURL: "http://114.115.183.226:6666",
    // baseURL: "http://localhost:6666",
    timeout: 10000, //超时时间
})
//添加请求与响应拦截器
request.interceptors.request.use((config) => {
    //获取token携带给服务器
    let userStore = useUserStore()
    if (userStore.token) {
        // config.headers.token = userStore.token
        config.headers.Authorization = `Bearer ${userStore.token}`
    }
    //返回配置对象
    return config
})
//响应拦截器
request.interceptors.response.use((response) => {
    //成功回调
    return response.data
}, (error) => {
    //失败回调
    let message = ''
    let status = error.response.status
    switch (status) {
        // 401: 未登录
        // 未登录则跳转登录页面，并携带当前页面的路径
        // 在登录成功后返回当前页面，这一步需要在登录页操作。
        case 401:
            message = '未登录'
            break
        // 403 token过期
        // 登录过期对用户进行提示
        // 清除本地token和清空vuex中token对象
        // 跳转登录页面
        case 403:
            message = '登录过期，请重新登录'
            break
        case 404:
            message = '网络请求不存在'
            break
        case 500:
            message = '服务器出现问题'
            break
        default:
            message = error.response.data.message
            break
    }
    ElMessage({
        type: 'error',
        message,
    })
    return Promise.reject(error)
})
export default request