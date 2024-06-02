import request from '@/util/request'

enum API {
  ALLPERMISSION_URL = '/api/sys/menu/list',
  ADDMENU_URL = '/api/sys/menu/add',
  UPDATE_URL = '/api/sys/menu/update',
  DELETEMENU_URL = '/api/sys/menu/delete',
}

export const reqAllPermission = () =>
  request.post<any, any>(API.ALLPERMISSION_URL)

export const reqAddOrUpdateMenu = (data: any) => {
  if (data.id) {
    return request.post<any, any>(API.UPDATE_URL, data)
  } else {
    return request.post<any, any>(API.ADDMENU_URL, data)
  }
}

export const reqRemoveMenu = (data: any) =>
  request.post<any, any>(API.DELETEMENU_URL, data)
