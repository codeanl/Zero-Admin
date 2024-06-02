import request from '@/util/request'

enum API {
    ALL_URL = '/api/pms/category/list',
    ADDMENU_URL = '/api/pms/category/add',
    UPDATE_URL = '/api/pms/category/update',
    DELETEMENU_URL = '/api/pms/category/delete',
}

export const reqAll = () =>
    request.post<any, any>(API.ALL_URL)

export const reqAddOrUpdate = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATE_URL, data)
    } else {
        return request.post<any, any>(API.ADDMENU_URL, data)
    }
}

export const reqRemove = (data: any) =>
    request.post<any, any>(API.DELETEMENU_URL, data)
