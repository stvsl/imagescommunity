<template>
  <div class="panel"></div>
</template>

<script>
export default {
  name: "UserAccountPanel",
  props: {
    url: {
      type: String,
      default: "https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg",
    },
    name: {
      type: String,
      default: "stvsl",
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
}
</style>
