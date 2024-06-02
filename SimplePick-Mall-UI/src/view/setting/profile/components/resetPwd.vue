<template>
    <el-form ref="pwdRef" :model="Pwd" :rules="rules" label-width="80px">
        <el-form-item label="旧密码" prop="oldPassword">
            <el-input v-model="Pwd.oldPassword" placeholder="请输入旧密码" type="password" show-password />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
            <el-input v-model="Pwd.newPassword" placeholder="请输入新密码" type="password" show-password />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
            <el-input v-model="Pwd.confirmPassword" placeholder="请确认新密码" type="password" show-password />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submit">保存</el-button>
            <el-button type="danger" @click="close">关闭</el-button>
        </el-form-item>
    </el-form>
</template>
 
<script setup lang="ts">
import { ref, reactive } from 'vue'
import { reqUpdatePwd } from '@/api/sys/user'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus';
let $router = useRouter()

let pwdRef = ref<any>()
const Pwd = reactive({
    oldPassword: '',
    newPassword: '',
    confirmPassword: ''
});

const equalToPassword = (rule: any, value: any, callback: any) => {
    console.log(rule)
    if (Pwd.newPassword !== value) {
        callback(new Error("两次输入的密码不一致"));
    } else {
        callback();
    }
};
const rules = ref({
    oldPassword: [{ required: true, message: "旧密码不能为空", trigger: "blur" }],
    newPassword: [{ required: true, message: "新密码不能为空", trigger: "blur" }, { min: 6, max: 20, message: "长度在 6 到 20 个字符", trigger: "blur" }],
    confirmPassword: [{ required: true, message: "确认密码不能为空", trigger: "blur" }, { required: true, validator: equalToPassword, trigger: "blur" }]
});
/** 提交按钮 */
let submit = async () => {
    await pwdRef.value.validate()
    let res = await reqUpdatePwd(Pwd)
    if (res.code == 200) {
        $router.push({ path: '/login' })
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
};
/** 关闭按钮 */
function close() {
    alert("off")
};
</script>