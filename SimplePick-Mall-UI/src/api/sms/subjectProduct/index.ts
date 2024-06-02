import request from '@/util/request'
enum API {
    ALLsubjectProduct_URL = '/api/sms/subjectProduct/list',
    DELETEsubjectProduct_URL = '/api/sms/subjectProduct/delete',
    ADD_URL = '/api/sms/subjectProduct/add',
    UPDATE_URL = '/api/sms/subjectProduct/update',
}
export const reqsubjectProductList = (current: number, pageSize: number, subjectId: number) =>
    request.get<any, any>(API.ALLsubjectProduct_URL, { params: { current, pageSize, subjectId } });

export const reqRemovesubjectProduct = (data: number) =>
    request.post<any, any>(API.DELETEsubjectProduct_URL, data)

export const reqAdd = (data: any) => {
    return request.post<any, any>(API.ADD_URL, data)
}

export const reqUpdate = (data: any) => {
    return request.post<any, any>(API.UPDATE_URL, data)
}