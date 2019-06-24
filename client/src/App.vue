

<template>
  <div id="app">
    <!-- main -->
    <Container>
      <header class="header limit-width">
        <el-button
          type="text"
          v-if="!profile.user_name && routesName !== 'login'"
          @click="openLoginDialog"
        >登录</el-button>
        <el-button
          type="text"
          v-if="profile.user_name && routesName !== 'login'"
          @click="logoutDialogShow = true"
        >{{profile.user_name}}</el-button>
      </header>
      <transition name="fade" mode="out-in">
        <keep-alive>
          <router-view @not-login="openLoginDialog"></router-view>
        </keep-alive>
      </transition>
      <router-view name="sidebar"></router-view>
    </Container>
    <!-- main:end -->
    <!-- login-dialog -->
    <el-dialog title="登录" :close-on-click-modal="false" :visible.sync="loginDialogShow" width="27%" center>
      <div class="login-container">
        <el-form :model="loginForm" ref="loginForm" label-width="70px">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="loginForm.username" placeholder="请输入用户名"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="username">
            <el-input v-model="loginForm.password" type="password" placeholder="请输入密码"></el-input>
            <!-- style="width: 420px" -->
          </el-form-item>
        </el-form>
        <div
          :class="['login-error-message', {
          'show': !!loginError
        }]"
        >{{loginError}}</div>
      </div>
      <span slot="footer" class="dialog-footer">
        <div>
          <el-button @click="loginDialogShow = false">取 消</el-button>
          <el-button type="primary" @click="login">确 定</el-button>
        </div>
      </span>
    </el-dialog>
    <!-- login-dialog:end -->
    <!-- logout-dalog -->
    <el-dialog title="退出登录"  :visible.sync="logoutDialogShow" width="27%" center>
      <div class="login-container">
        是否退出登录?
      </div>
      <span slot="footer" class="dialog-footer">
        <div>
          <el-button @click="logoutDialogShow = false">取 消</el-button>
          <el-button type="primary" @click="logout">确 定</el-button>
        </div>
      </span>
    </el-dialog>
    <!-- logout-dalog:end -->
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      profile: {},
      // 登录对话框显示
      loginDialogShow: false,
      // 登出对话框显示
      logoutDialogShow: false,
      // 登录信息
      loginForm: {
        username: "",
        password: ""
      },
      loginError: ""
    };
  },
  watch: {
    loginForm: {
      handler(v) {
        this.loginError = "";
      },
      deep: true
    }
  },
  computed: {
    routesName() {
      return this.$route.name;
    }
  },
  created() {
    this.getProfile();
  },
  methods: {
    goBack() {
      window.history.length > 1 ? this.$router.go(-1) : this.$router.push("/");
    },
    // 获取用户数据
    getProfile() {
      this.$api
        .get("/user/profile")
        .then(res => {
          let data = res.data;
          this.profile = data;
        })
        .catch(err => {
          this.openLoginDialog();
        });
    },
    // 打开登录弹窗
    openLoginDialog() {
      this.loginDialogShow = true;
    },
    // 登录
    login() {
      if (!this.loginForm.username) {
        this.loginError = "请填写用户名";
        return;
      }
      if (!this.loginForm.password) {
        this.loginError = "请填写密码";
        return;
      }

      this.$api
        .post("/user/login", {
          user_name: this.loginForm.username,
          password: this.loginForm.password
        })
        .then(resp => {
          this.$router.go()
        })
        .catch(err => {
          console.log("loginErr", JSON.stringify(err));
          if (!err.response) {
            this.loginError = "请求失败， 服务器错误";
            return;
          }
          if (err.response.data && err.response.data.message) {
            this.loginError = err.response.data.message;
            return;
          }
        });
    },
    logout() {
      this.$api.post('/user/logout')
      this.$router.go()
      this.logoutDialogShow = false
    }
  }
};
</script>

<style lang="stylus">

#app
  box-sizing border-box
  font-family Roboto
  .header
    box-sizing border-box 
    width 100%
    padding 16px 32px
    margin 0
    text-align right
  .login-error-message    
    padding-left 70px  
    color #f56c6c
    overflow hidden
    height 0
    transition height .1s
    &.show
      height 18px
  .dialog-footer
    text-align right


</style>


