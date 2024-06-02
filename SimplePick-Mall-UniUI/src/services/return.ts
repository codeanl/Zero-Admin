import { http } from '@/utils/http'

export const addreturnApply = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/returnApply/add',
        data: data,
    })
}

export const updatereturnApply = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/returnApply/update',
        data: data,
    })
}
export const deletereturnApply = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/returnApply/delete',
        data: data,
    })
}
export const listreturnApply = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/returnApply/list',
        data: data,
    })
}

export const listreturnReason = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/returnReason/list',
        data: data,
    })
}

export const returnApplyInfo = (data: any) => {
    return http<any[]>({
        method: 'POST',
        url: '/api/returnApply/info',
        data: data,
    })
}