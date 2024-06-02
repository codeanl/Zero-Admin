<template>
    <el-form ref="userRef" :model="userModel" label-width="80px">
        <el-form-item label="昵称" prop="nickname">
            <el-input v-model="userModel.nickname" maxlength="30" />
        </el-form-item>
        <el-form-item label="手机" prop="phone">
            <el-input v-model="userModel.phone" maxlength="30" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
            <el-input v-model="userModel.email" maxlength="30" />
        </el-form-item>
        <el-form-item label="性别" prop="gender">
            <el-radio-group v-model="userModel.gender">
                <el-radio label="0">保密</el-radio>
                <el-radio label="1">男</el-radio>
                <el-radio label="2">女</el-radio>
            </el-radio-group>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submit">保存</el-button>
        </el-form-item>
    </el-form>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue';
import { ElMessage } from 'element-plus';
import { reqAddOrUpdateUser } from '@/api/sys/user'
const props: any = defineProps({
    user: {
        type: Object,
    },

});
const userModel = reactive({
    nickname: '',
    phone: '',
    email: '',
    gender: ''
});
watch(() => props.user, (newUser) => {
    Object.assign(userModel, newUser);
});

let submit = async () => {
    let res = await reqAddOrUpdateUser(userModel)
    if (res.code == 200) {
        window.location.reload()
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}
</script>