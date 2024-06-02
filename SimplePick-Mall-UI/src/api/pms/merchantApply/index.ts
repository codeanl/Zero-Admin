import request from '@/util/request'

enum API {
    ALL_URL = '/api/pms/merchantsApply/list',
    ADDMENU_URL = '/api/pms/merchantsApply/add',
    UPDATE_URL = '/api/pms/merchantsApply/update',
    DELETEMENU_URL = '/api/pms/merchantsApply/delete',
}

export const reqAllMerchantApply = (current: number, pageSize: number, name: string, status: string, type: string) =>
    request.get<any, any>(API.ALL_URL, { params: { current, pageSize, name, status, type } })


export const reqAddOrUpdateMerchantApply = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATE_URL, data)
    } else {
        return request.post<any, any>(API.ADDMENU_URL, data)
    }
}

export const reqRemoveMerchantApply = (data: any) =>
    request.post<any, any>(API.DELETEMENU_URL, data)
