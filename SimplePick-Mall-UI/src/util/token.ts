//封装本地存储数据与读取数据方法
export const SET_TOKEN = (token: string) => {
    localStorage.setItem('TOKEN', token)
}
//本地获取
export const GET_TOKEN = () => {
    return localStorage.getItem('TOKEN')
}
//清除token
export const REMOVE_TOKEN = () => {
    localStorage.removeItem('TOKEN')
}
