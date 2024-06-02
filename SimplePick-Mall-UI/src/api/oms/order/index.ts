import request from '@/util/request'
enum API {
    ALLOrder_URL = '/api/oms/order/list',
    ADDOrder_URL = '/api/oms/order/add',
    UPDATEOrder_URL = '/api/oms/order/update',
    DELETEALLOrder_URL = '/api/oms/order/delete',
    OrderInfo = '/api/oms/order/info',
}
export const reqOrderAll = (current: number, pageSize: number, orderSn: string, memberUsername: string, status: string) =>
    request.get<any, any>(API.ALLOrder_URL, { params: { current, pageSize, orderSn, memberUsername, status } });

export const reqAddOrUpdateOrder = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATEOrder_URL, data)
    } else {
        return request.post<any, any>(API.ADDOrder_URL, data)
    }
}

export const reqDeleteOrder = (data: any) =>
    request.post(API.DELETEALLOrder_URL, data)

export const reqOrderInfo = (data: any) =>
    request.post(API.OrderInfo, data)