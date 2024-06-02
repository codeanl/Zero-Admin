<template>
    <!-- 上边搜索 -->
    <el-card>
        <el-form :inline="true">
            <el-form-item label="名称:">
                <el-input placeholder="请输入搜索的用户名" v-model="name"></el-input>
            </el-form-item>
            <el-form-item label="状态:">
                <el-select v-model="status" class="m-2" placeholder="请选择状态">
                    <el-option label="正常" value="1" />
                    <el-option label="禁用" value="0" />
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" size="default" @click="search" :disabled="name || status ? false : true">
                    搜索
                </el-button>
                <el-button size="default" @click="reset">重置</el-button>
            </el-form-item>
        </el-form>
    </el-card>
    <!--  -->
    <!--  -->
    <el-card>
        <el-button type="success" size="default" @click="add">
            添加
        </el-button>
        <el-button type="danger" size="default" @click="deleteSelect" :disabled="selectIdArr.length ? false : true">
            批量删除
        </el-button>
        <el-table border :data="listArr" @selection-change="selectChange" style="margin: 15px 0">
            <el-table-column type="selection" align="center" width="40px"></el-table-column>
            <el-table-column label="id" align="center" prop="id" width="50px"></el-table-column>
            <el-table-column label="名称" align="center" prop="name" show-overflow-tooltip></el-table-column>
            <el-table-column label="状态" align="center" prop="status" show-overflow-tooltip>
                <template #="{ row }">
                    <el-switch v-model="row.status" class="ml-2"
                        style="--el-switch-on-color: #13ce66; --el-switch-off-color: #5597ce" active-value="1"
                        inactive-value="0" @change="handleChange(row)" />
                </template>
            </el-table-column>
            <el-table-column label="操作" width="300px" align="center">
                <template #="{ row }">
                    <el-button type="primary" size="small" icon="Edit" @click="update(row)">
                        编辑
                    </el-button>
                    <el-popconfirm :title="`你确定删除${row.name}`" width="260px" @confirm="deletePlace(row.id)">
                        <template #reference>
                            <el-button type="danger" size="small" icon="Delete">
                                删除
                            </el-button>
                        </template>
                    </el-popconfirm>
                </template>
            </el-table-column>
        </el-table>
        <!-- 分页 -->
        <el-pagination v-model:current-page="pageNo" v-model:page-size="pageSize" :page-sizes="[5, 10, 15, 20]"
            :background="true" layout="prev, pager, next, jumper, -> , sizes, total" :total="total" @current-change="getHas"
            @size-change="handler" />
    </el-card>

    <!-- 抽屉  完成 添加｜修改 的窗口 -->
    <el-dialog v-model="drawer" :title="Params.id ? '更新' : '添加'">
        <el-form :model="Params" ref="formRef">
            <el-form-item label="名称" prop="name">
                <el-input placeholder="请您输入名称" v-model="Params.name"></el-input>
            </el-form-item>
            <el-form-item label="状态" prop="status">
                <el-select v-model="Params.status" class="m-2" placeholder="请选择状态">
                    <el-option label="禁用" value="0" />
                    <el-option label="正常" value="1" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <div style="flex: auto">
                <el-button @click="cancel">取消</el-button>
                <el-button type="primary" @click="save">确定</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, nextTick } from 'vue'
import { ElMessage } from 'element-plus';
import { reqreturnReasonAll, reqDeletereturnReason, reqAddOrUpdatereturnReason } from '@/api/oms/returnReason'
//setting仓库
import useLayoutSettingStore from '@/store/setting'
let settingStore = useLayoutSettingStore()
//默认页码
let pageNo = ref<number>(1)
//默认个数
let pageSize = ref<number>(10)
let total = ref<number>(0)
//数据列表
let listArr = ref<any>([])
//收集用户查找的关键字
let name = ref<string>('')
let status = ref<string>('')
//收集删除的id
let ids = ref<number[]>([])
//多选框选择的id
let selectIdArr = ref<any[]>([])
//复选框选择
const selectChange = (value: any) => {
    selectIdArr.value = value
}
//组件实例
let formRef = ref<any>()
//定义响应式数据 抽屉的显示隐藏
let drawer = ref<boolean>(false)
let Params = reactive<any>({
    id: 0,
    name: '',
    status: '',
})
//组件挂载完毕
onMounted(() => {
    getHas()
})
//获取信息
const getHas = async (pager = 1) => {
    pageNo.value = pager
    let res: any = await reqreturnReasonAll(
        pageNo.value,
        pageSize.value,
        name.value,
        status.value,
    )
    if (res.code == 200) {
        total.value = res.total
        listArr.value = res.data
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}
//下拉改变
const handler = () => {
    getHas()
}
//重置按钮
const reset = () => {
    settingStore.refresh = !settingStore.refresh
}
//搜索按钮
const search = () => {
    getHas()
    name.value = ''
    status.value = ''
}

// 删除按钮
const deletePlace = async (id: number) => {
    ids.value.push(id);
    const requestData: any = { ids: ids.value };
    let res: any = await reqDeletereturnReason(requestData);
    if (res.code == 200) {
        getHas(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1)
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}
//批量删除用户按钮
const deleteSelect = async () => {
    ids.value = selectIdArr.value.map((item) => {
        return item.id
    })
    const requestData: any = { ids: ids.value };
    let res: any = await reqDeletereturnReason(requestData);
    if (res.code === 200) {
        getHas(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1)
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}
// 编辑按钮
const update = (row: any) => {
    drawer.value = true
    Object.assign(Params, row)
    nextTick(() => {
        formRef.value.clearValidate('name')
        formRef.value.clearValidate('status')
    })
}
//添加按钮
const add = () => {
    drawer.value = true
    //存储收集已有的账号信息
    Object.assign(Params, {
        id: 0,
        name: '',
        status: '',
    })
    nextTick(() => {
        formRef.value.clearValidate('name')
        formRef.value.clearValidate('status')
    })
}
//取消按钮
const cancel = () => {
    drawer.value = false
}
//窗口保存按钮
const save = async () => {
    let res: any = await reqAddOrUpdatereturnReason(Params)
    if (res.code == 200) {
        drawer.value = false
        ElMessage({ type: 'success', message: res.message })
        // window.location.reload()
        getHas(listArr.value.length > 1 ? pageNo.value : pageNo.value - 1)
    } else {
        drawer.value = false
        ElMessage({ type: 'error', message: res.message })
    }
}

//状态修改用户
let handleChange = async (row: any) => {
    let data: any = {
        id: row.id as number,
        status: row.status
    }
    let res = await reqAddOrUpdatereturnReason(data)
    if (res.code === 200) {
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}
</script>

<style lang="scss" scoped></style>