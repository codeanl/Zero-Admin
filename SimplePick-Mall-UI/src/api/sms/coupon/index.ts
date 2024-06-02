import request from '@/util/request'
enum API {
    ALLcoupon_URL = '/api/sms/coupon/list',
    DELETEcoupon_URL = '/api/sms/coupon/delete',
    ADD_URL = '/api/sms/coupon/add',
    UPDATE_URL = '/api/sms/coupon/update',
}
export const reqcouponList = (current: number, pageSize: number, type: string, name: string, useType: string) =>
    request.get<any, any>(API.ALLcoupon_URL, { params: { current, pageSize, type, name, useType } });

export const reqRemovecoupon = (data: number) =>
    request.post<any, any>(API.DELETEcoupon_URL, data)

export const reqAddOrUpdate = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATE_URL, data)
    } else {
        return request.post<any, any>(API.ADD_URL, data)
    }
}