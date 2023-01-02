<!-- eslint-disable vue/no-multiple-template-root -->
<template>
  <section>
    <div class="imgdisplay">
      <header>
        <h2>{{ data.Name }}</h2>
        <h3>
          发布时间： {{ data.ReleaseTime }}
          <span class="left"
            ><font-awesome-icon icon="fa-solid fa-eye" />&nbsp; {{ data.Views }} &nbsp;
            <font-awesome-icon icon="fa-solid fa-heart" />&nbsp; {{ data.Likes }} &nbsp;
            <font-awesome-icon icon="fa-solid fa-star" />&nbsp; {{ data.Collects }} &nbsp;
          </span>
        </h3>
      </header>
      <article>
        <el-image
          :src="data.Uri"
          :zoom-rate="1.2"
          :preview-src-list="[data.Uri]"
          :initial-index="4"
          fit="cover"
        />
        <div class="instroduce">
          <p>简介:</p>
          <span>
            {{ data.Describe }}
            <br />
          </span>
          <div class="describetag">
            <el-button
              id="likebtn"
              color="var(--main-btn-bg-color)"
              :round="true"
              :onclick="like"
              >点赞</el-button
            >
            <el-button
              id="collectbtn"
              color="var(--main-btn-bg-color)"
              :onclick="favor"
              :round="true"
              >收藏</el-button
            >&nbsp;
            <div v-for="label in labels" :key="label">{{ label }}</div>
          </div>
        </div>
      </article>
    </div>
    <aside><ArtestCard :uuid="data.OwnerUuid"></ArtestCard></aside>
  </section>
  <section style="margin-top: -60px">
    <CommentsPanel :id="data.ImgId"></CommentsPanel>
  </section>
  <el-divider border-style="dotted" />
</template>
<!-- { { $route.params.id } } -->
<script>
import ArtestCard from "@/components/ArtestCard.vue";
import CommentsPanel from "@/components/CommentsPanel.vue";

export default {
  name: "ArtView",
  components: { ArtestCard, CommentsPanel },
  created() {
    this.loade();
  },
  data() {
    return {
      data: null,
      labels: ["标签1", "标签2", "标签3"],
    };
  },
  methods: {
    loade() {
      fetch("http://127.0.0.1:8521/api/art/getArt?id=" + this.$route.query.id, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            // data转换
            this.data = JSON.parse(data.arts);
            this.labels = this.data.Label.split(",");
          } else {
            alert("获取作品信息失败");
          }
        });
    },
    like() {
      fetch("http://127.0.0.1:8521/api/art/like?id=" + this.$route.query.id, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            this.data.Likes++;
            alert("点赞成功");
          } else {
            alert("获取作品信息失败");
          }
        });
    },
    favor() {
      fetch("http://127.0.0.1:8521/api/collect/collectit", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Imgid: this.data.ImgId,
          Uuid: this.$store.state.userInfo.uuid,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            this.data.Collects++;
            alert("收藏成功！");
          } else {
            alert("获取作品信息失败");
          }
        });
    },
  },
};
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
  width: 1400px;
  margin-top: 10px;
  margin-bottom: 10px;
  box-shadow: var(--main-box-shadow);
  text-align: left;
  overflow: hidden;
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
