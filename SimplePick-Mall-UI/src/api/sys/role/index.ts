
import request from '@/util/request'

enum API {
  ALLROLE_URL = '/api/sys/role/list',
  ADDROLE_URL = '/api/sys/role/add',
  UPDATEROLE_URL = '/api/sys/role/update',
  ALLPERMISSION_URL = '/api/sys/menu/list',
  MENUBYROLEID_URL = '/api/sys/role/queryMenuByRoleId',
  UPDATEROLEMMENU_URL = '/api/sys/role/updateRoleMenu',
  REMOVEROLE_URL = '/api/sys/role/delete',
}

export const reqAllRoleList = (current: number, pageSize: number, name: string) =>
  request.get<any, any>(API.ALLROLE_URL, { params: { current, pageSize, name } });

export const reqAddOrUpdateRole = (data: any) => {
  if (data.id) {
    return request.post<any, any>(API.UPDATEROLE_URL, data)
  } else {
    return request.post<any, any>(API.ADDROLE_URL, data)
  }
}

export const reqAllMenuList = () =>
  request.post<any, any>(API.ALLPERMISSION_URL)

export const reqMenuByRoleID = (data: any) =>
  request.post<any, any>(API.MENUBYROLEID_URL, data)

export const reqUpdateRoleMenu = (data: any) =>
  request.post<any, any>(API.UPDATEROLEMMENU_URL, data)


export const reqRemoveRole = (data: any) =>
  request.post<any, any>(API.REMOVEROLE_URL, data)

