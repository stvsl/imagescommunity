<template>
  <section>
    <div
      style="width: 90%; border-radius: 15px; box-shadow: var(--main-box-shadow-light)"
    >
      <div class="leavecomment">
        <el-avatar :src="circleUrl" :size="70"></el-avatar>
        <el-input
          v-model="textarea"
          maxlength="200"
          placeholder="说点什么吧～"
          show-word-limit
          type="textarea"
          :rows="3"
        />
        <el-button color="var(--main-btn-bg-color)" size="large" style="height: 95%"
          >发送</el-button
        >
      </div>
      <!-- 评论区 -->
      <div class="comment" v-for="comment in 9" :key="comment.id">
        <el-divider
          border-style="dotted"
          style="margin: auto; width: 80%; margin-block: 5px"
        />
        <el-col style="background-color: aqua">
          <el-row>
            <div>
              <el-avatar :src="comment.avatar" :size="50"></el-avatar>
              <div class="cname">11</div>
            </div>
            <div>
              <el-col>
                <div class="context">
                  这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：这是一串评论，它很长，有200字：
                </div>
                <h5 class="time">1212</h5>
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
    circleUrl: {
      type: String,
      default: "https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg",
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
  font-size: 20px;
}

.context {
  line-height: 40px;
  font-size: 20px;
  margin: auto;
  color: var(--main-text-color);
}

.time {
  font-size: 15px;
  margin: auto;
  color: var(--main-text-color-shade);
}
</style>
