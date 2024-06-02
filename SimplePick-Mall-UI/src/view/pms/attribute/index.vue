<template>
    <!-- 上边搜索 -->
    <el-card>
        <el-form :inline="true">
            <el-form-item label="名称:">
                <el-input placeholder="请输入搜索的用户名" v-model="name"></el-input>
            </el-form-item>
            <el-form-item label="所属属性:">
                <el-tree-select v-model="CateID" :data="CateListArr" check-strictly
                    :props="{ key: 'categoryId', label: 'name' }" node-key="id" :render-after-expand="false" />
            </el-form-item>
            <el-form-item label="类型:">
                <el-select v-model="type" class="m-2" placeholder="请选择类型">
                    <el-option label="属性" value="1" />
                    <el-option label="规格" value="2" />
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" size="default" @click="search"
                    :disabled="name.length || type.length || CateID ? false : true">
                    搜索
                </el-button>
                <el-button size="default" @click="reset">重置</el-button>
            </el-form-item>
        </el-form>
    </el-card>
    <!--  -->
    <el-card>
        <!-- 展示数据列表 -->
        <el-button type="success" size="default" icon="Plus" @click="add">添加</el-button>
        <el-table border style="margin:15px 0" :data="ListArr">
            <el-table-column label="编号" width="50px" align="center" prop="id"></el-table-column>
            <el-table-column label="属性名称" align="center" prop="name"></el-table-column>
            <el-table-column label="类型" align="center" prop="type">
                <template #="{ row }">
                    <template v-if="row.type === '1'">
                        <el-tag class="mx-1" type="success" effect="light">
                            属性
                        </el-tag>
                    </template>
                    <template v-if="row.type === '2'">
                        <el-tag class="mx-1" type="warning" effect="light">
                            规格
                        </el-tag>
                    </template>
                </template>
            </el-table-column>
            <el-table-column label="排序" align="center" prop="sort"></el-table-column>
            <el-table-column label="属性值" align="center" prop="value"></el-table-column>
            <el-table-column label="操作" align="center">
                <template #="{ row }">
                    <el-button type="primary" size="small" @click="update(row)" icon="Edit">编辑</el-button>
                    <el-popconfirm :title="`你确定删除${row.name}?`" width="200px" @confirm="deleteAttr(row.id)">
                        <template #reference>
                            <el-button type="danger" size="small" icon="Delete">删除</el-button>
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
    <!--  -->
    <el-dialog v-model="drawer">
        <el-form>
            <el-form-item label="所属分类">
                <el-tree-select v-model="Params.attributeCategoryID" :data="CateListArr" check-strictly
                    :props="{ key: 'attributeCategoryID', label: 'name' }" node-key="id" :render-after-expand="false" />
            </el-form-item>
            <el-form-item label="属性名称">
                <el-input placeholder="请输入属性名称" v-model="Params.name"></el-input>
            </el-form-item>
            <el-form-item label="属性值">
                <el-input placeholder="请输入属性值" v-model="Params.value"></el-input>
            </el-form-item>
            <el-form-item label="排序">
                <!-- <el-input placeholder="请输入排序" v-model="Params.sort"></el-input> -->
                排在第<el-input-number v-model="Params.sort" :min="1" size="small" controls-position="right" />位
            </el-form-item>
            <el-form-item label="类型">
                <el-select v-model="Params.type" class="m-2" placeholder="请选择类型">
                    <el-option label="属性" value="1" />
                    <el-option label="规格" value="2" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" size="default" @click="save">保存</el-button>
                <el-button type="primary" size="default" @click="cancel">取消</el-button>
            </span>
        </template>

    </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus';
//获取属性
import { reqAllAttribute, reqAddOrUpdate, reqRemove } from '@/api/pms/attribute'
import { reqAllattributeCategory } from '@/api/pms/attributeCategory'
//c存储已有的属性和属性值
let ListArr = ref<any[]>([])
//默认页码
let pageNo = ref<number>(1)
//默认个数
let pageSize = ref<number>(10)
let total = ref<number>(0)

//数据
let name = ref<string>('')
let type = ref<string>('')
let CateID = ref<any>()
// 
let drawer = ref<boolean>(false)
//分类
let CateListArr = ref<any>([]);
let ids = ref<number[]>([]);
//准备一个数组：将来存储对于的组件实例
//收集新增的数据
let Params = reactive<any>({
    id: 0,
    name: '',
    sort: 0,
    type: '',
    value: '',
    attributeCategoryID: 0,
})

//组件挂载完毕
onMounted(() => {
    getHas()
})
const getHas = async () => {
    let res: any = await reqAllAttribute(pageNo.value,
        pageSize.value,
        name.value,
        type.value,
        CateID.value,
    )
    if (res.code === 200) {
        ListArr.value = res.data
        total.value = res.total
    }
    let res1: any = await reqAllattributeCategory()
    if (res1.code === 200) {
        CateListArr.value = res1.data
    }
    ElMessage({ type: 'success', message: res.message })
    if (res.code === 400 || res1.code === 400) {
        ElMessage({ type: 'success', message: res.message })
    }
}
//下拉改变
const handler = () => {
    getHas()
}
// 
let add = () => {
    //清除数据 
    Object.assign(Params, {
        id: 0,
        name: '',
        sort: 0,
        type: '',
        value: '',
        attributeCategoryID: '',
    })
    drawer.value = true
}
//修改
let update = (row: any) => {
    drawer.value = true
    //将已有的属性对象给
    Object.assign(Params, JSON.parse(JSON.stringify(row)))
}


//取消按钮
let cancel = () => {
    drawer.value = false
}



//保存按钮
const save = async () => {
    Params.sort = parseInt(Params.sort)
    let res: any = await reqAddOrUpdate(Params)
    if (res.code === 200) {
        drawer.value = false
        getHas()
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}


//删除属性
let deleteAttr = async (id: number) => {
    ids.value.push(id);
    const requestData: any = { ids: ids.value };
    let res = await reqRemove(requestData);
    if (res.code === 200) {
        getHas()
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}

//搜索按钮
const search = () => {
    getHas()
    name.value = ''
    type.value = ''
}

//setting仓库
import useLayoutSettingStore from '@/store/setting'
let settingStore = useLayoutSettingStore()
//重置按钮
const reset = () => {
    settingStore.refresh = !settingStore.refresh
}
</script>

<style scoped lang="scss">
.form {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>