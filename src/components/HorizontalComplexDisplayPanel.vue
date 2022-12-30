<template>
  <div class="section">
    <div class="aside">
      <div>
        <h2>{{ title }}</h2>
        <h3>{{ subtitle }}</h3>
      </div>
    </div>
    <div class="arts">
      <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
        <div :style="{ display: 'flex', flexWrap: 'wrap' }">
          <el-col v-for="item in arts" :key="item" :span="3.5" style="flex-shrink: 0">
            <ArtCard
              :artid="item.ImgId"
              :artname="item.Name"
              :artauthor="item.OwnerUuid"
              :arturl="item.Uri"
            />
          </el-col>
        </div>
      </ul>
    </div>
  </div>
</template>

<script>
import ArtCard from "@/components/ArtCard.vue";

export default {
  name: "HorizontalSingleDisplayPanel",
  props: {
    title: {
      type: String,
      default: "标题标题",
    },
    subtitle: {
      type: String,
      default: "副标题",
    },
    arts: {
      type: Array,
      default: () => [
        {
          ImgId: 1,
          Name: "作品名称",
          OwnerUuid: "作者",
          Uri: "https://picsum.photos/200/300",
        },
      ],
    },
  },

  components: {
    ArtCard,
  },
};
</script>

<style scoped>
.section {
  margin-top: 15px;
  display: flex;
  height: 700px;
  overflow: hidden;
}

.aside {
  flex-grow: 1;
}

.aside div {
  width: 300px;
  height: 90px;
  max-width: 200px;
  margin-left: 35%;
  margin-top: 15px;
  border-left: 5px solid var(--main-theme-color);
  text-align: left;
  align-items: baseline;
}

.aside h2 {
  margin-left: 20px;
  padding-top: 5px;
}

.aside h3 {
  margin-left: 20px;
}

.arts {
  margin-top: 10px;
  margin-left: 120px;
  flex-grow: 5;
  height: 100%;
  display: flex;
  transition: 0.15s;
  overflow: hidden;
  margin-right: 150px;
}

.arts:hover {
  opacity: 0.9;
}
</style>
