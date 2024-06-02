import { http } from '@/utils/http'

export const addCart = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/cart/add',
        data: data,
    })
}

export const updateCart = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/cart/update',
        data: data,
    })
}
export const deleteCart = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/cart/delete',
        data: data,
    })
}
export const listCart = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/cart/list',
        data: data,
    })
}