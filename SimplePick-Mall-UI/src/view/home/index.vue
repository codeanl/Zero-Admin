<script setup lang="ts">
// 组件引用
import GithubCorner from "@/components/GithubCorner/index.vue";
import SvgIcon from "@/components/SvgIcon/index.vue";
import { getTime } from "@/util/time";
import useUserStore from "@/store/sys/user";
let userStore = useUserStore();
import Echarts from "./components/echarts.vue";
</script>

<template>
  <!-- 欢迎 -->
  <div>
    <el-card>
      <div class="box">
        <img :src="userStore.avatar" alt="" class="avatar" />
        <div class="footer">
          <h3 class="title">
            {{ getTime() }}好~{{ userStore.username }}，祝你开心每一天！
          </h3>
          <p class="subtitle">今天依旧是上班的好日子呀！</p>
        </div>
      </div>
    </el-card>
  </div>
  <!--  -->
  <div class="dashboard-container">
    <github-corner class="github-corner" />
    <!-- 数据 -->
    <el-row :gutter="40" class="card-panel__col">
      <el-col :xs="24" :sm="12" :lg="6" class="card-panel__col">
        <div class="card-panel">
          <div class="card-panel-icon-wrapper icon-people">
            <svg-icon icon-class="uv" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">访问数</div>
            <div class="card-panel-num">1000</div>
          </div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6" class="card-panel__col">
        <div class="card-panel">
          <div class="card-panel-icon-wrapper icon-message">
            <svg-icon icon-class="message" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">消息数</div>
            <div class="card-panel-num">1000</div>
          </div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6" class="card-panel__col">
        <div class="card-panel">
          <div class="card-panel-icon-wrapper icon-money">
            <svg-icon icon-class="money" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">收入金额</div>
            <div class="card-panel-num">1000</div>
          </div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6" class="card-panel__col">
        <div class="card-panel">
          <div class="card-panel-icon-wrapper icon-shopping">
            <svg-icon icon-class="shopping" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">订单数</div>
            <div class="card-panel-num">1000</div>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
  <!-- echarts -->
  <div class="echarts-container">
    <Echarts class="echarts"></Echarts>
    <Echarts class="echarts"></Echarts>
  </div>
</template>

<style lang="scss" scoped>
.box {
  display: flex;

  .avatar {
    width: 60px;
    height: 60px;
    border-radius: 50%;
  }

  .footer {
    margin-left: 20px;
    margin-top: 10px;

    .title {
      font-size: 25px;
      margin-bottom: 20px;
      font-weight: 900;

      .gradient {
        background-clip: text;
        -webkit-background-clip: text;
        /* 兼容WebKit浏览器（例如Chrome和Safari） */
        color: transparent;
        font-size: 24px;
        font-weight: bold;
      }
    }

    .subtitle {
      font-style: italic;
      color: #ccc;
      font-weight: 700;
    }
  }
}

.dashboard-container {
  padding: 24px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
    z-index: 99;
  }

  .box-center {
    margin: 0 auto;
    display: table;
  }

  .user-profile {
    .box-center {
      padding-top: 10px;
    }

    .user-role {
      padding-top: 10px;
      font-weight: 400;
      font-size: 14px;
    }

    .box-social {
      padding-top: 30px;

      .el-table {
        border-top: 1px solid #dfe6ec;
      }
    }

    .user-follow {
      padding-top: 20px;
    }
  }

  .card-panel__col {
    margin-bottom: 12px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);

    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }

      .icon-user {
        background: #e77541;
      }

      .icon-people {
        background: #40c9c6;
      }

      .icon-message {
        background: #36a3f7;
      }

      .icon-money {
        background: #f4516c;
      }

      .icon-shopping {
        background: #34bfa3;
      }
    }

    .icon-people {
      color: #40c9c6;

      .svg-icon {
        width: 4em !important;
        height: 4em !important;
      }
    }

    .icon-message {
      color: #36a3f7;

      .svg-icon {
        width: 4em !important;
        height: 4em !important;
      }
    }

    .icon-money {
      color: #f4516c;

      .svg-icon {
        width: 4em !important;
        height: 4em !important;
      }
    }

    .icon-shopping {
      color: #34bfa3;

      .svg-icon {
        width: 4em !important;
        height: 4em !important;
      }
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px 20px 0;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
        text-align: right;
      }
    }
  }

  .chart-container {
    background: #ffffff;
  }
}

.echarts-container {
  display: flex;

  .echarts {
    flex: 1;
  }
}
</style>
