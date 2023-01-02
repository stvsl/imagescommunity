<template>
  <section>
    <div
      style="width: 90%; border-radius: 15px; box-shadow: var(--main-box-shadow-light)"
    >
      <div class="leavecomment">
        <el-avatar :src="avatar" :size="70"></el-avatar>
        <el-input
          v-model="textarea"
          maxlength="200"
          placeholder="说点什么吧～"
          show-word-limit
          type="textarea"
          :rows="3"
        />
        <el-button
          color="var(--main-btn-bg-color)"
          size="large"
          style="height: 95%"
          @click="sendComment"
          >发送</el-button
        >
      </div>
      <!-- 评论区 -->
      <div class="comment" v-for="comment in comdata" :key="comment">
        <el-divider
          border-style="dotted"
          style="margin: auto; width: 80%; margin-block: 8px"
        />
        <el-col>
          <el-row>
            <div>
              <el-avatar :src="comment.Avatar" :size="50"></el-avatar>
              <div class="cname">{{ comment.Name }}</div>
            </div>
            <div style="width: 80%; margin-top: auto; margin-bottom: auto">
              <el-col>
                <div class="context">{{ comment.Comments }}</div>
                <h5 class="time">{{ comment.Time }}</h5>
              </el-col>
            </div>
          </el-row>
        </el-col>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "CommentsPanel",
  props: {
    id: {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      textarea: "",
      avatar: this.$store.state.userInfo.avatar,
      comdata: [],
    };
  },
  created() {
    this.getcomment();
  },
  methods: {
    sendComment() {
      console.log(this.id);
      fetch("http://127.0.0.1:8521/api/comment/upload", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Uuid: this.$store.state.userInfo.uuid,
          Imgid: this.id,
          Comments: this.textarea,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            this.$message({
              message: "评论成功",
              type: "success",
            });
          }
        });
    },
    getcomment() {
      fetch("http://127.0.0.1:8521/api/comment/get", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Id: this.id,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            this.comdata = JSON.parse(data.comments);
          }
        });
    },
  },
};
</script>

<style scoped>
section {
  display: flex;
  flex-direction: column;
}

.leavecomment {
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  justify-content: left;
  align-items: center;
}

.leavecomment .el-avatar {
  width: 80px;
  height: 70px;
}

.leavecomment > * {
  margin: 10px;
}

.comment {
  margin-left: 75px;
}

.cname {
  text-align: center;
  font-size: 16px;
}

.context {
  font-size: 16px;
  margin: auto;
  margin-left: 20px;
  text-align: left;
  color: var(--main-text-color);
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}

.time {
  font-size: 15px;
  margin: auto;
  margin-left: 20px;
  margin-top: 5px;
  text-align: left;
  color: var(--main-text-color-shade);
}
</style>
