<template>
  <div class="artcard">
    <el-card :body-style="{ padding: '0px' }" shadow="hover">
      <router-link :to="'/art?id=' + artid">
        <el-image :src="arturl" class="image" loading="lazy" />
        <div class="bottom">
          <div style="margin-left: 10px">
            <div>{{ artname }}</div>
            <div>{{ tartauthor }}</div>
          </div>
        </div>
      </router-link>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "ArtCard",
  created() {
    this.getUserName(this.artauthor);
  },
  props: {
    artid: {
      type: Number,
      default: 0,
    },
    artname: {
      type: String,
      default: "作品名",
    },
    artauthor: {
      type: String,
      default: "作者",
    },
    arturl: {
      type: String,
      default:
        "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
    },
  },
  data() {
    return {
      tartauthor: "作者",
    };
  },
  methods: {
    getUserName(artauthor) {
      fetch("http://127.0.0.1:8521/api/user/getname", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          uuid: artauthor,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          this.tartauthor = data.name;
        });
    },
  },
};
</script>

<style>
.artcard {
  width: 200px;
  margin: 10px;
  overflow: hidden;
}
.bottom {
  margin: 5px;
  line-height: 20px;
  text-align: left;
  font-size: 12px;
  color: var(--main-text-color);
}

.image {
  width: 200px;
  height: 250px;
  object-fit: cover;
}

img {
  object-fit: cover;
}
</style>
