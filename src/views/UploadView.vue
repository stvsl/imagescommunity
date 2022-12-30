<template>
  <div class="uploadpanel">
    <el-form
      :model="formData"
      ref="vForm"
      :rules="rules"
      label-position="left"
      label-width="100px"
      size="default"
      @submit.prevent
    >
      <div class="table-container">
        <table class="table-layout">
          <tbody>
            <el-upload
              class="upload-demo"
              drag
              list-type="picture"
              :auto-upload="false"
              :limit="1"
              :on-change="handleChange"
              ref="imgupload"
            >
              <el-icon class="el-icon--upload"><upload-filled /></el-icon>
              <div class="el-upload__text">拖拽文件到此处或 <em>点击上传图片</em></div>
              <template #tip>
                <div class="el-upload__tip">仅支持png/jpg格式且图片大小不得大于20MB</div>
              </template>
            </el-upload>
            <el-divider border-style="dotted" style="height: 10px" />
            <el-form-item label="作品名称：" prop="name" class="required">
              <el-input v-model="formData.name" type="text" clearable></el-input>
            </el-form-item>
            <el-form-item label="作品简介：" prop="describe" class="required">
              <el-input
                type="textarea"
                v-model="formData.describe"
                rows="6"
                :show-word-limit="true"
                :maxlength="200"
              >
              </el-input>
            </el-form-item>
            <el-form-item label="作品标签：" prop="tag" class="required">
              <el-input v-model="formData.tag" type="text" clearable></el-input>
              <el-alert
                title="标签可以输入多个，按 ‘,’分割多个部分，最多支持10个，超过部分将自动删除"
                type="info"
                :closable="true"
                :show-icon="false"
                effect="light"
                style="height: 30px; margin-top: 5px"
              >
              </el-alert>
            </el-form-item>
            <br />
            <el-form-item>
              <el-button type="danger" @click="submitForm">提交</el-button>
              &nbsp;
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </tbody>
        </table>
      </div>
    </el-form>
  </div>
</template>

<style scoped>
.uploadpanel {
  width: 80%;
  margin: auto;
  margin-top: 40px;
  margin-bottom: 270px;
}

.table-layout {
  width: 50%;
  border-collapse: collapse;
  border-spacing: 0;
  margin: auto;
}
</style>

<script>
import { defineComponent } from "vue";
import { UploadFilled } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
export default defineComponent({
  name: "UploadView",
  components: {
    UploadFilled,
  },
  data() {
    return {
      formData: {
        picture: [],
        name: "",
        describe: "",
        tag: "",
      },
      rules: {
        picture: [
          {
            required: true,
            message: "图片不可为空",
          },
        ],
        name: [
          {
            required: true,
            message: "作品名不可为空",
          },
        ],
        describe: [
          {
            required: true,
            message: "简介不可为空",
          },
        ],
        tag: [
          {
            required: true,
            message: "标签不可为空",
          },
        ],
      },
    };
  },
  methods: {
    submitForm() {
      this.$refs["vForm"].validate((valid) => {
        if (!valid) return;
        const file = this.formData.picture[0];
        const reader = new FileReader();
        reader.readAsDataURL(file.raw);
        reader.onload = () => {
          fetch("http://127.0.0.1:8521/api/art/upload", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              name: this.formData.name,
              describe: this.formData.describe,
              tag: this.formData.tag,
              picture: reader.result,
              uuid: this.$store.state.userInfo.uuid,
            }),
          })
            .then((response) => response.json())
            .then((data) => {
              if (data.status === "success") {
                ElMessage({
                  message: "上传成功",
                  type: "success",
                });
                this.resetForm();
              }
            });
        };
      });
    },
    resetForm() {
      this.$refs.vForm.resetFields();
      this.$refs.imgupload.clearFiles();
    },
    handleChange(file, fileList) {
      this.formData.picture = fileList;
    },
  },
});
</script>
