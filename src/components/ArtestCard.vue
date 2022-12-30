<template>
  <section>
    <RouterLink to="/user">
      <el-avatar :size="60" :src="circleUrl" />
      <h4>{{ artsname }}</h4>
    </RouterLink>
    <div>
      <span>UUID</span>
      <br />
      <span>{{ uuid }}</span>
      <br /><br />
      <span>作品数量：{{ num }}</span>
      <br />
    </div>
  </section>
</template>
<script>
export default {
  name: "ArtestCard",
  props: {
    uuid: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      artsname: "1",
      num: 0,
      circleUrl: "",
    };
  },

  created() {
    fetch("http://127.0.0.1:8521/api/user/getuser?uuid=" + this.uuid, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.status === "success") {
          // data转换
          let user = JSON.parse(data.user);
          this.artsname = user.Name;
          this.circleUrl = user.Avatar;
          this.num = user.Actnum;
        } else {
          alert("获取信息失败");
        }
      });
  },
};
</script>
<style scoped>
section {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  box-shadow: var(--main-box-shadow);
  border-radius: 20px;
  padding: 10px;
  margin-left: 20px;
  margin-top: 15px;
}
</style>
