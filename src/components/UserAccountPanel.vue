<template>
  <div class="panel">
    <div class="left">
      <div class="demo-image__preview">
        <el-image style="width: 100px; height: 100px" :src="url" fit="cover" />
      </div>
      <p>用户名</p>
      {{ name }}
    </div>
    <div class="right">
      <ul>
        <li>收藏数：{{ favorites }}</li>
        <li>作品数：{{ arts }}</li>
      </ul>
      <ul>
        <li>被浏览：{{ views }}</li>
        <li>被点赞：{{ likes }}</li>
        <li>被收藏：{{ collections }}</li>
      </ul>
    </div>
    <div class="right">
      <ul>
        <li>
          <p>
            <router-link to="/me">
              <el-button type="danger">个人空间</el-button>
            </router-link>
          </p>
        </li>
        <br />
        <li>
          <p>
            <el-button type="danger" @click="logout">退出登陆</el-button>
          </p>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "UserAccountPanel",
  methods: {
    logout() {
      this.$cookies.remove("uuid");
      this.$router.push("/login");
    },
  },
  props: {
    url: {
      type: String,
      default: "https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg",
    },
    name: {
      type: String,
      default: "stvsl",
    },
    // 收藏数，作品数，被浏览，被点赞，被收藏
    favorites: {
      type: Number,
      default: 0,
    },
    arts: {
      type: Number,
      default: 0,
    },
    views: {
      type: Number,
      default: 0,
    },
    likes: {
      type: Number,
      default: 0,
    },
    collections: {
      type: Number,
      default: 0,
    },
  },
};
</script>

<style>
@import "../assets/css/defaultcolor.css";
.panel {
  position: fixed;
  top: 68px;
  right: 10px;
  width: 400px;
  height: 180px;
  background-color: var(--main-bg-color);
  border-radius: 10px;
  box-shadow: var(--main-box-shadow);
  display: flex;
  z-index: 100;
}

.left {
  width: 120px;
  height: 90%;
  border-radius: 10px 0 0 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

.left .demo-image__preview {
  margin-top: 12px;
  width: 80px;
  height: 80px;
}

.left p {
  margin: 10px 0px 5px 0px;
  padding: 0;
  font-size: 8px;
  color: var(--main-text-color);
}

.left .demo-image__preview {
  border-radius: 50%;
  overflow: hidden;
  box-shadow: var(--big-.0box-shadow);
}

.right {
  width: 30%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
ul {
  list-style: none;
}

ul li {
  margin: 3px;
  font-size: 13px;
  color: var(--main-text-color);
}
</style>
