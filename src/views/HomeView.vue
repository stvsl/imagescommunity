<template>
  <el-row>
    <CarouselNavigation />
  </el-row>
  <el-divider border-style="dotted" />
  <el-col>
    <HorizontalSingleDisplayPanel />
  </el-col>
  <el-divider border-style="dotted" />
  <el-col>
    <HorizontalSingleDisplayPanel />
  </el-col>
  <el-divider border-style="dotted" />
  <el-col>
    <HorizontalSingleDisplayPanel />
  </el-col>
  <el-divider border-style="dotted" />
  <el-col> <HorizontalComplexDisplayPanel /> </el-col><el-divider border-style="dotted" />
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
};
</script>

<style scoped>
.el-divider {
  width: 70%;
  margin-left: 15%;
}
</style>
