<template>
    <div class="app-container">
        <el-row>
            <!-- <el-row :gutter="20"> -->
            <el-col :span="8" :xs="24">
                <el-card class="box-card">
                    <el-descriptions direction="horizontal" :column="1" :size="size" border title="个人信息">
                        <el-descriptions-item>
                            <el-upload class="avatar-uploader" action="/api/api/sys/upload" :show-file-list="false"
                                :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
                                <img v-if="Params.avatar" :src="Params.avatar" class="avatar" />
                                <el-icon v-else class="avatar-uploader-icon">
                                    <Plus />
                                </el-icon>
                            </el-upload>
                        </el-descriptions-item>
                        <el-descriptions-item label="用户名">{{ info?.userInfo?.username }}</el-descriptions-item>
                        <el-descriptions-item label="昵称">{{ info?.userInfo?.nickname }}</el-descriptions-item>
                        <el-descriptions-item label="电话">{{ info?.userInfo?.phone }}</el-descriptions-item>
                        <el-descriptions-item label="邮箱">{{ info?.userInfo?.email }}</el-descriptions-item>
                        <el-descriptions-item label="角色">
                            <el-tag v-for="item in info?.roles" :key="item.label" type="success" class="mx-2"
                                style=" margin-right: 5px;" effect="dark">
                                {{ item }}
                            </el-tag>
                        </el-descriptions-item>

                    </el-descriptions>
                </el-card>
            </el-col>
            <el-col :span="16" :xs="24">
                <el-card>
                    <template v-slot:header>
                        <div class="clearfix">
                            <span>基本资料</span>
                        </div>
                    </template>
                    <el-tabs v-model="activeTab">
                        <el-tab-pane label="基本资料" name="userinfo">
                            <userInfo :user="info?.userInfo" />
                        </el-tab-pane>
                        <el-tab-pane label="修改密码" name="resetPwd">
                            <resetPwd />
                        </el-tab-pane>
                    </el-tabs>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>
 
<script setup  lang="ts">
import { ref, onMounted, reactive } from 'vue'
import userInfo from "./components/userInfo.vue";
import resetPwd from "./components/resetPwd.vue";
import { reqUserInfo } from '@/api/user'
import { reqAddOrUpdateUser } from '@/api/sys/user'
import { ElMessage, UploadProps } from 'element-plus';
const size = ref('')
const activeTab = ref("userinfo");
//组件挂载完毕
onMounted(() => {
    getHas()
})
//获取信息
let info = ref<any>(null)
let Params = reactive<any>({
    id: 0,
    avatar: ''
})
const getHas = async () => {
    let res: any = await reqUserInfo()
    if (res.code == 200) {
        info.value = res.data
        Params.id = res.data.userInfo.id
        Params.avatar = res.data.userInfo.avatar
    }
}

let updateAvatar = async () => {
    let res = await reqAddOrUpdateUser(Params)
    if (res.code == 200) {
        window.location.reload()
        ElMessage({ type: 'success', message: res.message })
    } else {
        ElMessage({ type: 'error', message: res.message })
    }
}


//图片上传成功的钩子
const handleAvatarSuccess: UploadProps['onSuccess'] = (response: any) => {
    //上传返回的数据 图片url  uploadFile
    Params.avatar = response.data;
    updateAvatar()
}
//上传图片之前出发的钩子函数
const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile: any) => {
    if (
        rawFile.type === 'image/png' ||
        rawFile.type === 'image/jpeg' ||
        rawFile.type === 'image/gif'
    ) {
        if (rawFile.size / 1024 / 1024 < 4) {
            return true
        } else {
            ElMessage({
                type: 'error',
                message: '上传的文件大小应小于4M',
            })
        }
    } else {
        ElMessage({
            type: 'error',
            message: '上传的文件类型必须是PNG|JPG|GIF',
        })
        return false
    }
}
</script>
<style>
.avatar-uploader .avatar {
    width: 128px;
    height: 128px;
    display: block;
    border-radius: 50%;

}

.avatar-uploader .el-upload {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
    border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
}
</style>