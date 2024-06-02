<template>
  <div class="login_container">
    <el-row>
      <!--  -->
      <el-col :span="16"></el-col>
      <!-- 登录框 -->
      <el-col :span="8">
        <el-form
          class="login_form"
          :model="loginForm"
          :rules="rules"
          ref="loginForms"
        >
          <h1>欢迎您</h1>
          <h2>感谢使用定点帮扶农产品直销系统</h2>
          <el-form-item prop="username">
            <el-input :prefix-icon="User" v-model="loginForm.username">
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              type="password"
              :prefix-icon="Lock"
              v-model="loginForm.password"
              show-password
            ></el-input>
          </el-form-item>
          <el-button
            class="login_btn"
            type="primary"
            size="default"
            @click="login"
            >登录</el-button
          >
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { User, Lock } from "@element-plus/icons-vue";
import { reactive, ref } from "vue";
//表单数据
let loginForm = reactive({
  username: "admin",
  password: "123456",
});
//获取loginForm
let loginForms = ref();

//定义表单校验
const rules = {
  username: [
    { required: true, message: "请输入登陆名称", trigger: "blur" },
    { min: 3, max: 20, message: "长度在 3 到 10 个字符", trigger: "blur" },
  ],
  password: [
    { required: true, message: "请输入登陆密码", trigger: "blur" },
    { min: 6, max: 20, message: "长度在 6 到 15 个字符", trigger: "blur" },
  ],
};

//获取路由器
import { useRouter } from "vue-router";
let router = useRouter();

import { ElNotification } from "element-plus";

//获取当前时间段工具
import { getTime } from "@/util/time";

//存放用户数据 store
import useUserStore from "@/store/sys/user";
let userStore = useUserStore();

//登录按钮
const login = async () => {
  //保证全部表单校验通过再发请求
  await loginForms.value.validate();
  //   try {
  //保证登录成功
  let res = await userStore.userLogin(loginForm);
  let res2 = await userStore.userInfo();

  if (res == "ok") {
    if (res2 == "ok") {
      ElNotification({
        type: "success",
        message: "欢迎回来",
        title: `Hi, ${getTime()}好`,
      });
      //跳转首页
      router.push("/");
    } else {
      ElNotification({
        type: "error",
        message: res2,
        title: `Hi, ${getTime()}好`,
      });
    }
  } else {
    ElNotification({
      type: "error",
      message: res,
      title: `Hi, ${getTime()}好`,
    });
  }
};
</script>

<style scoped lang="scss">
.login_container {
  width: 100%;
  height: 100vh;
  background: url("@/assets/images/666.png") no-repeat;
  background-size: cover;

  .login_form {
    position: relative;
    width: 80%;
    top: 30vh;
    background: url("@/assets/images/login_form.png") no-repeat;
    background-size: cover;
    padding: 40px;

    h1 {
      color: rgb(83, 79, 79);
      font-size: 30px;
    }

    h2 {
      color: rgb(83, 78, 78);
      font-size: 20px;
      margin: 20px 0px;
    }

    .login_btn {
      width: 100%;
    }
  }
}
</style>
