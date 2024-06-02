import request from '@/util/request'
enum API {
    ALLPlACE_URL = '/api/pms/place/list',
    DELETEPLACE_URL = '/api/pms/place/delete',
    ADD_URL = '/api/pms/place/add',
    UPDATE_URL = '/api/pms/place/update',
}
export const reqPlaceList = (current: number, pageSize: number, name: string, place: string, phone: string, principal: string) =>
    request.get<any, any>(API.ALLPlACE_URL, { params: { current, pageSize, name, place, phone, principal } });

export const reqRemovePlace = (data: number) =>
    request.post<any, any>(API.DELETEPLACE_URL, data)

export const reqAddOrUpdate = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATE_URL, data)
    } else {
        return request.post<any, any>(API.ADD_URL, data)
    }
}