import request from '@/util/request'

enum API {
    ALLUSER_URL = '/api/sys/user/list',
    ADDUSER_URL = '/api/sys/user/add',
    UPDATEUSER_URL = '/api/sys/user/update',
    ALLROLEURL = '/api/sys/role/list',
    SETROLE_url = '/admin/acl/user/doAssignRole',
    DELETEUSER_URL = '/api/sys/user/delete',
    DELETEALLUSER_URL = '/admin/acl/user/batchRemove',
    ALLROLEBYUSER_URl = '/api/sys/role/byUserList',
    UPDATEPWD_URl = '/api/sys/user/updatePassword',
    resetPassword_url = '/api/sys/user/restartPassword',
    rbac_url = '/api/sys/user/rbac'
    
}

export const reqUserInfo = (current: number, pageSize: number, username: string, phone: string, nickname: string, status: string, gender: string, email: string) =>
    request.get<any, any>(API.ALLUSER_URL, { params: { current, pageSize, username, phone, nickname, status, gender, email } });

export const reqAddOrUpdateUser = (data: any) => {
    if (data.id) {
        return request.post<any, any>(API.UPDATEUSER_URL, data)
    } else {
        return request.post<any, any>(API.ADDUSER_URL, data)
    }
}

export const Rbac = (data: any) =>
    request.post<any, any>(API.rbac_url, data)

export const reqAllRole = () =>
    request.get<any, any>(API.ALLROLEURL)

export const reqSetUserRole = (data: any) =>
    request.post<any, any>(API.SETROLE_url, data)

export const reqRemoveUser = (data: number) =>
    request.post<any, any>(API.DELETEUSER_URL, data)

export const reqSelectUser = (idList: number[]) =>
    request.delete(API.DELETEALLUSER_URL, { data: idList })

export const reqMenuListByUser = (data: any) =>
    request.post<any, any>(API.ALLROLEBYUSER_URl, data)

export const reqUpdatePwd = (data: any) =>
    request.post<any, any>(API.UPDATEPWD_URl, data)

export const resetPassword = (data: any) =>
    request.post<any, any>(API.resetPassword_url, data)