import request from '@/util/request'
enum API {
    ALLLOG_URL = '/api/sys/loginLog/list',
    DELETELOGINLOG_URL = '/api/sys/loginLog/delete',
    ALLSYSTEMLOG_URL = '/api/sys/systemLog/list',
    DELETESYSTEMLOG_URL = '/api/sys/systemLog/delete',
}

export const reqLoginLog = (current: number, pageSize: number, status: string) =>
    request.get<any, any>(API.ALLLOG_URL, { params: { current, pageSize, status } });

export const reqRemoveLoginLog = (data: number) =>
    request.post<any, any>(API.DELETELOGINLOG_URL, data)

export const reqSystemLog = (current: number, pageSize: number, method: string) =>
    request.get<any, any>(API.ALLSYSTEMLOG_URL, { params: { current, pageSize, method } });

export const reqRemoveSystemLog = (data: number) =>
    request.post<any, any>(API.DELETESYSTEMLOG_URL, data)