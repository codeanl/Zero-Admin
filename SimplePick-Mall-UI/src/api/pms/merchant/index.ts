import request from '@/util/request'

enum API {
    ALL_URL = '/api/pms/merchants/list',
    ADDMENU_URL = '/api/pms/merchants/add',
    UPDATE_URL = '/api/pms/merchants/update',
    DELETEMENU_URL = '/api/pms/merchants/delete',
}

export const reqAll = (current: number, pageSize: number, name: string, phone: string, address: string, principal: string) =>
    request.get<any, any>(API.ALL_URL, { params: { current, pageSize, name, phone, address, principal } })


export const reqAddOrUpdate = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATE_URL, data)
    } else {
        return request.post<any, any>(API.ADDMENU_URL, data)
    }
}

export const reqRemove = (data: any) =>
    request.post<any, any>(API.DELETEMENU_URL, data)
