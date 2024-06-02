import request from '@/util/request'
enum API {
    ALLhomeAdvertise_URL = '/api/sms/homeAdvertise/list',
    DELETEhomeAdvertise_URL = '/api/sms/homeAdvertise/delete',
    ADD_URL = '/api/sms/homeAdvertise/add',
    UPDATE_URL = '/api/sms/homeAdvertise/update',
}
export const reqhomeAdvertiseList = (current: number, pageSize: number, name: string, status: string) =>
    request.get<any, any>(API.ALLhomeAdvertise_URL, { params: { current, pageSize, name, status } });

export const reqRemovehomeAdvertise = (data: number) =>
    request.post<any, any>(API.DELETEhomeAdvertise_URL, data)

export const reqAddOrUpdate = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATE_URL, data)
    } else {
        return request.post<any, any>(API.ADD_URL, data)
    }
}