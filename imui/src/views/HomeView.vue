<template>
  <div>
    <el-row>
      <CarouselNavigation />
    </el-row>
    <el-divider border-style="dotted" />
    <el-col>
      <HorizontalSingleDisplayPanel title="动漫专区" subtitle="分享快乐" :arts="ktarts" />
    </el-col>
    <el-divider border-style="dotted" />
    <el-col>
      <HorizontalSingleDisplayPanel title="摄影专区" subtitle="品味生活" :arts="zparts" />
    </el-col>
    <el-divider border-style="dotted" />
    <el-col>
      <HorizontalSingleDisplayPanel title="特效专区" subtitle="体验未来" :arts="tsarts" />
    </el-col>
    <el-divider border-style="dotted" />
    <el-col>
      <HorizontalComplexDisplayPanel
        :arts="parts"
        title="随机推荐"
        subtitle="随机推荐"
      /> </el-col
    ><el-divider border-style="dotted" />
  </div>
</template>

<script>
import CarouselNavigation from "@/components/CarouselNavigation.vue";
import HorizontalSingleDisplayPanel from "@/components/HorizontalSingleDisplayPanel.vue";
import HorizontalComplexDisplayPanel from "@/components/HorizontalComplexDisplayPanel.vue";
export default {
  name: "HomeView",
  components: {
    CarouselNavigation,
    HorizontalSingleDisplayPanel,
    HorizontalComplexDisplayPanel,
  },
  created() {
    // 检查是否登录
    if (this.$store.state.userInfo.isLogin !== true) {
      this.$router.push("/login");
      alert("您尚未登录，请先登录");
    }

    // 获取首页推荐
    fetch("http://127.0.0.1:8521/api/art/recommend", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.status === "success") {
          this.ktarts = JSON.parse(data.ktarts);
          this.tsarts = JSON.parse(data.tsarts);
          this.zparts = JSON.parse(data.zparts);
          this.parts = JSON.parse(data.parts);
        }
      });
  },
  data() {
    return {
      ktarts: [],
      tsarts: [],
      zparts: [],
      parts: [],
    };
  },
};
</script>

<style scoped>
.el-divider {
  width: 70%;
  margin-left: 15%;
}
</style>
