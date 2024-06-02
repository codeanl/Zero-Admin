// layout组件相关
import { defineStore } from 'pinia'
let useLayoutSettingStore = defineStore('LayoutSettingStore', {
    state: () => {
        return {
            fold: false, //菜单折叠
            refresh: false,
        }
    }
})

export default useLayoutSettingStore