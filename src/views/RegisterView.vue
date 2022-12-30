<template>
  <div class="registpanel">
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
            margin-top: -20px;
            background-color: rgba(255, 255, 255, 0.2);
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
              >注册</span
            >
          </template>
          <el-form-item label="头像" prop="avatar" class="required">
            <el-upload
              :file-list="avatarFileList"
              :headers="avatarUploadHeaders"
              list-type="picture-card"
              show-file-list
              :limit="1"
              :auto-upload="false"
              :on-change="handleChange"
              ref="avatarUpload"
              style="margin-left: 20px"
            >
              <el-icon class="el-icon-plus"> <Plus /> </el-icon
            ></el-upload>
          </el-form-item>
          <el-form-item label="昵称" prop="name" class="required">
            <el-input
              v-model="formData.name"
              class="input"
              type="text"
              clearable
            ></el-input>
          </el-form-item>
          <el-form-item label="性别" prop="sex" class="required">
            <el-radio-group v-model="formData.sex">
              <el-radio
                v-for="(item, index) in sexOptions"
                :key="index"
                :label="item.value"
                :disabled="item.disabled"
                style="display: inline; margin-left: 20px; margin-right: 0"
                >{{ item.label }}</el-radio
              >
            </el-radio-group>
          </el-form-item>
          <el-form-item label="电话号码" prop="tel" class="required">
            <el-input
              v-model="formData.tel"
              class="input"
              type="text"
              clearable
            ></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="passwd" class="required">
            <el-input
              v-model="formData.passwd"
              type="password"
              :show-password="true"
              class="input"
              clearable
            ></el-input>
          </el-form-item>
          <el-form-item label="确认密码" prop="passwd2" class="required">
            <el-input
              v-model="formData.passwd2"
              type="password"
              :show-password="true"
              class="input"
              clearable
            ></el-input>
          </el-form-item>
          <div class="static-content-item">
            <el-button
              @click="
                () => {
                  submitForm();
                }
              "
              >注册</el-button
            ><el-button
              @click="
                () => {
                  $refs.vForm.resetFields();
                }
              "
              >清除</el-button
            >
            <br />
            <router-link
              to="/login"
              style="color: var(--main-text-color-shade); font-size: 0.9em"
            >
              <br />
              已有账号？前往登录！
            </router-link>
          </div>
        </el-card>
      </div>
    </el-form>
  </div>
</template>
<script>
import { Plus } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";

export default {
  name: "LoginView",
  created() {
    if (this.$store.state.userInfo.isLogin) {
      this.$router.push("/");
      alert("您已经登录，无需再次登录" + this.$store.state.userInfo.isLogin);
    }
  },
  components: {
    Plus,
  },
  data() {
    return {
      formData: {
        avatar: [],
        name: "",
        sex: "2",
        tel: "",
        passwd: "",
        passwd2: "",
      },
      fileList: [],
      rules: {
        avatar: [
          {
            required: true,
            message: "头像不可为空",
          },
        ],
        name: [
          {
            required: true,
            message: "昵称不可为空",
          },
        ],
        sex: [
          {
            required: true,
            message: "性别不可为空",
          },
        ],
        tel: [
          {
            required: true,
            message: "密码不可为空",
          },
        ],
        passwd: [
          {
            required: true,
            message: "密码不可为空",
          },
        ],
        passwd2: [
          {
            required: true,
            message: "两次输入的密码不一致",
            validator: (rule, value, callback) => {
              if (value !== this.formData.passwd) {
                callback(new Error("两次输入的密码不一致"));
              } else {
                callback();
              }
            },
          },
        ],
      },
      sexOptions: [
        {
          label: "男",
          value: "0",
        },
        {
          label: "女",
          value: "1",
        },
        {
          label: "保密",
          value: "2",
        },
      ],
    };
  },
  methods: {
    submitForm() {
      this.$refs["vForm"].validate((valid) => {
        if (!valid) return;
        // 读取文件
        const file = this.formData.avatar[0];
        // 获取文件的 base64 格式(可发送给服务器的)
        const reader = new FileReader();
        reader.readAsDataURL(file.raw);
        reader.onload = () => {
          // 向服务器发送POST请求
          fetch("http://127.0.0.1:8521/api/user/register", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              name: this.formData.name,
              sex: this.formData.sex,
              tel: this.formData.tel,
              passwd: this.formData.passwd,
              avatar: reader.result,
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
        };
        ElMessage({
          message: "注册成功,自动登录",
          type: "success",
        });
      });
    },
    resetForm() {
      this.$refs.vForm.resetFields();
      this.$refs.avatarUpload.clearFiles();
    },
    handleChange(file, fileList) {
      this.formData.avatar = fileList;
    },
  },
};
</script>
<style scoped>
.registpanel {
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
  margin: 0 auto 0 auto;
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
