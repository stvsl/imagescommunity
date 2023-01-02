<template>
  <div>
    <h1>我的作品</h1>
    <HorizontalComplexDisplayPanelVue
      title="我的作品"
      subtitle="记录你的创作历程"
      style="height: 100%; min-height: 850px"
      :arts="arts"
    />
  </div>
</template>

<script>
import HorizontalComplexDisplayPanelVue from "@/components/HorizontalComplexDisplayPanel.vue";

export default {
  name: "UserView",
  created() {
    fetch("http://127.0.0.1:8521/api/art/findUserArt", {
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
  },
  data() {
    return { arts: [] };
  },
  components: {
    HorizontalComplexDisplayPanelVue,
  },
};
</script>

<style scoped></style>
