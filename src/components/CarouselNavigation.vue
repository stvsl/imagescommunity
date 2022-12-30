<template>
  <div class="carousel">
    <el-carousel :interval="4000" type="card" trigger="click"
      ><el-carousel-item v-for="item in images" :key="item"
        ><router-link :to="'/art?id=' + item.artid"
          ><div class="item">
            <img :src="item.url" alt="无图片" />
            <div class="content">
              <h1 justify="center">{{ item.artnames }}</h1>
              <h2>{{ item.artauthors }}</h2>
              <p>{{ item.artintroductions }}</p>
            </div>
          </div></router-link
        ></el-carousel-item
      ></el-carousel
    >
  </div>
</template>
<script>
export default {
  name: "CarouselNavigation",
  data() {
    return {
      images: [
        {
          url: require("@/assets/imgs/1.png"),
          artnames: "崩坏三三周年",
          artauthors: "stvsl",
          artintroductions: "崩坏三三周年官方壁纸",
          artid: 1000000003,
        },

        {
          url: require("@/assets/imgs/2.png"),
          artnames: "孤独摇滚",
          artauthors: "stvsl",
          artintroductions:
            "孤独摇滚孤独摇滚孤独摇滚孤独摇滚孤独摇滚孤独摇滚孤独摇滚孤独摇滚孤独摇滚孤独摇滚",
          artid: 1000000002,
        },

        {
          url: require("@/assets/imgs/3.png"),
          artnames: "崩坏三六周年",
          artauthors: "stvsl2",
          artintroductions: "崩坏三五周年官方壁纸",
          artid: 1000000006,
        },
      ],
    };
  },
};
</script>
<style scoped>
.el-carousel__item {
  border-radius: 30px;
  overflow: hidden;
}

.carousel {
  width: 75%;
  padding: 5px;
  margin: auto;
  background-color: var(--main-panel-bg-color);
  border-radius: 50px;
  box-shadow: var(--main-box-shadow-light);
}

.item {
  width: 100%;
}

img {
  width: auto;
  height: auto;
  max-width: 100%;
  max-height: 100%;
  object-fit: cover;
  object-position: 50% 50%;
  filter: brightness(0.9);
  border-radius: 30px;
  z-index: -1;
  transition: 0.15s;
}

.content {
  position: absolute;
  top: 40%;
  width: 100%;
  height: 70%;
  color: var(--main-text-color-white);
  text-align: left;
  align-items: inherit;
  border-radius: 30px;
  padding: 0 20px;
  z-index: 1;
  background: linear-gradient(to top, rgba(0, 0, 0, 0.6) 0%, rgba(0, 0, 0, 0) 100%);
  transition: 0.15s;
}

.content:hover {
  text-shadow: 0 0 2px var(--main-text-color-shade);
  background: linear-gradient(to top, rgba(0, 0, 0, 0.7) 0%, rgba(0, 0, 0, 0) 100%);
}

.content > * {
  margin: 0;
  padding: 0;
}

.content h1 {
  font-size: 2em;
  font-weight: 600;
  margin-top: 1.5em;
  margin-bottom: 0.2em;
}

.content h2 {
  font-size: 1.3em;
  font-weight: 400;
  margin-bottom: 0.5em;
}

.content p {
  font-size: 1em;
  font-weight: 300;
  margin-bottom: 1em;
}

img:hover {
  filter: brightness(1);
}
</style>
