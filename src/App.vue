<template>
  <header>
    <div class="logo">
      <router-link to="/">
        <font-awesome-icon icon="fa-brands fa-flickr" />
        &nbsp;图片社区
      </router-link>
    </div>
    <div class="center">
      <div class="search">
        <input type="text" placeholder="搜索所想" />
        <font-awesome-icon icon="fa-solid fa-magnifying-glass" />
      </div>
    </div>
    <nav>
      <div>
        <el-avatar
          id="show-usericon"
          class="img"
          @mouseover="showuseraccountpanel = true"
          @mouseleave="hideuseraccountpanel"
          :size="60"
          :src="usericon"
        >
          <img
            src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png"
          />
        </el-avatar>
        <div></div>
      </div>
      <div>
        <router-link to="/upload">
          <font-awesome-icon icon="fa-solid fa-arrow-up-from-bracket" /><br />投稿
        </router-link>
      </div>
      <div>
        <router-link to="/favorite"
          ><font-awesome-icon icon="fa-solid fa-star" /><br />收藏
        </router-link>
      </div>
      <div>
        <router-link to="/arts">
          <font-awesome-icon icon="fa-solid fa-image" /><br />作品
        </router-link>
      </div>
      <div>
        <router-link to="/user">
          <font-awesome-icon icon="fa-solid fa-user" /><br />我的
        </router-link>
      </div>
      <router-link class="bar" to="/more">
        <font-awesome-icon icon="fa-solid fa-bars" />
      </router-link>
    </nav>
  </header>
  <UserAccountPanel
    v-if="showuseraccountpanel || onuseraccountpanel"
    @mouseover="onuseraccountpanel = true"
    @mouseleave="
      () => {
        onuseraccountpanel = false;
        hideuseraccountpanel();
      }
    "
    :name="this.$store.state.userInfo.username"
    :url="this.$store.state.userInfo.avatar"
    :favorites="this.$store.state.userInfo.favorites"
    :arts="this.$store.state.userInfo.arts"
    :views="this.$store.state.userInfo.views"
    :likes="this.$store.state.userInfo.likes"
    :collections="this.$store.state.userInfo.collections"
  />

  <!-- 空隙填充 -->
  <div id="topemp"></div>

  <!-- 路由 -->
  <div>
    <router-view />
  </div>

  <!-- 回到头部 -->
  <div class="backtop" v-if="goTop" @click="backTop">
    <font-awesome-icon icon="fa-solid fa-arrow-up" />
  </div>

  <!-- 页脚 -->
  <footer>
    <div class="footerbox">
      <h1 style="font-family: logofont">图片社区课设项目</h1>
    </div>
    <div class="footerbox">
      <p>© 2022 图片社区课设项目</p>
      <p>地址：锦州市古塔区辽宁工业大学</p>
      <p>邮箱：aptx4869stvsl@outlook.com</p>
    </div>
    <div class="footerbox">
      <p>辽ICP备2022006009号-1</p>
      <p>版权所有: stvsl</p>
      <p>电话：+86 17641544869</p>
    </div>
    <div class="footerbox">
      <h1 style="font-family: logofont">stvsljl.com</h1>
    </div>
  </footer>
</template>

<script>
import UserAccountPanel from "./components/UserAccountPanel.vue";

export default {
  name: "App",
  data() {
    return {
      goTop: false,
      showuseraccountpanel: false,
      onuseraccountpanel: false,
      usericon: "https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg",
    };
  },
  methods: {
    backTop() {
      window.scrollTo({
        top: 0,
        behavior: "smooth",
      });
    },

    // 鼠标离开用户账户面板时，隐藏面板
    hideuseraccountpanel() {
      setTimeout(() => {
        this.showuseraccountpanel = false;
      }, 100);
    },
  },
  mounted() {
    window.addEventListener("scroll", () => {
      if (window.scrollY > 100) {
        this.goTop = true;
      } else {
        this.goTop = false;
      }
    });
  },
  components: { UserAccountPanel },
};
</script>

<style>
@import "./assets/css/defaultcolor.css";
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#topemp {
  height: 80px;
}

body {
  margin: 0;
  background-color: var(--main-bg-color);
  color: var(--main-text-color);
}
</style>

<style scoped>
header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 55px;
  background-color: var(--main-theme-color);
  display: flex;
  box-shadow: 0 0 1px 0 var(--main-text-color);
  z-index: 100;
}

.logo {
  width: 0;
  flex-grow: 1;
  margin: auto;
  margin-left: 20px;
  text-align: left;
  text-shadow: 0 0 1px var(--main-text-color-white);
  font: 2em logofont;
}

.logo a {
  color: var(--main-text-color-white);
}

header .center {
  flex-grow: 1;
  margin: auto;
}

.center .search {
  margin: auto;
  width: 70%;
  height: 30px;
  background-color: var(--main-text-color-white);
  border-radius: 7px;
  display: flex;
  align-items: center;
  padding: 1px 5px;
  opacity: 0.9;
  transition: 0.15s;
}

.center .search:hover {
  padding: 3px 8px;
  opacity: 1;
}

.search svg {
  font-size: 1em;
  color: var(--main-text-color-shade);
  padding: 5px;
}

.search input {
  width: 100%;
  border: none;
  outline: none;
  margin: auto;
  background-color: transparent;
  font-size: 1em;
  color: var(--main-text-color-shade);
}

nav {
  width: 0px;
  flex-grow: 1;
  padding: 0 20px;
  margin: auto;
  text-align: right;
  display: flex;
  justify-content: flex-end;
  opacity: 0.9;
}

nav div {
  padding: 0 18px;
  align-items: center;
}

nav div:hover {
  transform: scale(1.1);
}

nav div svg {
  width: 100%;
  font-size: 1.2em;
  color: var(--main-text-color-white);
}

nav div a {
  font-size: 0.8m;
  flex-grow: 3;
  color: var(--main-text-color-white);
}

.img {
  width: 42px;
  height: 42px;
  margin: auto;
  padding: 0;
  border-radius: 90px;
  float: left;
  box-shadow: var(--big-box-shadow-light);
}

.img:hover {
  box-shadow: 0 0 1px 0 var(--main-text-color-white);
}

.bar {
  width: 40px;
  height: 40px;
  color: var(--main-text-color-white);
}

.bar svg {
  margin: 3px;
  margin-left: 15px;
  font-size: 34px;
  transition: 0.15s;
}

.bar svg:hover {
  transform: rotate(90deg);
}

.backtop {
  position: fixed;
  display: flex;
  bottom: 170px;
  right: 70px;
  width: 50px;
  height: 50px;
  color: var(--main-text-color-shade);
  background-color: var(--main-theme-color-shade);
  border-radius: 10px;
  box-shadow: 0 0 2px 0 var(--main-text-color);
  justify-content: center;
  align-items: center;
  opacity: 0.8;
}

footer {
  width: 100%;
  height: auto;
  display: flex;
  background-color: rgb(80, 80, 80);
  box-shadow: 0 2px rgba(0, 0, 0, 0.2);
  justify-content: center;
}

.footerbox {
  flex-direction: column;
  justify-content: space-between;
  padding: 1em;
  margin: auto;
  color: rgba(255, 255, 255, 0.8);
}
</style>
