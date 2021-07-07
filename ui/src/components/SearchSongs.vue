<template>
  <div class="search-songs">
    <Songs :songs="songs" />
  </div>
</template>

<script>
import { SongService } from "@/services/SongService";
import Songs from "@/components/Songs.vue";

export default {
  name: "SearchSongs",
  components: {
    Songs,
  },
  data() {
    return {
      songs: [],
    };
  },
  watch: {
    "$route.params.search": function (search) {
      if (!search) {
        return;
      }
      this.findSongs(search);
    },
  },
  methods: {
    findSongs(search) {
      SongService.findSongs(search).then((songs) => (this.songs = songs));
    },
  },
};
</script>

<style scoped>
.search-songs {
  flex-grow: 10;
  background-color: #222831;
  color: #ffffff;
  padding: 1em;
}
</style>
