import { http } from '@/utils/http'

export const addOrder = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/order/add',
        data: data,
    })
}

export const updateOrder = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/order/update',
        data: data,
    })
}
export const deleteOrder = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/order/delete',
        data: data,
    })
}
export const listOrder = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/order/list',
        data: data,
    })
}

export const orderInfo = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/order/info',
        data: data,
    })
}