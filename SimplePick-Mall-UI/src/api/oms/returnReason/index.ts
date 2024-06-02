import request from '@/util/request'
enum API {
    ALLreturnReason_URL = '/api/oms/returnReason/list',
    ADDreturnReason_URL = '/api/oms/returnReason/add',
    UPDATEreturnReason_URL = '/api/oms/returnReason/update',
    DELETEALLreturnReason_URL = '/api/oms/returnReason/delete',
    returnReasonInfo = '/api/oms/returnReason/info',
}
export const reqreturnReasonAll = (current: number, pageSize: number, name: string, status: string) =>
    request.get<any, any>(API.ALLreturnReason_URL, { params: { current, pageSize, name, status } });

export const reqAddOrUpdatereturnReason = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATEreturnReason_URL, data)
    } else {
        return request.post<any, any>(API.ADDreturnReason_URL, data)
    }
}

export const reqDeletereturnReason = (data: any) =>
    request.post(API.DELETEALLreturnReason_URL, data)

export const reqreturnReasonInfo = (data: any) =>
    request.post(API.returnReasonInfo, data)