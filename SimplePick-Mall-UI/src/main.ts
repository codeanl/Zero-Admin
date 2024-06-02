import { createApp } from 'vue'
// import './style.css'
import App from './App.vue'
const app = createApp(App)

//引入全局scss
import '@/style/index.scss'

//引入ElementPlus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
//暗黑模式 
import 'element-plus/theme-chalk/dark/css-vars.css'
//国际化
app.use(ElementPlus, {
    locale: zhCn,
})
//echarts
import "echarts";
import ECharts from 'vue-echarts'
app.component('v-chart', ECharts)

//引入路由 
import router from './router'
app.use(router)

//引入element-ui 图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

//将el图标注册全局组件
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
// svg插件需要配置代码
import 'virtual:svg-icons-register'

//引入自定义插件全局
import globalComponent from '@/components'
app.use(globalComponent)

//引入pinia
import pinia from './store'
app.use(pinia)

//引入路由鉴权
import './permission'

//引入自定义指令文件
import { isHasButton } from './directive/has'
isHasButton(app)

app.mount('#app')