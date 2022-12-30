<template>
  <div>
    <div class="artmgr">
      <el-row v-for="o in arts" :key="o"
        ><el-card style="width: 100%; margin: 5px" :body-style="{ padding: '0px' }"
          ><el-row class="card"
            ><el-col :span="2"
              ><el-image
                style="width: 100px; height: 100px; margin: 10px"
                :src="o.Uri"
                :zoom-rate="1.2"
                :preview-src-list="[o.Uri]"
                :initial-index="4"
                fit="cover"
            /></el-col>
            <el-col :span="19"
              ><el-row style="text-align: left; margin-left: 10px"
                ><el-col :span="20"
                  ><h3 style="margin: 10px 0 6px 0; color: var(--main-text-color)">
                    {{ o.Name }}
                  </h3></el-col
                ><el-col :span="20" style="font-size: 12px"
                  ><font-awesome-icon icon="fa-solid fa-eye" />&nbsp; {{ o.Views }} &nbsp;
                  <font-awesome-icon icon="fa-solid fa-heart" />&nbsp;
                  {{ o.Likes }} &nbsp; <font-awesome-icon icon="fa-solid fa-star" />&nbsp;
                  {{ o.Collects }} &nbsp;</el-col
                >
                <el-col
                  style="
                    text-align: left;
                    font-size: 14px;
                    -webkit-line-clamp: 2;
                    -webkit-box-orient: vertical;
                    overflow: hidden;
                  "
                  ><p>{{ o.Describe }}</p></el-col
                >
              </el-row>
            </el-col>
            <el-col :span="2" style="margin: auto">
              <el-row>
                <el-button
                  class="button"
                  style="margin: 8px; width: 100px"
                  color="var(--main-btn-bg-color)"
                  @click="editit(o)"
                  >编 辑</el-button
                >
              </el-row>
              <el-row>
                <el-button
                  class="button"
                  style="margin: 8px; width: 100px"
                  color="var(--main-btn-bg-color)"
                  @click="deleteit(o)"
                  >删 除</el-button
                >
              </el-row>
            </el-col>
          </el-row></el-card
        >
      </el-row>
    </div>
    <div>
      <el-card>
        <el-dialog v-model="dialogVisible" title="作品修改">
          <el-form>
            <el-form-item label="作品名称">
              <el-input placeholder="请输入作品名称" v-model="editedname" />
            </el-form-item>
            <el-form-item label="作品简介">
              <el-input
                type="textarea"
                :rows="5"
                placeholder="请输入作品简介"
                v-model="editeddescribe"
              />
            </el-form-item>
          </el-form>
          <template #footer>
            <span class="dialog-footer">
              <el-button @click="dialogVisible = false">取消</el-button>
              <el-button type="primary" @click="updateit"> 确定 </el-button>
            </span>
          </template>
        </el-dialog>
      </el-card>
    </div>
  </div>
</template>
<script>
import { ElMessage } from "element-plus";
export default {
  name: "ArtsManageView",

  created() {
    // 延迟加载
    setTimeout(() => {
      fetch("http://127.0.0.1:8521/api/art/manage", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          uuid: this.$store.state.userInfo.uuid,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            console.log(data.arts);
            this.arts = JSON.parse(data.arts);
          }
        });
    }, 300);
  },
  data() {
    return {
      dialogVisible: false,
      arts: [],
      editedname: "",
      editeddescribe: "",
      editedid: "",
    };
  },
  methods: {
    editit(o) {
      this.editedname = o.Name;
      this.editeddescribe = o.Describe;
      this.editedid = o.ImgId;
      this.dialogVisible = true;
    },
    updateit() {
      fetch("http://127.0.0.1:8521/api/art/update", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Id: this.editedid,
          Name: this.editedname,
          Describe: this.editeddescribe,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            ElMessage({
              message: "修改成功",
              type: "success",
            });
            this.dialogVisible = false;
            setTimeout(() => {
              window.location.reload();
            }, 3000);
          } else {
            ElMessage({
              message: "修改失败",
              type: "error",
            });
          }
        });
    },
    deleteit(o) {
      fetch("http://127.0.0.1:8521/api/art/delete", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Id: o.ImgId,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            ElMessage({
              message: "修改成功",
              type: "success",
            });
            this.dialogVisible = false;
            setTimeout(() => {
              window.location.reload();
            }, 3000);
          } else {
            ElMessage({
              message: "修改失败",
              type: "error",
            });
          }
        });
    },
  },
  components: {},
};
</script>
<style scoped>
.artmgr {
  padding: 20px;
  width: 70%;
  min-height: 700px;
  align-items: center;
  margin: auto;
  justify-content: center;
  border-radius: 15px;
}

.card {
  width: 100%;
  height: 120px;
}
</style>
