/** 订单状态枚举 */
export enum OrderState {
  /** 待付款 */
  DaiFuKuan = 0,
  /** 待发货 */
  DaiFaHuo = 1,
  /** 待收货 */
  DaiShouHuo = 2,
  /** 待评价 */
  DaiPingJia = 3,
  /** 已完成 */
  YiWanCheng = 4,
  /** 已取消 */
  YiQuXiao = 5,
}
/** 订单状态列表 */
export const orderStateList = [
  { id: 0, text: '待付款' },
  { id: 1, text: '待发货' },
  { id: 2, text: '运输中' },
  { id: 3, text: '待提货' },
  { id: 4, text: '订单完成' },
  { id: 5, text: '无效订单' },
]
