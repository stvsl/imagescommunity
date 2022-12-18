<template>
  <section>
    <RouterLink to="/user">
      <el-avatar :size="60" :src="circleUrl" />
      <h4>1{{ artsname }}</h4>
    </RouterLink>
    <div>
      <span>作品数量：{{ num }}1</span>
    </div>
  </section>
</template>
<script>
export default {
  name: "ArtestCard",
  uuid: {
    type: Number,
    default: 0,
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
