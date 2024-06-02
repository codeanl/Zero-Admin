
import request from '@/util/request'
enum API {
  All_URL = '/api/pms/attributeCategory/list',
  ADD_URL = '/api/pms/attributeCategory/add',
  UPDATE_URL = '/api/pms/attributeCategory/update',
  DELETE_URL = '/api/pms/attributeCategory/delete',


}


export const reqAllattributeCategory = () =>
  request.get<any, any>(API.All_URL)

export const reqAddOrUpdate = (data: any) => {
  if (data.id) {
    return request.post<any, any>(API.UPDATE_URL, data)
  } else {
    return request.post<any, any>(API.ADD_URL, data)
  }
}

export const reqRemove = (data: number) =>
  request.post<any, any>(API.DELETE_URL, data)


