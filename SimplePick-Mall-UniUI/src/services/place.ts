import { http } from '@/utils/http'

export const addplace = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/place/add',
        data: data,
    })
}

export const updateplace = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/place/update',
        data: data,
    })
}
export const deleteplace = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/place/delete',
        data: data,
    })
}
export const listplace = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/place/list',
        data: data,
    })
}

export const placeInfo = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/place/info',
        data: data,
    })
}