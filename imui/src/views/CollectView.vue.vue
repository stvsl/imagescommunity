<template>
  <div>
    <div>
      <HorizontalComplexDisplayPanelVue
        title="收藏列表"
        subtitle="我的收藏"
        style="height: 100%; min-height: 850px"
        :arts="arts"
      />
    </div>
  </div>
</template>

<script>
import HorizontalComplexDisplayPanelVue from "@/components/HorizontalComplexDisplayPanel.vue";

export default {
  name: "CollectView",
  created() {
    fetch("http://127.0.0.1:8521/api/collect/get", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Uuid: this.$store.state.userInfo.uuid,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.status === "success") {
          this.arts = JSON.parse(data.collects);
        }
      });
  },
  data() {
    return { arts: [] };
  },
  props: {},
  methods: {},
  components: {
    HorizontalComplexDisplayPanelVue,
  },
};
</script>

<style scoped></style>
