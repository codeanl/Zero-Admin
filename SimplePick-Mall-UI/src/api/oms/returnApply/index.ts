import request from '@/util/request'
enum API {
    ALLReturnApply_URL = '/api/oms/ReturnApply/list',
    ADDReturnApply_URL = '/api/oms/ReturnApply/add',
    UPDATEReturnApply_URL = '/api/oms/ReturnApply/update',
    DELETEALLReturnApply_URL = '/api/oms/ReturnApply/delete',
    ReturnApplyInfo = '/api/oms/ReturnApply/info',
}
export const reqReturnApplyAll = (current: number, pageSize: number, status: string) =>
    request.get<any, any>(API.ALLReturnApply_URL, { params: { current, pageSize, status } });

export const reqAddOrUpdateReturnApply = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATEReturnApply_URL, data)
    } else {
        return request.post<any, any>(API.ADDReturnApply_URL, data)
    }
}

export const reqDeleteReturnApply = (data: any) =>
    request.post(API.DELETEALLReturnApply_URL, data)

export const reqReturnApplyInfo = (data: any) =>
    request.post(API.ReturnApplyInfo, data)