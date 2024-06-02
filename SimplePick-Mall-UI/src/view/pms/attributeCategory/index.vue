<template>
    <el-card>
        <el-button type="success" size="default" icon="Plus" @click="add">
            添加
        </el-button>
        <!--  -->
        <el-table style="margin: 15px 0" row-key="id" border :data="ListArr">
            <el-table-column prop="name" label="名称" />
            <el-table-column label="操作" width="260px">
                <template #="{ row }">
                    <el-button type="primary" size="small" @click="update(row)">
                        编辑
                    </el-button>
                    <el-popconfirm :title="`你确定要删除${row.name}?`" width="260px" @confirm="remove(row.id)">
                        <template #reference>
                            <el-button type="danger" size="small">
                                删除
                            </el-button>
                        </template>
                    </el-popconfirm>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
    <!--  -->
    <el-dialog v-model="dialogVisible" :title="Data.id ? '更新' : '添加'">
        <el-form ref="formRef">
            <el-form-item label="名称">
                <el-input placeholder="请你输入菜单的名称" v-model="Data.name"></el-input>
            </el-form-item>
            <el-form-item label="父级id">
                <el-tree-select v-model="Data.parentID" :data="ListArrWithRoot" check-strictly
                    :props="{ key: 'id', label: 'name' }" node-key="id" :render-after-expand="false" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="save">确定</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, computed } from 'vue';
import { ElTreeSelect, ElMessage } from 'element-plus';
import {
    reqAllattributeCategory,
    reqAddOrUpdate,
    reqRemove
} from '@/api/pms/attributeCategory';
//组件实例
let formRef = ref<any>()
let ListArr = ref<any>([]);
let dialogVisible = ref<boolean>(false);
let ids = ref<number[]>([]);

let Data = reactive<any>({
    id: 0,
    name: '',
    parentID: 0,
});

onMounted(() => {
    getHas();
});
const getHas = async () => {
    let res: any = await reqAllattributeCategory();
    if (res.code === 200) {
        ListArr.value = res.data;
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
};
const ListArrWithRoot = computed(() => {
    return [
        { id: 0, name: '无上级' },
        ...ListArr.value,
    ];
});

const add = () => {
    Object.assign(Data, {
        name: '',
        parentID: 0,
    });
    dialogVisible.value = true;
};

const update = (row: any) => {
    dialogVisible.value = true;
    Data.parentID = parseInt(row.parentID);
    Object.assign(Data, row);
};
const save = async () => {
    dialogVisible.value = false;
    let res: any = await reqAddOrUpdate(Data);
    if (res.code === 200) {
        dialogVisible.value = false;
        getHas();
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
};

const remove = async (id: number) => {
    ids.value.push(id);
    const requestData: any = { ids: ids.value };
    let res = await reqRemove(requestData);
    if (res.code == 200) {
        ids.value = [];
        getHas();
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
};
</script>

<style scoped></style>