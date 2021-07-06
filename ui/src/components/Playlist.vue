<template>
  <div class="playlist-overview">
    <h1 class="playlist-headline">{{ playlist.name }}</h1>
    <hr />
    <Songs :songs="songs" :playlistId="playlist.id" />
  </div>
</template>

<script>
import { useRoute } from "vue-router";
import { PlaylistService } from "@/services/PlaylistService";
import Songs from "./Songs.vue";

export default {
  name: "Playlist",
  components: {
    Songs,
  },
  data() {
    return {
      playlist: {},
      songs: [],
    };
  },
  watch: {
    "$route.params.id": function (id) {
      // Fetch playlist when component is mounted, but params.id is changed
      if (!id) {
        return;
      }
      this.fetchPlaylist(id);
    },
  },
  methods: {
    fetchPlaylist(id) {
      this.playlist = PlaylistService.fetchPlaylistById(id).then((playlist) => {
        this.playlist = playlist;
        this.songs = playlist.songs;
      });
    },
  },
  mounted() {
    const route = useRoute();
    this.fetchPlaylist(route.params.id);
  },
};
</script>

<style scoped>
.playlist-overview {
  flex-grow: 10;
  background-color: #222831;
  color: #ffffff;
  padding: 1em;
}

.playlist-headline {
  display: flex;
  justify-content: space-around;
  text-transform: uppercase;
  align-self: center;
  font-size: 40pt;
  margin-top: 0;
  margin-bottom: 0.4em;
}

hr {
  border-color: #00adb5;
  background-color: #00adb5;
  margin-bottom: 1em;
}
</style>
