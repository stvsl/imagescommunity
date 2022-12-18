import { createStore } from "vuex";

export default createStore({
  // 用户信息
  state: {
    userInfo: {
      //是否登录
      isLogin: false,
      // 用户名
      username: "",
      // 用户头像
      avatar: "",
      // 用户id
      id: "",
      // 用户uuid
      uuid: "",
      // 性别
      sex: "",
      // 电话
      phone: "",
      // 作品数
      arts: 0,
      // 总收藏数
      favorites: 0,
      // 被收藏
      collections: 0,
      // 被点赞数
      likes: 0,
      // 被浏览数
      views: 0,
    },
  },
  getters: {},
  mutations: {},
  actions: {
    setCookie(state, data) {
      state.userInfo.isLogin = true;
      state.userInfo.username = data.username;
      state.userInfo.avatar = data.avatar;
      state.userInfo.id = data.id;
      state.userInfo.uuid = data.uuid;
    },
  },
  modules: {},
});
