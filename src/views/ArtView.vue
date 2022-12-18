<template>
  <section>
    <div class="imgdisplay">
      <header>
        <h2>{{ name }}</h2>
        <h3>
          发布时间： {{ release_time }}
          <span class="left"
            ><font-awesome-icon icon="fa-solid fa-eye" />&nbsp; {{ views }} &nbsp;
            <font-awesome-icon icon="fa-solid fa-heart" />&nbsp; {{ likes }} &nbsp;
            <font-awesome-icon icon="fa-solid fa-star" />&nbsp; {{ collects }} &nbsp;
          </span>
        </h3>
      </header>
      <article>
        <el-image
          style="width: 100%"
          src="https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg"
          :zoom-rate="1.2"
          :preview-src-list="urls"
          :initial-index="4"
          fit="cover"
        />
        <div class="instroduce">
          <p>简介:</p>
          <span>
            {{ describe }}
            <br />
          </span>
          <div class="describetag">
            <el-button
              id="likebtn"
              color="var(--main-btn-bg-color)"
              :round="true"
              onclick="alert('点赞成功！')"
              >点赞</el-button
            >
            <el-button
              id="collectbtn"
              color="var(--main-btn-bg-color)"
              onclick="alert('收藏成功！')"
              :round="true"
              >收藏</el-button
            >&nbsp;
            <div v-for="label in labels" :key="label">{{ label }}</div>
          </div>
        </div>
      </article>
    </div>
    <aside><ArtestCard></ArtestCard></aside>
  </section>
  <section style="margin-top: -60px">
    <CommentsPanel></CommentsPanel>
  </section>
  <el-divider border-style="dotted" />
</template>
<!-- { { $route.params.id } } -->
<script>
import ArtestCard from "@/components/ArtestCard.vue";
import CommentsPanel from "@/components/CommentsPanel.vue";
const urls = ["https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg"];
const labels = ["标签1", "标签2", "标签3"];

export default {
  name: "ArtView",
  components: { ArtestCard, CommentsPanel },

  props: {
    name: {
      type: String,
      default: "作品名称",
    },

    release_time: {
      type: String,
      default: "2022-12-17",
    },

    views: {
      type: Number,
      default: 0,
    },

    likes: {
      type: Number,
      default: 0,
    },

    collects: {
      type: Number,
      default: 0,
    },

    urls: {
      type: Array,
      default: urls,
    },

    describe: {
      type: String,
      default: "这是一幅画",
    },
    labels: {
      type: Array,
      default: labels,
    },
  },

  methods: {},
};

let load = () => {
  // 向后端请求数据
  //TODO
  //   fetch("http://localhost:8080/api/artinfo=" + this.$route.params.id)
  //     .then((response) => response.json())
  //     .then((data) => {
  //       console.log(data);
  //     });
};

load();
</script>
<style scoped>
section {
  width: 86%;
  display: flex;
  margin: auto;
  margin-bottom: 120px;
}

section > * {
  margin: 10px;
}

aside {
  flex-grow: 1;
  flex-basis: 200px;
}

.imgdisplay {
  flex-grow: 4;
  background-size: cover;
  background-position: center;
  display: flex;
  flex-shrink: 0;
  flex-direction: column;
}

.imgdisplay header {
  width: 100%;
  height: 65px;
  text-align: left;
  padding-left: 15px;
}

.imgdisplay header h2 {
  color: var(--main-text-color);
  margin: 3px;
}

.imgdisplay header h3 {
  color: var(--main-text-color-shade);
  font-size: 1rem;
  font-weight: 500;
  margin: 3px;
}

span.left {
  float: right;
  margin-right: 15px;
  width: fit-content;
  display: inline-block;
}

article {
  border-radius: 10px;
  margin-top: 10px;
  margin-bottom: 10px;
  box-shadow: var(--main-box-shadow);
  text-align: left;
}

.instroduce {
  margin-left: 15px;
  color: var(--main-text-color);
  background-color: var(--main-bg-color);
  align-items: flex-start;
}

.instroduce p {
  font-size: 1.1rem;
  margin: 10px;
}

.instroduce span {
  display: block;
  width: 100%;
  margin-bottom: 10px;
  text-indent: 2em;
}

.artest {
  width: 100%;
  height: 100%;
  background-color: var(--main-bg-color);
  border-radius: 10px;
  box-shadow: var(--main-box-shadow);
}

.describetag {
  margin: 10px;
  width: 100%;
  /* 居中显示 */
  display: flex;
  justify-content: flex-start;
}

.describetag div {
  border-radius: 15px;
  margin-left: 10px;
  background-color: var(--main-theme-color-shade);
  background-attachment: fixed;
  white-space: nowrap; /*强制span不换行*/
  display: inline-block; /*将span当做块级元素对待*/
  overflow: hidden; /*超出宽度部分隐藏*/
  text-align: center;
  padding: 5px 10px 5px 10px;
  font-size: small;
}
</style>
