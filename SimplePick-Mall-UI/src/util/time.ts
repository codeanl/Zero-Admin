//获取时间段 上午下午
export const getTime = () => {
    let msg = ''
    let h = new Date().getHours()
    if (h <= 12) {
        msg = '早上'
    } else if (h <= 18) {
        msg = '下午'
    } else {
        msg = '晚上'
    }
    return msg
}