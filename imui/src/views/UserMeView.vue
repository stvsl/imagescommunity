<template>
  <div style="display: flex; justify-content: left; text-align: left">
    <el-form
      :model="formData"
      ref="vForm"
      :rules="rules"
      label-position="left"
      label-width="150px"
      size="default"
    >
      <div class="card-container">
        <el-card style="width: 80%" shadow="always">
          <template #header>
            <div class="clear-fix">
              <span>个人信息</span>
            </div>
          </template>
          <el-form-item label="我的头像" prop="avatar" class="required label-right-align">
            <el-upload
              :file-list="avatarFileList"
              :file-data="fileData"
              list-type="picture-card"
              show-file-list
              :limit="1"
              accept="image/*"
              :auto-upload="false"
              :on-change="handleChange"
            >
              <template #default>
                <el-icon class="el-icon-plus"> <Plus /> </el-icon> </template
            ></el-upload>
          </el-form-item>
          <el-row>
            <el-col :span="12" class="grid-cell">
              <el-form-item label="用户名" prop="name" class="required label-right-align">
                <el-input v-model="formData.name" type="text" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12" class="grid-cell"> </el-col>
            <el-col :span="12" class="grid-cell">
              <el-form-item label="性别" prop="sex" class="required label-right-align">
                <el-radio-group v-model="formData.sex">
                  <el-radio
                    v-for="(item, index) in sexOptions"
                    :key="index"
                    :label="item.value"
                    :disabled="item.disabled"
                    style="display: inline"
                    >{{ item.label }}</el-radio
                  >
                </el-radio-group>
              </el-form-item>
            </el-col>
            <el-col :span="12" class="grid-cell"> </el-col>
            <el-col :span="12" class="grid-cell">
              <el-form-item
                label="电话号码"
                prop="tel"
                class="required label-right-align"
              >
                <el-input v-model="formData.tel" type="text" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12" class="grid-cell"> </el-col>
            <el-col :span="12" class="grid-cell">
              <el-row
                ><el-col :span="6" class="grid-cell">
                  <div class="static-content-item">
                    <el-button type="danger" @click="submit">保存更改</el-button>
                  </div>
                </el-col>
                <el-col :span="6" class="grid-cell">
                  <div class="static-content-item">
                    <el-button type="danger" @click="logout">退出登录</el-button>
                  </div>
                </el-col>
                <el-col :span="6" class="grid-cell">
                  <div class="static-content-item">
                    <el-button type="primary" @click="dialogVisible = true"
                      >修改密码</el-button
                    >
                  </div>
                </el-col>
              </el-row>
            </el-col>
          </el-row>
        </el-card>
        <el-dialog v-model="dialogVisible" title="修改密码" width="30%">
          <el-form
            :model="passwdData"
            ref="vForm"
            :rules="rules"
            label-position="left"
            label-width="150px"
            size="default"
          >
            <el-form-item
              label="新密码"
              prop="newPassword"
              class="required label-right-align"
            >
              <el-input
                v-model="passwdData.newPassword"
                type="password"
                clearable
              ></el-input>
            </el-form-item>
            <el-form-item
              label="确认密码"
              prop="confirmPassword"
              class="required label-right-align"
            >
              <el-input
                v-model="passwdData.confirmPassword"
                type="password"
                clearable
              ></el-input>
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button type="primary" @click="submitpass">保存更改</el-button>
          </template>
        </el-dialog>
      </div>
    </el-form>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Plus } from "@element-plus/icons-vue";

export default defineComponent({
  components: {
    Plus,
  },
  props: {},
  data() {
    return {
      formData: {
        avatar: "",
        name: this.$store.state.userInfo.username,
        sex: this.$store.state.userInfo.sex,
        tel: this.$store.state.userInfo.tel,
      },
      passwdData: {
        newPassword: "",
        confirmPassword: "",
      },
      rules: {
        name: [
          {
            required: true,
            message: "字段值不可为空",
          },
        ],
        sex: [
          {
            required: true,
            message: "字段值不可为空",
          },
        ],
        tel: [
          {
            required: true,
            message: "字段值不可为空",
          },
          {
            pattern: /^[1][3-9][0-9]{9}$/,
            trigger: ["blur", "change"],
            message: "请输入正确的电话号码",
          },
        ],
      },
      sexOptions: [
        {
          label: "男",
          value: 0,
        },
        {
          label: "女",
          value: 1,
        },
        {
          label: "保密",
          value: 2,
        },
      ],
      avatarFileList: [],
      dialogVisible: false,
    };
  },
  methods: {
    submit() {
      this.$refs.vForm.validate((valid) => {
        if (valid) {
          const file = this.formData.avatar[0];
          // 获取文件的 base64 格式(可发送给服务器的)
          const reader = new FileReader();
          reader.readAsDataURL(file.raw);
          reader.onload = () => {
            fetch("http://127.0.0.1:8521/api/user/updateuser", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                Uuid: this.$store.state.userInfo.uuid,
                Name: this.formData.name,
                Sex: this.formData.sex,
                Tel: this.formData.tel,
                Avatar: reader.result,
              }),
            })
              .then((response) => response.json())
              .then((data) => {
                if (data.status === "success") {
                  this.$message.success("提交成功");
                }
              });
          };
        } else {
          this.$message.error("提交失败");
        }
      });
    },
    logout() {
      this.$cookies.remove("uuid");
      this.$router.push("/login");
    },
    handleChange(file, fileList) {
      this.formData.avatar = fileList;
    },
    submitpass() {
      if (this.passwdData.newPassword === this.passwdData.confirmPassword) {
        fetch("http://127.0.0.1:8521/api/user/updatepasswd", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            Uuid: this.$store.state.userInfo.uuid,
            Passwd: this.passwdData.newPassword,
          }),
        })
          .then((response) => response.json())
          .then((data) => {
            if (data.status === "success") {
              this.$message.success("提交成功");
            }
          });
      } else {
        this.$message.error("两次密码不一致");
      }
      this.dialogVisible = false;
    },
  },
});
</script>

<style scoped>
.card-container {
  margin: 40px 30% 20% 30%;
  width: 100%;
  min-width: 1400px;
}
</style>
