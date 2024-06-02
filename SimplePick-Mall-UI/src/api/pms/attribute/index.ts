
import request from '@/util/request'
enum API {
  All_URL = '/api/pms/attribute/list',
  ADD_URL = '/api/pms/attribute/add',
  UPDATE_URL = '/api/pms/attribute/update',
  DELETE_URL = '/api/pms/attribute/delete',


}


export const reqAllAttribute = (current: number, pageSize: number, name: string, type: string, attributeCategoryID: number) =>
  request.get<any, any>(API.All_URL, { params: { current, pageSize, name, type, attributeCategoryID } });


export const reqAddOrUpdate = (data: any) => {
  if (data.id) {
    return request.post<any, any>(API.UPDATE_URL, data)
  } else {
    return request.post<any, any>(API.ADD_URL, data)
  }
}

export const reqRemove = (data: number) =>
  request.post<any, any>(API.DELETE_URL, data)


