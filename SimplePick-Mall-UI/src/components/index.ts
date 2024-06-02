const allGlocalComponent: any = {}
//对外暴露插件对象
export default {
    install(app: any) {
        //注册全部
        Object.keys(allGlocalComponent).forEach(key => {
            app.component(key, allGlocalComponent[key])
        });
    }
}