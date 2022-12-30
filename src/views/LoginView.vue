<template>
  <div class="loginpanel">
    <div class="circle" style="--x: 0"></div>
    <div class="circle" style="--x: 1"></div>
    <div class="circle" style="--x: 2"></div>
    <div class="circle" style="--x: 3"></div>
    <div class="circle" style="--x: 4"></div>
    <h1
      style="
        color: var(--main-text-color-white);
        font-size: 5.5rem;
        font-family: logofont;
        text-shadow: var(--main-text-shadow-logo-big);
      "
    >
      图片分享社区
    </h1>
    <el-form
      :model="formData"
      ref="vForm"
      :rules="rules"
      label-position="center"
      label-width="100px"
      size="default"
      @submit.prevent
    >
      <div class="card-container">
        <el-card
          style="
            width: 560px;
            margin: 0 auto;
            margin-top: -120px;
            background-color: var(--main-theme-color-transparent);
            border: 0;
            backdrop-filter: blur(10px);
            border-radius: 30px;
          "
          shadow="always"
        >
          <template #header>
            <span
              style="
                color: var(--main-text-color-white);
                font-size: 1.8em;
                font-weight: 900;
                display: inline-block;
                text-align: left;
                text-shadow: var(--main-text-shadow);
              "
              >登 录</span
            >
          </template>
          <el-form-item class="input" label="电话号码：" prop="id">
            <el-input v-model="formData.id" type="text" clearable></el-input>
          </el-form-item>
          <el-form-item class="input" label="请输入密码:" prop="passwd">
            <el-input
              v-model="formData.passwd"
              type="password"
              show-password
              clearable
            ></el-input>
          </el-form-item>
          <div class="static-content-item">
            <el-button @click="submitForm">登录</el-button
            ><el-button @click="resetForm">清除</el-button>
            <br />
            <router-link
              to="/register"
              style="color: var(--main-text-color-shade); font-size: 0.9em"
            >
              <br />
              没有账号？立刻注册！
            </router-link>
          </div>
        </el-card>
      </div>
    </el-form>
  </div>
</template>
<script>
import { ElMessage } from "element-plus";
export default {
  name: "LoginView",
  created() {},

  data() {
    return {
      formData: {
        id: "",
        passwd: "",
      },
      rules: {
        id: [
          { required: true, message: "请输入ID", trigger: "blur" },
          { min: 11, max: 11, message: "长度在11个字符", trigger: "blur" },
        ],
        passwd: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, max: 50, message: "长度不符合要求", trigger: "blur" },
        ],
      },
    };
  },

  methods: {
    submitForm() {
      this.$refs["vForm"].validate((valid) => {
        if (!valid) return;
        fetch("http://127.0.0.1:8521/api/user/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            id: this.formData.id,
            passwd: this.formData.passwd,
          }),
        })
          .then((response) => response.json())
          .then((data) => {
            this.$store.commit("login", data);
            this.$cookies.set("username", data.username);
            this.$cookies.set("avatar", data.avatar);
            this.$cookies.set("id", data.id);
            this.$cookies.set("uuid", data.uuid);
            this.$cookies.set("sex", data.sex);
            this.$cookies.set("tel", data.tel);
            this.$cookies.set("arts", data.actnum);
            this.$router.push("/");
          });
        ElMessage({
          message: "登录成功，欢迎回来！",
          type: "success",
        });
      });
    },

    resetForm() {
      this.$refs["vForm"].resetFields();
    },
  },
};
</script>
<style scoped>
.loginpanel {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url(../assets/imgs/loginbg.jpg);
  background-size: cover;
  z-index: 100;
  display: flex;
  flex-direction: column;
}

/* 背景圆样式 */
.circle {
  position: fixed;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(5px);
  box-shadow: 0 25px 45px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.5);
  border-right: 1px solid rgba(255, 255, 255, 0.2);
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  filter: hue-rotate(calc(var(--x) * 70deg));
  animation: animate 25s linear infinite;
  animation-delay: calc(var(--x) * -1s);
  z-index: 101;
}

/* 背景圆动画 */
@keyframes animate {
  100% {
    transform: translateY(50px) translateX(0px);
  }
  75% {
    transform: translateY(40px) translateX(50px);
  }
  50% {
    transform: translateY(-50px) translateX(0px);
  }
  25% {
    transform: translateY(50px) translateX(-50px);
  }
  0% {
    transform: translateY(50px) translateX(0px);
  }
}

.circle:nth-child(1) {
  top: 50px;
  right: 160px;
  width: 120px;
  height: 120px;
}

.circle:nth-child(2) {
  top: 250px;
  left: 290px;
  width: 140px;
  height: 140px;
}

.circle:nth-child(3) {
  bottom: 160px;
  right: 280px;
  width: 120px;
  height: 120px;
}

.circle:nth-child(4) {
  bottom: 80px;
  left: 100px;
  width: 80px;
  height: 80px;
}

.circle:nth-child(5) {
  top: 80px;
  left: 140px;
  width: 80px;
  height: 80px;
}

.input {
  width: 80%;
  margin: 0 auto 40px auto;
  text-align: right;
  align-items: center;
}

.static-content-item {
  width: 80%;
  margin: 0 auto;
  text-align: center;
  align-items: center;
}

.static-content-item .el-button {
  width: 25%;
  height: 40px;
  margin: auto 10px auto 10px;
  text-align: center;
  align-items: center;
  background-color: var(--main-btn-bg-color);
  color: var(--main-text-color-white);
  border: 0;
  border-radius: 12px;
  font-size: 1.2em;
  font-weight: 900;
  opacity: 0.8;
}

.static-content-item .el-button:hover {
  opacity: 1;
}

.static-content-item .el-button:active {
  opacity: 0.8;
}

.el-input {
  width: 90%;
  height: 40px;
  opacity: 0.5;
}

.card-container {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
