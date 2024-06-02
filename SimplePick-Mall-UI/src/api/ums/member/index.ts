import request from '@/util/request'
enum API {
    ALLMember_URL = '/api/ums/member/list',
    ADDMember_URL = '/api/ums/member/add',
    UPDATEMember_URL = '/api/ums/member/update',
    DELETEALLMember_URL = '/api/ums/member/delete',
}
export const reqMemberAll = (current: number, pageSize: number, username: string, phone: string, status: string,nickname:string,gender:string) =>
    request.get<any, any>(API.ALLMember_URL, { params: { current, pageSize, username, phone, status ,nickname,gender} });

export const reqAddOrUpdateMember = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATEMember_URL, data)
    } else {
        return request.post<any, any>(API.ADDMember_URL, data)
    }
}

export const reqDeleteMember = (data: any) =>
    request.post(API.DELETEALLMember_URL, data)