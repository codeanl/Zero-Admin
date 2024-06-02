import request from '@/util/request'
enum API {
    ALLsubject_URL = '/api/sms/subject/list',
    DELETEsubject_URL = '/api/sms/subject/delete',
    ADD_URL = '/api/sms/subject/add',
    UPDATE_URL = '/api/sms/subject/update',
}
export const reqsubjectList = (current: number, pageSize: number, name: string, status: string) =>
    request.get<any, any>(API.ALLsubject_URL, { params: { current, pageSize, name, status } });

export const reqRemovesubject = (data: number) =>
    request.post<any, any>(API.DELETEsubject_URL, data)

export const reqAddOrUpdate = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATE_URL, data)
    } else {
        return request.post<any, any>(API.ADD_URL, data)
    }
}