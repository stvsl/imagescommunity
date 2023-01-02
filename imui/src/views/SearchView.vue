<template>
  <div>
    <h1>{{ keywords }}的搜索结果</h1>
    <HorizontalComplexDisplayPanelVue
      title="搜索结果"
      :subtitle="keywords"
      style="height: 100%; min-height: 850px"
      :arts="arts"
    />
  </div>
</template>

<script>
import HorizontalComplexDisplayPanelVue from "@/components/HorizontalComplexDisplayPanel.vue";
export default {
  name: "SearchView",
  created() {
    fetch("http://127.0.0.1:8521/api/art/searchArt", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Tag: this.keywords,
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
    return {
      search: "",
      arts: [],
      keywords: this.$route.query.keyword,
    };
  },
  components: {
    HorizontalComplexDisplayPanelVue,
  },
};
</script>
