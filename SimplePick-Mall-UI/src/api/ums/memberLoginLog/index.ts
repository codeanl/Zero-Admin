import request from '@/util/request'
enum API {
    ALLMemberLoginLog_URL = '/api/ums/MemberLoginLog/list',
    ADDMemberLoginLog_URL = '/api/ums/MemberLoginLog/add',
    UPDATEMemberLoginLog_URL = '/api/ums/MemberLoginLog/update',
    DELETEALLMemberLoginLog_URL = '/api/ums/MemberLoginLog/delete',
}
export const reqMemberLoginLogAll = (current: number, pageSize: number, userId: number) =>
    request.get<any, any>(API.ALLMemberLoginLog_URL, { params: { current, pageSize, userId } });

export const reqDeleteMemberLoginLog = (data: any) =>
    request.post(API.DELETEALLMemberLoginLog_URL, data)